package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// ClusterBranchesController implements the cluster-branches resource.
type ClusterBranchesController struct {
	*goa.Controller
}

// NewClusterBranchesController creates a cluster-branches controller.
func NewClusterBranchesController(service *goa.Service) *ClusterBranchesController {
	return &ClusterBranchesController{Controller: service.NewController("ClusterBranchesController")}
}

// List runs the list action.
func (c *ClusterBranchesController) List(ctx *app.ListClusterBranchesContext) error {
	// ClusterBranchesController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionClusterBranches{}
	return ctx.OK(res)
	// ClusterBranchesController_List: end_implement
}
