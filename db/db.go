package db

import (
	"database/sql"
	"fmt"
	"os"
)

var dbConnection *sql.DB

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
func connectDB() *sql.DB {
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
