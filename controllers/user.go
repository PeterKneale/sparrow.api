package controllers

import (
	"fmt"

	"github.com/simplicate/sparrow.api/app"
	"github.com/simplicate/sparrow.api/models"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ErrDatabaseError is the error returned when a db query fails.
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

// UserController implements the User resource.
type UserController struct {
	*goa.Controller
	userdb *models.UserDataDB
}

// NewUserController creates a User controller.
func NewUserController(service *goa.Service, db *models.UserDataDB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		userdb:     db,
	}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	user := models.UserData{}
	user.FirstName = ctx.Payload.FirstName
	user.LastName = ctx.Payload.LastName

	err := c.userdb.Add(ctx.Context, &user)
	if err != nil {
		return ErrDatabaseError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.UserHref(user.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	_, err := c.userdb.Get(ctx.Context, ctx.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NotFound()
		}
		return ErrDatabaseError(err)
	}

	c.userdb.Delete(ctx, ctx.ID)
	return ctx.NoContent()
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	users, err := c.userdb.List(ctx)
	if err != nil {
		return ErrDatabaseError(err)
	}

	mode := ctx.Params.Get("mode")
	if mode == "tiny" {
		res := make(app.UserTinyCollection, len(users))
		for i, user := range users {
			res[i] = ToUserTiny(user)
		}
		return ctx.OKTiny(res)
	}

	res := make(app.UserCollection, len(users))
	for i, user := range users {
		res[i] = ToUserMedia(user)
	}

	return ctx.OK(res)
}

// Read runs the read action.
func (c *UserController) Read(ctx *app.ReadUserContext) error {
	user, err := c.userdb.Get(ctx.Context, ctx.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NotFound()
		}
		return ErrDatabaseError(err)
	}

	return ctx.OK(ToUserMedia(user))
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {

	user, err := c.userdb.Get(ctx.Context, ctx.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NotFound()
		}
		return ErrDatabaseError(err)
	}

	if ctx.Payload.FirstName != nil {
		user.FirstName = *ctx.Payload.FirstName
	}
	if ctx.Payload.LastName != nil {
		user.LastName = *ctx.Payload.LastName
	}

	err = c.userdb.Update(ctx, user)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// ToUserMedia from UserData
func ToUserMedia(u *models.UserData) *app.User {
	name := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return &app.User{
		ID:        u.ID,
		Name:      &name,
		Href:      app.UserHref(u.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

// ToUserTiny from UserData
func ToUserTiny(user *models.UserData) *app.UserTiny {
	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	return &app.UserTiny{
		ID:   user.ID,
		Href: app.UserHref(user.ID),
		Name: &name,
	}
}

// ToUserLink from UserData
func ToUserLink(user *models.UserData) *app.UserLink {
	return &app.UserLink{
		ID:   user.ID,
		Href: app.UserHref(user.ID),
	}
}

/*

import (
	"fmt"
	"strings"
)

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/accounts/%v", paramaccountID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v", paramuserID)
}

*/
