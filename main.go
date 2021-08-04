package main

import (
	"log"

	"github.com/joho/godotenv"

	"timetracker/db"
	"timetracker/routes"
)

func main() {
	initDotEnv()

	// open a new db connection and defer closing until the end of main()
	db.ConnectDB()
	defer db.CloseDB()

	routes.ListenAndServe()
}

// Loads the .env file into our system
func initDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
