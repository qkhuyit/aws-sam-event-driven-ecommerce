package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &Formatter{
			TimestampFormat: "2006-01-02 15:04:05.0000",
			LogFormat:       "[%lvl%]: %time% - %msg%",
		},
	}
	return log
}
