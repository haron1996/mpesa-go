package funcs

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/haron1996/mpesa-go/utils"
)

type accessToken struct {
	AccessToken string
	ExpiresIn   string
}

func newAccessToken(aToken, expiresIn string) *accessToken {
	return &accessToken{AccessToken: aToken, ExpiresIn: expiresIn}
}

func RequestAccessToken(config utils.Config) (*accessToken, error) {

	consumerKey := config.ConsumerKey
	consumerSecret := config.ConsumerSecret

	// Create a basic authentication string by combining the client ID and client secret
	authString := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))

	// Prepare the request
	requestURL := safaricomBaseURL + oauthTokenEndpoint

	// init get reauest
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request with error: %v", err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Basic "+authString)

	// Make the request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request with error: %v", err)
	}

	defer resp.Body.Close()

	// Read the response body
	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode response with error: %v", err)
	}

	// check and get access token
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return nil, errors.New("could not find access_token in response")
	}

	// check ang get expiry
	expires_in, ok := result["expires_in"].(string)
	if !ok {
		return nil, errors.New("could not find expires_in in response")
	}

	token := newAccessToken(accessToken, expires_in)

	return token, nil
}
