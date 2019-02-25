// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "yawarad": snapshot TestHelpers
//
// Command:
// $ goagen
// --design=github.com/itamartempel/yawarad/design
// --out=$(GOPATH)/src/github.com/itamartempel/yawarad
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/itamartempel/yawarad/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// ListSnapshotOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListSnapshotOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SnapshotController, clusterName *string, fromTime *int, limit int, skip *int, toTime *int) (http.ResponseWriter, *app.CountableCollectionClusterSnapshots) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if clusterName != nil {
		sliceVal := []string{*clusterName}
		query["cluster_name"] = sliceVal
	}
	if fromTime != nil {
		sliceVal := []string{strconv.Itoa(*fromTime)}
		query["from_time"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	if skip != nil {
		sliceVal := []string{strconv.Itoa(*skip)}
		query["skip"] = sliceVal
	}
	if toTime != nil {
		sliceVal := []string{strconv.Itoa(*toTime)}
		query["to_time"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/cluster-branching/snapshots/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if clusterName != nil {
		sliceVal := []string{*clusterName}
		prms["cluster_name"] = sliceVal
	}
	if fromTime != nil {
		sliceVal := []string{strconv.Itoa(*fromTime)}
		prms["from_time"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	if skip != nil {
		sliceVal := []string{strconv.Itoa(*skip)}
		prms["skip"] = sliceVal
	}
	if toTime != nil {
		sliceVal := []string{strconv.Itoa(*toTime)}
		prms["to_time"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SnapshotTest"), rw, req, prms)
	listCtx, _err := app.NewListSnapshotContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CountableCollectionClusterSnapshots
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.CountableCollectionClusterSnapshots)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CountableCollectionClusterSnapshots", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowSnapshotNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowSnapshotNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SnapshotController, snapshotID string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/cluster-branching/snapshots/%v", snapshotID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["snapshot_id"] = []string{fmt.Sprintf("%v", snapshotID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SnapshotTest"), rw, req, prms)
	showCtx, _err := app.NewShowSnapshotContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ShowSnapshotOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowSnapshotOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SnapshotController, snapshotID string) (http.ResponseWriter, *app.ClusterForBranching) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/cluster-branching/snapshots/%v", snapshotID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["snapshot_id"] = []string{fmt.Sprintf("%v", snapshotID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SnapshotTest"), rw, req, prms)
	showCtx, _err := app.NewShowSnapshotContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ClusterForBranching
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ClusterForBranching)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ClusterForBranching", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
