package rest

import (
	"net/http"
)

type RestError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewUnauthorizedRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewConflictError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "conflict",
		Code:    http.StatusConflict,
	}
}

func NewTooManyRequestsError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "too_many_requests",
		Code:    http.StatusTooManyRequests,
	}
}

func NewUnprocessableEntityError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "unprocessable_entity",
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewPreconditionFailedError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "precondition_failed",
		Code:    http.StatusPreconditionFailed,
	}
}

func NewServiceUnavailableError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "service_unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}

func NewGatewayTimeoutError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "gateway_timeout",
		Code:    http.StatusGatewayTimeout,
	}
}

func NewRequestEntityTooLargeError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "request_entity_too_large",
		Code:    http.StatusRequestEntityTooLarge,
	}
}
