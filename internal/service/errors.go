package service

import "errors"

var (
	ErrInvalidEmail   = errors.New("user email is invalid")
	ErrInvalidPassword = errors.New("password should contain 8 characters: at least 1 symbol, 1 uppercase letter, 1 lowercase letter, 1 number")
)
