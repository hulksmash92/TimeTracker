package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"timetracker/github"
	"timetracker/helpers"
)

// Recovers the application from a runtime error
func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		next.ServeHTTP(w, r)
	})
}

type SpaHandler struct {
	staticPath string
	indexPath  string
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		// file does not exist, so serve index.html and exist the function
		defaultFile := filepath.Join(h.staticPath, h.indexPath)
		http.ServeFile(w, r, defaultFile)
		return
	} else if err != nil {
		// if we got an error thats not file doesn't exist stating the
		// return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// Checks if the route requires authentication and then checks if the user is currently logged in
func CheckAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if routeRequiresAuthentication(r) && !isLoggedIn(r) {
			resp := map[string]interface{}{
				"message": "Access denied, please login and try again",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resp)
		}

		next.ServeHTTP(w, r)
	})
}

// Checks if the route requires authentication
func routeRequiresAuthentication(r *http.Request) bool {
	if !strings.Contains(r.RequestURI, "/api/") {
		return false
	}

	unauthedRoutes := []string{"/api/github/url", "/api/github/login"}
	return !helpers.StrArrayContains(unauthedRoutes, r.RequestURI)
}

// Checks if the user is currently logged in
func isLoggedIn(r *http.Request) bool {
	token, err := parseTokenFromCookie(r)
	if err != nil {
		println(err)
		return false
	}

	_, err = github.CheckToken(token)
	if err != nil {
		println(err)
		return false
	}

	return true
}

// Parses the login token from the LoginData HTTP cookie if it exists
func parseTokenFromCookie(r *http.Request) (string, error) {
	var token string
	cookie, err := r.Cookie("LoginData")
	if err != nil {
		return token, err
	}

	token = cookie.Value
	return token, nil
}
