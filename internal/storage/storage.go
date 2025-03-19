package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	errURLExists   = errors.New("url exists")
)
