package service

import (
	"context"
	"strings"

	"github.com/Levap123/adverts/internal/entity"
	"github.com/Levap123/adverts/internal/repository"
)

type Advert struct {
	repo repository.AdvertRepo
}

func NewAdvert(repo repository.AdvertRepo) *Advert {
	return &Advert{
		repo: repo,
	}
}

func (a *Advert) Create(ctx context.Context, title, body string, price, userId int) (int, error) {
	banWords := []string{"nigger", "faggot", "simp"}

	for ind := range banWords {
		if strings.Contains(title, banWords[ind]) {
			return 0, ErrInorrectTitle
		}
	}
	
	return a.repo.Create(ctx, title, body, price, userId)
}

func (a *Advert) GetAll(ctx context.Context, userId int) ([]entity.Advert, error) {
	return a.repo.GetAll(ctx, userId)
}

func (a *Advert) Get(ctx context.Context, advertId int) (entity.Advert, error) {
	return a.repo.Get(ctx, advertId)
}
