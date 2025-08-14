package errors

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type APIError struct {
	Code     int               `json:"code"`
	HTTPCode int               `json:"-"`
	Message  string            `json:"message"`
	Template string            `json:"-"`
	Args     map[string]string `json:"args"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf(e.Template, e.Args)
}

func NewErrInternalServerError(err error) *APIError {
	return &APIError{
		Code:     1500,
		HTTPCode: http.StatusInternalServerError,
		Message:  "Internal server error",
		Template: "Internal server error: {error}",
		Args:     map[string]string{"error": err.Error()},
	}
}

func NewErrParsingBody(err error) *APIError {
	return &APIError{
		Code:     1401,
		HTTPCode: http.StatusBadRequest,
		Message:  "Failed to parse body",
		Template: "Failed to parse body: {error}",
		Args:     map[string]string{"error": err.Error()},
	}
}

func NewErrParsingURI(err error) *APIError {
	return &APIError{
		Code:     1402,
		HTTPCode: http.StatusBadRequest,
		Message:  "Failed to parse URI",
		Template: "Failed to parse URI: {error}",
		Args:     map[string]string{"error": err.Error()},
	}
}

func NewErrEmailIsTaken(email string) *APIError {
	return &APIError{
		Code:     1403,
		HTTPCode: http.StatusBadRequest,
		Message:  "Email is already taken",
		Template: "Email {email} is already taken",
		Args:     map[string]string{"email": email},
	}
}

func NewErrUserNotFound(userID uuid.UUID) *APIError {
	return &APIError{
		Code:     1404,
		HTTPCode: http.StatusNotFound,
		Message:  "User not found",
		Template: "User {userID} not found",
		Args:     map[string]string{"userID": userID.String()},
	}
}

func NewErrRecordNotFound(recordID uuid.UUID) *APIError {
	return &APIError{
		Code:     1405,
		HTTPCode: http.StatusNotFound,
		Message:  "Record not found",
		Template: "Record {recordID} not found",
		Args:     map[string]string{"recordID": recordID.String()},
	}
}

func NewErrMissingAuthorizationToken() *APIError {
	return &APIError{
		Code:     1406,
		HTTPCode: http.StatusUnauthorized,
		Message:  "Missing authorization token",
		Template: "Missing authorization token",
	}
}

func NewErrInvalidAuthorizationToken() *APIError {
	return &APIError{
		Code:     1406,
		HTTPCode: http.StatusUnauthorized,
		Message:  "Invalid authorization token",
		Template: "Invalid authorization token",
	}
}

func NewErrInvalidRecordType(recordType string) *APIError {
	return &APIError{
		Code:     1408,
		HTTPCode: http.StatusBadRequest,
		Message:  "Invalid record type",
		Template: "Invalid record type {recordType}",
		Args: map[string]string{
			"recordType": recordType,
		},
	}
}
