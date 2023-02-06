package repository

import (
	"context"

	"github.com/Levap123/adverts/internal/entity"
	"github.com/Levap123/adverts/internal/repository/postgres"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	AuthRepo
	AdvertRepo
	BetRepo
}

func NewRepostory(db *pgx.Conn) *Repository {
	return &Repository{
		AuthRepo:   postgres.NewAuth(db),
		AdvertRepo: postgres.NewAdvert(db),
		BetRepo:    postgres.NewBet(db),
	}
}

type AuthRepo interface {
	Create(ctx context.Context, email, password string) (int, error)
	Get(ctx context.Context, email string) (entity.User, error)
}

type AdvertRepo interface {
	Create(ctx context.Context, title, body string, price, userId int) (int, error)
	GetAll(ctx context.Context, userId int) ([]entity.Advert, error)
	Get(ctx context.Context, advertId int) (entity.Advert, error)
}

type BetRepo interface {
	Create(ctx context.Context, userId, advertId, betPrice int) (int, error)
	Update(ctx context.Context, userId, advertId, betPrice int) (int, error)
	GetPrice(ctx context.Context, userId, advertId int) (int, error)
	GetAdvertPrice(ctx context.Context, userId, advertId int) (int, error)
}
