package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// RecaptchaResponse :
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func main() {
	token := "set token from client here"
	const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {"set your secret key here"}, "remoteip": {"127.0.0.1"}, "response": {token}})
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r := RecaptchaResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}
	fmt.Printf("%v", r)
}
