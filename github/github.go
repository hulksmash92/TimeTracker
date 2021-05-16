package github

import (
	"errors"
	"os"
)

func LoginUrl() (string, error) {
	scopes := os.Getenv("GITHUB_SCOPES")
	if scopes == "" {
		scopes = "user:email"
	}
	clientId := os.Getenv("GITHUB_CLIENT_ID")
	loginUrl := os.Getenv("GITHUB_URL_LOGIN")

	if clientId == "" || loginUrl == "" {
		return "", errors.New("oops something went wrong")
	}

	// Build the final URL
	url := loginUrl + "?scopes=" + scopes + "&client_id=" + clientId

	return url, nil
}
