package controllers

import "github.com/goadesign/goa"

// Js2Controller implements the js2 resource.
type Js2Controller struct {
	*goa.Controller
}

// NewJs2Controller creates a js2 controller.
func NewJs2Controller(service *goa.Service) *Js2Controller {
	return &Js2Controller{Controller: service.NewController("Js2Controller")}
}
