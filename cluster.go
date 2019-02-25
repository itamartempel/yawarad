package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
)

// ClusterController implements the cluster resource.
type ClusterController struct {
	*goa.Controller
}

// NewClusterController creates a cluster controller.
func NewClusterController(service *goa.Service) *ClusterController {
	return &ClusterController{Controller: service.NewController("ClusterController")}
}

// List runs the list action.
func (c *ClusterController) List(ctx *app.ListClusterContext) error {
	// ClusterController_List: start_implement

	// Put your logic here

	res := &app.CountableCollectionClusterForBranching{}
	return ctx.OK(res)
	// ClusterController_List: end_implement
}

// Show runs the show action.
func (c *ClusterController) Show(ctx *app.ShowClusterContext) error {
	// ClusterController_Show: start_implement

	// Put your logic here

	res := &app.ClusterForBranching{}
	return ctx.OK(res)
	// ClusterController_Show: end_implement
}
