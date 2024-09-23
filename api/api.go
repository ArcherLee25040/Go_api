package api

import (
	"encoding/json"
	"net/http"
)

// Coint Balance Params
type CoinBalanceParms struct {
	Username string
}

// Coin Balance response
type CoinBalanceResponse struct {
	//Success response , Usually 200
	Code int

	// Account Balance
	Balance float64
}

// Error Response
type Error struct {
	//Error code
	Code int

	//Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Connect-Type:", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)

	}
	internalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexcepted Error Occurred.", http.StatusInternalServerError)

	}
)
