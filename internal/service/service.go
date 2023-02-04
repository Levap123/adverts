package service

import (
	"context"

	"github.com/Levap123/adverts/internal/entity"
	"github.com/Levap123/adverts/internal/repository"
)

type Service struct {
	AuthService
	AdvertService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService:   NewAuth(repo.AuthRepo),
		AdvertService: NewAdvert(repo.AdvertRepo),
	}
}

type AuthService interface {
	Create(ctx context.Context, email, password string) (int, error)
	GetTokens(ctx context.Context, email, password string) (string, string, error)
}

type AdvertService interface {
	Create(ctx context.Context, title, body string, price, userId int) (int, error)
	GetAll(ctx context.Context, userId int) ([]entity.Advert, error)
	Get(ctx context.Context, advertId int) (entity.Advert, error)
}
