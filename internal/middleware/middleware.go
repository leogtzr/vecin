package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

func RateLimitMiddlewareAdapter(limiter *rate.Limiter, next http.Handler) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
