package routes

import "net/http"

// Checks if the user is authenticated and has a valid session
func isAuthenticated(w http.ResponseWriter, r *http.Request) {
	isLoggedIn := isLoggedIn(r)

	resp := map[string]interface{}{
		"success": isLoggedIn,
	}
	apiResponse(resp, w)
}

// Signs a user out of the application by removing any auth related cookies
func signOut(w http.ResponseWriter, r *http.Request) {
	emptyCookie := &http.Cookie{
		Name:   tokenCookieName,
		MaxAge: -1,
	}
	http.SetCookie(w, emptyCookie)

	resp := map[string]interface{}{
		"success": true,
	}
	apiResponse(resp, w)
}
