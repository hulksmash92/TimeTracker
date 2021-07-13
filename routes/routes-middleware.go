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

// API routes that allow unauthenticated access
var unauthedRoutes = []string{
	"/api/github/url",
	"/api/github/login",
	"/api/auth/isAuthenticated",
}

// Middleware function for recovering the application from a panicking request
func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Using the defer keyword will cause this function to only run
		// after the surrounding function returns
		defer func() {
			// Recover the application from an error when panic() is called
			// by a go routine grab the error so we can do something with it,
			// if no error has occurred, err will be nil.
			err := recover()

			// if the error has a value check it and return something based on this error to the user
			if err != nil {
				log.Println(err)

				// Create a very generic error message for now
				// TODO: Check the actual error returned by recover() and return message a bit more specific
				//       to the error, like "Invalid access token" for revoked or expired tokens
				resp := map[string]interface{}{
					"message": "Internal server error",
				}

				// return the above error message to the caller
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(resp)
			}
		}()

		// Forward onto the next HTTP handler
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
	cookie, err := r.Cookie(tokenCookieName)
	if err != nil {
		return token, err
	}
	token = cookie.Value
	return token, nil
}
