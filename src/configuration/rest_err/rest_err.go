// Package rest_err defines structures and methods for handling REST API errors.
package rest_err

import "net/http"

// RestErr represents a standard error structure for API responses.
// Fields:
// - Message: The human-readable error message.
// - Err: A short description of the error type.
// - Code: HTTP status code associated with the error.
// - Causes: A slice of detailed error causes.
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

// Causes represents the details of a specific error cause.
// Fields:
// - Field: The field name related to the error.
// - Message: A description of the error.
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error implements the error interface for RestErr.
func (r *RestErr) Error() string {
	return r.Message
}

// NewRestError creates a new RestErr with custom details.
func NewRestError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestError creates a RestErr for bad request scenarios.
func NewBadRequestError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestValidationError creates a RestErr for validation errors.
func NewBadRequestValidationError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request validation",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// NewInternalServerError creates a RestErr for internal server errors.
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFoundError creates a RestErr for not found errors.
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found error",
		Code:    http.StatusNotFound,
	}
}

// NewForbiddenError creates a RestErr for forbidden access errors.
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden error",
		Code:    http.StatusForbidden,
	}
}
