package models

import "errors"

var (
	// ErrInternalServerError ...
	ErrInternalServerError = errors.New("Internal server error")
	// ErrNotFound ...
	ErrNotFound = errors.New("Your request item not found")
	// ErrConflict ...
	ErrConflict = errors.New("Your item already existing")
	// ErrBadParamInput ...
	ErrBadParamInput = errors.New("Given param is invalid")
)
