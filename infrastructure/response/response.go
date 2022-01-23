package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response interface {
	StatusCode() int
	Data() interface{}
}

type SuccessResponse struct {
	statusCode int
	data       interface{}
}

func (e SuccessResponse) StatusCode() int { return e.statusCode }

func (e SuccessResponse) Data() interface{} { return e.data }

func (e SuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.data)
}

func Success(statusCode int, data interface{}) Response {
	return SuccessResponse{
		statusCode: statusCode,
		data:       data,
	}
}

func OK(data interface{}) Response {
	return Success(http.StatusOK, data)
}

func Created(data interface{}) Response {
	return Success(http.StatusCreated, data)
}

type ErrorResponse struct {
	statusCode int
	message    string
}

func (e ErrorResponse) StatusCode() int { return e.statusCode }

func (e ErrorResponse) Data() interface{} { return e.message }

func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"message": e.message,
	})
}

func Error(statusCode int, message string, err error) Response {
	return ErrorResponse{
		statusCode: statusCode,
		message:    fmt.Sprintf("%s: %v", message, err),
	}
}

func BadRequest(message string, err error) Response {
	return Error(http.StatusBadRequest, message, err)
}

func NotFound(message string, err error) Response {
	return Error(http.StatusNotFound, message, err)
}

func InternalServerError(message string, err error) Response {
	return Error(http.StatusInternalServerError, message, err)
}

func Forbidden(message string, err error) Response {
	return Error(http.StatusForbidden, message, err)
}

func Unauthorized(message string, err error) Response {
	return Error(http.StatusUnauthorized, message, err)
}
