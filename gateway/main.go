//go:generate goagen bootstrap -d api/public/design

package main

import (
	"fmt"

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
	url := fmt.Sprintf("dbname=postgres user=postgres password=password123 sslmode=disable port=%d host=%s", 5432, "postgres.33c58199.svc.dockerapp.io")
	fmt.Println(url)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	//db.DropTable(&models.UserData{})
	//db.DropTable(&models.AccountData{})
	//db.AutoMigrate(&models.UserData{})
	//db.AutoMigrate(&models.AccountData{})

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

	// Mount "Account" controller
	c := controllers.NewAccountController(service, accountdb)
	app.MountAccountController(service, c)

	// Mount "User" controller
	c2 := controllers.NewUserController(service, userdb)
	app.MountUserController(service, c2)

	// Mount "health" controller
	c3 := controllers.NewHealthController(service)
	app.MountHealthController(service, c3)

	// Mount "js1" controller
	c4 := controllers.NewJs1Controller(service)
	app.MountJs1Controller(service, c4)
	// Mount "js2" controller
	c5 := controllers.NewJs2Controller(service)
	app.MountJs2Controller(service, c5)
	// Mount "swagger" controller
	c6 := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, c6)
	// Mount "web1" controller
	c7 := controllers.NewWeb1Controller(service)
	app.MountWeb1Controller(service, c7)
	// Mount "web2" controller
	c8 := controllers.NewWeb2Controller(service)
	app.MountWeb2Controller(service, c8)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
