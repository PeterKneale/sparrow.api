package controllers

import "github.com/goadesign/goa"

// Js1Controller implements the js1 resource.
type Js1Controller struct {
	*goa.Controller
}

// NewJs1Controller creates a js1 controller.
func NewJs1Controller(service *goa.Service) *Js1Controller {
	return &Js1Controller{Controller: service.NewController("Js1Controller")}
}
