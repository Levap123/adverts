package lg

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(logrus.StandardLogger().Formatter)
}

func InitConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}

func NewLogger() (*logrus.Logger, error) {
	out, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &logrus.Logger{
		Out: out,
	}, nil
}
