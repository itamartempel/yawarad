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
		Attribute("backup_hostname", String, "The current backup server of the cluster", func() {
			Format("hostname")
			Example("db-mysql-my-cluster0a.42.wixprod.net")
		})
		View("default", func() {
			Attribute("id")
			Attribute("cluster_type")
			Attribute("name")
			Attribute("servers")
			Attribute("master_hostname")
			Attribute("backup_hostname")
		})

	})
})
