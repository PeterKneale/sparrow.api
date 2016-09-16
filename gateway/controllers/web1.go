package controllers

import "github.com/goadesign/goa"

// Web1Controller implements the web1 resource.
type Web1Controller struct {
	*goa.Controller
}

// NewWeb1Controller creates a web1 controller.
func NewWeb1Controller(service *goa.Service) *Web1Controller {
	return &Web1Controller{Controller: service.NewController("Web1Controller")}
}
