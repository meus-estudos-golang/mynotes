package models

import "errors"

var (
	ErrNoRecord           = errors.New("record not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicatedEmail    = errors.New("duplicated email")
)
