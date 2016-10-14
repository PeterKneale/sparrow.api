package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("StorageGroup", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the postgres relational store")
		Model("UserData", func() {
			Alias("users")
			RendersTo(UserMedia)
			Description("This is the user model")

			Field("ID", gorma.Integer, func() { // Required for CRUD getters to take a PK argument!
				PrimaryKey()
				Description("This is the ID PK field")
			})

			Field("first_name", gorma.String)
			Field("last_name", gorma.String)

			Field("CreatedAt", gorma.Timestamp)
			Field("UpdatedAt", gorma.Timestamp)         // Shown for demonstration
			Field("DeletedAt", gorma.NullableTimestamp) // These are added by default
		})

		Model("AccountData", func() {
			Alias("accounts")
			RendersTo(UserMedia)
			Description("This is the account model")

			Field("ID", gorma.Integer, func() { // Required for CRUD getters to take a PK argument!
				PrimaryKey()
				Description("This is the ID PK field")
			})

			Field("name", gorma.String)

			Field("CreatedAt", gorma.Timestamp)
			Field("UpdatedAt", gorma.Timestamp)         // Shown for demonstration
			Field("DeletedAt", gorma.NullableTimestamp) // These are added by default
		})
	})
})
