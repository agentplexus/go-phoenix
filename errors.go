package phoenix

import "errors"

// Sentinel errors for the Phoenix SDK.
var (
	// ErrMissingURL is returned when the API URL is not configured.
	ErrMissingURL = errors.New("phoenix: missing API URL")

	// ErrMissingAPIKey is returned when the API key is required but not provided.
	ErrMissingAPIKey = errors.New("phoenix: missing API key")

	// ErrProjectNotFound is returned when a project cannot be found.
	ErrProjectNotFound = errors.New("phoenix: project not found")

	// ErrTraceNotFound is returned when a trace cannot be found.
	ErrTraceNotFound = errors.New("phoenix: trace not found")

	// ErrSpanNotFound is returned when a span cannot be found.
	ErrSpanNotFound = errors.New("phoenix: span not found")

	// ErrDatasetNotFound is returned when a dataset cannot be found.
	ErrDatasetNotFound = errors.New("phoenix: dataset not found")

	// ErrExperimentNotFound is returned when an experiment cannot be found.
	ErrExperimentNotFound = errors.New("phoenix: experiment not found")

	// ErrPromptNotFound is returned when a prompt cannot be found.
	ErrPromptNotFound = errors.New("phoenix: prompt not found")

	// ErrInvalidInput is returned when input validation fails.
	ErrInvalidInput = errors.New("phoenix: invalid input")
)

// APIError represents an error returned by the Phoenix API.
type APIError struct {
	StatusCode int
	Message    string
	Details    string
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return "phoenix: API error (" + e.Message + "): " + e.Details
	}
	return "phoenix: API error: " + e.Message
}

// IsNotFound returns true if the error indicates a resource was not found.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 404
	}
	return errors.Is(err, ErrProjectNotFound) ||
		errors.Is(err, ErrTraceNotFound) ||
		errors.Is(err, ErrSpanNotFound) ||
		errors.Is(err, ErrDatasetNotFound) ||
		errors.Is(err, ErrExperimentNotFound) ||
		errors.Is(err, ErrPromptNotFound)
}

// IsUnauthorized returns true if the error indicates an authentication failure.
func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsForbidden returns true if the error indicates access is forbidden.
func IsForbidden(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 403
	}
	return false
}

// IsRateLimited returns true if the error indicates rate limiting.
func IsRateLimited(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 429
	}
	return false
}
