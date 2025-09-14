package errors

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

// APIError represents an error with HTTP/GRPC mapping and templated message.
type APIError struct {
	Code     int               `json:"code"`
	GRPCCode codes.Code        `json:"-"`
	HTTPCode int               `json:"-"`
	Message  string            `json:"message"`
	Template string            `json:"-"`
	Args     map[string]string `json:"args"`
}

// Error returns a formatted error message.
func (e *APIError) Error() string {
	return fmt.Sprintf(e.Template, e.Args)
}

// NewErrInternalServerError creates a 500 Internal Server Error APIError.
func NewErrInternalServerError(err error) *APIError {
	return &APIError{
		Code:     1500,
		HTTPCode: http.StatusInternalServerError,
		GRPCCode: codes.Internal,
		Message:  "Internal server error",
		Template: "Internal server error: {error}",
		Args:     map[string]string{"error": err.Error()},
	}
}

// NewErrParsingBody creates a 400 error for body parsing failures.
func NewErrParsingBody(err error) *APIError {
	return &APIError{
		Code:     1401,
		HTTPCode: http.StatusBadRequest,
		GRPCCode: codes.InvalidArgument,
		Message:  "Failed to parse body",
		Template: "Failed to parse body: {error}",
		Args:     map[string]string{"error": err.Error()},
	}
}

// NewErrParsingURI creates a 400 error for URI parsing failures.
func NewErrParsingURI(err error) *APIError {
	return &APIError{
		Code:     1402,
		HTTPCode: http.StatusBadRequest,
		GRPCCode: codes.InvalidArgument,
		Message:  "Failed to parse URI",
		Template: "Failed to parse URI: {error}",
		Args:     map[string]string{"error": err.Error()},
	}
}

// NewErrEmailIsTaken creates a 400 error for duplicate email.
func NewErrEmailIsTaken(email string) *APIError {
	return &APIError{
		Code:     1403,
		HTTPCode: http.StatusBadRequest,
		GRPCCode: codes.AlreadyExists,
		Message:  "Email is already taken",
		Template: "Email {email} is already taken",
		Args:     map[string]string{"email": email},
	}
}

// NewErrUserNotFound creates a 404 error for missing user.
func NewErrUserNotFound(userParam string) *APIError {
	return &APIError{
		Code:     1404,
		HTTPCode: http.StatusNotFound,
		GRPCCode: codes.NotFound,
		Message:  "User not found",
		Template: "User {userParam} not found",
		Args:     map[string]string{"userParam": userParam},
	}
}

// NewErrRecordNotFound creates a 404 error for missing record.
func NewErrRecordNotFound(recordID uuid.UUID) *APIError {
	return &APIError{
		Code:     1405,
		HTTPCode: http.StatusNotFound,
		GRPCCode: codes.NotFound,
		Message:  "Record not found",
		Template: "Record {recordID} not found",
		Args:     map[string]string{"recordID": recordID.String()},
	}
}

// NewErrMissingAuthorizationToken creates a 401 error for missing token.
func NewErrMissingAuthorizationToken() *APIError {
	return &APIError{
		Code:     1406,
		HTTPCode: http.StatusUnauthorized,
		GRPCCode: codes.Unauthenticated,
		Message:  "Missing authorization token",
		Template: "Missing authorization token",
	}
}

// NewErrInvalidAuthorizationToken creates a 401 error for invalid token.
func NewErrInvalidAuthorizationToken() *APIError {
	return &APIError{
		Code:     1406,
		HTTPCode: http.StatusUnauthorized,
		GRPCCode: codes.Unauthenticated,
		Message:  "Invalid authorization token",
		Template: "Invalid authorization token",
	}
}

// NewErrInvalidRecordType creates a 400 error for invalid record type.
func NewErrInvalidRecordType(recordType string) *APIError {
	return &APIError{
		Code:     1408,
		HTTPCode: http.StatusBadRequest,
		GRPCCode: codes.InvalidArgument,
		Message:  "Invalid record type",
		Template: "Invalid record type {recordType}",
		Args: map[string]string{
			"recordType": recordType,
		},
	}
}

// NewErrSignup creates a 400 error for invalid signup flow.
func NewErrSignup() *APIError {
	return &APIError{
		Code:     1409,
		HTTPCode: http.StatusBadRequest,
		GRPCCode: codes.InvalidArgument,
		Message:  "Signup failed",
		Template: "Signup failed",
	}
}

// NewErrLogin creates a 400 error for invalid login flow.
func NewErrLogin() *APIError {
	return &APIError{
		Code:     1410,
		HTTPCode: http.StatusBadRequest,
		GRPCCode: codes.InvalidArgument,
		Message:  "Login failed",
		Template: "Login failed",
	}
}
