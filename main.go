package main

import (
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
	fileServer := http.FileServer(http.Dir("./ClientApp/dist"))
	http.Handle("/", http.StripPrefix("/", fileServer))

	port := ":" + os.Getenv("PORT")
	log.Printf("Listening on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
