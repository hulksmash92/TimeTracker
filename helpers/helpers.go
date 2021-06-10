package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Handles an error
func HandleError(err error) {
	if err != nil {
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

// Generates the PostgreSQL connection string
func generatePsqlConnStr() string {
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	user := os.Getenv("PSQL_USER")
	pass := os.Getenv("PSQL_PASS")
	db := os.Getenv("PSQL_DB")
	ssl := os.Getenv("PSQL_SSL")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, db, ssl)
}

// Connects the PostgreSQL DB and returns the open connection
func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", generatePsqlConnStr())
	HandleError(err)
	return db
}
