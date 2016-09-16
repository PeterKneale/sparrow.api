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

	displayEnvironmentVariables()

	var db *gorm.DB
	var userdb *models.UserDataDB
	var accountdb *models.AccountDataDB

	var err error
	url := fmt.Sprintf("dbname=postgres user=postgres password=password123 sslmode=disable port=%d host=%s", 5432, "postgres.33c58199.svc.dockerapp.io")
	fmt.Println(url)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	//db.DropTable(&models.UserData{})
	//db.DropTable(&models.AccountData{})
	db.AutoMigrate(&models.UserData{})
	db.AutoMigrate(&models.AccountData{})

	userdb = models.NewUserDataDB(db)
	accountdb = models.NewAccountDataDB(db)

	db.DB().SetMaxOpenConns(50)

	// Create service
	service := goa.New("Public API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

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

func displayEnvironmentVariables() {

	fmt.Println("Environment variables:")
	for _, e := range os.Environ() {
		fmt.Println(e)
	}

}
