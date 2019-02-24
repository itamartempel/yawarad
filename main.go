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

	// Mount "cluster-branching" controller
	c := NewClusterBranchingController(service)
	app.MountClusterBranchingController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
