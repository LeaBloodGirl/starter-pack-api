package middleware

import (
	"net/http"
	"starter-pack-api/internal/logger"
)

type LoggingData struct {
	LogDebugStored *logger.Logger
	LogStatStored  *logger.Logger
}

func (l LoggingData) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
