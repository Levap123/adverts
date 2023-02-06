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

const BetTable = "bets"

func (b *Bet) Create(ctx context.Context, userId, advertId, betPrice int) (int, error) {
	tx, err := b.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("repo - create bet - %w", err)
	}
	defer tx.Rollback(ctx)
	var betId int
	query := fmt.Sprintf("INSERT INTO %s (user_id, advert_id, bet_price) VALUES ($1, $2, $3) RETURNING id", BetTable)
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
	query := fmt.Sprintf("UPDATE %s SET bet_price = $1, user_id = $2 WHERE advert_id = $3", BetTable)
	row := tx.QueryRow(ctx, query, userId, advertId, betPrice)
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
	query := fmt.Sprintf("SELECT bet_price FROM %s WHERE user_id = $1 and advert_id = $2", BetTable)
	row := tx.QueryRow(ctx, query, userId, advertId)
	if err := row.Scan(&price); err != nil {
		return 0, fmt.Errorf("repo - get price bet - %w", err)
	}
	return price, tx.Commit(ctx)
}
