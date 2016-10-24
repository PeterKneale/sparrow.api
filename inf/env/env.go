package env

import (
	"fmt"
	"os"

	"github.com/simplicate/sparrow.api/inf/log"
)

const ENV_DB_HOST string = "DB_PORT_6379_TCP_ADDR"
const ENV_DB_PORT string = "DB_PORT_6379_TCP_PORT"
const ENV_DB_DATABASE string = "DB_DATABASE"
const ENV_DB_USERNAME string = "DB_USERNAME"
const ENV_DB_PASSWORD string = "DB_PASSWORD"

// Get config value from environment
func Get(key string, defaultValue string) string {
	var value, success = os.LookupEnv(key)
	if success {
		msg := fmt.Sprintf("Looking for: %s and found %s", key, value)
		log.Info(log.ACTION_API_HOST, log.ACTION_HOST_CONFIGURE, msg)
		return value
	}

	if defaultValue != "" {
		msg := fmt.Sprintf("Looking for: %s but not found. Using default %s", key, defaultValue)
		log.Info(log.ACTION_API_HOST, log.ACTION_HOST_CONFIGURE, msg)
		return defaultValue
	}

	msg := fmt.Sprintf("Looking for: %s but not found and no default provided", key)
	log.Fatal(log.ACTION_API_HOST, log.ACTION_HOST_CONFIGURE, msg)
	panic(msg)
}
