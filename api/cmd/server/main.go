package main

import (
	"log"
	"net/http"
	"os"

	apihttp "github.com/TheSCodex/sezzle-calculator/internal/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:5173"
	}

	mux := http.NewServeMux()

	apihttp.RegisterCalculatorRoutes(mux)

	handler := apihttp.CorsMiddleware(mux, allowedOrigin)

	log.Printf("Server running on :%s", port)

	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
