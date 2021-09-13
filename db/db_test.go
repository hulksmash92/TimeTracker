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

	expected := "host=localhost port=5432 user=postgres password=password123 dbname=timetracker sslmode=enable"

	result := generatePsqlConnStr()

	if result != expected {
		t.Errorf("Expected '%s' to equal '%s'", result, expected)
	}
}

// Tests the Pagination.GetPageSize() func
func Test_Pagination_GetPageSize(t *testing.T) {
	// Create an object for testing
	p := Pagination{
		PageIndex: 0,
		PageSize:  5,
		Sort:      "created",
		SortDesc:  true,
	}

	// Test for returning pagesize when > 0
	if pageSize := p.GetPageSize(); pageSize != 5 {
		t.Errorf("Expected Pagination.GetPageSize() to return 5 when PageSize is 5, but got %d", pageSize)
	}

	// Test for returning default page size of 10 when PageSize is 0
	p.PageSize = 0
	if pageSize := p.GetPageSize(); pageSize != 10 {
		t.Errorf("Expected Pagination.GetPageSize() to return 10 when PageSize is 0, but got %d", pageSize)
	}
}

// Tests the Pagination.Offset() function
func Test_Pagination_Offset(t *testing.T) {
	// Create an initial object will a PageIndex that should return an offset value
	p := Pagination{
		PageIndex: 2,
		PageSize:  5,
		Sort:      "created",
		SortDesc:  true,
	}

	// Check that the offset returned is PageIndex * PageSize when PageIndex > 0
	// fail when returned value is not 10 when PageIndex = 2 and PageSize is 5
	if offset := p.Offset(); offset != 10 {
		t.Errorf("Expected Pagination.Offset() to return 10 when PageSize is 5 and PageIndex is 2, but got %d", offset)
	}

	// Set the PageIndex to 0 so we can test for the 0 offset edge case
	p.PageIndex = 0

	// Check that the returned offset is 0, if 0 then fail
	if offset := p.Offset(); offset != 0 {
		t.Errorf("Expected Pagination.Offset() to return 0 when on first page, but got %d", offset)
	}
}

// Tests that Pagination.SortDirection() returns the correct values
func Test_Pagination_SortDirection(t *testing.T) {
	// Initialise a new Pagination structure with sort direction set to descending
	p := Pagination{
		PageIndex: 2,
		PageSize:  5,
		Sort:      "created",
		SortDesc:  true,
	}

	// Check that the `DESC` is returned when SortDesc is set to true, else fail
	if dir := p.SortDirection(); dir != "DESC" {
		t.Errorf("Expected Pagination.SortDirection() to return `DESC` when SortDesc is true but got %s", dir)
	}

	// Change sort to ascending to check for the returned `ASC` edge case
	p.SortDesc = false

	// Check that the ASC is returned when SortDesc is set to false, else fail
	if dir := p.SortDirection(); dir != "ASC" {
		t.Errorf("Expected Pagination.SortDirection() to return `ASC` when SortDesc is false but got %s", dir)
	}
}
