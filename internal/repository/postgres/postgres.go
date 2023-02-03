package postgres

import (
	"context"
	"fmt"

	"github.com/Levap123/adverts/configs"
	"github.com/jackc/pgx/v5"
)

func InitDB(ctx context.Context, conf configs.PostgresConf) (*pgx.Conn, error) {
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s", conf.Username, conf.Password, conf.Host, conf.DBName)
	conn, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
