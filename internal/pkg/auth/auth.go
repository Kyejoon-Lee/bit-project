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

	base, err := url.Parse("https://kauth.kakao.com/oauth/token")
	if err != nil {
		log.Println(err)
	}

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println(err)
	}
	sendByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	json.Unmarshal(sendByte, &cfg.KakaoJWKs)
}
