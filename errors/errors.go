package errors

import "fmt"

type APIError struct {
	Message    string
	StatusCode int
	Response   map[string]interface{}
	Err        error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error [%d]: %s", e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

func (e *APIError) GetResponse() map[string]interface{} {
	return e.Response
}

type AuthenticationError struct {
	*APIError
}

func NewAuthenticationError(message string, response map[string]interface{}) *AuthenticationError {
	return &AuthenticationError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 401,
			Response:   response,
		},
	}
}

type RateLimitError struct {
	*APIError
	RetryAfter *int
}

func NewRateLimitError(message string, retryAfter *int, response map[string]interface{}) *RateLimitError {
	return &RateLimitError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 429,
			Response:   response,
		},
		RetryAfter: retryAfter,
	}
}

func (e *RateLimitError) GetRetryAfter() *int {
	return e.RetryAfter
}

type InvalidRequestError struct {
	*APIError
}

func NewInvalidRequestError(message string, response map[string]interface{}) *InvalidRequestError {
	return &InvalidRequestError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 400,
			Response:   response,
		},
	}
}

type NotFoundError struct {
	*APIError
}

func NewNotFoundError(message string, response map[string]interface{}) *NotFoundError {
	return &NotFoundError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 404,
			Response:   response,
		},
	}
}

func NewAPIError(message string, statusCode int, response map[string]interface{}, err error) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
		Response:   response,
		Err:        err,
	}
}
