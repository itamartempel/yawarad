package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// RequestController implements the request resource.
type RequestController struct {
	*goa.Controller
}

// NewRequestController creates a request controller.
func NewRequestController(service *goa.Service) *RequestController {
	return &RequestController{Controller: service.NewController("RequestController")}
}

// List runs the list action.
func (c *RequestController) List(ctx *app.ListRequestContext) error {
	// RequestController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionRequests{}
	return ctx.OK(res)
	// RequestController_List: end_implement
}

// Show runs the show action.
func (c *RequestController) Show(ctx *app.ShowRequestContext) error {
	// RequestController_Show: start_implement

	// Put your logic here

	res := &app.BranchRequest{}
	return ctx.OK(res)
	// RequestController_Show: end_implement
}
