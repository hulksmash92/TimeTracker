package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"
	"timetracker/models"
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
		fmt.Printf("Github user %s does not exist in the db", *ct.User.Login)
		user = db.CreateUser(*ct.User)
		fmt.Printf("User created for %s in the db", *ct.User.Login)
	} else {
		user = db.GetUserByGitHubLogin(*ct.User.Login)
	}

	// 2: Set a cookie containing the user's token
	//    that we can use for future request
	isDev := os.Getenv("HOSTING_ENV") == "Development"
	expires := 30 * 24 * time.Hour
	cookie := &http.Cookie{
		Name:     "LoginData",
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
