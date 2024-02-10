package handler

import (
	"encoding/json"
	"net/http"
)

type JsonBody struct {
	HttpStatus string      `json:"status"`
	HttpCode   int         `json:"code,omitempty"`
	Payload    interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func JsonResponse(w http.ResponseWriter, body JsonBody) {
	response, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(body.HttpCode)
	w.Write(response)
}
