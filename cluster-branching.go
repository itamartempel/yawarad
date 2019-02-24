package main

import (
	"github.com/goadesign/goa"
	"github.com/itamartempel/yawarad/app"
	"golang.org/x/net/websocket"
	"io"
)

// ClusterBranchingController implements the cluster-branching resource.
type ClusterBranchingController struct {
	*goa.Controller
}

// NewClusterBranchingController creates a cluster-branching controller.
func NewClusterBranchingController(service *goa.Service) *ClusterBranchingController {
	return &ClusterBranchingController{Controller: service.NewController("ClusterBranchingController")}
}

// CreateBranch runs the create-branch action.
func (c *ClusterBranchingController) CreateBranch(ctx *app.CreateBranchClusterBranchingContext) error {
	// ClusterBranchingController_CreateBranch: start_implement

	// Put your logic here

	return nil
	// ClusterBranchingController_CreateBranch: end_implement
}

// ListBranches runs the list-branches action.
func (c *ClusterBranchingController) ListBranches(ctx *app.ListBranchesClusterBranchingContext) error {
	// ClusterBranchingController_ListBranches: start_implement

	// Put your logic here

	res := app.ClusterBranchCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListBranches: end_implement
}

// ListClusterBranches runs the list-cluster-branches action.
func (c *ClusterBranchingController) ListClusterBranches(ctx *app.ListClusterBranchesClusterBranchingContext) error {
	// ClusterBranchingController_ListClusterBranches: start_implement

	// Put your logic here

	res := app.ClusterBranchCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListClusterBranches: end_implement
}

// ListClusterSnapshots runs the list-cluster-snapshots action.
func (c *ClusterBranchingController) ListClusterSnapshots(ctx *app.ListClusterSnapshotsClusterBranchingContext) error {
	// ClusterBranchingController_ListClusterSnapshots: start_implement

	// Put your logic here

	res := app.ClusterSnapshotCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListClusterSnapshots: end_implement
}

// ListClusters runs the list-clusters action.
func (c *ClusterBranchingController) ListClusters(ctx *app.ListClustersClusterBranchingContext) error {
	// ClusterBranchingController_ListClusters: start_implement

	// Put your logic here

	res := app.ClusterForBranchingCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListClusters: end_implement
}

// ListRequestTasks runs the list-request-tasks action.
func (c *ClusterBranchingController) ListRequestTasks(ctx *app.ListRequestTasksClusterBranchingContext) error {
	// ClusterBranchingController_ListRequestTasks: start_implement

	// Put your logic here

	res := app.BranchRequestTaskCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListRequestTasks: end_implement
}

// ListRequests runs the list-requests action.
func (c *ClusterBranchingController) ListRequests(ctx *app.ListRequestsClusterBranchingContext) error {
	// ClusterBranchingController_ListRequests: start_implement

	// Put your logic here

	res := app.BranchRequestCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListRequests: end_implement
}

// ListSnapshotBranches runs the list-snapshot-branches action.
func (c *ClusterBranchingController) ListSnapshotBranches(ctx *app.ListSnapshotBranchesClusterBranchingContext) error {
	// ClusterBranchingController_ListSnapshotBranches: start_implement

	// Put your logic here

	res := app.ClusterBranchCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListSnapshotBranches: end_implement
}

// ListSnapshots runs the list-snapshots action.
func (c *ClusterBranchingController) ListSnapshots(ctx *app.ListSnapshotsClusterBranchingContext) error {
	// ClusterBranchingController_ListSnapshots: start_implement

	// Put your logic here

	res := app.ClusterSnapshotCollection{}
	return ctx.OK(res)
	// ClusterBranchingController_ListSnapshots: end_implement
}

// ShowBranche runs the show-branche action.
func (c *ClusterBranchingController) ShowBranche(ctx *app.ShowBrancheClusterBranchingContext) error {
	// ClusterBranchingController_ShowBranche: start_implement

	// Put your logic here

	res := &app.ClusterForBranching{}
	return ctx.OK(res)
	// ClusterBranchingController_ShowBranche: end_implement
}

