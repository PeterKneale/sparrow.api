package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Sparrow API", func() {
	Title("Sparrow API")
	Description("Sparrow API")
	Host("api.webapitester.com")
	Scheme("http")

	Consumes("application/json")
	Produces("application/json")

	Origin("*", func() { Methods("GET,POST,PUT,DELETE,OPTIONS") })

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

var _ = Resource("meta", func() {
	Action("alive", func() {
		Routing(GET("/health/alive"))
		Description("Perform aliveness check.")
		Response(OK, "text/plain")
	})
	Action("ready", func() {
		Routing(GET("/health/ready"))
		Description("Perform readiness check.")
		Response(OK, "text/plain")
	})
	Action("root", func() {
		Routing(GET("/"))
		Description("Perform root check.")
		Response(OK, "text/plain")
	})
})
