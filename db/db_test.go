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

func Test_Pagination_GetPageSize(t *testing.T) {
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

	// Test for returning default page size of 20 when PageSize is 0
	p.PageSize = 0
	if pageSize := p.GetPageSize(); pageSize != 20 {
		t.Errorf("Expected Pagination.GetPageSize() to return 20 when PageSize is 0, but got %d", pageSize)
	}
}

func Test_Pagination_Offset(t *testing.T) {
	p := Pagination{
		PageIndex: 2,
		PageSize:  5,
		Sort:      "created",
		SortDesc:  true,
	}

	// Test for returning pagesize when > 0
	if offset := p.Offset(); offset != 10 {
		t.Errorf("Expected Pagination.Offset() to return 10 when PageSize is 5 and PageIndex is 2, but got %d", offset)
	}

	// Test for returning default page size of 20 when PageSize is 0
	p.PageIndex = 0
	if offset := p.Offset(); offset != 0 {
		t.Errorf("Expected Pagination.Offset() to return 0 when on first page, but got %d", offset)
	}
}

func Test_Pagination_SortDirection(t *testing.T) {
	p := Pagination{
		PageIndex: 2,
		PageSize:  5,
		Sort:      "created",
		SortDesc:  true,
	}

	// Test for returning pagesize when > 0
	if dir := p.SortDirection(); dir != "DESC" {
		t.Errorf("Expected Pagination.SortDirection() to return `DESC` when SortDesc is true but got %s", dir)
	}

	// Test for returning default page size of 20 when PageSize is 0
	p.SortDesc = false
	if dir := p.SortDirection(); dir != "ASC" {
		t.Errorf("Expected Pagination.SortDirection() to return `ASC` when SortDesc is false but got %s", dir)
	}
}
