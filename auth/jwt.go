// auth/jwt.go
package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
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

func (j *JWT) Authenticate(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method used in the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		// Return the secret key used for signing the token
		return []byte(j.jwtConfig.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	return userID, nil
}
