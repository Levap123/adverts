package lg

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(logrus.StandardLogger().Formatter)
}

func NewLogger() (*logrus.Logger, error) {
	out, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &logrus.Logger{
		Out: out,
	}, nil
}
