package routes

import "net/http"

// Gets the github URL for logging into this app with GitHub
func isAuthenticated(w http.ResponseWriter, r *http.Request) {
	isLoggedIn := isLoggedIn(r)

	resp := map[string]interface{}{
		"success": isLoggedIn,
	}
	apiResponse(resp, w)
}
