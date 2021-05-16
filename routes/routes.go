package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"timetracker/helpers"

	"github.com/gorilla/mux"
)

// Initialises the HTTP router and pipeline, then listen and serves application
func ListenAndServe() {
	router := configureRouter()
	port := ":" + os.Getenv("PORT")
	log.Printf("Listening on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Configures the Mux Router for serving the front end and the API
func configureRouter() *mux.Router {
	router := mux.NewRouter()

	// Add any additional middleware
	router.Use(PanicHandler)

	// Configure any API routes
	router.HandleFunc("/api/github/url", getGitHubLoginUrl).Methods(http.MethodGet)
	router.HandleFunc("/api/github/login", getGitHubAccessToken).Methods(http.MethodPost)

	// Configure the static file serving for the SPA
	// This must be configured after API routes to stop any /api/
	// requests being redirected to our SPA
	spa := SpaHandler{staticPath: "./ClientApp/dist/", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	return router
}

// Reads the body of the http request to a byte array, handles any errors that may occur
func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleError(err)

	return body
}

// Correctly formats and encodes an API response
func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(call)
}
