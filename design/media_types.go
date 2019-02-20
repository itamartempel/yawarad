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

var BranchRequest = MediaType("application/vnd.branch-request+json", func() {
	Description("A request")
	Attributes(func() {
		Attribute("id", UUID, "Request Id")
		Attribute("requestor", String, "The user that request the branch", func() {
			Format("email")
			Example("iisraelly@wix.com")
		})
		Attribute("jira_ticket", String, "The Jira ticket that related to this request", func() {
			Example("DBS-38329")
		})
	})
	View("default", func() {
		Attribute("id")
		Attribute("requestor")
		Attribute("jira_ticket")
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
		Attribute("request", BranchRequest, "The request that created this burch")

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
		Attribute("request")
		Attribute("cluster_name")
		Attribute("created_at")
		Attribute("expire_at")
		Attribute("rollforword_till")
		Attribute("type")
		Attribute("status")
	})
})
