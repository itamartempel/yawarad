// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "yawarad": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/itamartempel/yawarad/design
// --out=$(GOPATH)/src/github.com/itamartempel/yawarad
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"time"
)

// A request (default view)
//
// Identifier: application/vnd.branch-request+json; view=default
type BranchRequest struct {
	// Branch id
	BranchID *uuid.UUID `form:"branch_id,omitempty" json:"branch_id,omitempty" yaml:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// The Jira ticket that related to this request
	JiraTicket *string `form:"jira_ticket,omitempty" json:"jira_ticket,omitempty" yaml:"jira_ticket,omitempty" xml:"jira_ticket,omitempty"`
	// Request Id
	RequestID *uuid.UUID `form:"request_id,omitempty" json:"request_id,omitempty" yaml:"request_id,omitempty" xml:"request_id,omitempty"`
	// The request type
	RequestType *string `form:"request_type,omitempty" json:"request_type,omitempty" yaml:"request_type,omitempty" xml:"request_type,omitempty"`
	// The user that request the branch
	Requestor *string              `form:"requestor,omitempty" json:"requestor,omitempty" yaml:"requestor,omitempty" xml:"requestor,omitempty"`
	Tasks     []*BranchRequestTask `form:"tasks,omitempty" json:"tasks,omitempty" yaml:"tasks,omitempty" xml:"tasks,omitempty"`
}

// Validate validates the BranchRequest media type instance.
func (mt *BranchRequest) Validate() (err error) {
	if mt.RequestType != nil {
		if !(*mt.RequestType == "create_branch" || *mt.RequestType == "update_branch" || *mt.RequestType == "extend_expration" || *mt.RequestType == "expire_now") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.request_type`, *mt.RequestType, []interface{}{"create_branch", "update_branch", "extend_expration", "expire_now"}))
		}
	}
	if mt.Requestor != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Requestor); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.requestor`, *mt.Requestor, goa.FormatEmail, err2))
		}
	}
	for _, e := range mt.Tasks {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBranchRequest decodes the BranchRequest instance encoded in resp body.
func (c *Client) DecodeBranchRequest(resp *http.Response) (*BranchRequest, error) {
	var decoded BranchRequest
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A sigle task of operation (default view)
//
// Identifier: application/vnd.branch-request-task+json; view=default
type BranchRequestTask struct {
	// The description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// The time that the task was ended
	EndAt *time.Time `form:"end_at,omitempty" json:"end_at,omitempty" yaml:"end_at,omitempty" xml:"end_at,omitempty"`
	// Request task Id
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The last log messages
	LastLogMessage *string `form:"last_log_message,omitempty" json:"last_log_message,omitempty" yaml:"last_log_message,omitempty" xml:"last_log_message,omitempty"`
	// The task name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// The progress of this task in presentage
	Progress int `form:"progress" json:"progress" yaml:"progress" xml:"progress"`
	// The sequence order position of this specific task
	Seq *int `form:"seq,omitempty" json:"seq,omitempty" yaml:"seq,omitempty" xml:"seq,omitempty"`
	// The time that the task was started
	StartAt *time.Time `form:"start_at,omitempty" json:"start_at,omitempty" yaml:"start_at,omitempty" xml:"start_at,omitempty"`
	// The state of the current task
	State string `form:"state" json:"state" yaml:"state" xml:"state"`
}

// Validate validates the BranchRequestTask media type instance.
func (mt *BranchRequestTask) Validate() (err error) {
	if mt.Progress < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.progress`, mt.Progress, 0, true))
	}
	if mt.Progress > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.progress`, mt.Progress, 100, false))
	}
	if !(mt.State == "pendding" || mt.State == "running" || mt.State == "finish_successfully" || mt.State == "failed") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.state`, mt.State, []interface{}{"pendding", "running", "finish_successfully", "failed"}))
	}
	return
}

