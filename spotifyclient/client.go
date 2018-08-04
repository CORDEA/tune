package spotifyclient

import (
	"encoding/json"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"io/ioutil"
)

const CredentialFileName = "credential.txt"
const redirectURL = "http://localhost:8080/callback"

var (
	Authenticator = spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadCurrentlyPlaying)
)

func NewClient() (*spotify.Client, error) {
	file, err := ioutil.ReadFile(CredentialFileName)
	if err != nil {
		return nil, err
	}
	token := oauth2.Token{}
	if err = json.Unmarshal(file, &token); err != nil {
		return nil, err
	}
	client := Authenticator.NewClient(&token)
	return &client, nil
}
