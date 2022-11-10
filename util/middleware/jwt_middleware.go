package middleware

import (
	"github.com/golang-jwt/jwt"
	"music-api-go/config"
	"time"
)

func CreateToken(username, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(12 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.TokenSecret))
}
