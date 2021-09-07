package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
)

// Handles the user api calls by calling the method used to handle
// the appropriate HTTP Verb
func userRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

// Handles the get user api calls
func getUser(w http.ResponseWriter, r *http.Request) {
	token, err := parseTokenFromCookie(r)
	helpers.HandleError(err)

	ct, err := github.CheckToken(token)
	helpers.HandleError(err)

	user := db.GetUserByGitHubLogin(*ct.User.Login)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Structure of the values a user can update in the update user model
type UpdateUserReqBody struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// Handles the update user api calls
func updateUser(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)
	body := readBody(r)
	var newValues UpdateUserReqBody
	err := json.Unmarshal(body, &newValues)
	helpers.HandleError(err)

	// Call the user profile update func
	db.UpdateUserProfile(userId, newValues.Name, newValues.Email)

	resp := map[string]interface{}{
		"success": true,
	}
	apiResponse(resp, w)
}

// Deletes the selected user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the ID for the user in the DB
	// Check if it exists
	userId := getUserId(r)
	if userId == 0 {
		helpers.HandleError(errors.New("User does not exist"))
	}

	// Delete the user from the database
	db.DeleteUser(userId)

	// Log the user out from their current session as it will no longer be valid
	logoutUser(w, r)

	// send success signal back to caller if all statements
	// above succeeded, the error handler should not have been called
	resp := map[string]interface{}{"success": true}
	apiResponse(resp, w)
}
