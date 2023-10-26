package utils

import (
	"encoding/json"
	"net/http"
)

func Response(responseWriter http.ResponseWriter, statusCode int) {

	responseWriter.Header().Add("Content-Type", "appliction/json")
	responseWriter.WriteHeader(statusCode)
}

func ResponseAsBody(responseWriter http.ResponseWriter, statusCode int, responseBody any) {

	Response(responseWriter, statusCode)

	if err := json.NewEncoder(responseWriter).Encode(responseBody); err != nil {

	}
}

func ErrorResponse(responseWriter http.ResponseWriter, statusCode int, err error) {

	responseError := struct{ Error string }{err.Error()}
	ResponseAsBody(responseWriter, statusCode, responseError)
}
