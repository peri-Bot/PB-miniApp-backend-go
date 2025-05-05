package errors

import "errors"

// Pre-defined application errors for common scenarios.
// These can be checked using errors.Is().
var (
	// ErrNotFound indicates a requested resource was not found.
	ErrNotFound = errors.New("resource not found")

	// ErrUnauthorized indicates missing or invalid authentication credentials.
	ErrUnauthorized = errors.New("unauthorized: authentication required or invalid")

	// ErrForbidden indicates the authenticated user does not have permission.
	ErrForbidden = errors.New("forbidden: permission denied")

	// ErrValidation indicates invalid input data.
	// Often, more specific validation errors with details are preferred.
	ErrValidation = errors.New("validation failed: invalid input data")

	// ErrConflict indicates an attempt to create a resource that already exists
	// or violates a unique constraint.
	ErrConflict = errors.New("conflict: resource already exists or violates constraint")

	// ErrInternalServer indicates an unexpected error occurred on the server.
	// Specific details should be logged, but this generic error returned to the client.
	ErrInternalServer = errors.New("internal server error")

	// ErrBadRequest indicates a malformed or invalid request from the client.
	ErrBadRequest = errors.New("bad request: invalid request format or parameters")
)

// Consider adding custom error types if you need to attach more context,
// for example, validation errors with specific field details:
/*
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}
*/
