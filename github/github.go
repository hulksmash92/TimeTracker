package github

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"

	"timetracker/helpers"

	"github.com/google/go-github/v35/github"
)

// Gets the URL this application uses to log user's into the app using their GitHub creds
func LoginUrl() (string, error) {
	clientId := getClientID()
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
	clientId := getClientID()
	clientSecret := getClientSecret()
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

	if !helpers.StrArrayContains(tokenResData["scope"], "user:email") {
		return token, errors.New("Invalid token scopes")
	}

	// If no errors have been found, return the parsed response to the caller
	token = tokenResData.Get("access_token")
	return token, nil
}

// Checks the access token given to a user
func CheckToken(token string) (*github.Authorization, error) {
	clientID := getClientID()
	client, ctx := getBasicAuthClient()
	auth, _, err := client.Authorizations.Check(ctx, clientID, token)

	if auth == nil || err != nil {
		return nil, err
	}

	return auth, nil
}

// Gets a new github client with basic auth configured
func getBasicAuthClient() (*github.Client, context.Context) {
	ctx := context.Background()
	bat := &github.BasicAuthTransport{
		Username: getClientID(),
		Password: getClientSecret(),
	}
	return github.NewClient(bat.Client()), ctx
}

// Gets a new authenticated github oauth client
func getOauthClient(token string) (*github.Client, context.Context) {
	oauthToken := &oauth2.Token{
		AccessToken: token,
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(oauthToken)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)
	return ghClient, ctx
}

// Gets the GitHub OAuth client ID
func getClientID() string {
	return os.Getenv("GITHUB_CLIENT_ID")
}

// Gets the client secret for the GitHub OAuth client
func getClientSecret() string {
	return os.Getenv("GITHUB_CLIENT_SECRET")
}
