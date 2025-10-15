package router

import (
	"net/http"

	"github.com/knipers/rate-limiter/internal/config"
	"github.com/knipers/rate-limiter/internal/limiter"
	"github.com/knipers/rate-limiter/internal/limiter/strategy"
	"github.com/knipers/rate-limiter/internal/middleware"
)

func NewRouter(cfg *config.Config) (http.Handler, error) {
	rs, err := strategy.NewRedisStrategy(cfg)
	if err != nil {
		return nil, err
	}

	rateLimiter := limiter.NewRateLimiter(cfg, rs)

	mux := http.NewServeMux()
	mux.Handle("/", middleware.NewRateLimiterMiddleware(rateLimiter)(
		http.HandlerFunc(indexHandler),
	))

	return mux, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Requisição aceita!"))
}
