package constants

import (
	"errors"
	"net/http"
)

// CodedError is an error wrapper which wraps errors with http status codes.
type CodedError struct {
	Err  error
	Code int
}

func (ce *CodedError) Error() string {
	return ce.Err.Error()
}

func NewCodedError(msg string, code int) *CodedError {
	return &CodedError{errors.New(msg), http.StatusNotFound}
}

var (
	ErrDBNotFound  = &CodedError{errors.New("Not found in db"), http.StatusNotFound}
	ErrLinkExpired = &CodedError{errors.New("Link is expired"), http.StatusConflict}
)
