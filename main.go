//go:generate goagen bootstrap -d github.com/itamartempel/yawarad/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/itamartempel/yawarad/app"
)

func main() {
	// Create service
	service := goa.New("yawarad")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "branche" controller
	c := NewBrancheController(service)
	app.MountBrancheController(service, c)
	// Mount "cluster" controller
	c2 := NewClusterController(service)
	app.MountClusterController(service, c2)
	// Mount "cluster-branches" controller
	c3 := NewClusterBranchesController(service)
	app.MountClusterBranchesController(service, c3)
	// Mount "cluster-snapshots" controller
	c4 := NewClusterSnapshotsController(service)
	app.MountClusterSnapshotsController(service, c4)
	// Mount "request" controller
	c5 := NewRequestController(service)
	app.MountRequestController(service, c5)
	// Mount "request-changes" controller
	c6 := NewRequestChangesController(service)
	app.MountRequestChangesController(service, c6)
	// Mount "request-task" controller
	c7 := NewRequestTaskController(service)
	app.MountRequestTaskController(service, c7)
	// Mount "snapshot" controller
	c8 := NewSnapshotController(service)
	app.MountSnapshotController(service, c8)
	// Mount "snapshot-branches" controller
	c9 := NewSnapshotBranchesController(service)
	app.MountSnapshotBranchesController(service, c9)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
