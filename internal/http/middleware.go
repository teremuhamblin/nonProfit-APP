package http

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/teremuhamblin/nonProfit-APP/internal/config"
	"github.com/teremuhamblin/nonProfit-APP/internal/security"
	"go.uber.org/zap"
)

func CommonMiddlewares(logger *zap.Logger, cfg *config.Config) []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		httprate.LimitByIP(100, 60),
	}
}

func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(auth, "Bearer ")
			claims, err := security.ParseToken(secret, tokenStr)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = SetUserID(ctx, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
