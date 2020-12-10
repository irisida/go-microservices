package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrErr     string `json:"error, omitempty"`
}

func (e *apiError) Status() int {
	return e.ErrStatus
}

func (e *apiError) Message() string {
	return e.ErrMessage
}

func (e *apiError) Error() string {
	return e.ErrErr
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: message,
	}
}

func NewBadRequestApiError(message string) ApiError {
	return &apiError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
	}
}

// NewApiError for untrapped errors
func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		ErrStatus:  statusCode,
		ErrMessage: message,
	}
}
