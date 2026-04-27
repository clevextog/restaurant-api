package user

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailValidation    = errors.New("not valid email")
	ErrPasswordValidation = errors.New("not valid password")
)
