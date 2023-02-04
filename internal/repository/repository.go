package repository

import (
	"context"

	"github.com/Levap123/adverts/internal/entity"
	"github.com/Levap123/adverts/internal/repository/postgres"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	AuthRepo
}

func NewRepostory(db *pgx.Conn) *Repository {
	return &Repository{
		AuthRepo: postgres.NewAuth(db),
	}
}

type AuthRepo interface {
	Create(ctx context.Context, email, password string) (int, error)
	Get(ctx context.Context, email string) (entity.User, error)
}
