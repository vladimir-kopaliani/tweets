package middlewares

import (
	"fmt"
	"net/http"

	apperrors "github.com/vladimir-kopaliani/tweets/auth_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"

	"github.com/golang-jwt/jwt/v5"
)

func CheckAuthentication(next func(http.ResponseWriter, *http.Request) error, jwtSecret string) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		userContext, ok := r.Context().Value(entities.UserContextKey).(entities.UserContext)
		if !ok {
			return apperrors.ErrExtractUserContext
		}

		if userContext.AcessToken == "" {
			return apperrors.ErrAuthenticationRequired
		}

		token, err := jwt.Parse(userContext.AcessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			return fmt.Errorf("error parsing JWT token: %w", err)
		}

		if !token.Valid {
			return apperrors.ErrJWTInvalid
		}

		return next(w, r)
	}
}
