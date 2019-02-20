package design

const (
	DefaultResultLimit int = 1000
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
	var branchStatus = []string{"active", "grace", "expired", "deleted"}

	result := make([]interface{}, len(branchStatus))
	for i, v := range branchStatus {
		result[i] = v
	}
	return result

}
