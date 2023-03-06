package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"
)

func InjectUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authHeader  = r.Header.Get("Authorization")
			accessToken = strings.TrimPrefix(authHeader, "Bearer ")
		)

		ctx := context.WithValue(
			r.Context(),
			entities.UserContextKey,
			entities.UserContext{
				AcessToken: accessToken,
			},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
