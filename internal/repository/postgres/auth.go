package postgres

import "github.com/jackc/pgx/v5"

type Auth struct {
	db *pgx.Conn
}

func NewAuth(db *pgx.Conn) *Auth {
	return &Auth{
		db: db,
	}
}

func (a *Auth) Create(email, password string) (int, error) {
}
