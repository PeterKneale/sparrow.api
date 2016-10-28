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

// This is the account model
type AccountData struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	Name      string
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m AccountData) TableName() string {

	return "accounts"
}

// AccountDataDB is the implementation of the storage interface for
// AccountData.
type AccountDataDB struct {
	Db *gorm.DB
}

// NewAccountDataDB creates a new storage type.
func NewAccountDataDB(db *gorm.DB) *AccountDataDB {
	return &AccountDataDB{Db: db}
}

// DB returns the underlying database.
func (m *AccountDataDB) DB() interface{} {
	return m.Db
}

// AccountDataStorage represents the storage interface.
type AccountDataStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*AccountData, error)
	Get(ctx context.Context, id int) (*AccountData, error)
	Add(ctx context.Context, accountdata *AccountData) error
	Update(ctx context.Context, accountdata *AccountData) error
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
func (m *AccountDataDB) TableName() string {

	return "accounts"
}

// CRUD Functions

// Get returns a single AccountData as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AccountDataDB) Get(ctx context.Context, id int) (*AccountData, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountData", "get"}, time.Now())

	var native AccountData
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of AccountData
func (m *AccountDataDB) List(ctx context.Context) ([]*AccountData, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountData", "list"}, time.Now())

	var objs []*AccountData
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *AccountDataDB) Add(ctx context.Context, model *AccountData) error {
	defer goa.MeasureSince([]string{"goa", "db", "accountData", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding AccountData", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *AccountDataDB) Update(ctx context.Context, model *AccountData) error {
	defer goa.MeasureSince([]string{"goa", "db", "accountData", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating AccountData", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDataDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "accountData", "delete"}, time.Now())

	var obj AccountData

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting AccountData", "error", err.Error())
		return err
	}

	return nil
}
