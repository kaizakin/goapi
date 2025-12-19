package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct{
	code int
	message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		code: code,
		message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var(// this is the syntax to group all var declarations its the same as declaring variables multiple times just used () instead
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	RequestInternalHandler = func(w http.ResponseWriter){
		writeError(w,"An unexpected error occured!!", http.StatusInternalServerError)
	}
)