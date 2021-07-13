package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

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
	dbConn, err := sql.Open("postgres", generatePsqlConnStr())
	HandleError(err)
	return dbConn
}

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
