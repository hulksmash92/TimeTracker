package helpers

import (
	"errors"
	"testing"
)

// Tests that the HandleError() function panics the
// application when an error is passed in
func Test_HandlerError(t *testing.T) {
	// Create a function that will be deferred the end of execution in this functon.
	// Will act as a try/catch in other languages and allows us to check if the
	// function caused a call to panic()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code did not panic")
		}
	}()

	// Throw an error that panics the application
	HandleError(errors.New("Some test error"))
}

// Tests that the StrArrayContains() function correctly checks string
// arrays/slices if a given value exists
func Test_StrArrayContains(t *testing.T) {
	// Initialise an array/slice of string values for testing
	arr := []string{"Chewbacca", "BB8", "R2D2", "Vader", "Obi-Wan"}

	// Test output for item that should be in the array
	if exists := StrArrayContains(arr, "Chewbacca"); !exists {
		t.Errorf("Array should contain 'Chewbacca'")
	}

	// Test the output for an item that shouldn't be in the array
	if exists := StrArrayContains(arr, "C3P0"); exists {
		t.Errorf("Array should not contain C3P0")
	}
}
