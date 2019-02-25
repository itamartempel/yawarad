// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "yawarad": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/itamartempel/yawarad/design
// --out=$(GOPATH)/src/github.com/itamartempel/yawarad
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// BrancheController is the controller interface for the Branche actions.
type BrancheController interface {
	goa.Muxer
	Create(*CreateBrancheContext) error
	List(*ListBrancheContext) error
	Show(*ShowBrancheContext) error
	Update(*UpdateBrancheContext) error
}

// MountBrancheController "mounts" a Branche resource controller on the given service.
func MountBrancheController(service *goa.Service, ctrl BrancheController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateBrancheContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateBranchPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/api/v1/cluster-branching/branches/", ctrl.MuxHandler("create", h, unmarshalCreateBranchePayload))
	service.LogInfo("mount", "ctrl", "Branche", "action", "Create", "route", "POST /api/v1/cluster-branching/branches/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListBrancheContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/branches/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Branche", "action", "List", "route", "GET /api/v1/cluster-branching/branches/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowBrancheContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/branches/:branch_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Branche", "action", "Show", "route", "GET /api/v1/cluster-branching/branches/:branch_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateBrancheContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateBranchPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/api/v1/cluster-branching/branches/", ctrl.MuxHandler("update", h, unmarshalUpdateBranchePayload))
	service.LogInfo("mount", "ctrl", "Branche", "action", "Update", "route", "PUT /api/v1/cluster-branching/branches/")
}

// unmarshalCreateBranchePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateBranchePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createBranchPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateBranchePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateBranchePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateBranchPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ClusterController is the controller interface for the Cluster actions.
type ClusterController interface {
	goa.Muxer
	List(*ListClusterContext) error
	Show(*ShowClusterContext) error
}

// MountClusterController "mounts" a Cluster resource controller on the given service.
func MountClusterController(service *goa.Service, ctrl ClusterController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListClusterContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/clusters/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Cluster", "action", "List", "route", "GET /api/v1/cluster-branching/clusters/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowClusterContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/clusters/:cluster_name", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Cluster", "action", "Show", "route", "GET /api/v1/cluster-branching/clusters/:cluster_name")
}

// ClusterBranchesController is the controller interface for the ClusterBranches actions.
type ClusterBranchesController interface {
	goa.Muxer
	List(*ListClusterBranchesContext) error
}

// MountClusterBranchesController "mounts" a ClusterBranches resource controller on the given service.
func MountClusterBranchesController(service *goa.Service, ctrl ClusterBranchesController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListClusterBranchesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/clusters/:cluster_name/branches", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "ClusterBranches", "action", "List", "route", "GET /api/v1/cluster-branching/clusters/:cluster_name/branches")
}

// ClusterSnapshotsController is the controller interface for the ClusterSnapshots actions.
type ClusterSnapshotsController interface {
	goa.Muxer
	List(*ListClusterSnapshotsContext) error
}

// MountClusterSnapshotsController "mounts" a ClusterSnapshots resource controller on the given service.
func MountClusterSnapshotsController(service *goa.Service, ctrl ClusterSnapshotsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListClusterSnapshotsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/clusters/:cluster_name/snapshots", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "ClusterSnapshots", "action", "List", "route", "GET /api/v1/cluster-branching/clusters/:cluster_name/snapshots")
}

// RequestController is the controller interface for the Request actions.
type RequestController interface {
	goa.Muxer
	List(*ListRequestContext) error
	Show(*ShowRequestContext) error
}

// MountRequestController "mounts" a Request resource controller on the given service.
func MountRequestController(service *goa.Service, ctrl RequestController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRequestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Request", "action", "List", "route", "GET /api/v1/cluster-branching/requests/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRequestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/:request_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Request", "action", "Show", "route", "GET /api/v1/cluster-branching/requests/:request_id")
}

// RequestChangesController is the controller interface for the RequestChanges actions.
type RequestChangesController interface {
	goa.Muxer
	Subscribe(*SubscribeRequestChangesContext) error
}

// MountRequestChangesController "mounts" a RequestChanges resource controller on the given service.
func MountRequestChangesController(service *goa.Service, ctrl RequestChangesController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSubscribeRequestChangesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Subscribe(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/:request_id/ws", ctrl.MuxHandler("subscribe", h, nil))
	service.LogInfo("mount", "ctrl", "RequestChanges", "action", "Subscribe", "route", "GET /api/v1/cluster-branching/requests/:request_id/ws")
}

// RequestTaskController is the controller interface for the RequestTask actions.
type RequestTaskController interface {
	goa.Muxer
	List(*ListRequestTaskContext) error
	Log(*LogRequestTaskContext) error
	Show(*ShowRequestTaskContext) error
	StreamLog(*StreamLogRequestTaskContext) error
}

// MountRequestTaskController "mounts" a RequestTask resource controller on the given service.
func MountRequestTaskController(service *goa.Service, ctrl RequestTaskController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRequestTaskContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/:request_id/tasks", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "RequestTask", "action", "List", "route", "GET /api/v1/cluster-branching/requests/:request_id/tasks")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLogRequestTaskContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Log(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/:request_id/tasks/:task_name/log", ctrl.MuxHandler("log", h, nil))
	service.LogInfo("mount", "ctrl", "RequestTask", "action", "Log", "route", "GET /api/v1/cluster-branching/requests/:request_id/tasks/:task_name/log")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRequestTaskContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/:request_id/tasks/:task_name", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "RequestTask", "action", "Show", "route", "GET /api/v1/cluster-branching/requests/:request_id/tasks/:task_name")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStreamLogRequestTaskContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.StreamLog(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/requests/:request_id/tasks/:task_name/log/ws", ctrl.MuxHandler("stream-log", h, nil))
	service.LogInfo("mount", "ctrl", "RequestTask", "action", "StreamLog", "route", "GET /api/v1/cluster-branching/requests/:request_id/tasks/:task_name/log/ws")
}

// SnapshotController is the controller interface for the Snapshot actions.
type SnapshotController interface {
	goa.Muxer
	List(*ListSnapshotContext) error
	Show(*ShowSnapshotContext) error
}

// MountSnapshotController "mounts" a Snapshot resource controller on the given service.
func MountSnapshotController(service *goa.Service, ctrl SnapshotController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListSnapshotContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/snapshots/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Snapshot", "action", "List", "route", "GET /api/v1/cluster-branching/snapshots/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowSnapshotContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/snapshots/:snapshot_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Snapshot", "action", "Show", "route", "GET /api/v1/cluster-branching/snapshots/:snapshot_id")
}

// SnapshotBranchesController is the controller interface for the SnapshotBranches actions.
type SnapshotBranchesController interface {
	goa.Muxer
	List(*ListSnapshotBranchesContext) error
}

// MountSnapshotBranchesController "mounts" a SnapshotBranches resource controller on the given service.
func MountSnapshotBranchesController(service *goa.Service, ctrl SnapshotBranchesController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListSnapshotBranchesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/cluster-branching/snapshots/:snapshot_id/branches", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "SnapshotBranches", "action", "List", "route", "GET /api/v1/cluster-branching/snapshots/:snapshot_id/branches")
}