package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// ClusterSnapshotsController implements the cluster-snapshots resource.
type ClusterSnapshotsController struct {
	*goa.Controller
}

// NewClusterSnapshotsController creates a cluster-snapshots controller.
func NewClusterSnapshotsController(service *goa.Service) *ClusterSnapshotsController {
	return &ClusterSnapshotsController{Controller: service.NewController("ClusterSnapshotsController")}
}

// List runs the list action.
func (c *ClusterSnapshotsController) List(ctx *app.ListClusterSnapshotsContext) error {
	// ClusterSnapshotsController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionClusterSnapshots{}
	return ctx.OK(res)
	// ClusterSnapshotsController_List: end_implement
}
