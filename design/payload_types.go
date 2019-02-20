package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CreateBranchPayload = Type("CreateBranchPayload", func() {
	Description("CreateBranchPayload defind a new branch request")
	Attribute("snapshot_id", String, "Snapshot image id that this branch will create from")
	Attribute("team", String, "The Team that this branch will create for", func() {
		Example("Booking")
	})
	Attribute("cluster_name", String, "Cluster name that the branch will create from", func() {
		Example("my_cluster")
	})
	Attribute("requestor", String, "The user that request the branch", func() {
		Format("email")
		Example("iisraelly@wix.com")
	})
	Attribute("cluster_name", String, "Cluster name that the branch will created from", func() {
		Example("my_cluster")
	})
	Attribute("type", String, "The type of this branch", func() {
		Enum(getBranchTypes()...)
	})
	Required("snapshot", "team", "cluster_name", "requestor", "cluster_name", "type")
})
