// auth/jwt.go
package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sarahrajabazdeh/DreamPilot/config"
)

type JWT struct{ jwtConfig config.JWTConfig }

func NewJWT(jwtConfig config.JWTConfig) *JWT {
	return &JWT{
		jwtConfig: jwtConfig,
	}
}

func (j *JWT) Generate(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(j.jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
