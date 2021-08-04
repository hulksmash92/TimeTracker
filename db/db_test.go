package db

import (
	"os"
	"testing"
)

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
