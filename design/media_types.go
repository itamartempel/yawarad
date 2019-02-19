package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ClusterForBranching = MediaType("application/vnd.cluster-for-branching+json", func() {
	Description("A Database Cluster with all the relevant data for creating new cluster branch from a snapshot")
	Attributes(func() {
		Attribute("id", UUID, "Uniq Id from Shapherd")
		Attribute("name", String, "Cluster name")
		Attribute("cluster_type", String, "The cluster type", func() {
			Enum(getAvailableClusterType()...)
		})
		Attribute("servers", ArrayOf(String, func() {
			Format("hostname")
		}), "List of server in cluster", func() {
			Example([5]string{"db-mysql-my-cluster0a.42.wixprod.net",
				"db-mysql-my-cluster1b.42.wixprod.net",
				"db-mysql-my-cluster0a.96.wixprod.net",
				"db-mysql-my-cluster1b.96.wixprod.net",
				"db-mysql-my-cluster0a.42.wixprod.net"})
		})
		Attribute("master_hostname", String, "The current RW master of the cluster", func() {
			Format("hostname")
			Example("db-mysql-my-cluster0a.42.wixprod.net")
		})
		Attribute("latest_snapshot_id", String, "The latest snapshot thas was taken to this specific cluster", func() {
			Example("Image_4946347")
		})
		Attribute("top_ten_latest_snapshots")

	})
	View("default", func() {
		Attribute("id")
		Attribute("cluster_type")
		Attribute("name")
		Attribute("servers")
		Attribute("master_hostname")
		Attribute("latest_snapshot_id")
	})
	View("less", func() {
		Attribute("id")
		Attribute("cluster_type")
		Attribute("name")

	})
})

var ClusterSnapshot = MediaType("application/vnd.cluster-snapshot+json", func() {
	Description("A single sky snapshot for a specific cluster")
	Attributes(func() {
		Attribute("id", String, "Snapshot image id", func() {
			Example("Image_4946347")
		})

	})
	View("default", func() {
		Attribute("id")
	})
})
