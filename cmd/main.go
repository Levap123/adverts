package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/adverts/configs"
	"github.com/Levap123/adverts/internal/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := configs.InitConfigs(); err != nil {
		logrus.Fatal(err)
	}
	app, err := app.NewApp()
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		logrus.Infof("server is listening on http://0.0.0.0%s\n", viper.GetString("server_addr"))
		if err := app.Run(); err != nil {
			logrus.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
