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

	// ErrGameNotFound is a specific error for when a game is not found.
	ErrGameNotFound = errors.New("game not found")

	// ErrPlayerNotFoundInGame is for when a player isn't part of a game.
	ErrPlayerNotFoundInGame = errors.New("player not found in game")

	// ErrGameNotJoinable indicates a game cannot be joined (e.g., ongoing, full).
	ErrGameNotJoinable = errors.New("game is not joinable")

	// ErrGameNotOngoing indicates an action cannot be performed because the game is not ongoing.
	ErrGameNotOngoing = errors.New("game is not ongoing")

	// ErrNoMoreNumbers indicates all numbers have been drawn.
	ErrNoMoreNumbers = errors.New("no more numbers to draw")

	// ErrInvalidBingoClaim indicates a bingo claim was not valid.
	ErrInvalidBingoClaim = errors.New("invalid bingo claim")
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
