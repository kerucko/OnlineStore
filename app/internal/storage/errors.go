package storage

import "errors"

var (
	ErrNotExist     = errors.New("Not exist")
	ErrAlreadyExist = errors.New("Already exist")
)
