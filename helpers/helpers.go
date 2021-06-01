package helpers

// Handles an error
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Checks if an array contains the specified item
func StrArrayContains(arr []string, v string) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}
