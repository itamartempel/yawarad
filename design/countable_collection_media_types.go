package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CountableCollectionOfClusterForBranching = MediaType("application/vnd.countable-collection-cluster-for-branching+json", func() {
	Description("Countable Collection of Clusters")
	Attributes(func() {
		Attribute("count", Integer, "The total Count Of Results", func() {
			Example(606)
		})
		Attribute("limit", Integer, "The limit of result from the total result set", func() {
			Example(DefaultResultLimit)
		})
		Attribute("skip", Integer, "The count of skiped results from the total result set", func() {
			Example(0)
		})

		Attribute("clusters", CollectionOf(ClusterForBranching), "The slice of result from the result set")

	})
	View("default", func() {
		Attribute("count")
		Attribute("limit")
		Attribute("skip")
		Attribute("clusters")
	})
})

var CountableCollectionOfClusterSnapshots = MediaType("application/vnd.countable-collection-cluster-snapshots+json", func() {
	Description("Countable Collection of Clusters")
	Attributes(func() {
		Attribute("count", Integer, "The total Count Of Results", func() {
			Example(1029)
		})
		Attribute("limit", Integer, "The limit of result from the total result set", func() {
			Example(DefaultResultLimit)
		})
		Attribute("skip", Integer, "The count of skiped results from the total result set", func() {
			Example(0)
		})

		Attribute("snapshots", CollectionOf(ClusterSnapshot), "The slice of result from the result set")

	})
	View("default", func() {
		Attribute("count")
		Attribute("limit")
		Attribute("skip")
		Attribute("snapshots")
	})
})

var CountableCollectionOfClusterBranches = MediaType("application/vnd.countable-collection-cluster-branches+json", func() {
	Description("Countable Collection of Clusters")
	Attributes(func() {
		Attribute("count", Integer, "The total Count Of Results", func() {
			Example(798)
		})
		Attribute("limit", Integer, "The limit of result from the total result set", func() {
			Example(DefaultResultLimit)
		})
		Attribute("skip", Integer, "The count of skiped results from the total result set", func() {
			Example(0)
		})

		Attribute("branches", CollectionOf(ClusterBranch), "The slice of result from the result set")

	})
	View("default", func() {
		Attribute("count")
		Attribute("limit")
		Attribute("skip")
		Attribute("branches")
	})
})

var CountableCollectionOfRequests = MediaType("application/vnd.countable-collection-requests+json", func() {
	Description("Countable Collection of Clusters")
	Attributes(func() {
		Attribute("count", Integer, "The total Count Of Results", func() {
			Example(4321)
		})
		Attribute("limit", Integer, "The limit of result from the total result set", func() {
			Example(DefaultResultLimit)
		})
		Attribute("skip", Integer, "The count of skiped results from the total result set", func() {
			Example(0)
		})

		Attribute("requests", CollectionOf(BranchRequest), "The slice of result from the result set")

	})
	View("default", func() {
		Attribute("count")
		Attribute("limit")
		Attribute("skip")
		Attribute("requests")
	})
})
