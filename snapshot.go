package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// SnapshotController implements the snapshot resource.
type SnapshotController struct {
	*goa.Controller
}

// NewSnapshotController creates a snapshot controller.
func NewSnapshotController(service *goa.Service) *SnapshotController {
	return &SnapshotController{Controller: service.NewController("SnapshotController")}
}

// List runs the list action.
func (c *SnapshotController) List(ctx *app.ListSnapshotContext) error {
	// SnapshotController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionClusterSnapshots{}
	return ctx.OK(res)
	// SnapshotController_List: end_implement
}

// Show runs the show action.
func (c *SnapshotController) Show(ctx *app.ShowSnapshotContext) error {
	// SnapshotController_Show: start_implement

	// Put your logic here

	res := &app.ClusterForBranching{}
	return ctx.OK(res)
	// SnapshotController_Show: end_implement
}
