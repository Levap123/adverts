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
	"github.com/Levap123/adverts/internal/validator"
	"github.com/Levap123/adverts/pkg/json"
	"github.com/Levap123/adverts/pkg/lg"
	"github.com/spf13/viper"
)

type App struct {
	server *http.Server
}

const MB = 1024 * 1024

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

	JSON := new(json.JSONSerializer)

	lg, err := lg.NewLogger()
	if err != nil {
		return nil, fmt.Errorf("new app - logger: %w", err)
	}

	validator := new(validator.Validator)

	repos := repository.NewRepostory(db)

	service := service.NewService(repos)

	handler := handler.NewHandler(service, JSON, lg, validator)

	routes := handler.InitRoutes()

	server := InitServer(routes, configs.ServerConf{
		Addr:      viper.GetString("server_addr"),
		RWTimeout: viper.GetInt("rw_timeout"),
		HeaderMBs: viper.GetInt("header_mbs"),
	})

	return &App{
		server: server,
	}, nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}

func InitServer(routes http.Handler, confs configs.ServerConf) *http.Server {
	return &http.Server{
		ReadTimeout:    time.Second * time.Duration(confs.RWTimeout),
		WriteTimeout:   time.Second * time.Duration(confs.RWTimeout),
		MaxHeaderBytes: MB * confs.HeaderMBs,
		Handler:        routes,
		Addr:           confs.Addr,
	}
}
