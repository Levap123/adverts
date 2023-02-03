package service

import (
	"context"
	"fmt"
	"net/mail"
	"unicode"

	"github.com/Levap123/adverts/internal/repository"
	"github.com/Levap123/adverts/pkg/crypt"
)

type Auth struct {
	repo repository.AuthRepo
}

func NewAuth(repo repository.AuthRepo) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (ar *Auth) Create(ctx context.Context, email, password string) (int, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return 0, fmt.Errorf("service - create user - %w", ErrInvalidEmail)
	}
	if !ar.isPasswordValid(password) {
		return 0, fmt.Errorf("service - create - user - %w", ErrInvalidPassword)
	}
	password, err := crypt.GeneratePasswordHash(password)
	if err != nil {
		return 0, fmt.Errorf("service - create user - %w", err)
	}
	userID, err := ar.repo.Create(ctx, email, password)
	if err != nil {
		return 0, fmt.Errorf("service - create user - %w", err)
	}
	return userID, nil
}

func (ar *Auth) isPasswordValid(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case !unicode.IsLetter(char) && !unicode.IsDigit(char):
			hasSpecial = true
		}
	}

	return len(password) >= 8 && hasUpper && hasLower && hasDigit && hasSpecial
}
