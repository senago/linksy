package constants

import (
	"errors"
	"net/http"
)

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
	ErrInternal    = &CodedError{errors.New("Internal server error"), http.StatusInternalServerError}

	ErrDuplicateHash = errors.New("Duplicate hash value")
)
