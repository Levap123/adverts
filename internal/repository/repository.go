package repository

import (
	"context"

	"github.com/Levap123/adverts/internal/entity"
	"github.com/Levap123/adverts/internal/repository/postgres"
	goredis "github.com/Levap123/adverts/internal/repository/redis"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	AuthRepo
	AdvertRepo
	BetRepo
	BetTimeoutRepo
}

func NewRepostory(db *pgx.Conn, cl *redis.Client) *Repository {
	return &Repository{
		AuthRepo:       postgres.NewAuth(db),
		AdvertRepo:     postgres.NewAdvert(db),
		BetRepo:        postgres.NewBet(db),
		BetTimeoutRepo: goredis.NewBetTimeout(cl),
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
	GetEmail(ctx context.Context, advertId int) (string, error)
}

type BetRepo interface {
	Create(ctx context.Context, userId, advertId, betPrice int) (int, error)
	Update(ctx context.Context, userId, advertId, betPrice int) (int, error)
	GetPrice(ctx context.Context, advertId int) (int, error)
	GetAdvertPrice(ctx context.Context, advertId int) (int, error)
	IsActive(ctx context.Context, advertId int) (bool, error)
}

type BetTimeoutRepo interface {
	ListenProducer(ch chan int, errCh chan error)
}
