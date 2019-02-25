// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "yawarad": branche Resource Client
//
// Command:
// $ goagen
// --design=github.com/itamartempel/yawarad/design
// --out=$(GOPATH)/src/github.com/itamartempel/yawarad
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateBranchePath computes a request path to the create action of branche.
func CreateBranchePath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/branches/")
}

// Create a new branch
func (c *Client) CreateBranche(ctx context.Context, path string, payload *CreateBranchPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateBrancheRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateBrancheRequest create the request corresponding to the create action endpoint of the branche resource.
func (c *Client) NewCreateBrancheRequest(ctx context.Context, path string, payload *CreateBranchPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ListBranchePath computes a request path to the list action of branche.
func ListBranchePath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/branches/")
}

// Retrive branches
func (c *Client) ListBranche(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Response, error) {
	req, err := c.NewListBrancheRequest(ctx, path, clusterName, fromTime, limit, requestor, skip, status, team, toTime, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBrancheRequest create the request corresponding to the list action endpoint of the branche resource.
func (c *Client) NewListBrancheRequest(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if clusterName != nil {
		values.Set("cluster_name", *clusterName)
	}
	if fromTime != nil {
		tmp19 := strconv.Itoa(*fromTime)
		values.Set("from_time", tmp19)
	}
	if limit != nil {
		tmp20 := strconv.Itoa(*limit)
		values.Set("limit", tmp20)
	}
	if requestor != nil {
		values.Set("requestor", *requestor)
	}
	if skip != nil {
		tmp21 := strconv.Itoa(*skip)
		values.Set("skip", tmp21)
	}
	if status != nil {
		values.Set("status", *status)
	}
	if team != nil {
		values.Set("team", *team)
	}
	if toTime != nil {
		tmp22 := strconv.Itoa(*toTime)
		values.Set("to_time", tmp22)
	}
	if type_ != nil {
		values.Set("type", *type_)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowBranchePath computes a request path to the show action of branche.
func ShowBranchePath(branchID string) string {
	param0 := branchID

	return fmt.Sprintf("/api/v1/cluster-branching/branches/%s", param0)
}

// Retrieve single cluster branch by id.
func (c *Client) ShowBranche(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBrancheRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBrancheRequest create the request corresponding to the show action endpoint of the branche resource.
func (c *Client) NewShowBrancheRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateBranchePath computes a request path to the update action of branche.
func UpdateBranchePath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/branches/")
}

// Update a new branch
func (c *Client) UpdateBranche(ctx context.Context, path string, payload *UpdateBranchPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateBrancheRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateBrancheRequest create the request corresponding to the update action endpoint of the branche resource.
func (c *Client) NewUpdateBrancheRequest(ctx context.Context, path string, payload *UpdateBranchPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}