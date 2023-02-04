package service

import (
	"context"
	"fmt"

	"github.com/Levap123/adverts/internal/repository"
	"github.com/Levap123/adverts/pkg/crypt"
	"github.com/Levap123/adverts/pkg/jwt"
)

type Auth struct {
	repo repository.AuthRepo
}

func NewAuth(repo repository.AuthRepo) *Auth {
	return &Auth{
		repo: repo,
	}
}

const (
	AccessType  = "access"
	RefreshType = "refresh"
)

func (a *Auth) Create(ctx context.Context, email, password string) (int, error) {
	password, err := crypt.GeneratePasswordHash(password)
	if err != nil {
		return 0, fmt.Errorf("service - create user - %w", err)
	}
	userID, err := a.repo.Create(ctx, email, password)
	if err != nil {
		return 0, fmt.Errorf("service - create user - %w", err)
	}
	return userID, nil
}

func (a *Auth) GetTokens(ctx context.Context, email, password string) (string, string, error) {
	user, err := a.repo.Get(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("service - get user - %w", err)
	}
	if err := crypt.ComparePassword(password, user.Password); err != nil {
		return "", "", fmt.Errorf("service - get user - %w", ErrInvalidPassword)
	}
	accessToken, err := jwt.GenerateJwt(user.ID, 1, AccessType)
	if err != nil {
		return "", "", fmt.Errorf("service - get user - %w", err)
	}
	refreshToken, err := jwt.GenerateJwt(user.ID, 24, RefreshType)
	if err != nil {
		return "", "", fmt.Errorf("service - get user - %w", err)
	}
	return accessToken, refreshToken, nil
}
