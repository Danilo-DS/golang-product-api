package utils

import (
	"encoding/json"
	"net/http"
)

// Response return response without body with status code only
func Response(responseWriter http.ResponseWriter, statusCode int) {

	responseWriter.Header().Add("Content-Type", "appliction/json")
	responseWriter.WriteHeader(statusCode)
}

// ResponseAsBody return response with body and status code
func ResponseAsBody(responseWriter http.ResponseWriter, statusCode int, responseBody any) {

	Response(responseWriter, statusCode)

	if err := json.NewEncoder(responseWriter).Encode(responseBody); err != nil {
		ErrorResponse(responseWriter, http.StatusInternalServerError, err)
	}
}

// ErrorResponse return error response with body and status code referring to the error
func ErrorResponse(responseWriter http.ResponseWriter, statusCode int, err error) {

	responseError := struct{ Error string }{err.Error()}
	ResponseAsBody(responseWriter, statusCode, responseError)
}
