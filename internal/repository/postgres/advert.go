package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Advert struct {
	db *pgx.Conn
}

func NewAdvert(db *pgx.Conn) *Advert {
	return &Advert{
		db: db,
	}
}

const advertTable = "adverts"

func (a *Advert) Create(ctx context.Context, title, body string, price, userId int) (int, error) {
	var advertId int
	query := fmt.Sprintf("INSERT INTO %s (title, body, price, user_id) VALUES ($1, $2, $3, $4) RETURNING id", advertTable)
	row := a.db.QueryRow(ctx, query, title, body, price, userId)
	if err := row.Scan(&advertId); err != nil {
		return 0, fmt.Errorf("repo - create advert - %w", err)
	}
	return advertId, nil
}
