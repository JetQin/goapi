package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Logger returns a middleware that logs basic request info using the provided logger.
func Logger(logger *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.WithFields(log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"remote": r.RemoteAddr,
			}).Info("incoming request")
			next.ServeHTTP(w, r)
		})
	}
}
