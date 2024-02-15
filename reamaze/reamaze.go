package reamaze

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/mail"
)

// Reamaze Client
type Client struct {
	baseURL    string
	auth       string
	httpClient *http.Client
}

// NewClient creates a new re:amaze client that uses the given workspace with user and pass.
func NewClient(email, apiToken, brand string) (*Client, error) {
	// Checking email
	if len(email) == 0 {
		return nil, errors.New("email address cannot be empty")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, errors.New("incorrect email address")
	}
	// checking if apiToken isn't empty
	if len(apiToken) == 0 {
		return nil, errors.New("apiToken cannot be empty")
	}
	// checking if brand isn't empty
	if len(brand) == 0 {
		return nil, errors.New("brand cannot be empty")
	}
	// encoding email:apiToken using base64
	sEnc := base64.StdEncoding.EncodeToString([]byte(email + ":" + apiToken))
	return &Client{
		baseURL:    "https://" + brand + ".reamaze.io",
		auth:       sEnc,
		httpClient: &http.Client{},
	}, nil
}

// reamazeRequset is a wrapper on http client to authenticate and set proper headers
func (c *Client) reamazeRequest(method string, endpoint string, payload []byte) ([]byte, error) {
	// Setting up request
	req, err := http.NewRequest(method, c.baseURL+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	// Setting HTTP headers
	req.Header.Add("Authorization", "Basic "+c.auth)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	// Executing request to reamaze
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// Checking if we have response status code within acceptable numbers 200-299
	if res.StatusCode >= 300 || res.StatusCode < 200 {
		return nil, errors.New(res.Status)
	}
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bodyData, nil
}
