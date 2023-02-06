package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Bet struct {
	DB *pgx.Conn
}

func NewBet(DB *pgx.Conn) *Bet {
	return &Bet{
		DB: DB,
	}
}

const betTable = "bets"

func (b *Bet) Create(ctx context.Context, userId, advertId, betPrice int) (int, error) {
	tx, err := b.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("repo - create bet - %w", err)
	}
	defer tx.Rollback(ctx)
	var betId int
	query := fmt.Sprintf("INSERT INTO %s (user_id, advert_id, bet_price) VALUES ($1, $2, $3) RETURNING id", betTable)
	row := tx.QueryRow(ctx, query, userId, advertId, betPrice)
	if err := row.Scan(&betId); err != nil {
		return 0, fmt.Errorf("repo - create bet - %w", err)
	}
	return betId, tx.Commit(ctx)
}

func (b *Bet) Update(ctx context.Context, userId, advertId, betPrice int) (int, error) {
	tx, err := b.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("repo - update bet - %w", err)
	}
	defer tx.Rollback(ctx)
	var betId int
	query := fmt.Sprintf("UPDATE %s SET bet_price = $1, user_id = $2 WHERE advert_id = $3 RETURNING id", betTable)
	fmt.Println(advertId)
	row := tx.QueryRow(ctx, query, betPrice, userId, advertId)
	if err := row.Scan(&betId); err != nil {
		return 0, fmt.Errorf("repo - update bet - %w", err)
	}
	return betId, tx.Commit(ctx)
}

func (b *Bet) GetPrice(ctx context.Context, userId, advertId int) (int, error) {
	tx, err := b.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("repo - get price bet - %w", err)
	}
	defer tx.Rollback(ctx)
	var price int
	query := fmt.Sprintf("SELECT bet_price FROM %s WHERE user_id = $1 and advert_id = $2", betTable)
	row := tx.QueryRow(ctx, query, userId, advertId)
	if err := row.Scan(&price); err != nil {
		return 0, fmt.Errorf("repo - get price bet - %w", err)
	}
	return price, tx.Commit(ctx)
}

func (b *Bet) GetAdvertPrice(ctx context.Context, userId, advertId int) (int, error) {
	tx, err := b.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("repo - get price advert - %w", err)
	}
	defer tx.Rollback(ctx)
	var price int
	query := fmt.Sprintf("SELECT price FROM %s WHERE user_id = $1 and id = $2", advertTable)
	row := tx.QueryRow(ctx, query, userId, advertId)
	if err := row.Scan(&price); err != nil {
		return 0, fmt.Errorf("repo - get price advert - %w", err)
	}
	return price, tx.Commit(ctx)
}

func (b *Bet) IsActive(ctx context.Context, userId, advertId int) (bool, error) {
	tx, err := b.DB.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("repo - get status bet - %w", err)
	}
	defer tx.Rollback(ctx)
	var status string
	query := fmt.Sprintf("SELECT status FROM %s WHERE user_id = $1 and id = $2", advertTable)
	row := tx.QueryRow(ctx, query, userId, advertId)
	if err := row.Scan(&status); err != nil {
		return false, fmt.Errorf("repo - get status bet - %w", err)
	}
	fmt.Println(status)
	return status == activeStatus, tx.Commit(ctx)
}
