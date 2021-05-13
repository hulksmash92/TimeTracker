package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

// Sets up the middleware pipeline
func initMiddleware() {
	setupStaticFileServer()
}

// Setups the route to serve the static files for the front-end
func setupStaticFileServer() {
	fileServer := http.FileServer(http.Dir("./ClientApp/dist"))
	http.Handle("/", http.StripPrefix("/", fileServer))
}

// Helper function to allow the api to recover from runtime error/panic
func panicHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()

		if err != nil {
			log.Println(err)

			resp := map[string]interface{}{
				"message": "Internal server error",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resp)
		}
	}()
}
