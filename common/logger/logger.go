package logger

import (
	"github.com/sirupsen/logrus"

	"frendler/processor/config"
)

var log = logrus.New()

func init() {
	cfg := config.Get()
	level, err := logrus.ParseLevel(cfg.Logger.LogLevel)
	if err != nil {
		panic(err)
	}

	log.SetLevel(level)

	log.SetFormatter(&logrus.JSONFormatter{})

	logger := initLogger(cfg.Logger)
	logrus.SetOutput(logger.Writer())
	logrus.SetLevel(logger.Level)
	logrus.SetFormatter(logger.Formatter)

}

func initLogger(cfg config.LoggerConf) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Warnf("Invalid log level '%s', defaulting to 'info'", cfg.LogLevel)
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	if cfg.LogFormat == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return logger
}
