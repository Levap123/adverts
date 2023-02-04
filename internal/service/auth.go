package service

import (
	"context"
	"fmt"

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
