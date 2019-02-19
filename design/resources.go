package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Media Types

// GET - api/v1/cluster-branching/clusters?cluster_type
// GET - api/v1/cluster-branching/clusters/{:cluster_name}
// GET - api/v1/cluster-branching/clusters/{:cluster_name}/snapshots?from_time&to_time
// GET - api/v1/cluster-branching/clusters/{:cluster_name}/branches?expire
// GET - api/v1/cluster-branching/snapshots?&limit&skip&cluster&from_time&to_time
// GET - api/v1/cluster-branching/snapshots/{:snapshots_name}
// GET - api/v1/cluster-branching/snapshots/{:snapshots_name}/branches
// GET - api/v1/cluster-branching/branches?&limit&skip&status&cluster&user&team
// POST - api/v1/cluster-branching/branches # create new request to create a new branche
// GET - api/v1/cluster-branching/branches/{:branch_id}
// PUT - api/v1/cluster-branching/branches/{:branch_id}  # create new request to update the branche
// GET - api/v1/cluster-branching/request?status&cluster&branch&from_date&to_date&limit&skip
// GET - api/v1/cluster-branching/request/{:request_id}
// WS - api/v1/cluster-branching/request/{:request_id}
// GET - api/v1/cluster-branching/request/{:request_id}/steps
// GET - api/v1/cluster-branching/request/{:request_id}/steps/{step_name}
// GET - api/v1/cluster-branching/request/{:request_id}/steps/{step_name}/log
// WS - api/v1/cluster-branching/request/{:request_id}/steps/{step_name}/log
var _ = Resource("cluster-branching", func() {
	BasePath("/v1/cluster-branching")
	Action("list-clusters", func() {
		Routing(
			GET("/clusters"),
		)
		Params(func() {
			Param("cluster_type")
		})
		Description("Retrieve all Available cluster that can be branch.")
		Response(OK, CollectionOf(ClusterForBranching))
	})
	Action("show-cluster", func() {

	})
})
