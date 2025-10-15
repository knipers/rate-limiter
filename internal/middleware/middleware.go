package middleware

import (
	"log"
	"net/http"

	"github.com/knipers/rate-limiter/internal/limiter"
)

func NewRateLimiterMiddleware(rl *limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ok, err := rl.Allow(r)
			if err != nil {
				log.Printf("Rate limiter error: %v \n", err)
				http.Error(w, "An internal server error occurred.", http.StatusInternalServerError)
				return
			}
			if !ok {
				log.Println("Rate limit exceeded")
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
