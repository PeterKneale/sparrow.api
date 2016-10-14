//go:generate goagen bootstrap -d api/public/design

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/simplicate/sparrow.api/app"
	"github.com/simplicate/sparrow.api/controllers"
	"github.com/simplicate/sparrow.api/models"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	var userdb *models.UserDataDB
	var accountdb *models.AccountDataDB

	db := getDatabase()

	db.AutoMigrate(&models.UserData{})
	db.AutoMigrate(&models.AccountData{})

	userdb = models.NewUserDataDB(db)
	accountdb = models.NewAccountDataDB(db)

	// Create service
	service := goa.New("API")
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount middleware
	app.MountAccountController(service, controllers.NewAccountController(service, accountdb))
	app.MountUserController(service, controllers.NewUserController(service, userdb))
	app.MountHealthController(service, controllers.NewHealthController(service))
	app.MountJsController(service, controllers.NewJsController(service))
	app.MountSwaggerController(service, controllers.NewSwaggerController(service))
	app.MountWebController(service, controllers.NewWebController(service))

	// Start service
	if err := service.ListenAndServe("0.0.0.0:8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

func getDatabase() *gorm.DB {
	connection := getConnection()
	for attempt := 0; attempt < 10; attempt++ {
		fmt.Printf("Trying: %s. Attempt %d.\n", connection, attempt)
		db, err := gorm.Open("postgres", connection)
		if err == nil {
			db.LogMode(true)
			db.DB().SetMaxOpenConns(50)
			return db
		}
		time.Sleep(2000 * time.Millisecond)
	}
	panic("Cannot connect to " + connection)
}

func getConnection() string {
	host := getEnv("DB_HOST")
	database := getEnv("DB_DATABASE")
	user := getEnv("DB_USER")
	password := getEnv("DB_PASSWORD")
	port := getEnv("DB_PORT")
	sslmode := "disable"
	return fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%s host=%s", database, user, password, sslmode, port, host)
}

func getEnv(key string) string {
	var value, success = os.LookupEnv(key)
	if !success {
		panic("Cannot find " + key + " in environment variables.\n")
	}
	fmt.Printf("Looking for: %s and found %s.\n", key, value)
	return value
}
