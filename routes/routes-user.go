package routes

import (
	"encoding/json"
	"net/http"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	token, err := parseTokenFromCookie(r)
	helpers.HandleError(err)

	ct, err := github.CheckToken(token)
	helpers.HandleError(err)

	user := db.GetUserByGitHubLogin(*ct.User.Login)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
