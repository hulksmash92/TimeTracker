package db

import (
	"database/sql"
	"fmt"
	"os"
)

// Global database connection, will be reused when a funcs calls connectDB()
var dbConnection *sql.DB

// Defines pagination params
type Pagination struct {
	PageSize  uint
	PageIndex uint
	Sort      string
	SortDesc  bool
}

// Gets the pagesize
func (p Pagination) GetPageSize() uint {
	if p.PageSize == 0 {
		return 20
	}
	return p.PageSize
}

// Gets the num of rows to offset by
func (p Pagination) Offset() uint {
	if p.PageIndex == 0 {
		return 0
	}
	return p.PageIndex * p.GetPageSize()
}

// Gets the direction of sort based on SortDesc val
func (p Pagination) SortDirection() string {
	if p.SortDesc {
		return "DESC"
	}
	return "ASC"
}

// Generates a PostgreSQL connection string from details defined in the
// environment variables starting with `PSQL_`
func generatePsqlConnStr() string {
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	user := os.Getenv("PSQL_USER")
	pass := os.Getenv("PSQL_PASS")
	db := os.Getenv("PSQL_DB")
	ssl := os.Getenv("PSQL_SSL")

	// Build our final connection string using the
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, db, ssl)
}

// Connects to the PostgreSQL DB and returns the open connection,
// handling any errors that may occur
func ConnectDB() *sql.DB {
	var err error

	if dbConnection != nil {
		err = dbConnection.Ping()
	} else {
		dbConnection, err = sql.Open("postgres", generatePsqlConnStr())
	}

	if err != nil {
		panic(err)
	}

	return dbConnection
}

// Closes the open database connection
func CloseDB() {
	if dbConnection != nil {
		dbConnection.Close()
	}
}
