package env

import (
	"fmt"
	"os"

	"github.com/simplicate/sparrow.api/inf/log"
)

const ENV_DB_HOST string = "DB_HOST"
const ENV_DB_PORT string = "DB_PORT"
const ENV_DB_DATABASE string = "DB_DATABASE"
const ENV_DB_USERNAME string = "DB_USERNAME"
const ENV_DB_PASSWORD string = "DB_PASSWORD"

// Get config value from environment
func Get(key string, defaultValue string) string {
	var value, success = os.LookupEnv(key)
	if success {
		log.Info(fmt.Sprintf("looking for: %s and found %s", key, value))
		return value
	}

	if defaultValue != "" {
		log.Info(fmt.Sprintf("looking for: %s but not found. using default %s", key, defaultValue))
		return defaultValue
	}

	msg := fmt.Sprintf("looking for: %s but not found and no default provided", key)
	log.Fatal(msg)
	panic(msg)
}
