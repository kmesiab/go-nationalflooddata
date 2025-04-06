package client

import (
	"fmt"
	"net/http"
)

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Status   int            `json:"status"`  // Error status code
	Message  string         `json:"message"` // Error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

type InvalidRequestError struct {
	*ErrorResponse
}

func (e *InvalidRequestError) Error() string {
	return fmt.Sprintf("Invalid request: %s", e.ErrorResponse.Message)
}

type AuthenticationError struct {
	*ErrorResponse
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication error: %s", e.ErrorResponse.Message)
}

type NoDataAvailableError struct {
	*ErrorResponse
}

func (e *NoDataAvailableError) Error() string {
	return fmt.Sprintf("No data available: %s", e.ErrorResponse.Message)
}

type LocationNotFoundError struct {
	*ErrorResponse
}

func (e *LocationNotFoundError) Error() string {
	return fmt.Sprintf("Location not found: %s", e.ErrorResponse.Message)
}

type ParcelNotFoundError struct {
	*ErrorResponse
}

func (e *ParcelNotFoundError) Error() string {
	return fmt.Sprintf("Parcel not found: %s", e.ErrorResponse.Message)
}

type InternalServerError struct {
	*ErrorResponse
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Internal server error: %s", e.ErrorResponse.Message)
}
