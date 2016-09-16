package controllers

import "github.com/goadesign/goa"

// Web2Controller implements the web2 resource.
type Web2Controller struct {
	*goa.Controller
}

// NewWeb2Controller creates a web2 controller.
func NewWeb2Controller(service *goa.Service) *Web2Controller {
	return &Web2Controller{Controller: service.NewController("Web2Controller")}
}
