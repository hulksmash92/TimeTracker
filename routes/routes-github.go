package routes

import (
	"encoding/json"
	"net/http"
	"timetracker/github"
	"timetracker/helpers"
)

type GHTokenReqBody struct {
	sessionCode string
}

// Gets the github URL for logging into this app with GitHub
func getGitHubLoginUrl(w http.ResponseWriter, r *http.Request) {
	loginUrl, err := github.LoginUrl()
	helpers.HandleError(err)
	resp := map[string]interface{}{
		"data": loginUrl,
	}
	apiResponse(resp, w)
}

// Gets the users access token
func getGitHubAccessToken(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var fmtBody GHTokenReqBody
	err := json.Unmarshal(body, &fmtBody)
	helpers.HandleError(err)

	token, err := github.GetAccessToken(fmtBody.sessionCode)
	helpers.HandleError(err)

	// Return the response as is for now
	// TODO 1: Set a cookie containing the user's token
	// TODO 2: Create a new user if its this user's first time logging into our application
	// TODO 3: Return the users details for and their settings
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
