package logging_helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

// SetupLogger sets up the global logger configuration.
func SetupLogger() *logrus.Logger {
	log := logrus.New()

	// Output to stdout instead of the default stderr
	log.Out = os.Stdout

	// Set log format to JSON for structured logging
	log.SetFormatter(&logrus.JSONFormatter{})

	// Set log level (you can change this based on the environment)
	log.SetLevel(logrus.InfoLevel)

	return log
}
