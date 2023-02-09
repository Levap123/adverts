package lg

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	logrus.SetFormatter(logrus.StandardLogger().Formatter)
}

func NewLogger() (*logrus.Logger, error) {
	f, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	logger := &logrus.Logger{
		Out:   io.MultiWriter(os.Stdout, f),
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	logger.Info("231123")
	return logger, nil
}
