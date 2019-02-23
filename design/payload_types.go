package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CreateBranchPayload = Type("CreateBranchPayload", func() {
	Description("CreateBranchPayload defind a new branch request")
	Attribute("snapshot_id", String, "Snapshot image id that this branch will create from", func() {
		Example("Image_4946347")
	})
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
	Required("snapshot_id", "team", "cluster_name", "requestor", "cluster_name", "type")
})

var UpdateBranchPayload = Type("UpdateBranchPayload", func() {
	Description("UpdateBranchPayload defind an update request on existing branch")
	Attribute("operation", String, "The operation to be done on the existing branch", func() {
		Enum(getBranchUpdateOperations()...)
	})
	Attribute("comment", String, "Comment of this operation")
	Required("operation")
})
