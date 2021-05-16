package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

	// Configure the static file serving for the SPA
	spa := SpaHandler{staticPath: "./ClientApp/dist/", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	// Configure any API routes
	router.HandleFunc("/api/test", testRoute).Methods(http.MethodGet)

	return router
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"message": "Success",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
