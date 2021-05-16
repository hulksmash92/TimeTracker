package helpers

// Handles an error
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
