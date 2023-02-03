package service

import "github.com/Levap123/adverts/internal/repository"

type Service struct {
	AuthService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		AuthService: NewAuth(repo.AuthRepo),
	}
}

type AuthService interface {
	Create(email, password string) (int, error)
}
