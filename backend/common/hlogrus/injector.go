package hlogrus

import (
	"fmt"
	"log"
	"os"

	"github.com/Ralphbaer/hubla/backend/common/console"
	"github.com/Ralphbaer/hubla/backend/common/hlog"
	"github.com/sirupsen/logrus"
)

// InitializeLogger initializes our log layer and returns it
func InitializeLogger() hlog.Logger {
	fmt.Println(console.Title("InitializeLogger"))

	logger := logrus.New()

	var f logrus.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	if os.Getenv("LOG_FORMAT") == "json" {
		f = &logrus.JSONFormatter{}
	}

	logger.SetFormatter(f)

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Println("Logrus error:", err.Error())
		level = logrus.InfoLevel
	}

	log.Printf("Log level is (%v)\n", level)
	log.Printf("Logger is (%T)\n", logger)

	logger.SetLevel(level)
	w := logger.Writer()
	log.SetOutput(w)

	fmt.Println(console.Line(console.DefaultLineSize))

	return &LogrusLogger{
		Logger: logger,
	}
}
