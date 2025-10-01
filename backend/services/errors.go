package services

import "net/http"

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewServiceError(code int, msg string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: msg,
	}
}

var (
	ErrUnauthorized = NewServiceError(http.StatusUnauthorized, "User not authenticated")
	ErrForbidden    = NewServiceError(http.StatusForbidden, "Access denied")
	ErrNotFound     = NewServiceError(http.StatusNotFound, "Resource not found")
)
