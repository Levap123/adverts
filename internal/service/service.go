package service

import (
	"context"

	"github.com/Levap123/adverts/internal/repository"
)

type Service struct {
	AuthService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuth(repo.AuthRepo),
	}
}

type AuthService interface {
	Create(ctx context.Context, email, password string) (int, error)
	GetTokens(ctx context.Context, email, password string) (string, string, error)
}
