package design

import (
	"fmt"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Media Types

// var _ = Resource("cluster-branching", func() {
// 	BasePath("/v1/cluster-branching")
// 	// GET - api/v1/cluster-branching/clusters?cluster_type
// 	Action("list-clusters", func() {
// 		Routing(
// 			GET("/clusters"),
// 		)
// 		Params(func() {
// 			Param("cluster_type", func() {
// 				Enum(getAvailableClusterTypeForBranching()...)
// 			})
// 			Param("cluster_name", String, "Filter by cluster name (contains)", func() {
// 				Example("html_editor")
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d)", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")
// 		})
// 		Description("Retrieve all Available cluster that can be branch.")
// 		Response(OK, CountableCollectionOfClusterForBranching)
// 	})
// 	// GET - api/v1/cluster-branching/clusters/{:cluster_name}
// 	Action("show-cluster", func() {
// 		Routing(
// 			GET("/clusters/:cluster_name"),
// 		)
// 		Params(func() {
// 			Param("cluster_name", String, "The cluster name to show", func() {
// 				Example("html_editor")
// 			})
// 		})
// 		Description("Retrieve single cluster to be branch by name.")
// 		Response(OK, ClusterForBranching)
// 		Response(NotFound, func() {
// 			Description("If cluster name does not exists")
// 		})
// 	})
// 	// GET - api/v1/cluster-branching/clusters/{:cluster_name}/snapshots?from_time&to_time
// 	Action("list-cluster-snapshots", func() {
// 		Routing(
// 			GET("/clusters/:cluster_name/snapshots"),
// 		)
// 		Params(func() {
// 			Param("cluster_name", String, "The cluster name to show", func() {
// 				Example("html_editor")
// 			})
// 			Param("from_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created after the giving value (Epoch time in milliseconds)", func() {
// 				Example(1550657633000)
// 			})
// 			Param("to_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created before the giving value (Epoch time in milliseconds)", func() {
// 				Example(1547979233000)
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d)", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")
// 		})
// 		Description("Retrive snapshots of a specific cluster")
// 		Response(OK, CountableCollectionOfClusterSnapshots)
// 		Response(NotFound, func() {
// 			Description("If cluster name does not exists")
// 		})

// 	})

// 	// GET - api/v1/cluster-branching/clusters/{:cluster_name}/branches?status&team&requestor&type&from_time&to_time
// 	Action("list-cluster-branches", func() {
// 		Routing(
// 			GET("/clusters/:cluster_name/branches"),
// 		)
// 		Params(func() {
// 			Param("cluster_name", String, "The cluster name to show", func() {
// 				Example("html_editor")
// 			})
// 			Param("from_time", Integer, "Filter cluster branch by retriving all branch that was created after the giving value (Epoch time in milliseconds)", func() {
// 				Example(1550657633000)
// 			})
// 			Param("to_time", Integer, "Filter cluster branch by retriving all branch that was created before the giving value (Epoch time in milliseconds)", func() {
// 				Example(1547979233000)
// 			})
// 			Param("status", String, "Filter branches by branch status", func() {
// 				Enum(getBranchStatus()...)
// 			})
// 			Param("team", String, "Filter branches by team name (contains)")
// 			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
// 			Param("type", String, "Filter branches by branch type", func() {
// 				Enum(getBranchTypes()...)
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")

// 		})
// 		Description("Retrive branches of a specific cluster")
// 		Response(OK, CountableCollectionOfClusterBranches)
// 		Response(NotFound, func() {
// 			Description("If cluster name does not exists")
// 		})

