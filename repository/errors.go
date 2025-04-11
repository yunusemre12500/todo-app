package repository

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNoRecords     = errors.New("no records")
	ErrNotFound      = errors.New("not found")
)
