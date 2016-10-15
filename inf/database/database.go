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

		log.Debug(log.INF_DATABASE, log.ACTION_DATABASE_CONNECT, "connecting to database")

		db, err := gorm.Open("postgres", connection)

		if err == nil {
			db.LogMode(true)
			db.DB().SetMaxOpenConns(50)
			return db
		}

		log.Warn(log.INF_DATABASE, log.ACTION_DATABASE_CONNECT, "failed to connect to database")

		time.Sleep(2000 * time.Millisecond)
	}
	log.Fatal(log.INF_DATABASE, log.ACTION_DATABASE_CONNECT, "failed to connect to database")
	panic("failed to connect:" + connection)
}

func getConnection() string {
	host := env.Get("DB_HOST", "localhost")
	database := env.Get("DB_DATABASE", "postgres")
	user := env.Get("DB_USER", "postgres")
	password := env.Get("DB_PASSWORD", "password")
	port := env.Get("DB_PORT", "5432")
	sslmode := "disable"

	format := "dbname=%s user=%s password=%s sslmode=%s port=%s host=%s"
	connection := fmt.Sprintf(format, database, user, password, sslmode, port, host)
	connectionAnon := fmt.Sprintf(format, database, user, "xxxxxxxxx", sslmode, port, host)

	msg := fmt.Sprintf("database connection: %s", connectionAnon)
	log.Info(log.ACTION_API_HOST, log.ACTION_HOST_CONFIGURE, msg)
	return connection
}