// 	})
// 	// GET - api/v1/cluster-branching/snapshots?&limit&skip&cluster&from_time&to_time
// 	Action("list-snapshots", func() {
// 		Routing(
// 			GET("/snapshots"),
// 		)
// 		Params(func() {
// 			Param("cluster_name", String, "Filter by cluster name (contains)", func() {
// 				Example("html_editor")
// 			})
// 			Param("from_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created after the giving value (Epoch time in milliseconds)", func() {
// 				Example(1550657633000)
// 			})
// 			Param("to_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created before the giving value (Epoch time in milliseconds)", func() {
// 				Example(1547979233000)
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d)", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")
// 		})
// 		Description("Retrive snapshots")
// 		Response(OK, CountableCollectionOfClusterSnapshots)

// 	})
// 	// GET - api/v1/cluster-branching/snapshots/{:snapshot_id}
// 	Action("show-snapshot", func() {
// 		Routing(
// 			GET("/snapshots/:snapshot_id"),
// 		)
// 		Params(func() {
// 			Param("snapshot_id", String, "The snapshot id to show")
// 		})
// 		Description("Retrieve single cluster snapshot by id.")
// 		Response(OK, ClusterForBranching)
// 		Response(NotFound, func() {
// 			Description("If snapshot id does not exists")
// 		})
// 	})
// 	// GET - api/v1/cluster-branching/snapshots/{:snapshot_id}/branches
// 	Action("list-snapshot-branches", func() {
// 		Routing(
// 			GET("/snapshots/:snapshot_id/branches"),
// 		)
// 		Params(func() {
// 			Param("snapshot_id", String, "The snapshot id to show")
// 			Param("from_time", Integer, "Filter cluster branch by retriving all branch that was created after the giving value (Epoch time in milliseconds)", func() {
// 				Example(1550657633000)
// 			})
// 			Param("to_time", Integer, "Filter cluster branch by retriving all branch that was created before the giving value (Epoch time in milliseconds)", func() {
// 				Example(1547979233000)
// 			})
// 			Param("status", String, "Filter branches by branch status", func() {
// 				Enum(getBranchStatus()...)
// 			})
// 			Param("team", String, "Filter branches by team name (contains)")
// 			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
// 			Param("type", String, "Filter branches by branch type", func() {
// 				Enum(getBranchTypes()...)
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")

// 		})
// 		Description("Retrive branches that was created from a specific snapshot")
// 		Response(OK, CountableCollectionOfClusterBranches)
// 		Response(NotFound, func() {
// 			Description("If snapshot id does not exists")
// 		})

// 	})
// 	// GET - api/v1/cluster-branching/branches?&limit&skip&status&cluster&user&team
// 	Action("list-branches", func() {
// 		Routing(
// 			GET("/branches"),
// 		)
// 		Params(func() {
// 			Param("cluster_name", String, "Filter branch by cluster name (contains)", func() {
// 				Example("html_editor")
// 			})
// 			Param("from_time", Integer, "Filter cluster branch by retriving all branch that was created after the giving value (Epoch time in milliseconds)", func() {
// 				Example(1550657633000)
// 			})
// 			Param("to_time", Integer, "Filter cluster branch by retriving all branch that was created before the giving value (Epoch time in milliseconds)", func() {
// 				Example(1547979233000)
// 			})
// 			Param("status", String, "Filter branches by branch status", func() {
// 				Enum(getBranchStatus()...)
// 			})
// 			Param("team", String, "Filter branches by team name (contains)")
// 			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
// 			Param("type", String, "Filter branches by branch type", func() {
// 				Enum(getBranchTypes()...)
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")

// 		})
// 		Description("Retrive branches")
// 		Response(OK, CountableCollectionOfClusterBranches)

// 	})
// 	// GET - api/v1/cluster-branching/branches/{:branch_id}
// 	Action("show-branche", func() {
// 		Routing(
// 			GET("/branches/:branch_id"),
// 		)
// 		Params(func() {
// 			Param("branch_id", String, "The branch id to show")
// 		})
// 		Description("Retrieve single cluster branch by id.")
// 		Response(OK, ClusterForBranching)
// 		Response(NotFound, func() {
// 			Description("If branch id does not exists")
// 		})
// 	})
// 	// POST - api/v1/cluster-branching/branches # create new request to create a new branche
// 	Action("create-branch", func() {
// 		Routing(
// 			POST("/branches"),
// 		)
// 		Description("Create a new branch")
// 		Payload(CreateBranchPayload)
// 		Response(Created, BranchRequest)
// 		Response(BadRequest)
// 		Response(UnprocessableEntity)

