package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Levap123/adverts/configs"
	"github.com/Levap123/adverts/internal/repository"
	"github.com/Levap123/adverts/internal/repository/postgres"
	"github.com/Levap123/adverts/internal/service"
	handler "github.com/Levap123/adverts/internal/transport"
	"github.com/spf13/viper"
)

type App struct {
	handler *handler.Handler
	service *service.Service
	repos   *repository.Repository
	server  *http.Server
}

func NewApp() (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	db, err := postgres.InitDB(ctx, configs.PostgresConf{
		DBName:   viper.GetString("db_name"),
		Username: viper.GetString("db_username"),
		Password: viper.GetString("db_password"),
		Host:     viper.GetString("db_host"),
	})
	if err != nil {
		return nil, fmt.Errorf("new app - open db: %w", err)
	}
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("new app - ping db: %w", err)
	}
	return &App{}, nil
}
