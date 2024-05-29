package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"bit-project/gateway/config"

	log "github.com/sirupsen/logrus"
)

func GetKakaoJWK() {
	cfg := config.GetConfig()

	base, err := url.Parse("https://kauth.kakao.com/.well-known/jwks.json")
	if err != nil {
		log.Error("Failed to parse URL:", err)
		return
	}

	resp, err := http.Get(base.String())
	if err != nil {
		log.Error("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed

	sendByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body:", err)
		return
	}

	var jwks struct {
		Keys []config.Key `json:"keys"`
	}

	err = json.Unmarshal(sendByte, &jwks)
	if err != nil {
		log.Error("Failed to unmarshal JSON:", err)
		return
	}

	cfg.KakaoJWKs = jwks.Keys // Assign the keys to the config

	log.Infof("Server config: %#v", *cfg)
}