// 	})
// 	// PUT - api/v1/cluster-branching/branches/{:branch_id}  # create new request to update the branche
// 	Action("update-branch", func() {
// 		Routing(
// 			PUT("/branches"),
// 		)
// 		Description("Create a new branch")
// 		Payload(UpdateBranchPayload)
// 		Response(Accepted, BranchRequest)
// 		Response(BadRequest)
// 		Response(UnprocessableEntity)

// 	})
// 	// GET - api/v1/cluster-branching/request?status&cluster&branch&from_date&to_date&limit&skip
// 	Action("list-requests", func() {
// 		Routing(
// 			GET("/requests"),
// 		)
// 		Params(func() {
// 			Param("cluster_name", String, "Filter request by cluster name (contains)", func() {
// 				Example("html_editor")
// 			})
// 			Param("from_time", Integer, "Filter requests that was created after the giving value (Epoch time in milliseconds)", func() {
// 				Example(1550657633000)
// 			})
// 			Param("to_time", Integer, "Filter request that was created before the giving value (Epoch time in milliseconds)", func() {
// 				Example(1547979233000)
// 			})
// 			Param("state", String, "Filter requests by request state", func() {
// 				Enum(getRequestStates()...)
// 			})
// 			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
// 			Param("type", String, "Filter request by request type", func() {
// 				Enum(getRequestType()...)
// 			})
// 			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
// 				Maximum(DefaultResultLimit)
// 				Default(DefaultResultLimit)
// 			})
// 			Param("skip", Integer, "For pagenation, skip results by giving value")

// 		})
// 		Description("Retrive requests")
// 		Response(OK, CountableCollectionOfRequests)
// 	})
// 	// GET - api/v1/cluster-branching/request/{:request_id}
// 	Action("show-request", func() {
// 		Routing(
// 			GET("/requests/:request_id"),
// 		)
// 		Params(func() {
// 			Param("request_id", String, "The request id to show")
// 		})
// 		Description("Retrieve single request by id.")
// 		Response(OK, BranchRequest)
// 		Response(NotFound, func() {
// 			Description("If request id does not exists")
// 		})
// 	})
// 	// WS - api/v1/cluster-branching/request/{:request_id}
// 	Action("subscribe-request-changes", func() {
// 		Routing(
// 			GET("/requests/:request_id/ws"),
// 		)
// 		Scheme("ws")
// 		Params(func() {
// 			Param("request_id", String, "The request id to show")
// 		})
// 		Description("Subscribe to any changes that will happend to a single request by websocket")
// 		Response(SwitchingProtocols)
// 	})
// 	// GET - api/v1/cluster-branching/request/{:request_id}/tasks
// 	Action("list-request-tasks", func() {
// 		Routing(
// 			GET("/requests/:request_id/tasks"),
// 		)
// 		Params(func() {
// 			Param("request_id", String, "The request id to show")
// 		})
// 		Description("Retrive request tasks")
// 		Response(OK, CollectionOf(BranchRequestTask))
// 	})
// 	// GET - api/v1/cluster-branching/request/{:request_id}/tasks/{task_name}
// 	Action("show-request-task", func() {
// 		Routing(
// 			GET("/requests/:request_id/tasks/:task_name"),
// 		)
// 		Params(func() {
// 			Param("request_id", String, "The request id to show")
// 			Param("task_name", String, "The request task name to show")
// 		})
// 		Description("Retrive a single request task by name")
// 		Response(OK, BranchRequestTask)
// 	})
// 	// GET - api/v1/cluster-branching/request/{:request_id}/tasks/{task_name}/log
// 	Action("show-request-task-log", func() {
// 		Routing(
// 			GET("/requests/:request_id/tasks/:task_name/log"),
// 		)
// 		Params(func() {
// 			Param("request_id", String, "The request id to show")
// 			Param("task_name", String, "The request task name to show")
// 		})
// 		Description("Retrive a single request task by id")
// 		Response(OK, BranchRequestTask)
// 	})
// 	// WS - api/v1/cluster-branching/request/{:request_id}/tasks/{task_name}/log
// 	Action("subscribe-request-task-log-stream", func() {
// 		Routing(
// 			GET("/requests/:request_id/tasks/:task_name/log/ws"),
// 		)
// 		Params(func() {
// 			Param("request_id", String, "The request id to show")
// 			Param("task_name", String, "The request task name to show")
// 		})
// 		Scheme("ws")
// 		Description("Retrive a single request task by id")
// 		Response(SwitchingProtocols)
// 	})

