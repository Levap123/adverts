package service

import "errors"

var (
	ErrInvalidPassword = errors.New("password is not correct")
	ErrInorrectTitle   = errors.New("incorrect title")
)
