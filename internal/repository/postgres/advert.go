package postgres

import (
	"context"
	"fmt"

	"github.com/Levap123/adverts/internal/entity"
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

const (
	advertTable   = "adverts"
	activeStatus  = "active"
	passvieStatus = "passvie"
)

func (a *Advert) Create(ctx context.Context, title, body string, price, userId int) (int, error) {
	var advertId int
	query := fmt.Sprintf("INSERT INTO %s (title, body, price, status, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", advertTable)
	row := a.db.QueryRow(ctx, query, title, body, price, activeStatus, userId)
	if err := row.Scan(&advertId); err != nil {
		return 0, fmt.Errorf("repo - create advert - %w", err)
	}
	return advertId, nil
}

func (a *Advert) GetAll(ctx context.Context, userId int) ([]entity.Advert, error) {
	var adverts []entity.Advert
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo - get all adverts - %w", err)
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", advertTable)
	rows, err := tx.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("repo - get all adverts - %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var buffer entity.Advert
		if err := rows.Scan(&buffer.ID, &buffer.Title, &buffer.Body, &buffer.Price, &buffer.Status, &buffer.UserID); err != nil {
			return nil, fmt.Errorf("repo - get all adverts - %w", err)
		}
		adverts = append(adverts, buffer)
	}
	return adverts, tx.Commit(ctx)
}

func (a *Advert) Get(ctx context.Context, advertId int) (entity.Advert, error) {
	var advert entity.Advert
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", advertTable)
	row := a.db.QueryRow(ctx, query, advertId)
	if err := row.Scan(&advert.ID, &advert.Title, &advert.Body, &advert.Price, &advert.Status, &advert.UserID); err != nil {
		return entity.Advert{}, fmt.Errorf("repo - get one advert - %w", err)
	}
	return advert, nil
}

func (a *Advert) GetEmail(ctx context.Context, advertId int) (string, error) {
	var email string
	query := fmt.Sprintf("SELECT u.email FROM %s u JOIN %s a ON u.id = a.user_id WHERE a.id = $1", userTable, advertTable)
	row := a.db.QueryRow(ctx, query, advertId)
	if err := row.Scan(&email); err != nil {
		return "", fmt.Errorf("repo - get one advert - %w", err)
	}
	return email, nil
}
