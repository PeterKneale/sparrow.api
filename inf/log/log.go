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
func Debug(inf, action, message string) {
	logrus.WithFields(logrus.Fields{
		"app":    "api",
		"inf":    inf,
		"action": action,
	}).Debug(message)
}

// Info level logging
func Info(inf, action, message string) {
	logrus.WithFields(logrus.Fields{
		"app":    "api",
		"inf":    inf,
		"action": action,
	}).Info(message)
}

// Warn level logging
func Warn(inf, action, message string) {
	logrus.WithFields(logrus.Fields{
		"app":    "api",
		"inf":    inf,
		"action": action,
	}).Warn(message)
}

// Fatal level logging
func Fatal(inf, action, message string) {
	logrus.WithFields(logrus.Fields{
		"app":    "api",
		"inf":    inf,
		"action": action,
	}).Fatal(message)
}

const INF_HOST = "HOST"
const ACTION_HOST_CONFIGURE = "CONFIGURE"

const INF_DATABASE = "DATABASE"
const ACTION_DATABASE_CONNECT = "CONNECT"
const ACTION_DATABASE_EXECUTE = "EXECUTE"

const INF_API = "API"
const ACTION_API_HOST = "EXECUTE"
