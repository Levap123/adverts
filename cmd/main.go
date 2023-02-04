package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Levap123/adverts/internal/app"
	"github.com/Levap123/adverts/pkg/lg"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := lg.InitConfigs(); err != nil {
		logrus.Fatal(err)
	}
	app, err := app.NewApp()
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		if err := app.Run(); err != nil {
			logrus.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := app.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
