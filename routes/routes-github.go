package routes

import (
	"net/http"
	"timetracker/github"
	"timetracker/helpers"
)

// Gets the github URL for logging into this app with GitHub
func getGitHubLoginUrl(w http.ResponseWriter, r *http.Request) {
	loginUrl, err := github.LoginUrl()
	helpers.HandleError(err)
	resp := map[string]interface{}{
		"data": loginUrl,
	}
	apiResponse(resp, w)
}
