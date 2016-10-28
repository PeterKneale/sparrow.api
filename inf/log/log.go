package log

import (
	"os"

	"github.com/Sirupsen/logrus"
)

//Init initialises the logger
func Init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

// Debug level logging
func Debug(message string) {
	logrus.WithFields(logrus.Fields{
		"app":       "sparrow",
		"component": "api",
	}).Debug(message)
}

// Info level logging
func Info(message string) {
	logrus.WithFields(logrus.Fields{
		"app":       "sparrow",
		"component": "api",
	}).Info(message)
}

// Warn level logging
func Warn(message string) {
	logrus.WithFields(logrus.Fields{
		"app":       "sparrow",
		"component": "api",
	}).Warn(message)
}

// Fatal level logging
func Fatal(message string) {
	logrus.WithFields(logrus.Fields{
		"app":       "sparrow",
		"component": "api",
	}).Fatal(message)
}
