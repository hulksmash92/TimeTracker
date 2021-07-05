package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
	"timetracker/models"

	"github.com/gorilla/mux"
)

// Defines the structure of the access token request body
type GHTokenReqBody struct {
	SessionCode string `json:sessionCode`
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

	// Grab the access token
	token, err := github.GetAccessToken(fmtBody.SessionCode)
	helpers.HandleError(err)

	// call the check token method to get our logged in users details
	ct, err := github.CheckToken(token)
	helpers.HandleError(err)

	// 1: Create a new user if its this user's
	//    first time logging into our application
	//    or get the existing users details

	var user models.User
	if !db.GitHubUserExists(*ct.User.Login) {
		user = db.CreateUser(*ct.User)
	} else {
		user = db.GetUserByGitHubLogin(*ct.User.Login)
	}

	// 2: Set a cookie containing the user's token
	//    that we can use for future request
	cookieName := "LoginData"
	isDev := os.Getenv("HOSTING_ENV") == "Development"
	expires := 30 * 24 * time.Hour
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(expires),
		MaxAge:   0,
		Secure:   !isDev,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Unparsed: []string{},
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)

	// 3: Return the users details for and their settings
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
