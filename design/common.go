package design

func getAvailableClusterType() []interface{} {
	var availableClusterTypes = []string{"mysql", "mongodb"}

	result := make([]interface{}, len(availableClusterTypes))
	for i, v := range availableClusterTypes {
		result[i] = v
	}
	return result

}
