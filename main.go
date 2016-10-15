//go:generate goagen bootstrap -d api/public/design

package main

import (
	"fmt"

	"github.com/simplicate/sparrow.api/app"
	"github.com/simplicate/sparrow.api/controllers"
	"github.com/simplicate/sparrow.api/inf/database"
	"github.com/simplicate/sparrow.api/inf/env"
	"github.com/simplicate/sparrow.api/inf/log"
	"github.com/simplicate/sparrow.api/models"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	_ "github.com/lib/pq"
)

func main() {

	log.Init()
	var userdb *models.UserDataDB
	var accountdb *models.AccountDataDB

	db := database.GetDatabase()

	db.AutoMigrate(&models.UserData{})
	db.AutoMigrate(&models.AccountData{})

	userdb = models.NewUserDataDB(db)
	accountdb = models.NewAccountDataDB(db)

	// Create service
	service := goa.New("api")
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount middleware
	app.MountAccountController(service, controllers.NewAccountController(service, accountdb))
	app.MountUserController(service, controllers.NewUserController(service, userdb))
	app.MountHealthController(service, controllers.NewHealthController(service))
	app.MountSwaggerController(service, controllers.NewSwaggerController(service))

	// Start service
	port := env.Get("API_PORT", "80")
	address := fmt.Sprintf("0.0.0.0:%s", port)

	if err := service.ListenAndServe(address); err != nil {
		service.LogError("startup", "err", err)
	}
}
