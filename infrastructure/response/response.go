package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response is a HTTP response with an optional Data body and a StatusCode.
type Response interface {
	// StatusCode returns the HTTP status code of the response
	StatusCode() int

	// Data returns the response body data
	Data() interface{}
}

// SuccessResponse is a Response specific for successful HTTP responses.
type SuccessResponse struct {
	statusCode int
	data       interface{}
}

// StatusCode returns the HTTP status code of the response
func (e SuccessResponse) StatusCode() int { return e.statusCode }

// Data returns the response body data
func (e SuccessResponse) Data() interface{} { return e.data }

// MarshalJSON is a custom JSON serialization function for SuccessResponse
func (e SuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.data)
}

// Success returns a new SuccessResponse with a statusCode and some data that should go in the response body
func Success(statusCode int, data interface{}) Response {
	return SuccessResponse{
		statusCode: statusCode,
		data:       data,
	}
}

// OK returns a new SuccessResponse with a http.StatusOK status code
func OK(data interface{}) Response {
	return Success(http.StatusOK, data)
}

// Created returns a new SuccessResponse with a http.StatusCreated status code
func Created(data interface{}) Response {
	return Success(http.StatusCreated, data)
}

// ErrorResponse is a Response specific for failure HTTP responses.
type ErrorResponse struct {
	statusCode int
	message    string
}

// StatusCode returns the HTTP status code of the response
func (e ErrorResponse) StatusCode() int { return e.statusCode }

// Data returns the response body data
func (e ErrorResponse) Data() interface{} { return e.message }

// MarshalJSON is a custom JSON serialization function for ErrorResponse
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"message": e.message,
	})
}

// Error returns a new ErrorResponse with a statusCode, and a message and error that should go in the response body
func Error(statusCode int, message string, err error) Response {
	return ErrorResponse{
		statusCode: statusCode,
		message:    fmt.Sprintf("%s: %v", message, err),
	}
}

// BadRequest returns a new ErrorResponse with a http.StatusBadRequest status code
func BadRequest(message string, err error) Response {
	return Error(http.StatusBadRequest, message, err)
}

// NotFound returns a new ErrorResponse with a http.StatusNotFound status code
func NotFound(message string, err error) Response {
	return Error(http.StatusNotFound, message, err)
}

// InternalServerError returns a new ErrorResponse with a http.StatusInternalServerError status code
func InternalServerError(message string, err error) Response {
	return Error(http.StatusInternalServerError, message, err)
}

// Forbidden returns a new ErrorResponse with a http.StatusForbidden status code
func Forbidden(message string, err error) Response {
	return Error(http.StatusForbidden, message, err)
}

// Unauthorized returns a new ErrorResponse with a http.StatusUnauthorized status code
func Unauthorized(message string, err error) Response {
	return Error(http.StatusUnauthorized, message, err)
}
