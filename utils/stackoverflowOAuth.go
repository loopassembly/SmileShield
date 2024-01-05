package utils

import (
	"bytes"
	"cypher-server/initializers"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type StackOverflowOauthToken struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
	Expires_in    int    `json:"expires_in"`
}

func GetStackOverflowOauthToken(code string) (*StackOverflowOauthToken, error) {
	const rootURL = "https://stackoverflow.com/oauth/access_token/json"

	config, _ := initializers.LoadConfig(".")

	values := url.Values{}
	values.Add("client_id", config.StackoverflowClientID)
	values.Add("client_secret", config.StackOverflowClientSecret)
	values.Add("code", code)
	values.Add("redirect_uri", config.StackOverflowOAuthRedirectUrl)
	values.Add("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", rootURL, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var tokenBody StackOverflowOauthToken
	if err := json.NewDecoder(res.Body).Decode(&tokenBody); err != nil {
		return nil, err
	}

	return &tokenBody, nil
}



type StackOverflowUserResult struct {
    DisplayName  string `json:"display_name"` //name
    ProfileImage string `json:"profile_image"` //profile_name
    Email        string `json:"email"` // email
    ProfileURL   string `json:"profile_url"`// profile uri
}

func GetStackOverflowUser(access_token string) (*StackOverflowUserResult, error) {
	rootURL := "https://api.stackexchange.com/2.3/me"
	params := map[string]string{
		"site":           "stackoverflow",
		"access_token":   access_token,
		"key":            "your_stackoverflow_api_key", // replace with your Stack Overflow API key
	}

	req, err := http.NewRequest("GET", rootURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var userBody StackOverflowUserResult
	if err := json.NewDecoder(res.Body).Decode(&userBody); err != nil {
		return nil, err
	}

	return &userBody, nil
}
