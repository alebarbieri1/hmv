package response

import (
	"encoding/json"
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

func Error(statusCode int, message string) Response {
	return ErrorResponse{
		statusCode: statusCode,
		message:    message,
	}
}

func BadRequest(message string) Response {
	return Error(http.StatusBadRequest, message)
}

func NotFound(message string) Response {
	return Error(http.StatusNotFound, message)
}

func InternalServerError(message string) Response {
	return Error(http.StatusInternalServerError, message)
}

func Forbidden(message string) Response {
	return Error(http.StatusForbidden, message)
}

func Unauthorized(message string) Response {
	return Error(http.StatusUnauthorized, message)
}
