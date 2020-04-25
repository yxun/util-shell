package sh

import (
	"github.com/sirupsen/logrus"
)

type Event struct {
	id      int
	message string
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	return standardLogger
}
