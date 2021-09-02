package routes

import (
	"encoding/json"
	"net/http"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
	"timetracker/models"
)

// Handles the user api calls by calling the method used to handle
// the appropriate HTTP Verb
func userRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPatch:
		updateUser(w, r)
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
	Name          *string                `json:"name,omitempty"`
	Email         *string                `json:"email,omitempty"`
	ApiClients    []*models.ApiClient    `json:"apiClients,omitempty"`
	Organisations []*models.Organisation `json:"organisations,omitempty"`
}

// Handles the update user api calls
func updateUser(w http.ResponseWriter, r *http.Request) {
	//userId := getUserId(r)

	resp := map[string]interface{}{
		"success": true,
	}
	apiResponse(resp, w)
}
