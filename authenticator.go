package main

import (
	"encoding/json"
	"fmt"
	"github.com/zmb3/spotify"
	"io/ioutil"
	"net/http"
	"os"
)

const redirectURL = "http://localhost:8080/callback"
const credentialFileName = "credential.txt"

var (
	auth  = spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadCurrentlyPlaying)
	state = "gpzn7t"
)

func handleCallback(writer http.ResponseWriter, request *http.Request) {
	token, err := auth.Token(state, request)
	if err != nil {
		http.Error(writer, "Failed to get token.", http.StatusForbidden)
		return
	}
	if s := request.FormValue("state"); s != state {
		http.Error(writer, "Failed to get state or did not agree with sent.", http.StatusForbidden)
		return
	}
	rawJson, err := json.Marshal(token)
	if err != nil {
		http.Error(writer, "Failed to build a json from struct.", http.StatusInternalServerError)
		return
	}
	ioutil.WriteFile(credentialFileName, rawJson, os.ModePerm)

}

func main() {
	url := auth.AuthURL(state)
	fmt.Println("url ", url)

	http.HandleFunc("/callback", handleCallback)
	http.ListenAndServe(":8080", nil)
}
