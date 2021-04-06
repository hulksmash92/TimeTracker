package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	initDotEnv()

	port := ":" + os.Getenv("PORT")

	log.Printf("Listening on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(port, initRouter()))
}

func initDotEnv() { // Loads the .env file into out system
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func initRouter() *mux.Router {
	// Initialise out mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)

	return router
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}
