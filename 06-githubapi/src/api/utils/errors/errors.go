package errors

import "net/http"

// APIError - error interface
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrErr     string `json:"error,omitempty"`
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

// NewNotFoundAPIError - raised on a not found case
func NewNotFoundAPIError(message string) APIError {
	return &apiError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
	}
}

// NewInternalServerError - raised on an internal server error
func NewInternalServerError(message string) APIError {
	return &apiError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: message,
	}
}

// NewBadRequestAPIError - raised for a bad request
func NewBadRequestAPIError(message string) APIError {
	return &apiError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
	}
}

// NewAPIError for untrapped errors
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		ErrStatus:  statusCode,
		ErrMessage: message,
	}
}
