package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"bit-project/gateway/config"
	"bit-project/gateway/internal/domain"
	"bit-project/gateway/internal/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService domain.UserService
}

type LoginModel struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	cfg := config.GetConfig()

	ex := c.Query("code")
	query := url.Values{}
	query.Add("grant_type", "authorization_code")
	query.Add("code", ex)
	query.Add("redirect_uri", "http://127.0.0.1:9091/login")
	query.Add("client_id", cfg.ClientID)
	query.Add("client_secret", cfg.ClientSecret)

	base, err := url.Parse("https://kauth.kakao.com/oauth/token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	base.RawQuery = query.Encode()
	b := bytes.NewBufferString("")
	resp, err := http.Post(base.String(), "application/x-www-form-urlencoded", b)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	mod := LoginModel{}
	sendByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = json.Unmarshal(sendByte, &mod)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(mod.IdToken, jwt.MapClaims{})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	kid, ok := parsedToken.Header["kid"].(string)
	if !ok {
		log.Error("kid not found in JWT header")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "kid not found in JWT header"})
		return
	}

	jwk, err := utils.FindJWKByKID(kid, cfg.KakaoJWKs)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pubKey, err := utils.CreateRSAKeyFromJWK(jwk.N, jwk.E)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parsed, err := jwt.Parse(mod.IdToken, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || !parsed.Valid {
		log.Error("invalid token claims")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token claims"})
		return
	}
	h.userService.Login(c.Request.Context(), claims)
}
