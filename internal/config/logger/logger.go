package logger

import (
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	lgr := logrus.New()

	lgr.SetLevel(logrus.Level(config.GetInt("LOG_LEVEL")))
	lgr.SetFormatter(&logrus.JSONFormatter{})

	return lgr
}
