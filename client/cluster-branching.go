// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "yawarad": cluster-branching Resource Client
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
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"strconv"
)

// CreateBranchClusterBranchingPath computes a request path to the create-branch action of cluster-branching.
func CreateBranchClusterBranchingPath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/branches")
}

// Create a new branch
func (c *Client) CreateBranchClusterBranching(ctx context.Context, path string, payload *CreateBranchPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateBranchClusterBranchingRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateBranchClusterBranchingRequest create the request corresponding to the create-branch action endpoint of the cluster-branching resource.
func (c *Client) NewCreateBranchClusterBranchingRequest(ctx context.Context, path string, payload *CreateBranchPayload, contentType string) (*http.Request, error) {
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

// ListBranchesClusterBranchingPath computes a request path to the list-branches action of cluster-branching.
func ListBranchesClusterBranchingPath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/branches")
}

// Retrive branches
func (c *Client) ListBranchesClusterBranching(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Response, error) {
	req, err := c.NewListBranchesClusterBranchingRequest(ctx, path, clusterName, fromTime, limit, requestor, skip, status, team, toTime, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBranchesClusterBranchingRequest create the request corresponding to the list-branches action endpoint of the cluster-branching resource.
func (c *Client) NewListBranchesClusterBranchingRequest(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Request, error) {
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

// ListClusterBranchesClusterBranchingPath computes a request path to the list-cluster-branches action of cluster-branching.
func ListClusterBranchesClusterBranchingPath(clusterName string) string {
	param0 := clusterName

	return fmt.Sprintf("/api/v1/cluster-branching/clusters/%s/branches", param0)
}

// Retrive branches of a specific cluster
func (c *Client) ListClusterBranchesClusterBranching(ctx context.Context, path string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Response, error) {
	req, err := c.NewListClusterBranchesClusterBranchingRequest(ctx, path, fromTime, limit, requestor, skip, status, team, toTime, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClusterBranchesClusterBranchingRequest create the request corresponding to the list-cluster-branches action endpoint of the cluster-branching resource.
func (c *Client) NewListClusterBranchesClusterBranchingRequest(ctx context.Context, path string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if fromTime != nil {
		tmp23 := strconv.Itoa(*fromTime)
		values.Set("from_time", tmp23)
	}
	if limit != nil {
		tmp24 := strconv.Itoa(*limit)
		values.Set("limit", tmp24)
	}
	if requestor != nil {
		values.Set("requestor", *requestor)
	}
	if skip != nil {
		tmp25 := strconv.Itoa(*skip)
		values.Set("skip", tmp25)
	}
	if status != nil {
		values.Set("status", *status)
	}
	if team != nil {
		values.Set("team", *team)
	}
	if toTime != nil {
		tmp26 := strconv.Itoa(*toTime)
		values.Set("to_time", tmp26)
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

// ListClusterSnapshotsClusterBranchingPath computes a request path to the list-cluster-snapshots action of cluster-branching.
func ListClusterSnapshotsClusterBranchingPath(clusterName string) string {
	param0 := clusterName

	return fmt.Sprintf("/api/v1/cluster-branching/clusters/%s/snapshots", param0)
}

// Retrive snapshots of a specific cluster
func (c *Client) ListClusterSnapshotsClusterBranching(ctx context.Context, path string, fromTime *int, limit *int, skip *int, toTime *int) (*http.Response, error) {
	req, err := c.NewListClusterSnapshotsClusterBranchingRequest(ctx, path, fromTime, limit, skip, toTime)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClusterSnapshotsClusterBranchingRequest create the request corresponding to the list-cluster-snapshots action endpoint of the cluster-branching resource.
func (c *Client) NewListClusterSnapshotsClusterBranchingRequest(ctx context.Context, path string, fromTime *int, limit *int, skip *int, toTime *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if fromTime != nil {
		tmp27 := strconv.Itoa(*fromTime)
		values.Set("from_time", tmp27)
	}
	if limit != nil {
		tmp28 := strconv.Itoa(*limit)
		values.Set("limit", tmp28)
	}
	if skip != nil {
		tmp29 := strconv.Itoa(*skip)
		values.Set("skip", tmp29)
	}
	if toTime != nil {
		tmp30 := strconv.Itoa(*toTime)
		values.Set("to_time", tmp30)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListClustersClusterBranchingPath computes a request path to the list-clusters action of cluster-branching.
func ListClustersClusterBranchingPath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/clusters")
}

// Retrieve all Available cluster that can be branch.
func (c *Client) ListClustersClusterBranching(ctx context.Context, path string, clusterType *string) (*http.Response, error) {
	req, err := c.NewListClustersClusterBranchingRequest(ctx, path, clusterType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClustersClusterBranchingRequest create the request corresponding to the list-clusters action endpoint of the cluster-branching resource.
func (c *Client) NewListClustersClusterBranchingRequest(ctx context.Context, path string, clusterType *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if clusterType != nil {
		values.Set("cluster_type", *clusterType)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListRequestTasksClusterBranchingPath computes a request path to the list-request-tasks action of cluster-branching.
func ListRequestTasksClusterBranchingPath(requestID string) string {
	param0 := requestID

	return fmt.Sprintf("/api/v1/cluster-branching/requests/%s/tasks", param0)
}

// Retrive request tasks
func (c *Client) ListRequestTasksClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListRequestTasksClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRequestTasksClusterBranchingRequest create the request corresponding to the list-request-tasks action endpoint of the cluster-branching resource.
func (c *Client) NewListRequestTasksClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListRequestsClusterBranchingPath computes a request path to the list-requests action of cluster-branching.
func ListRequestsClusterBranchingPath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/requests")
}

// Retrive requests
func (c *Client) ListRequestsClusterBranching(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, requestor *string, skip *int, state *string, toTime *int, type_ *string) (*http.Response, error) {
	req, err := c.NewListRequestsClusterBranchingRequest(ctx, path, clusterName, fromTime, limit, requestor, skip, state, toTime, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRequestsClusterBranchingRequest create the request corresponding to the list-requests action endpoint of the cluster-branching resource.
func (c *Client) NewListRequestsClusterBranchingRequest(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, requestor *string, skip *int, state *string, toTime *int, type_ *string) (*http.Request, error) {
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
		tmp31 := strconv.Itoa(*fromTime)
		values.Set("from_time", tmp31)
	}
	if limit != nil {
		tmp32 := strconv.Itoa(*limit)
		values.Set("limit", tmp32)
	}
	if requestor != nil {
		values.Set("requestor", *requestor)
	}
	if skip != nil {
		tmp33 := strconv.Itoa(*skip)
		values.Set("skip", tmp33)
	}
	if state != nil {
		values.Set("state", *state)
	}
	if toTime != nil {
		tmp34 := strconv.Itoa(*toTime)
		values.Set("to_time", tmp34)
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

// ListSnapshotBranchesClusterBranchingPath computes a request path to the list-snapshot-branches action of cluster-branching.
func ListSnapshotBranchesClusterBranchingPath(snapshotID string) string {
	param0 := snapshotID

	return fmt.Sprintf("/api/v1/cluster-branching/snapshots/%s/branches", param0)
}

// Retrive branches that was created from a specific snapshot
func (c *Client) ListSnapshotBranchesClusterBranching(ctx context.Context, path string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Response, error) {
	req, err := c.NewListSnapshotBranchesClusterBranchingRequest(ctx, path, fromTime, limit, requestor, skip, status, team, toTime, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSnapshotBranchesClusterBranchingRequest create the request corresponding to the list-snapshot-branches action endpoint of the cluster-branching resource.
func (c *Client) NewListSnapshotBranchesClusterBranchingRequest(ctx context.Context, path string, fromTime *int, limit *int, requestor *string, skip *int, status *string, team *string, toTime *int, type_ *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if fromTime != nil {
		tmp35 := strconv.Itoa(*fromTime)
		values.Set("from_time", tmp35)
	}
	if limit != nil {
		tmp36 := strconv.Itoa(*limit)
		values.Set("limit", tmp36)
	}
	if requestor != nil {
		values.Set("requestor", *requestor)
	}
	if skip != nil {
		tmp37 := strconv.Itoa(*skip)
		values.Set("skip", tmp37)
	}
	if status != nil {
		values.Set("status", *status)
	}
	if team != nil {
		values.Set("team", *team)
	}
	if toTime != nil {
		tmp38 := strconv.Itoa(*toTime)
		values.Set("to_time", tmp38)
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

// ListSnapshotsClusterBranchingPath computes a request path to the list-snapshots action of cluster-branching.
func ListSnapshotsClusterBranchingPath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/snapshots")
}

// Retrive snapshots
func (c *Client) ListSnapshotsClusterBranching(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, skip *int, toTime *int) (*http.Response, error) {
	req, err := c.NewListSnapshotsClusterBranchingRequest(ctx, path, clusterName, fromTime, limit, skip, toTime)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSnapshotsClusterBranchingRequest create the request corresponding to the list-snapshots action endpoint of the cluster-branching resource.
func (c *Client) NewListSnapshotsClusterBranchingRequest(ctx context.Context, path string, clusterName *string, fromTime *int, limit *int, skip *int, toTime *int) (*http.Request, error) {
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
		tmp39 := strconv.Itoa(*fromTime)
		values.Set("from_time", tmp39)
	}
	if limit != nil {
		tmp40 := strconv.Itoa(*limit)
		values.Set("limit", tmp40)
	}
	if skip != nil {
		tmp41 := strconv.Itoa(*skip)
		values.Set("skip", tmp41)
	}
	if toTime != nil {
		tmp42 := strconv.Itoa(*toTime)
		values.Set("to_time", tmp42)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowBrancheClusterBranchingPath computes a request path to the show-branche action of cluster-branching.
func ShowBrancheClusterBranchingPath(branchID string) string {
	param0 := branchID

	return fmt.Sprintf("/api/v1/cluster-branching/branches/%s", param0)
}

// Retrieve single cluster branch by id.
func (c *Client) ShowBrancheClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBrancheClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBrancheClusterBranchingRequest create the request corresponding to the show-branche action endpoint of the cluster-branching resource.
func (c *Client) NewShowBrancheClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowClusterClusterBranchingPath computes a request path to the show-cluster action of cluster-branching.
func ShowClusterClusterBranchingPath(clusterName string) string {
	param0 := clusterName

	return fmt.Sprintf("/api/v1/cluster-branching/clusters/%s", param0)
}

// Retrieve single cluster to be branch by name.
func (c *Client) ShowClusterClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowClusterClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowClusterClusterBranchingRequest create the request corresponding to the show-cluster action endpoint of the cluster-branching resource.
func (c *Client) NewShowClusterClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowRequestClusterBranchingPath computes a request path to the show-request action of cluster-branching.
func ShowRequestClusterBranchingPath(requestID string) string {
	param0 := requestID

	return fmt.Sprintf("/api/v1/cluster-branching/requests/%s", param0)
}

// Retrieve single request by id.
func (c *Client) ShowRequestClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowRequestClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowRequestClusterBranchingRequest create the request corresponding to the show-request action endpoint of the cluster-branching resource.
func (c *Client) NewShowRequestClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowRequestTaskClusterBranchingPath computes a request path to the show-request-task action of cluster-branching.
func ShowRequestTaskClusterBranchingPath(requestID string, taskName string) string {
	param0 := requestID
	param1 := taskName

	return fmt.Sprintf("/api/v1/cluster-branching/requests/%s/tasks/%s", param0, param1)
}

// Retrive a single request task by name
func (c *Client) ShowRequestTaskClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowRequestTaskClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowRequestTaskClusterBranchingRequest create the request corresponding to the show-request-task action endpoint of the cluster-branching resource.
func (c *Client) NewShowRequestTaskClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowRequestTaskLogClusterBranchingPath computes a request path to the show-request-task-log action of cluster-branching.
func ShowRequestTaskLogClusterBranchingPath(requestID string, taskName string) string {
	param0 := requestID
	param1 := taskName

	return fmt.Sprintf("/api/v1/cluster-branching/requests/%s/tasks/%s/log", param0, param1)
}

// Retrive a single request task by id
func (c *Client) ShowRequestTaskLogClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowRequestTaskLogClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowRequestTaskLogClusterBranchingRequest create the request corresponding to the show-request-task-log action endpoint of the cluster-branching resource.
func (c *Client) NewShowRequestTaskLogClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowSnapshotClusterBranchingPath computes a request path to the show-snapshot action of cluster-branching.
func ShowSnapshotClusterBranchingPath(snapshotID string) string {
	param0 := snapshotID

	return fmt.Sprintf("/api/v1/cluster-branching/snapshots/%s", param0)
}

// Retrieve single cluster snapshot by id.
func (c *Client) ShowSnapshotClusterBranching(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowSnapshotClusterBranchingRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowSnapshotClusterBranchingRequest create the request corresponding to the show-snapshot action endpoint of the cluster-branching resource.
func (c *Client) NewShowSnapshotClusterBranchingRequest(ctx context.Context, path string) (*http.Request, error) {
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

// SubscribeRequestChangesClusterBranchingPath computes a request path to the subscribe-request-changes action of cluster-branching.
func SubscribeRequestChangesClusterBranchingPath(requestID string) string {
	param0 := requestID

	return fmt.Sprintf("/api/v1/cluster-branching/requests/%s/ws", param0)
}

// Subscribe to any changes that will happend to a single request by websocket
func (c *Client) SubscribeRequestChangesClusterBranching(ctx context.Context, path string) (*websocket.Conn, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "ws"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	url_ := u.String()
	cfg, err := websocket.NewConfig(url_, url_)
	if err != nil {
		return nil, err
	}
	return websocket.DialConfig(cfg)
}

// SubscribeRequestTaskLogStreamClusterBranchingPath computes a request path to the subscribe-request-task-log-stream action of cluster-branching.
func SubscribeRequestTaskLogStreamClusterBranchingPath(requestID string, taskName string) string {
	param0 := requestID
	param1 := taskName

	return fmt.Sprintf("/api/v1/cluster-branching/requests/%s/tasks/%s/log/ws", param0, param1)
}

// Retrive a single request task by id
func (c *Client) SubscribeRequestTaskLogStreamClusterBranching(ctx context.Context, path string) (*websocket.Conn, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "ws"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	url_ := u.String()
	cfg, err := websocket.NewConfig(url_, url_)
	if err != nil {
		return nil, err
	}
	return websocket.DialConfig(cfg)
}

// UpdateBranchClusterBranchingPath computes a request path to the update-branch action of cluster-branching.
func UpdateBranchClusterBranchingPath() string {

	return fmt.Sprintf("/api/v1/cluster-branching/branches")
}

// Create a new branch
func (c *Client) UpdateBranchClusterBranching(ctx context.Context, path string, payload *UpdateBranchPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateBranchClusterBranchingRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateBranchClusterBranchingRequest create the request corresponding to the update-branch action endpoint of the cluster-branching resource.
func (c *Client) NewUpdateBranchClusterBranchingRequest(ctx context.Context, path string, payload *UpdateBranchPayload, contentType string) (*http.Request, error) {
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
