// middleware.go
package router

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/sarahrajabazdeh/DreamPilot/auth"
	"github.com/sarahrajabazdeh/DreamPilot/config"
	"github.com/sarahrajabazdeh/DreamPilot/dreamerr"
	"github.com/sarahrajabazdeh/DreamPilot/model"
	"github.com/sarahrajabazdeh/DreamPilot/utility"
)

// JWTMiddleware is the JWT authentication middleware
func JWTMiddleware(jwtConfig config.TokenConfig) func(http.Handler) http.Handler {
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

func checkAuthMdlw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var tokenStr string

		if tokenStr = r.Header.Get(model.TokenHeader); tokenStr == "" {
			handleError(r, w, dreamerr.ErrMissingToken(), dreamerr.LogMessageErrorResponse)
			return
		}

		if _, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Token.Secret), nil
		}); err != nil {
			handleError(r, w, dreamerr.ErrInvalidToken(), dreamerr.LogMessageErrorResponse)
			return
		}

		// Serve the request
		next.ServeHTTP(w, r)
	})
}

func handleError(r *http.Request, w http.ResponseWriter, err error, errLogType string) {
	err = dreamerr.PropagateError(err, 2)

	_, ok := r.Context().Value(model.GetCtxKeyID()).(uuid.UUID)
	if !ok {
		panic(dreamerr.ErrServerError)
	}

	appErr, ok := err.(*dreamerr.DreamError)
	if !ok {
		appErr = dreamerr.ErrServerError()
	}

	// appErr.Log(errLogType)

	json := appErr.MarhsalJSON()
	w.Header().Set("Content-Type", utility.MimeTypeJSON)
	w.WriteHeader(appErr.Status())
	_, _ = w.Write(json)
}
