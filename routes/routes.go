package routes

import (
	"log"
	"net/http"
	"os"
)

// Initialises the HTTP router and pipeline, then listen and serves application
func ListenAndServe() {
	initMiddleware()
	initApiRoutes()

	port := ":" + os.Getenv("PORT")
	log.Printf("Listening on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(port, http.HandlerFunc(panicHandler)))
}

func initApiRoutes() {
	http.HandleFunc("/api/test", testRoute)
}

func testRoute(w http.ResponseWriter, r *http.Request) {

}
