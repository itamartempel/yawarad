// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "yawarad": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/itamartempel/yawarad/design
// --out=$(GOPATH)/src/github.com/itamartempel/yawarad
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// BrancheHref returns the resource href.
func BrancheHref(branchID interface{}) string {
	parambranchID := strings.TrimLeftFunc(fmt.Sprintf("%v", branchID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/cluster-branching/branches/%v", parambranchID)
}

// ClusterHref returns the resource href.
func ClusterHref(clusterName interface{}) string {
	paramclusterName := strings.TrimLeftFunc(fmt.Sprintf("%v", clusterName), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/cluster-branching/clusters/%v", paramclusterName)
}

// RequestHref returns the resource href.
func RequestHref(requestID interface{}) string {
	paramrequestID := strings.TrimLeftFunc(fmt.Sprintf("%v", requestID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/cluster-branching/requests/%v", paramrequestID)
}

// RequestTaskHref returns the resource href.
func RequestTaskHref(requestID, taskName interface{}) string {
	paramrequestID := strings.TrimLeftFunc(fmt.Sprintf("%v", requestID), func(r rune) bool { return r == '/' })
	paramtaskName := strings.TrimLeftFunc(fmt.Sprintf("%v", taskName), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/cluster-branching/requests/%v/tasks/%v", paramrequestID, paramtaskName)
}

// SnapshotHref returns the resource href.
func SnapshotHref(snapshotID interface{}) string {
	paramsnapshotID := strings.TrimLeftFunc(fmt.Sprintf("%v", snapshotID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/cluster-branching/snapshots/%v", paramsnapshotID)
}