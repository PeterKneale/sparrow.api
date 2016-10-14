package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Account", func() {

	DefaultMedia(AccountMedia)

	BasePath("/")

	Action("read", func() {
		Routing(GET("/account"))
		Description("read current account")
		Response(OK, AccountMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(PUT("/account"))
		Description("Update account")
		Payload(UpdateAccountPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// UpdateAccountPayload is supplied when creating a new user
var UpdateAccountPayload = Type("UpdateAccountPayload", func() {
	Attribute("name")
})

// AccountMedia that a user belongs to
var AccountMedia = MediaType("application/vnd.account+json", func() {

	Description("An account")

	Attributes(func() {
		Attribute("name", String, "Account name", func() {
			MinLength(1)
			MaxLength(255)
		})
		Required("name")
	})

	View("default", func() {
		Attribute("name")
	})
})
