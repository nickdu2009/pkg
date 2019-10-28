package xgin

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	httpLogger "github.com/nickxb/gin-http-logger"
	"github.com/sirupsen/logrus"
)

func AccessLogger(file string, level string, excludePaths ...string) gin.HandlerFunc {
	accessLogger := logrus.New()
	accessLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	l, err := logrus.ParseLevel(level)
	if err != nil {
		panic(fmt.Sprintf("parse log levle %s %v\n", level, err))
	}

	accessLogger.SetLevel(l)
	w, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		panic(fmt.Sprintf("open access log file %s %v\n", file, err))
	}
	accessLogger.SetOutput(w)

	alc := httpLogger.AccessLoggerConfig{
		LogrusLogger:   accessLogger,
		BodyLogPolicy:  httpLogger.LogAllBodies,
		MaxBodyLogSize: 1024 * 16, //16k
		DropSize:       1024 * 10, //10k
	}

	alc.ExcludePaths = map[string]bool{}
	for _, excludePath := range excludePaths {
		alc.ExcludePaths[excludePath] = true
	}

	return httpLogger.New(alc)
}
