package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Levap123/adverts/configs"
	"github.com/jackc/pgx/v5"
)

func InitDB(ctx context.Context, conf configs.PostgresConf) (*pgx.Conn, error) {
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s", conf.Username, conf.Password, conf.Host, conf.DBName)
	db, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := createTables(db, ctx); err != nil {
		return nil, err
	}

	return db, nil
}
