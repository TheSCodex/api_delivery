package http

import (
	"encoding/json"
	"github.com/TheSCodex/sezzle-calculator/internal/calculator"
	"net/http"
)

type CalculateRequest struct {
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Operation string  `json:"operation"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {

	var request CalculateRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	result, err := calculator.Calculate(request.A, request.B, request.Operation)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response := CalculateResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
