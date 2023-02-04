package jwt

import "errors"

var (
	ErrExpired       = errors.New("token expired")
	ErrInvalidClaims = errors.New("invalid claims")
)