// })

var _ = Resource("cluster", func() {
	BasePath("/v1/cluster-branching/clusters")
	// GET - api/v1/cluster-branching/clusters?cluster_type
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("cluster_type", func() {
				Enum(getAvailableClusterTypeForBranching()...)
			})
			Param("cluster_name", String, "Filter by cluster name (contains)", func() {
				Example("html_editor")
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d)", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")
		})
		Description("Retrieve all Available cluster that can be branch.")
		Response(OK, CountableCollectionOfClusterForBranching)
	})
	// GET - api/v1/cluster-branching/clusters/{:cluster_name}
	Action("show", func() {
		Routing(
			GET("/:cluster_name"),
		)
		Params(func() {
			Param("cluster_name", String, "The cluster name to show", func() {
				Example("html_editor")
			})
		})
		Description("Retrieve single cluster to be branch by name.")
		Response(OK, ClusterForBranching)
		Response(NotFound, func() {
			Description("If cluster name does not exists")
		})
	})
})

var _ = Resource("cluster-snapshots", func() {
	Parent("cluster")
	// GET - api/v1/cluster-branching/clusters/{:cluster_name}/snapshots?from_time&to_time
	Action("list", func() {
		Routing(
			GET("/snapshots"),
		)
		Params(func() {
			Param("from_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created after the giving value (Epoch time in milliseconds)", func() {
				Example(1550657633000)
			})
			Param("to_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created before the giving value (Epoch time in milliseconds)", func() {
				Example(1547979233000)
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d)", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")
		})
		Description("Retrive snapshots of a specific cluster")
		Response(OK, CountableCollectionOfClusterSnapshots)
		Response(NotFound, func() {
			Description("If cluster name does not exists")
		})

	})

})

