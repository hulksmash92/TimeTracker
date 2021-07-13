package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
	"timetracker/models"

	"github.com/gorilla/mux"
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

// Structure of the access token request body
type GHTokenReqBody struct {
	// Code returned in the redirect URL after a user
	// has logged in with GitHub
	SessionCode string `json:"sessionCode"`
}

// Gets logged in user's access token from GitHub's auth server using the
// session code returned when they signed in with GitHub.
// The user's details are then either retreived from the database
// using their github user id, or a new user is created if this is their
// first time logging into the application.
func getGitHubAccessToken(w http.ResponseWriter, r *http.Request) {
	// Use the built in ioutil from io/ioutil to
	// read the request body into a []byte
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleError(err)

	// Decode the JSON request body to our GHTokenReqBody
	// so we can use the session code
	var tokenReqBody GHTokenReqBody
	err = json.Unmarshal(body, &tokenReqBody)
	helpers.HandleError(err)

	// 1. Grab the access token from GitHub using the session code
	accessToken, err := github.GetAccessToken(tokenReqBody.SessionCode)
	helpers.HandleError(err)

	// 2. Call the check token method with our new access token
	//    to get the logged in users details
	checkTokenResult, err := github.CheckToken(accessToken)
	helpers.HandleError(err)

	// 3: Check if the user exists using their GitHub user id, and either:
	//   - Create a new user record if this is their first time logging in
	//   - Get the existing users details

	var user models.User
	if !db.GitHubUserExists(*checkTokenResult.User.Login) {
		user = db.CreateUser(*checkTokenResult.User)
	} else {
		user = db.GetUserByGitHubLogin(*checkTokenResult.User.Login)
	}

	// 4: Set a cookie containing the user's token
	//    that we can use for future request, only
	//    set the Secure attribute to true if not in
	//    development mode
	isDev := os.Getenv("HOSTING_ENV") == "Development"
	tokenCookieExpires := 30 * 24 * time.Hour
	tokenCookie := &http.Cookie{
		Name:     tokenCookieName,
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(tokenCookieExpires),
		MaxAge:   0,
		Secure:   !isDev,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Unparsed: []string{},
	}
	http.SetCookie(w, tokenCookie)
	w.WriteHeader(http.StatusOK)

	// 5: Return the users details to the caller
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Searches for github repos
func searchGitHubRepos(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")
	if searchQuery == "" {
		apiResponse(map[string]interface{}{}, w)
	}

	for _, c := range r.Cookies() {
		fmt.Println(c)
	}

	token, err := parseTokenFromCookie(r)
	helpers.HandleError(err)

	res, err := github.SearchForRepos(token, searchQuery)
	helpers.HandleError(err)
	apiResponse(map[string]interface{}{"data": res}, w)
}

// Gets items that relate to a github repo
func getGitHubRepoItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemType := vars["itemType"]

	if itemType == "branch" || itemType == "commit" {
		owner := vars["owner"]
		repo := vars["repo"]
		token, err := parseTokenFromCookie(r)
		helpers.HandleError(err)

		var res interface{}

		if itemType == "branch" {
			res, err = github.GetBranches(token, owner, repo)
			helpers.HandleError(err)
		} else {
			from, err := time.Parse(dtParamLayout, r.URL.Query().Get("from"))
			if err != nil {
				from = time.Now().AddDate(0, 0, -7)
			}
			to, err := time.Parse(dtParamLayout, r.URL.Query().Get("to"))
			if err != nil {
				to = time.Now()
			}

			res, err = github.GetCommits(token, owner, repo, from, to)
			helpers.HandleError(err)
		}

		apiResponse(map[string]interface{}{"data": res}, w)
	} else {
		helpers.HandleError(errors.New("Invalid itemType param"))
	}
}
