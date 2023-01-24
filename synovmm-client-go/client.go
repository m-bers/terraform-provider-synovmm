package synovmm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "http://localhost:5000"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	SID        string
	Auth       AuthStruct
}

type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Data struct {
		Did string `json:"did"`
		Sid string `json:"sid"`
	} `json:"data"`
	Success bool `json:"success"`
}

func NewClient(host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},

		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.SID = ar.Data.Sid

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	if authToken != nil {
		q := req.URL.Query()
		q.Add("_sid", *authToken)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	return body, nil
}
