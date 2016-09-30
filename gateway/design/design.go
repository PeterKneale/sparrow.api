package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Gateway API", func() {
	Title("Gateway API")
	Description("Gateway API")
	Host("localhost")
	Scheme("http")
	Consumes("application/json")
	Origin("*", func() { Methods("GET,POST,PUT,DELETE") })
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

var _ = Resource("health", func() {
	Action("alive", func() {
		Routing(GET("/alive"))
		Description("Perform health check.")
		Response(OK, "text/plain")
	})
})
