//************************************************************************//
// API "Public API": Model Helpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/sparrow.api/design
// --out=C:\dev\go\src\github.com\simplicate\mango\gateway
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/simplicate/sparrow.api/app"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListUser returns an array of view: default.
func (m *UserDataDB) ListUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*UserData
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing UserData", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserDataToUser())
	}

	return objs
}

// UserDataToUser loads a UserData and builds the default view of media type User.
func (m *UserData) UserDataToUser() *app.User {
	userdata := &app.User{}
	userdata.FirstName = m.FirstName
	userdata.ID = m.ID
	userdata.LastName = m.LastName

	return userdata
}

// OneUser loads a UserData and builds the default view of media type User.
func (m *UserDataDB) OneUser(ctx context.Context, id int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native UserData
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting UserData", "error", err.Error())
		return nil, err
	}

	view := *native.UserDataToUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserLink returns an array of view: link.
func (m *UserDataDB) ListUserLink(ctx context.Context) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*UserData
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing UserData", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserDataToUserLink())
	}

	return objs
}

// UserDataToUserLink loads a UserData and builds the link view of media type User.
func (m *UserData) UserDataToUserLink() *app.UserLink {
	userdata := &app.UserLink{}
	userdata.ID = m.ID

	return userdata
}

// OneUserLink loads a UserData and builds the link view of media type User.
func (m *UserDataDB) OneUserLink(ctx context.Context, id int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native UserData
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting UserData", "error", err.Error())
		return nil, err
	}

	view := *native.UserDataToUserLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserTiny returns an array of view: tiny.
func (m *UserDataDB) ListUserTiny(ctx context.Context) []*app.UserTiny {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listusertiny"}, time.Now())

	var native []*UserData
	var objs []*app.UserTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing UserData", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserDataToUserTiny())
	}

	return objs
}

// UserDataToUserTiny loads a UserData and builds the tiny view of media type User.
func (m *UserData) UserDataToUserTiny() *app.UserTiny {
	userdata := &app.UserTiny{}
	userdata.ID = m.ID

	return userdata
}

// OneUserTiny loads a UserData and builds the tiny view of media type User.
func (m *UserDataDB) OneUserTiny(ctx context.Context, id int) (*app.UserTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneusertiny"}, time.Now())

	var native UserData
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting UserData", "error", err.Error())
		return nil, err
	}

	view := *native.UserDataToUserTiny()
	return &view, err
}
