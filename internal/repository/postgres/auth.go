package postgres

import (
	"context"
	"fmt"

	"github.com/Levap123/adverts/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Auth struct {
	db *pgx.Conn
}

func NewAuth(db *pgx.Conn) *Auth {
	return &Auth{
		db: db,
	}
}

const userTable = "users"

func (a *Auth) Create(ctx context.Context, email, password string) (int, error) {
	var userID int
	query := fmt.Sprintf("INSERT INTO %s (email, password) VALUES ($1, $2) RETURNING id", userTable)
	row := a.db.QueryRow(ctx, query, email, password)
	if err := row.Scan(&userID); err != nil {
		return 0, fmt.Errorf("repo - create user - %w", err)
	}
	return userID, nil
}

func (a *Auth) Get(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", userTable)
	row := a.db.QueryRow(ctx, query, email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return entity.User{}, fmt.Errorf("repo - get user - %w", err)
	}
	return user, nil
}
