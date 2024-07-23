package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

var store *sessions.CookieStore

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		log.Fatal("SESSION_KEY is not set in the environment variables")
	}

	// TODO: improve the store.
	store = sessions.NewCookieStore([]byte(sessionKey))
}

func GetSessionStore() *sessions.CookieStore {
	return store
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		if session.Values["user_id"] == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
