package main

import (
	"log"

	"github.com/joho/godotenv"

	"timetracker/routes"
)

func main() {
	initDotEnv()
	routes.ListenAndServe()
}

// Loads the .env file into our system
func initDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
