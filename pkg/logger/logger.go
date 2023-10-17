package logger

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/configuration"
)

var filepath = "./pkg/logger/.log"

var logger_instance = logrus.New()

type loggerService struct {
	logger *logrus.Logger
}

func Init(configService *configuration.Config) *logrus.Logger {

	// setting the format of the logs to be a JSON one
	logger_instance.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// getting the log level set in the configuration file
	logLevel, err := logrus.ParseLevel(strconv.Itoa(configService.Logger.Level))
	// If the log level in conf file can't be parsed, log level should be the default info level
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	// setting the log level
	logger_instance.SetLevel(logLevel)

	if configService.Logger.Env == "local" { // If we want to throw logs into a local file

		logger_instance.SetOutput(os.Stdout)
		// setting it to a file writer
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger_instance.Out = file
		} else {
			logger_instance.Info("Failed to log to file, using default stderr")
		}
	}
	return logger_instance
}
