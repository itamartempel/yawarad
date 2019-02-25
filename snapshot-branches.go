package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// SnapshotBranchesController implements the snapshot-branches resource.
type SnapshotBranchesController struct {
	*goa.Controller
}

// NewSnapshotBranchesController creates a snapshot-branches controller.
func NewSnapshotBranchesController(service *goa.Service) *SnapshotBranchesController {
	return &SnapshotBranchesController{Controller: service.NewController("SnapshotBranchesController")}
}

// List runs the list action.
func (c *SnapshotBranchesController) List(ctx *app.ListSnapshotBranchesContext) error {
	// SnapshotBranchesController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionClusterBranches{}
	return ctx.OK(res)
	// SnapshotBranchesController_List: end_implement
}
