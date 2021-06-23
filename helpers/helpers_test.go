package helpers

import (
	"errors"
	"os"
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

func Test_generatePsqlConnStr(t *testing.T) {
	// Set some env values when testing
	os.Setenv("PSQL_HOST", "localhost")
	os.Setenv("PSQL_PORT", "5432")
	os.Setenv("PSQL_USER", "postgres")
	os.Setenv("PSQL_PASS", "password123")
	os.Setenv("PSQL_DB", "timetracker")
	os.Setenv("PSQL_SSL", "disable")

	expected := "host=localhost port=5432 user=postgres password=password123 dbname=timetracker sslmode=disable"

	result := generatePsqlConnStr()

	if result != expected {
		t.Errorf("Expected '%s' to equal '%s'", result, expected)
	}
}
