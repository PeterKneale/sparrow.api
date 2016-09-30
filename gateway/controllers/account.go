package controllers

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/simplicate/sparrow.api/gateway/app"
	"github.com/simplicate/sparrow.api/gateway/models"
)

// AccountController implements the Account resource.
type AccountController struct {
	*goa.Controller
	accountdb *models.AccountDataDB
}

// NewAccountController creates a Account controller.
func NewAccountController(service *goa.Service, db *models.AccountDataDB) *AccountController {
	return &AccountController{
		Controller: service.NewController("AccountController"),
		accountdb:  db,
	}
}

// Read runs the read action.
func (c *AccountController) Read(ctx *app.ReadAccountContext) error {
	account, err := c.accountdb.Get(ctx.Context, 1)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NotFound()
		}
		return ErrDatabaseError(err)
	}

	return ctx.OK(ToAccountMedia(account))
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {

	account, err := c.accountdb.Get(ctx.Context, 1)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NotFound()
		}
		return ErrDatabaseError(err)
	}

	if ctx.Payload.Name != nil {
		account.Name = *ctx.Payload.Name
	}

	err = c.accountdb.Update(ctx, account)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// ToAccountMedia from AccountData
func ToAccountMedia(a *models.AccountData) *app.Account {
	return &app.Account{
		Name: a.Name,
	}
}
