package github

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"timetracker/helpers"
)

// Gets the URL this application uses to log user's into the app using their GitHub creds
func LoginUrl() (string, error) {
	clientId := os.Getenv("GITHUB_CLIENT_ID")
	scopes := os.Getenv("GITHUB_SCOPES")
	if scopes == "" {
		scopes = "user:email"
	}
	loginUrl := os.Getenv("GITHUB_URL_LOGIN")

	if clientId == "" || loginUrl == "" {
		return "", errors.New("oops something went wrong")
	}

	// Build the final URL
	url := loginUrl + "?scopes=" + scopes + "&client_id=" + clientId

	return url, nil
}

// Gets the access token for github
func GetAccessToken(sessionCode string) (string, error) {
	clientId := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	data := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {sessionCode},
	}
	tokenUrl := os.Getenv("GITHUB_URL_TOKEN")

	var token string

	// Make a request to github to get the user's access token
	// Check for any errors in the request
	resp, err := http.PostForm(tokenUrl, data)
	if err != nil {
		return token, err
	}

	// Parse the body to a byte[] and check for any errors
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	// Response is returned a query string format, so parse this to url.Values
	// and check for any errors in the parsing
	tokenRes := string(body[:])
	tokenResData, err := url.ParseQuery(tokenRes)
	if err != nil {
		return token, err
	}

	// Check for any errors in the token response body
	tokenErr := tokenResData.Get("error")
	if tokenErr != "" {
		return token, errors.New(tokenErr)
	}

	scopes := tokenResData["scope"]

	if !helpers.StrArrayContains(scopes, "user:read") {
		return token, errors.New("Invalid token scopes")
	}

	// If no errors have been found, return the parsed response to the caller
	token = tokenResData.Get("access_token")
	return token, nil
}
