package github

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
func GetAccessToken(sessionCode string) (url.Values, error) {
	clientId := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	data := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {sessionCode},
	}
	tokenUrl := os.Getenv("GITHUB_URL_TOKEN")

	// Make a request to github to get the user's access token
	// Check for any errors in the request
	resp, err := http.PostForm(tokenUrl, data)
	if err != nil {
		return nil, err
	}

	// Parse the body to a byte[] and check for any errors
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Response is returned a query string format, so parse this to url.Values
	// and check for any errors in the parsing
	tokenRes := string(body[:])
	tokenResData, err := url.ParseQuery(tokenRes)
	if err != nil {
		return nil, err
	}

	// Check for any errors in the token response body
	tokenErr := tokenResData.Get("error")
	if tokenErr != "" {
		return nil, errors.New(tokenErr)
	}

	// If no errors have been found, return the parsed response to the caller
	return tokenResData, nil
}
