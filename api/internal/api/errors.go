package api

import (
	"encoding/json"
	"net/http"
)

//Error ...
type Error struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

//HandleErrorAsJSON ...
func HandleErrorAsJSON(w http.ResponseWriter, e Error) {
	data, err := json.Marshal(e)
	if err != nil {
		HandleErrorAsPlainText(w, e.Status, "marshalling went wrong in HandleErrorAsJSON; "+e.Msg)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(data)
}

//HandleErrorAsPlainText ...
func HandleErrorAsPlainText(w http.ResponseWriter, code int, msg string) {
	http.Error(w, msg, code)
}
