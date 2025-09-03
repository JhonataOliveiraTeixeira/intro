package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct{
	Error string `json:"error,omitempty"`
	Data any	`json:"data,omitempty"`
}

func SendJSON(w http.ResponseWriter, resp Response, status int){
	data, err := json.Marshal(resp)
	if err != nil{
		slog.Error("Failed to marshal json Data", "Error", err)
		SendJSON(w, Response{Error: "Somenthing went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data);err!= nil{
		slog.Error("Failed to write response to client", "Error", err)
		return
	}

}