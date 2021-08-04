package helpers

import (
	"errors"
	"testing"
)

func TestHandlerError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code did not panic")
		}
	}()
	HandleError(errors.New("Some test error"))
}

func TestStrArrayContains(t *testing.T) {
	arr := []string{"Chewbacca", "BB8", "R2D2", "Vader", "Obi-Wan"}

	// Test output for item that should be in the array
	if res := StrArrayContains(arr, "Chewbacca"); !res {
		t.Errorf("Array should contain 'Chewbacca'")
	}

	// Test the output for an item that shouldn't be in the array
	if res := StrArrayContains(arr, "C3P0"); res {
		t.Errorf("Array should not contain C3P0")
	}
}
