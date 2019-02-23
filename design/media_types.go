package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ClusterForBranching = MediaType("application/vnd.cluster-for-branching+json", func() {
	Description("A Database Cluster with all the relevant data for creating new cluster branch from a snapshot")
	Attributes(func() {
		Attribute("id", UUID, "Uniq Id from Shapherd")
		Attribute("name", String, "Cluster name", func() {
			Example("my_cluster")
		})
		Attribute("cluster_type", String, "The cluster type", func() {
			Enum(getAvailableClusterTypeForBranching()...)
		})
		Attribute("servers", ArrayOf(String, func() {
			Format("hostname")
		}), "List of server in cluster", func() {
			Example([4]string{"db-mysql-my-cluster0a.42.wixprod.net",
				"db-mysql-my-cluster1b.42.wixprod.net",
				"db-mysql-my-cluster0a.96.wixprod.net",
				"db-mysql-my-cluster1b.96.wixprod.net"})
		})
		Attribute("master_hostname", String, "The current RW master of the cluster", func() {
			Format("hostname")
			Example("db-mysql-my-cluster0a.42.wixprod.net")
		})
		Attribute("latest_snapshot_id", String, "The latest snapshot thas was taken to this specific cluster", func() {
			Example("Image_4946347")
		})
		Attribute("top_ten_latest_snapshots", CollectionOf(ClusterSnapshot))
		Attribute("count_of_active_branches", Integer, "Count of active (none expire) branches", func() {
			Example(3)
		})

	})
	View("default", func() {
		Attribute("id")
		Attribute("cluster_type")
		Attribute("name")
		Attribute("servers")
		Attribute("master_hostname")
		Attribute("latest_snapshot_id")
		Attribute("top_ten_latest_snapshots")
		Attribute("count_of_active_branches")
	})
})

var ClusterSnapshot = MediaType("application/vnd.cluster-snapshot+json", func() {
	Description("A single sky snapshot for a specific cluster")
	Attributes(func() {
		Attribute("id", String, "Snapshot image id", func() {
			Example("Image_4946347")
		})
		Attribute("class", String, "Sky job class", func() {
			Enum(getSkyJobClasses()...)
		})
		Attribute("cluster_name", String, "Cluster name that the snapshot related to", func() {
			Example("my_cluster")
		})
		Attribute("created_at", DateTime, "The time that the snapshot was taken")
		Attribute("rollforword_till", DateTime, "The latest time that the specific snapshot can be rollforword (point in time restor) (not sure if we can calculate this)")
		Attribute("image_size_mb", Number, "Size of the image snapshot in MB", func() {
			Example(95367.4)
		})

	})
	View("default", func() {
		Attribute("id")
		Attribute("class")
		Attribute("cluster_name")
		Attribute("created_at")
		Attribute("rollforword_till")
		Attribute("image_size_mb")
	})
})

var ClusterBranch = MediaType("application/vnd.cluster-branch+json", func() {
	Description("A single cluster branch")
	Attributes(func() {
		Attribute("id", UUID, "Branch id")
		Attribute("snapshot", ClusterSnapshot, "Snapshot image that this branch was created from")
		Attribute("team", String, "The Team that this branch was created for", func() {
			Example("Booking")
		})
		Attribute("cluster_name", String, "Cluster name that the branch was created from", func() {
			Example("my_cluster")
		})
		Attribute("created_at", DateTime, "The time that the branch was created")
		Attribute("expire_at", DateTime, "The time that the branch will expire")
		Attribute("rollforword_till", DateTime, "Till what point in time this branch was created")
		Attribute("type", String, "The type of this branch", func() {
			Enum(getBranchTypes()...)
		})
		Attribute("status", String, "The status of the branch", func() {
			Enum(getBranchStatus()...)
		})

	})
	View("default", func() {
		Attribute("id")
		Attribute("snapshot")
		Attribute("team")
		Attribute("cluster_name")
		Attribute("created_at")
		Attribute("expire_at")
		Attribute("rollforword_till")
		Attribute("type")
		Attribute("status")
	})
})

var BranchRequestTask = MediaType("application/vnd.branch-request-task+json", func() {
	Description("A sigle task of operation")
	Attributes(func() {
		Attribute("id", UUID, "Request task Id")
		Attribute("seq", Integer, "The sequence order position of this specific task")
		Attribute("progress", Integer, "The progress of this task in presentage", func() {
			Default(0)
			Minimum(0)
			Maximum(100)
		})
		Attribute("last_log_message", String, "The last log messages")
		Attribute("name", String, "The task name")
		Attribute("description", String, "The description of the task")
		Attribute("start_at", DateTime, "The time that the task was started")
		Attribute("end_at", DateTime, "The time that the task was ended")
		Attribute("state", String, "The state of the current task", func() {
			Enum("pendding", "running", "finish_successfully", "failed")
			Default("pendding")
		})
	})
	View("default", func() {
		Attribute("id")
		Attribute("seq")
		Attribute("progress")
		Attribute("last_log_message")
		Attribute("name")
		Attribute("description")
		Attribute("start_at")
		Attribute("end_at")
		Attribute("state")
	})
})

var BranchRequest = MediaType("application/vnd.branch-request+json", func() {
	Description("A request")
	Attributes(func() {
		Attribute("request_id", UUID, "Request Id")
		Attribute("requestor", String, "The user that request the branch", func() {
			Format("email")
			Example("iisraelly@wix.com")
		})
		Attribute("jira_ticket", String, "The Jira ticket that related to this request", func() {
			Example("DBS-38329")
		})
		Attribute("branch_id", UUID, "Branch id")
		Attribute("request_type", String, "The request type", func() {
			Enum(getRequestType()...)
		})
		Attribute("tasks", ArrayOf(BranchRequestTask))
		Attribute("state", String, "The state of the request", func() {
			Enum(getRequestStates()...)
		})
	})
	View("default", func() {
		Attribute("request_id")
		Attribute("requestor")
		Attribute("jira_ticket")
		Attribute("branch_id")
		Attribute("request_type")
		Attribute("tasks")
	})
})
