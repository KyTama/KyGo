package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ResponseStruct model response
type ResponseStruct struct {
	Message string		`json:"message"`
	Data interface{}	`json:"data,omitempty"`
	Time string
}

// HTTPResponse to give http response
func HTTPResponse(w http.ResponseWriter, message string, data interface{}, statusCode int, timeIn string, body io.ReadCloser) {
	responseTemplate := ResponseStruct{
		Message: message,
		Data:    data,
		Time:    timeIn,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	resp, err := json.Marshal(responseTemplate)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
}
