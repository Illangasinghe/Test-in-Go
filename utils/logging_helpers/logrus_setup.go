package logging_helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

// SetupLogger sets up the logger using Logrus
func SetupLogger() *logrus.Logger {
	log := logrus.New()

	// Output logs to stdout
	log.Out = os.Stdout

	// Set the log format to JSON
	log.SetFormatter(&logrus.JSONFormatter{})

	// Set the log level (you can set this dynamically based on environment)
	log.SetLevel(logrus.InfoLevel)

	return log
}