// DecodeBranchRequestTask decodes the BranchRequestTask instance encoded in resp body.
func (c *Client) DecodeBranchRequestTask(resp *http.Response) (*BranchRequestTask, error) {
	var decoded BranchRequestTask
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Branch-Request-TaskCollection is the media type for an array of Branch-Request-Task (default view)
//
// Identifier: application/vnd.branch-request-task+json; type=collection; view=default
type BranchRequestTaskCollection []*BranchRequestTask

// Validate validates the BranchRequestTaskCollection media type instance.
func (mt BranchRequestTaskCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBranchRequestTaskCollection decodes the BranchRequestTaskCollection instance encoded in resp body.
func (c *Client) DecodeBranchRequestTaskCollection(resp *http.Response) (BranchRequestTaskCollection, error) {
	var decoded BranchRequestTaskCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Branch-RequestCollection is the media type for an array of Branch-Request (default view)
//
// Identifier: application/vnd.branch-request+json; type=collection; view=default
type BranchRequestCollection []*BranchRequest

// Validate validates the BranchRequestCollection media type instance.
func (mt BranchRequestCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBranchRequestCollection decodes the BranchRequestCollection instance encoded in resp body.
func (c *Client) DecodeBranchRequestCollection(resp *http.Response) (BranchRequestCollection, error) {
	var decoded BranchRequestCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A single cluster branch (default view)
//
// Identifier: application/vnd.cluster-branch+json; view=default
type ClusterBranch struct {
	// Cluster name that the branch was created from
	ClusterName *string `form:"cluster_name,omitempty" json:"cluster_name,omitempty" yaml:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// The time that the branch was created
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	// The time that the branch will expire
	ExpireAt *time.Time `form:"expire_at,omitempty" json:"expire_at,omitempty" yaml:"expire_at,omitempty" xml:"expire_at,omitempty"`
	// Branch id
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Till what point in time this branch was created
	RollforwordTill *time.Time `form:"rollforword_till,omitempty" json:"rollforword_till,omitempty" yaml:"rollforword_till,omitempty" xml:"rollforword_till,omitempty"`
	// Snapshot image that this branch was created from
	Snapshot *ClusterSnapshot `form:"snapshot,omitempty" json:"snapshot,omitempty" yaml:"snapshot,omitempty" xml:"snapshot,omitempty"`
	// The status of the branch
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// The Team that this branch was created for
	Team *string `form:"team,omitempty" json:"team,omitempty" yaml:"team,omitempty" xml:"team,omitempty"`
	// The type of this branch
	Type *string `form:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the ClusterBranch media type instance.
func (mt *ClusterBranch) Validate() (err error) {
	if mt.Snapshot != nil {
		if err2 := mt.Snapshot.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Status != nil {
		if !(*mt.Status == "creating" || *mt.Status == "updateing" || *mt.Status == "active" || *mt.Status == "grace" || *mt.Status == "expired" || *mt.Status == "deleted") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"creating", "updateing", "active", "grace", "expired", "deleted"}))
		}
	}
	if mt.Type != nil {
		if !(*mt.Type == "dev" || *mt.Type == "dr") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.type`, *mt.Type, []interface{}{"dev", "dr"}))
		}
	}
	return
}

// DecodeClusterBranch decodes the ClusterBranch instance encoded in resp body.
func (c *Client) DecodeClusterBranch(resp *http.Response) (*ClusterBranch, error) {
	var decoded ClusterBranch
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Cluster-BranchCollection is the media type for an array of Cluster-Branch (default view)
//
// Identifier: application/vnd.cluster-branch+json; type=collection; view=default
type ClusterBranchCollection []*ClusterBranch

// Validate validates the ClusterBranchCollection media type instance.
func (mt ClusterBranchCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeClusterBranchCollection decodes the ClusterBranchCollection instance encoded in resp body.
func (c *Client) DecodeClusterBranchCollection(resp *http.Response) (ClusterBranchCollection, error) {
	var decoded ClusterBranchCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A Database Cluster with all the relevant data for creating new cluster branch from a snapshot (default view)
//
// Identifier: application/vnd.cluster-for-branching+json; view=default
type ClusterForBranching struct {
	// The cluster type
	ClusterType *string `form:"cluster_type,omitempty" json:"cluster_type,omitempty" yaml:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Count of active (none expire) branches
	CountOfActiveBranches *int `form:"count_of_active_branches,omitempty" json:"count_of_active_branches,omitempty" yaml:"count_of_active_branches,omitempty" xml:"count_of_active_branches,omitempty"`
	// Uniq Id from Shapherd
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The latest snapshot thas was taken to this specific cluster
	LatestSnapshotID *string `form:"latest_snapshot_id,omitempty" json:"latest_snapshot_id,omitempty" yaml:"latest_snapshot_id,omitempty" xml:"latest_snapshot_id,omitempty"`
	// The current RW master of the cluster
	MasterHostname *string `form:"master_hostname,omitempty" json:"master_hostname,omitempty" yaml:"master_hostname,omitempty" xml:"master_hostname,omitempty"`
	// Cluster name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// List of server in cluster
	Servers               []string                  `form:"servers,omitempty" json:"servers,omitempty" yaml:"servers,omitempty" xml:"servers,omitempty"`
	TopTenLatestSnapshots ClusterSnapshotCollection `form:"top_ten_latest_snapshots,omitempty" json:"top_ten_latest_snapshots,omitempty" yaml:"top_ten_latest_snapshots,omitempty" xml:"top_ten_latest_snapshots,omitempty"`
}

// Validate validates the ClusterForBranching media type instance.
func (mt *ClusterForBranching) Validate() (err error) {
	if mt.ClusterType != nil {
		if !(*mt.ClusterType == "mysql" || *mt.ClusterType == "mongodb") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.cluster_type`, *mt.ClusterType, []interface{}{"mysql", "mongodb"}))
		}
	}
	if mt.MasterHostname != nil {
		if err2 := goa.ValidateFormat(goa.FormatHostname, *mt.MasterHostname); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.master_hostname`, *mt.MasterHostname, goa.FormatHostname, err2))
		}
	}
	for _, e := range mt.Servers {
		if err2 := goa.ValidateFormat(goa.FormatHostname, e); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.servers[*]`, e, goa.FormatHostname, err2))
		}
	}
	if err2 := mt.TopTenLatestSnapshots.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeClusterForBranching decodes the ClusterForBranching instance encoded in resp body.
func (c *Client) DecodeClusterForBranching(resp *http.Response) (*ClusterForBranching, error) {
	var decoded ClusterForBranching
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Cluster-For-BranchingCollection is the media type for an array of Cluster-For-Branching (default view)
//
// Identifier: application/vnd.cluster-for-branching+json; type=collection; view=default
type ClusterForBranchingCollection []*ClusterForBranching

// Validate validates the ClusterForBranchingCollection media type instance.
func (mt ClusterForBranchingCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeClusterForBranchingCollection decodes the ClusterForBranchingCollection instance encoded in resp body.
func (c *Client) DecodeClusterForBranchingCollection(resp *http.Response) (ClusterForBranchingCollection, error) {
	var decoded ClusterForBranchingCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A single sky snapshot for a specific cluster (default view)
//
// Identifier: application/vnd.cluster-snapshot+json; view=default
type ClusterSnapshot struct {
	// Sky job class
	Class *string `form:"class,omitempty" json:"class,omitempty" yaml:"class,omitempty" xml:"class,omitempty"`
	// Cluster name that the snapshot related to
	ClusterName *string `form:"cluster_name,omitempty" json:"cluster_name,omitempty" yaml:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// The time that the snapshot was taken
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	// Snapshot image id
	ID *string `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Size of the image snapshot in MB
	ImageSizeMb *float64 `form:"image_size_mb,omitempty" json:"image_size_mb,omitempty" yaml:"image_size_mb,omitempty" xml:"image_size_mb,omitempty"`
	// The latest time that the specific snapshot can be rollforword (point in time restor) (not sure if we can calculate this)
	RollforwordTill *time.Time `form:"rollforword_till,omitempty" json:"rollforword_till,omitempty" yaml:"rollforword_till,omitempty" xml:"rollforword_till,omitempty"`
}

// Validate validates the ClusterSnapshot media type instance.
func (mt *ClusterSnapshot) Validate() (err error) {
	if mt.Class != nil {
		if !(*mt.Class == "snapshot" || *mt.Class == "OnVault") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.class`, *mt.Class, []interface{}{"snapshot", "OnVault"}))
		}
	}
	return
}

