package design

const (
	DefaultResultLimit int = 100
)

func getAvailableClusterTypeForBranching() []interface{} {
	var availableClusterTypes = []string{"mysql", "mongodb"}

	result := make([]interface{}, len(availableClusterTypes))
	for i, v := range availableClusterTypes {
		result[i] = v
	}
	return result

}

func getSkyJobClasses() []interface{} {
	var skyJobClasses = []string{"snapshot", "OnVault"}

	result := make([]interface{}, len(skyJobClasses))
	for i, v := range skyJobClasses {
		result[i] = v
	}
	return result

}

func getBranchTypes() []interface{} {
	var branchTypes = []string{"dev", "dr"}

	result := make([]interface{}, len(branchTypes))
	for i, v := range branchTypes {
		result[i] = v
	}
	return result

}
func getBranchStatus() []interface{} {
	var branchStatus = []string{"creating", "updateing", "active", "grace", "expired", "deleted"}

	result := make([]interface{}, len(branchStatus))
	for i, v := range branchStatus {
		result[i] = v
	}
	return result

}

func getBranchUpdateOperations() []interface{} {
	var branchUpdateOperations = []string{"extend_expiration", "expire_now"}

	result := make([]interface{}, len(branchUpdateOperations))
	for i, v := range branchUpdateOperations {
		result[i] = v
	}
	return result

}
func getRequestType() []interface{} {
	var requestType = []string{"create_branch", "update_branch", "extend_expration", "expire_now"}

	result := make([]interface{}, len(requestType))
	for i, v := range requestType {
		result[i] = v
	}
	return result

}

func getRequestStates() []interface{} {
	var requestStetes = []string{"running", "finish_successfully", "failed"}

	result := make([]interface{}, len(requestStetes))
	for i, v := range requestStetes {
		result[i] = v
	}
	return result

}
