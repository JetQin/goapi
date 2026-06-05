package api

import (
	"encoding/json"
	"net/http"
)

// Coint Balance Params
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {

	// Success code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

type Error struct {

	// Error code, Usually 400 or 500
	Code int

	// Error message
	Message string
}

func writeErrorResponse(w http.ResponseWriter, message string, code int) {

	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHander = func(w http.ResponseWriter, err error) {
		writeErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
