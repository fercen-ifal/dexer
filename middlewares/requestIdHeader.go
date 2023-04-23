package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ContextKey string

func RequestIDHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const key ContextKey = "X-Request-Id"
		requestId := uuid.NewString()

		w.Header().Set("X-Request-Id", requestId)
		ctx := context.WithValue(r.Context(), key, requestId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
