package domain

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidName  = errors.New("invalid name")
)
