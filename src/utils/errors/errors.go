package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIError interface {
	Message() string
	Status() int
	Error() string
}

type apiError struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status_code"`
	Err        string `json:"error"`
}

func (e *apiError) Message() string {
	return e.ErrMessage
}

func (e *apiError) Status() int {
	return e.ErrStatus
}

func (e *apiError) Error() string {
	return e.Err
}

func NewError(message string) error {
	return errors.New(message)
}

func NewErrorFromBytes(bytes []byte) (APIError, error) {
	var apiErr apiError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &apiErr, nil
}

func NewBadRequestError(message string) APIError {
	return &apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		Err:        "bad_request",
	}
}

func NewNotFoundError(message string) APIError {
	return &apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		Err:        "not_found",
	}
}

func NewInternalServerError(message string) APIError {
	return &apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		Err:        "internal_server_error",
	}
}
