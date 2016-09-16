package controllers

import (
	"github.com/goadesign/goa"
	"github.com/simplicate/mango/gateway/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Alive runs the alive action.
func (c *HealthController) Alive(ctx *app.AliveHealthContext) error {
	// HealthController_Alive: start_implement

	// Put your logic here

	// HealthController_Alive: end_implement
	return nil
}
