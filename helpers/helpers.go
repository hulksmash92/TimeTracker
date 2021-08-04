package helpers

// Checks the passed in error value and calls the panic() func if err has a value
func HandleError(err error) {
	if err != nil {
		// panic the application to alert something went wrong
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
