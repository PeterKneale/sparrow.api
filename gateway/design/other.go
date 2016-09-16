package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("web", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/web", "web/index.html")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "web/js")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "web/swagger/swagger.json")
})
