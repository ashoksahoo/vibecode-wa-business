package errors

import (
	"fmt"
	"net/http"
)

// Error codes
const (
	ErrInvalidRequest     = "invalid_request"
	ErrUnauthorized       = "unauthorized"
	ErrForbidden          = "forbidden"
	ErrNotFound           = "not_found"
	ErrConflict           = "conflict"
	ErrInternalServer     = "internal_server_error"
	ErrInvalidPhoneNumber = "invalid_phone_number"
	ErrInvalidMessageType = "invalid_message_type"
	ErrWhatsAppAPI        = "whatsapp_api_error"
	ErrRateLimitExceeded  = "rate_limit_exceeded"
	ErrValidationFailed   = "validation_failed"
	ErrDatabaseError      = "database_error"
	ErrMediaUploadFailed  = "media_upload_failed"
	ErrTemplateNotFound   = "template_not_found"
	ErrAPIKeyExpired      = "api_key_expired"
	ErrAPIKeyInvalid      = "api_key_invalid"
)

// AppError represents an application error with additional context
type AppError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Details    map[string]interface{} `json:"details,omitempty"`
	StatusCode int                    `json:"-"`
	Err        error                  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new application error
func NewAppError(code, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Details:    make(map[string]interface{}),
	}
}

// WithError adds an underlying error
func (e *AppError) WithError(err error) *AppError {
	e.Err = err
	return e
}

// WithDetails adds details to the error
func (e *AppError) WithDetails(details map[string]interface{}) *AppError {
	e.Details = details
	return e
}

// WithDetail adds a single detail to the error
func (e *AppError) WithDetail(key string, value interface{}) *AppError {
	if e.Details == nil {
		e.Details = make(map[string]interface{})
	}
	e.Details[key] = value
	return e
}

// NewBadRequest creates a 400 Bad Request error
func NewBadRequest(message string) *AppError {
	return NewAppError(ErrInvalidRequest, message, http.StatusBadRequest)
}

// NewBadRequestWithDetails creates a 400 error with details
func NewBadRequestWithDetails(message string, details map[string]interface{}) *AppError {
	return NewAppError(ErrInvalidRequest, message, http.StatusBadRequest).WithDetails(details)
}

// NewUnauthorized creates a 401 Unauthorized error
func NewUnauthorized(message string) *AppError {
	if message == "" {
		message = "Authentication required"
	}
	return NewAppError(ErrUnauthorized, message, http.StatusUnauthorized)
}

// NewForbidden creates a 403 Forbidden error
func NewForbidden(message string) *AppError {
	if message == "" {
		message = "Access forbidden"
	}
	return NewAppError(ErrForbidden, message, http.StatusForbidden)
}

// NewNotFound creates a 404 Not Found error
func NewNotFound(resource, id string) *AppError {
	message := fmt.Sprintf("%s not found", resource)
	err := NewAppError(ErrNotFound, message, http.StatusNotFound)
	if id != "" {
		err.WithDetail("id", id)
	}
	return err
}

// NewConflict creates a 409 Conflict error
func NewConflict(message string) *AppError {
	return NewAppError(ErrConflict, message, http.StatusConflict)
}

// NewInternalError creates a 500 Internal Server Error
func NewInternalError(err error) *AppError {
	appErr := NewAppError(ErrInternalServer, "An internal error occurred", http.StatusInternalServerError)
	if err != nil {
		appErr.Err = err
	}
	return appErr
}

// NewWhatsAppError creates an error for WhatsApp API failures
func NewWhatsAppError(err error) *AppError {
	appErr := NewAppError(ErrWhatsAppAPI, "WhatsApp API error", http.StatusBadGateway)
	if err != nil {
		appErr.Err = err
		appErr.WithDetail("whatsapp_error", err.Error())
	}
	return appErr
}

// NewValidationError creates a validation error with field details
func NewValidationError(validationErrors map[string]string) *AppError {
	err := NewAppError(ErrValidationFailed, "Validation failed", http.StatusBadRequest)
	details := make(map[string]interface{})
	for k, v := range validationErrors {
		details[k] = v
	}
	return err.WithDetails(details)
}

// NewRateLimitError creates a rate limit exceeded error
func NewRateLimitError() *AppError {
	return NewAppError(ErrRateLimitExceeded, "Rate limit exceeded", http.StatusTooManyRequests)
}

// NewInvalidPhoneNumberError creates an invalid phone number error
func NewInvalidPhoneNumberError(phone string) *AppError {
	return NewAppError(
		ErrInvalidPhoneNumber,
		"Invalid phone number format",
		http.StatusBadRequest,
	).WithDetail("phone", phone)
}

// NewDatabaseError creates a database error
func NewDatabaseError(err error) *AppError {
	return NewAppError(
		ErrDatabaseError,
		"Database operation failed",
		http.StatusInternalServerError,
	).WithError(err)
}

// IsNotFound checks if the error is a not found error
func IsNotFound(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == ErrNotFound
	}
	return false
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == ErrValidationFailed
	}
	return false
}

// IsUnauthorized checks if the error is an unauthorized error
func IsUnauthorized(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == ErrUnauthorized
	}
	return false
}
