package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// BrancheController implements the branche resource.
type BrancheController struct {
	*goa.Controller
}

// NewBrancheController creates a branche controller.
func NewBrancheController(service *goa.Service) *BrancheController {
	return &BrancheController{Controller: service.NewController("BrancheController")}
}

// Create runs the create action.
func (c *BrancheController) Create(ctx *app.CreateBrancheContext) error {
	// BrancheController_Create: start_implement

	// Put your logic here

	return nil
	// BrancheController_Create: end_implement
}

// List runs the list action.
func (c *BrancheController) List(ctx *app.ListBrancheContext) error {
	// BrancheController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionClusterBranches{}
	return ctx.OK(res)
	// BrancheController_List: end_implement
}

// Show runs the show action.
func (c *BrancheController) Show(ctx *app.ShowBrancheContext) error {
	// BrancheController_Show: start_implement

	// Put your logic here

	res := &app.ClusterForBranching{}
	return ctx.OK(res)
	// BrancheController_Show: end_implement
}

// Update runs the update action.
func (c *BrancheController) Update(ctx *app.UpdateBrancheContext) error {
	// BrancheController_Update: start_implement

	// Put your logic here

	return nil
	// BrancheController_Update: end_implement
}
