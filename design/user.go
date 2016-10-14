package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("User", func() {

	DefaultMedia(UserMedia)

	BasePath("/")

	Action("list", func() {
		Routing(GET("/users"))
		Description("List users")
		Response(OK, func() {
			Media(CollectionOf(UserMedia, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(BadRequest, ErrorMedia)
	})

	Action("read", func() {
		Routing(GET("/users/:id"))
		Description("read user")
		Params(func() {
			Param("id", Integer, "User id")
		})
		Response(OK, UserMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(POST("/users"))
		Description("Create new user")
		Payload(CreateUserPayload)
		Response(Created, "/users/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(PUT("/users/:id"))
		Description("Update user")
		Params(func() {
			Param("id", Integer, "User id")
		})
		Payload(UpdateUserPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(DELETE("/users/:id"))
		Params(func() {
			Param("id", Integer, "User id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// CreateUserPayload is supplied when creating a new user
var CreateUserPayload = Type("CreateUserPayload", func() {
	Attribute("first_name")
	Attribute("last_name")
	Required("first_name")
	Required("last_name")
})

// UpdateUserPayload is supplied when creating a new user
var UpdateUserPayload = Type("UpdateUserPayload", func() {
	Attribute("first_name")
	Attribute("last_name")
})

// UserMedia that a user belongs to
var UserMedia = MediaType("application/vnd.user+json", func() {

	Description("A user")

	Attributes(func() {
		Attribute("id", Integer, "User id")
		Attribute("href", String, "API href of user", func() {
			Example("/users/1")
		})
		Attribute("first_name", String, "First name", func() {
			MinLength(1)
			MaxLength(255)
		})
		Attribute("last_name", String, "Last name", func() {
			MinLength(1)
			MaxLength(255)
		})
		Attribute("name", String, "name", func() {
			MinLength(1)
			MaxLength(255)
		})
		Required("id", "href", "first_name", "last_name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("name")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
