package service

import "errors"

var (
	ErrInvalidPassword = errors.New("password is not correct")
	ErrInorrectTitle   = errors.New("incorrect title")
	ErrPriceSmall      = errors.New("bet price is to small")
)
