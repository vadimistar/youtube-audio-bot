package objectstorage

import "errors"

var (
	ErrUnknownEndpoint = errors.New("unknown endpoint")
	ErrLoadConfig      = errors.New("load config")
)