var _ = Resource("cluster-branches", func() {
	Parent("cluster")
	// GET - api/v1/cluster-branching/clusters/{:cluster_name}/branches?status&team&requestor&type&from_time&to_time
	Action("list", func() {
		Routing(
			GET("/branches"),
		)
		Params(func() {
			Param("from_time", Integer, "Filter cluster branch by retriving all branch that was created after the giving value (Epoch time in milliseconds)", func() {
				Example(1550657633000)
			})
			Param("to_time", Integer, "Filter cluster branch by retriving all branch that was created before the giving value (Epoch time in milliseconds)", func() {
				Example(1547979233000)
			})
			Param("status", String, "Filter branches by branch status", func() {
				Enum(getBranchStatus()...)
			})
			Param("team", String, "Filter branches by team name (contains)")
			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
			Param("type", String, "Filter branches by branch type", func() {
				Enum(getBranchTypes()...)
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")

		})
		Description("Retrive branches of a specific cluster")
		Response(OK, CountableCollectionOfClusterBranches)
		Response(NotFound, func() {
			Description("If cluster name does not exists")
		})

	})

})

var _ = Resource("snapshot", func() {
	BasePath("/v1/cluster-branching/snapshots")
	// GET - api/v1/cluster-branching/snapshots?&limit&skip&cluster&from_time&to_time
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("cluster_name", String, "Filter by cluster name (contains)", func() {
				Example("html_editor")
			})
			Param("from_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created after the giving value (Epoch time in milliseconds)", func() {
				Example(1550657633000)
			})
			Param("to_time", Integer, "Filter cluster snapshot by retriving all snapshot that was created before the giving value (Epoch time in milliseconds)", func() {
				Example(1547979233000)
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d)", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")
		})
		Description("Retrive snapshots")
		Response(OK, CountableCollectionOfClusterSnapshots)

	})
	// GET - api/v1/cluster-branching/snapshots/{:snapshot_id}
	Action("show", func() {
		Routing(
			GET("/:snapshot_id"),
		)
		Params(func() {
			Param("snapshot_id", String, "The snapshot id to show")
		})
		Description("Retrieve single cluster snapshot by id.")
		Response(OK, ClusterForBranching)
		Response(NotFound, func() {
			Description("If snapshot id does not exists")
		})
	})
})

var _ = Resource("snapshot-branches", func() {
	Parent("snapshot")
	// GET - api/v1/cluster-branching/snapshots/{:snapshot_id}/branches
	Action("list", func() {
		Routing(
			GET("/branches"),
		)
		Params(func() {
			Param("from_time", Integer, "Filter cluster branch by retriving all branch that was created after the giving value (Epoch time in milliseconds)", func() {
				Example(1550657633000)
			})
			Param("to_time", Integer, "Filter cluster branch by retriving all branch that was created before the giving value (Epoch time in milliseconds)", func() {
				Example(1547979233000)
			})
			Param("status", String, "Filter branches by branch status", func() {
				Enum(getBranchStatus()...)
			})
			Param("team", String, "Filter branches by team name (contains)")
			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
			Param("type", String, "Filter branches by branch type", func() {
				Enum(getBranchTypes()...)
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")

		})
		Description("Retrive branches that was created from a specific snapshot")
		Response(OK, CountableCollectionOfClusterBranches)
		Response(NotFound, func() {
			Description("If snapshot id does not exists")
		})

	})
})

var _ = Resource("branche", func() {
	BasePath("/v1/cluster-branching/branches")
	// GET - api/v1/cluster-branching/branches?&limit&skip&status&cluster&user&team
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("cluster_name", String, "Filter branch by cluster name (contains)", func() {
				Example("html_editor")
			})
			Param("from_time", Integer, "Filter cluster branch by retriving all branch that was created after the giving value (Epoch time in milliseconds)", func() {
				Example(1550657633000)
			})
			Param("to_time", Integer, "Filter cluster branch by retriving all branch that was created before the giving value (Epoch time in milliseconds)", func() {
				Example(1547979233000)
			})
			Param("status", String, "Filter branches by branch status", func() {
				Enum(getBranchStatus()...)
			})
			Param("team", String, "Filter branches by team name (contains)")
			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
			Param("type", String, "Filter branches by branch type", func() {
				Enum(getBranchTypes()...)
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")

		})
		Description("Retrive branches")
		Response(OK, CountableCollectionOfClusterBranches)

	})
	// GET - api/v1/cluster-branching/branches/{:branch_id}
	Action("show", func() {
		Routing(
			GET("/:branch_id"),
		)
		Params(func() {
			Param("branch_id", String, "The branch id to show")
		})
		Description("Retrieve single cluster branch by id.")
		Response(OK, ClusterForBranching)
		Response(NotFound, func() {
			Description("If branch id does not exists")
		})
	})
	// POST - api/v1/cluster-branching/branches # create new request to create a new branche
	Action("create", func() {
		Routing(
			POST("/"),
		)
		Description("Create a new branch")
		Payload(CreateBranchPayload)
		Response(Created, BranchRequest)
		Response(BadRequest)
		Response(UnprocessableEntity)

	})
	// PUT - api/v1/cluster-branching/branches/{:branch_id}  # create new request to update the branche
	Action("update", func() {
		Routing(
			PUT("/"),
		)
		Description("Update a new branch")
		Payload(UpdateBranchPayload)
		Response(Accepted, BranchRequest)
		Response(BadRequest)
		Response(UnprocessableEntity)

	})
})

