package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
