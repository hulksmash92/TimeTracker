package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	initDotEnv()
	initRouter()
}

func initDotEnv() { // Loads the .env file into out system
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func initRouter() {
	http.HandleFunc("/", homePage)
	port := ":" + os.Getenv("PORT")

	log.Printf("Listening on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}
