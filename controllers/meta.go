package controllers

import (
	"github.com/goadesign/goa"
	"github.com/simplicate/sparrow.api/app"
	"github.com/simplicate/sparrow.api/models"
)

// ErrNotReady is the error returned when the api is not yet ready
var ErrNotReady = goa.NewErrorClass("api_not_ready", 500)

// MetaController implements the meta resource.
type MetaController struct {
	*goa.Controller
	accountdb *models.AccountDataDB
	userdb    *models.UserDataDB
}

// NewMetaController creates a meta controller.
func NewMetaController(service *goa.Service, userdb *models.UserDataDB, accountdb *models.AccountDataDB) *MetaController {
	return &MetaController{
		Controller: service.NewController("MetaController"),
		userdb:     userdb,
		accountdb:  accountdb,
	}
}

// Alive runs the alive action.
func (c *MetaController) Alive(ctx *app.AliveMetaContext) error {
	return ctx.OK(nil)
}

// Ready runs the ready action.
func (c *MetaController) Ready(ctx *app.ReadyMetaContext) error {
	_, err := c.accountdb.List(ctx.Context)
	if err != nil {
		return ErrNotReady(err)
	}
	return ctx.OK(nil)
}

// Root runs the root action.
func (c *MetaController) Root(ctx *app.RootMetaContext) error {
	return ctx.OK(nil)
}
