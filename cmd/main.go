package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/knipers/rate-limiter/internal/config"
	"github.com/knipers/rate-limiter/internal/router"
)

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	mux, err := router.NewRouter(cfg)
	if err != nil {
		log.Fatal("Error when trying to create router:", err)
	}

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
