package logger

import (
	"os"

	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
	"github.com/sirupsen/logrus"
)

type ProcessIdentifierHook struct{}

func (h *ProcessIdentifierHook) Fire(entry *logrus.Entry) error {
	entry.Data["pid"] = os.Getpid()
	return nil
}

func (h *ProcessIdentifierHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func New() *logrus.Logger {
	lgr := logrus.New()

	lgr.SetLevel(logrus.Level(config.GetInt("LOG_LEVEL")))
	lgr.SetFormatter(&logrus.JSONFormatter{})
	lgr.AddHook(&ProcessIdentifierHook{})

	return lgr
}