var _ = Resource("request", func() {
	BasePath("/v1/cluster-branching/requests")
	// GET - api/v1/cluster-branching/request?status&cluster&branch&from_date&to_date&limit&skip
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("cluster_name", String, "Filter request by cluster name (contains)", func() {
				Example("html_editor")
			})
			Param("from_time", Integer, "Filter requests that was created after the giving value (Epoch time in milliseconds)", func() {
				Example(1550657633000)
			})
			Param("to_time", Integer, "Filter request that was created before the giving value (Epoch time in milliseconds)", func() {
				Example(1547979233000)
			})
			Param("state", String, "Filter requests by request state", func() {
				Enum(getRequestStates()...)
			})
			Param("requestor", String, "Filter branches by username of the reuqstor(contains)")
			Param("type", String, "Filter request by request type", func() {
				Enum(getRequestType()...)
			})
			Param("limit", Integer, fmt.Sprintf("limit the result set by giving value (default is %d )", DefaultResultLimit), func() {
				Maximum(DefaultResultLimit)
				Default(DefaultResultLimit)
			})
			Param("skip", Integer, "For pagenation, skip results by giving value")

		})
		Description("Retrive requests")
		Response(OK, CountableCollectionOfRequests)
	})
	// GET - api/v1/cluster-branching/request/{:request_id}
	Action("show", func() {
		Routing(
			GET("/:request_id"),
		)
		Params(func() {
			Param("request_id", String, "The request id to show")
		})
		Description("Retrieve single request by id.")
		Response(OK, BranchRequest)
		Response(NotFound, func() {
			Description("If request id does not exists")
		})
	})

})

var _ = Resource("request-changes", func() {
	Parent("request")
	// WS - api/v1/cluster-branching/request/{:request_id}
	Action("subscribe", func() {
		Routing(
			GET("/ws"),
		)
		Scheme("ws")
		Description("Subscribe to any changes that will happend to a single request by websocket")
		Response(SwitchingProtocols)
	})
})

var _ = Resource("request-task", func() {
	Parent("request")
	Action("list", func() {
		Routing(
			GET("/tasks"),
		)
		Description("Retrive request tasks")
		Response(OK, CollectionOf(BranchRequestTask))
	})
	// GET - api/v1/cluster-branching/request/{:request_id}/tasks/{task_name}
	Action("show", func() {
		Routing(
			GET("/tasks/:task_name"),
		)
		Params(func() {
			Param("task_name", String, "The request task name to show")
		})
		Description("Retrive a single request task by name")
		Response(OK, BranchRequestTask)
	})
	// GET - api/v1/cluster-branching/request/{:request_id}/tasks/{task_name}/log
	Action("log", func() {
		Routing(
			GET("/tasks/:task_name/log"),
		)
		Params(func() {
			Param("task_name", String, "The request task name to show")
		})
		Description("Retrive a single request task by id")
		Response(OK, BranchRequestTask)
	})
	// WS - api/v1/cluster-branching/request/{:request_id}/tasks/{task_name}/log
	Action("stream-log", func() {
		Routing(
			GET("/tasks/:task_name/log/ws"),
		)
		Params(func() {
			Param("task_name", String, "The request task name to show")
		})
		Scheme("ws")
		Description("Retrive a single request task by id")
		Response(SwitchingProtocols)
	})
})
