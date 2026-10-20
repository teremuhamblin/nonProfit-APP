package middleware

import (
    "net/http"
    "strings"

    "github.com/teremuhamblin/nonProfit-APP/internal/security"
)

func JWTAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "missing authorization header", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "invalid authorization header", http.StatusUnauthorized)
            return
        }

        claims, err := security.ValidateToken(parts[1])
        if err != nil {
            http.Error(w, "invalid token", http.StatusUnauthorized)
            return
        }

        // Inject user ID into context
        ctx := r.Context()
        ctx = context.WithValue(ctx, "userID", claims.UserID)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
