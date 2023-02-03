package main

import (
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
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}
