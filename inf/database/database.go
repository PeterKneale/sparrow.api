package database

import (
	"fmt"
	"time"

	"github.com/simplicate/sparrow.api/inf/env"
	"github.com/simplicate/sparrow.api/inf/log"

	"github.com/jinzhu/gorm"
)

func GetDatabase() *gorm.DB {
	connection := getConnection()
	for attempt := 0; attempt < 10; attempt++ {

		log.Debug("connecting to database")

		db, err := gorm.Open("postgres", connection)

		if err == nil {
			db.LogMode(true)
			db.DB().SetMaxOpenConns(50)
			return db
		}

		log.Warn("failed to connect to database")

		time.Sleep(2000 * time.Millisecond)
	}
	log.Fatal("failed to connect to database")
	panic("failed to connect:" + connection)
}

func getConnection() string {
	host := env.Get(env.ENV_DB_HOST, "localhost")
	port := env.Get(env.ENV_DB_PORT, "5432")
	database := env.Get(env.ENV_DB_DATABASE, "postgres")
	user := env.Get(env.ENV_DB_USERNAME, "postgres")
	password := env.Get(env.ENV_DB_PASSWORD, "password")
	sslmode := "disable"

	format := "dbname=%s user=%s password=%s sslmode=%s port=%s host=%s"
	connection := fmt.Sprintf(format, database, user, password, sslmode, port, host)
	connectionAnon := fmt.Sprintf(format, database, user, "xxxxxxxxx", sslmode, port, host)

	log.Info(fmt.Sprintf("db connection: %s", connectionAnon))
	return connection
}
