package repository

import (
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
	Create(email, password string) (int, error)
}