// ShowCluster runs the show-cluster action.
func (c *ClusterBranchingController) ShowCluster(ctx *app.ShowClusterClusterBranchingContext) error {
	// ClusterBranchingController_ShowCluster: start_implement

	// Put your logic here

	res := &app.ClusterForBranching{}
	return ctx.OK(res)
	// ClusterBranchingController_ShowCluster: end_implement
}

// ShowRequest runs the show-request action.
func (c *ClusterBranchingController) ShowRequest(ctx *app.ShowRequestClusterBranchingContext) error {
	// ClusterBranchingController_ShowRequest: start_implement

	// Put your logic here

	res := &app.BranchRequest{}
	return ctx.OK(res)
	// ClusterBranchingController_ShowRequest: end_implement
}

// ShowRequestTask runs the show-request-task action.
func (c *ClusterBranchingController) ShowRequestTask(ctx *app.ShowRequestTaskClusterBranchingContext) error {
	// ClusterBranchingController_ShowRequestTask: start_implement

	// Put your logic here

	res := &app.BranchRequestTask{}
	return ctx.OK(res)
	// ClusterBranchingController_ShowRequestTask: end_implement
}

// ShowRequestTaskLog runs the show-request-task-log action.
func (c *ClusterBranchingController) ShowRequestTaskLog(ctx *app.ShowRequestTaskLogClusterBranchingContext) error {
	// ClusterBranchingController_ShowRequestTaskLog: start_implement

	// Put your logic here

	res := &app.BranchRequestTask{}
	return ctx.OK(res)
	// ClusterBranchingController_ShowRequestTaskLog: end_implement
}

// ShowSnapshot runs the show-snapshot action.
func (c *ClusterBranchingController) ShowSnapshot(ctx *app.ShowSnapshotClusterBranchingContext) error {
	// ClusterBranchingController_ShowSnapshot: start_implement

	// Put your logic here

	res := &app.ClusterForBranching{}
	return ctx.OK(res)
	// ClusterBranchingController_ShowSnapshot: end_implement
}

// SubscribeRequestChanges runs the subscribe-request-changes action.
func (c *ClusterBranchingController) SubscribeRequestChanges(ctx *app.SubscribeRequestChangesClusterBranchingContext) error {
	c.SubscribeRequestChangesWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// SubscribeRequestChangesWSHandler establishes a websocket connection to run the subscribe-request-changes action.
func (c *ClusterBranchingController) SubscribeRequestChangesWSHandler(ctx *app.SubscribeRequestChangesClusterBranchingContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// ClusterBranchingController_SubscribeRequestChanges: start_implement

		// Put your logic here

		ws.Write([]byte("subscribe-request-changes cluster-branching"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// ClusterBranchingController_SubscribeRequestChanges: end_implement
	}
} // SubscribeRequestTaskLogStream runs the subscribe-request-task-log-stream action.
func (c *ClusterBranchingController) SubscribeRequestTaskLogStream(ctx *app.SubscribeRequestTaskLogStreamClusterBranchingContext) error {
	c.SubscribeRequestTaskLogStreamWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// SubscribeRequestTaskLogStreamWSHandler establishes a websocket connection to run the subscribe-request-task-log-stream action.
func (c *ClusterBranchingController) SubscribeRequestTaskLogStreamWSHandler(ctx *app.SubscribeRequestTaskLogStreamClusterBranchingContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// ClusterBranchingController_SubscribeRequestTaskLogStream: start_implement

		// Put your logic here

		ws.Write([]byte("subscribe-request-task-log-stream cluster-branching"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// ClusterBranchingController_SubscribeRequestTaskLogStream: end_implement
	}
} // UpdateBranch runs the update-branch action.
func (c *ClusterBranchingController) UpdateBranch(ctx *app.UpdateBranchClusterBranchingContext) error {
	// ClusterBranchingController_UpdateBranch: start_implement

	// Put your logic here

	return nil
	// ClusterBranchingController_UpdateBranch: end_implement
}
