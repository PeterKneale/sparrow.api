package controllers

import "github.com/goadesign/goa"

// WebController implements the web resource.
type WebController struct {
	*goa.Controller
}

// NewWebController creates a web controller.
func NewWebController(service *goa.Service) *WebController {
	return &WebController{Controller: service.NewController("WebController")}
}
