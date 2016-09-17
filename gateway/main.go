//go:generate goagen bootstrap -d api/public/design

package main

import (
	"fmt"
	"os"

	"github.com/simplicate/mango/gateway/app"
	"github.com/simplicate/mango/gateway/controllers"
	"github.com/simplicate/mango/gateway/models"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	var db *gorm.DB
	var userdb *models.UserDataDB
	var accountdb *models.AccountDataDB
	var err error

	db, err = gorm.Open("postgres", getConnection())
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxOpenConns(50)

	//db.DropTable(&models.UserData{})
	//db.DropTable(&models.AccountData{})
	db.AutoMigrate(&models.UserData{})
	db.AutoMigrate(&models.AccountData{})

	userdb = models.NewUserDataDB(db)
	accountdb = models.NewAccountDataDB(db)

	// Create service
	service := goa.New("Public API")
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
	if err := service.ListenAndServe("0.0.0.0:80"); err != nil {
		service.LogError("startup", "err", err)
	}
}

func getConnection() string {
	user := getConfigValue("DB_ENV_POSTGRES_USER", "postgres")
	password := getConfigValue("DB_ENV_POSTGRES_PASSWORD", "password123")
	database := getConfigValue("DB_ENV_POSTGRES_DB", "postgres")
	host := getConfigValue("DB_PORT_5432_TCP_ADDR", "postgres.33c58199.svc.dockerapp.io")
	port := getConfigValue("DB_PORT_5432_TCP_PORT", "5432")
	sslmode := "disable"
	connection := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%s host=%s", database, user, password, sslmode, port, host)
	fmt.Printf("Connecting: %s", connection)
	return connection
}

func getConfigValue(key string, defaultValue string) string {
	var value, success = os.LookupEnv(key)
	if !success {
		fmt.Printf("Looking for: %s but not found. Using default value %s\n", key, defaultValue)
		return defaultValue
	}
	fmt.Printf("Looking for: %s and found %s", key, defaultValue)
	return value
}
