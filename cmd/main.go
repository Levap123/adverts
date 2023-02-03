package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(logrus.StandardLogger().Formatter)
	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
