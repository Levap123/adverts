package service

import "github.com/Levap123/adverts/internal/repository"

type AuthRepo struct {
	repo repository.AuthRepo
}

func NewAuth(repo repository.AuthRepo) *AuthRepo {
	return &AuthRepo{
		repo: repo,
	}
}

func (ar *AuthRepo) Create(email, password string) (int, error) {
	
}
