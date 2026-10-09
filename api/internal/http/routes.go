package http

import (
	"net/http"
)

func RegisterCalculatorRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /calculate", Calculate)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
}