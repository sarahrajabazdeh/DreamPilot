// middleware.go
package middleware

import (
	"fmt"
	"net/http"

	"github.com/sarahrajabazdeh/DreamPilot/auth"
	"github.com/sarahrajabazdeh/DreamPilot/config"
)

// JWTMiddleware is the JWT authentication middleware
func JWTMiddleware(jwtConfig config.JWTConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")

			jwtManager := auth.NewJWT(jwtConfig)

			// Authenticate the token
			_, err := jwtManager.Authenticate(tokenString)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	// This is a protected endpoint, and the user is authenticated here.

	userID := r.Context().Value("userID").(string)
	fmt.Fprintf(w, "Private endpoint accessed by user with ID: %s", userID)
}