// DecodeClusterSnapshot decodes the ClusterSnapshot instance encoded in resp body.
func (c *Client) DecodeClusterSnapshot(resp *http.Response) (*ClusterSnapshot, error) {
	var decoded ClusterSnapshot
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Cluster-SnapshotCollection is the media type for an array of Cluster-Snapshot (default view)
//
// Identifier: application/vnd.cluster-snapshot+json; type=collection; view=default
type ClusterSnapshotCollection []*ClusterSnapshot

// Validate validates the ClusterSnapshotCollection media type instance.
func (mt ClusterSnapshotCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeClusterSnapshotCollection decodes the ClusterSnapshotCollection instance encoded in resp body.
func (c *Client) DecodeClusterSnapshotCollection(resp *http.Response) (ClusterSnapshotCollection, error) {
	var decoded ClusterSnapshotCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Countable Collection of Clusters (default view)
//
// Identifier: application/vnd.countable-collection-cluster-branches+json; view=default
type CountableCollectionClusterBranches struct {
	// The slice of result from the result set
	Branches ClusterBranchCollection `form:"branches,omitempty" json:"branches,omitempty" yaml:"branches,omitempty" xml:"branches,omitempty"`
	// The total Count Of Results
	Count *int `form:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty" xml:"count,omitempty"`
	// The limit of result from the total result set
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty" xml:"limit,omitempty"`
	// The count of skiped results from the total result set
	Skip *int `form:"skip,omitempty" json:"skip,omitempty" yaml:"skip,omitempty" xml:"skip,omitempty"`
}

// Validate validates the CountableCollectionClusterBranches media type instance.
func (mt *CountableCollectionClusterBranches) Validate() (err error) {
	if err2 := mt.Branches.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeCountableCollectionClusterBranches decodes the CountableCollectionClusterBranches instance encoded in resp body.
func (c *Client) DecodeCountableCollectionClusterBranches(resp *http.Response) (*CountableCollectionClusterBranches, error) {
	var decoded CountableCollectionClusterBranches
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Countable Collection of Clusters (default view)
//
// Identifier: application/vnd.countable-collection-cluster-for-branching+json; view=default
type CountableCollectionClusterForBranching struct {
	// The slice of result from the result set
	Clusters ClusterForBranchingCollection `form:"clusters,omitempty" json:"clusters,omitempty" yaml:"clusters,omitempty" xml:"clusters,omitempty"`
	// The total Count Of Results
	Count *int `form:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty" xml:"count,omitempty"`
	// The limit of result from the total result set
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty" xml:"limit,omitempty"`
	// The count of skiped results from the total result set
	Skip *int `form:"skip,omitempty" json:"skip,omitempty" yaml:"skip,omitempty" xml:"skip,omitempty"`
}

// Validate validates the CountableCollectionClusterForBranching media type instance.
func (mt *CountableCollectionClusterForBranching) Validate() (err error) {
	if err2 := mt.Clusters.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeCountableCollectionClusterForBranching decodes the CountableCollectionClusterForBranching instance encoded in resp body.
func (c *Client) DecodeCountableCollectionClusterForBranching(resp *http.Response) (*CountableCollectionClusterForBranching, error) {
	var decoded CountableCollectionClusterForBranching
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Countable Collection of Clusters (default view)
//
// Identifier: application/vnd.countable-collection-cluster-snapshots+json; view=default
type CountableCollectionClusterSnapshots struct {
	// The total Count Of Results
	Count *int `form:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty" xml:"count,omitempty"`
	// The limit of result from the total result set
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty" xml:"limit,omitempty"`
	// The count of skiped results from the total result set
	Skip *int `form:"skip,omitempty" json:"skip,omitempty" yaml:"skip,omitempty" xml:"skip,omitempty"`
	// The slice of result from the result set
	Snapshots ClusterSnapshotCollection `form:"snapshots,omitempty" json:"snapshots,omitempty" yaml:"snapshots,omitempty" xml:"snapshots,omitempty"`
}

// Validate validates the CountableCollectionClusterSnapshots media type instance.
func (mt *CountableCollectionClusterSnapshots) Validate() (err error) {
	if err2 := mt.Snapshots.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeCountableCollectionClusterSnapshots decodes the CountableCollectionClusterSnapshots instance encoded in resp body.
func (c *Client) DecodeCountableCollectionClusterSnapshots(resp *http.Response) (*CountableCollectionClusterSnapshots, error) {
	var decoded CountableCollectionClusterSnapshots
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Countable Collection of Clusters (default view)
//
// Identifier: application/vnd.countable-collection-requests+json; view=default
type CountableCollectionRequests struct {
	// The total Count Of Results
	Count *int `form:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty" xml:"count,omitempty"`
	// The limit of result from the total result set
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty" xml:"limit,omitempty"`
	// The slice of result from the result set
	Requests BranchRequestCollection `form:"requests,omitempty" json:"requests,omitempty" yaml:"requests,omitempty" xml:"requests,omitempty"`
	// The count of skiped results from the total result set
	Skip *int `form:"skip,omitempty" json:"skip,omitempty" yaml:"skip,omitempty" xml:"skip,omitempty"`
}

// Validate validates the CountableCollectionRequests media type instance.
func (mt *CountableCollectionRequests) Validate() (err error) {
	if err2 := mt.Requests.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeCountableCollectionRequests decodes the CountableCollectionRequests instance encoded in resp body.
func (c *Client) DecodeCountableCollectionRequests(resp *http.Response) (*CountableCollectionRequests, error) {
	var decoded CountableCollectionRequests
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
