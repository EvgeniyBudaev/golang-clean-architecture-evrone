package middleware

import (
	"context"
	"net/http"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
)

func AddContextMiddleware(log logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
