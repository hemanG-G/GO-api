package api

import (
	"encoding/json"
	"net/http"
)

// param
type CoinBalanceParam struct {
	Username string
}

// Successful Response
type CoinBalanceResponse struct {
	//success code (200)
	Code int

	// coin balance value
	Balance int64
}

// Error Response
type Error struct {
	//error code
	Code int

	//error message
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int) {

	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An Unexpected Error Occured ", http.StatusInternalServerError)
	}
)
