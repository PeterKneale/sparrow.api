//************************************************************************//
// API "Sparrow API": Models
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/sparrow.api/design
// --out=$(GOPATH)\src\github.com\simplicate\sparrow.api\
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/simplicate/sparrow.api/app"
	"golang.org/x/net/context"
)

// This is the user model
type UserData struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	FirstName string
	LastName  string
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m UserData) TableName() string {

	return "users"
}

// UserDataDB is the implementation of the storage interface for
// UserData.
type UserDataDB struct {
	Db *gorm.DB
}

// NewUserDataDB creates a new storage type.
func NewUserDataDB(db *gorm.DB) *UserDataDB {
	return &UserDataDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDataDB) DB() interface{} {
	return m.Db
}

// UserDataStorage represents the storage interface.
type UserDataStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*UserData, error)
	Get(ctx context.Context, id int) (*UserData, error)
	Add(ctx context.Context, userdata *UserData) error
	Update(ctx context.Context, userdata *UserData) error
	Delete(ctx context.Context, id int) error

	ListUser(ctx context.Context) []*app.User
	OneUser(ctx context.Context, id int) (*app.User, error)

	ListUserLink(ctx context.Context) []*app.UserLink
	OneUserLink(ctx context.Context, id int) (*app.UserLink, error)

	ListUserTiny(ctx context.Context) []*app.UserTiny
	OneUserTiny(ctx context.Context, id int) (*app.UserTiny, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDataDB) TableName() string {

	return "users"
}

// CRUD Functions

// Get returns a single UserData as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDataDB) Get(ctx context.Context, id int) (*UserData, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userData", "get"}, time.Now())

	var native UserData
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of UserData
func (m *UserDataDB) List(ctx context.Context) ([]*UserData, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userData", "list"}, time.Now())

	var objs []*UserData
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserDataDB) Add(ctx context.Context, model *UserData) error {
	defer goa.MeasureSince([]string{"goa", "db", "userData", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding UserData", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserDataDB) Update(ctx context.Context, model *UserData) error {
	defer goa.MeasureSince([]string{"goa", "db", "userData", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating UserData", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDataDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "userData", "delete"}, time.Now())

	var obj UserData

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting UserData", "error", err.Error())
		return err
	}

	return nil
}
