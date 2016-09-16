package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("web1", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/web1", "web1/index.html")
})

var _ = Resource("web2", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/web2", "web2/index.html")
})

var _ = Resource("js1", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js1/*filepath", "web1/js1")
})

var _ = Resource("js2", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js2/*filepath", "web2/js2")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "web1/swagger/swagger.json")
})
