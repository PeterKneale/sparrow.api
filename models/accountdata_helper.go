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
func (m *AccountDataDB) ListUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*AccountData
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing AccountData", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountDataToUser())
	}

	return objs
}

// AccountDataToUser loads a AccountData and builds the default view of media type User.
func (m *AccountData) AccountDataToUser() *app.User {
	accountdata := &app.User{}
	accountdata.ID = m.ID
	accountdata.Name = &m.Name

	return accountdata
}

// OneUser loads a AccountData and builds the default view of media type User.
func (m *AccountDataDB) OneUser(ctx context.Context, id int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native AccountData
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting AccountData", "error", err.Error())
		return nil, err
	}

	view := *native.AccountDataToUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserLink returns an array of view: link.
func (m *AccountDataDB) ListUserLink(ctx context.Context) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*AccountData
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing AccountData", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountDataToUserLink())
	}

	return objs
}

// AccountDataToUserLink loads a AccountData and builds the link view of media type User.
func (m *AccountData) AccountDataToUserLink() *app.UserLink {
	accountdata := &app.UserLink{}
	accountdata.ID = m.ID

	return accountdata
}

// OneUserLink loads a AccountData and builds the link view of media type User.
func (m *AccountDataDB) OneUserLink(ctx context.Context, id int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native AccountData
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting AccountData", "error", err.Error())
		return nil, err
	}

	view := *native.AccountDataToUserLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserTiny returns an array of view: tiny.
func (m *AccountDataDB) ListUserTiny(ctx context.Context) []*app.UserTiny {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listusertiny"}, time.Now())

	var native []*AccountData
	var objs []*app.UserTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing AccountData", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountDataToUserTiny())
	}

	return objs
}

// AccountDataToUserTiny loads a AccountData and builds the tiny view of media type User.
func (m *AccountData) AccountDataToUserTiny() *app.UserTiny {
	accountdata := &app.UserTiny{}
	accountdata.ID = m.ID
	accountdata.Name = &m.Name

	return accountdata
}

// OneUserTiny loads a AccountData and builds the tiny view of media type User.
func (m *AccountDataDB) OneUserTiny(ctx context.Context, id int) (*app.UserTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneusertiny"}, time.Now())

	var native AccountData
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting AccountData", "error", err.Error())
		return nil, err
	}

	view := *native.AccountDataToUserTiny()
	return &view, err
}
