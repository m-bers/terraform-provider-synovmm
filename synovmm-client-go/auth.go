package synovmm

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
)

func (c *Client) SignIn() (*AuthResponse, error) {

	api := "SYNO.API.Auth"
	version := "3"
	method := "login"
	format := "sid"
	session := "dsm_info"

	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/webapi/auth.cgi?api=%s&version=%s&method=%s&account=%s&passwd=%s&format=%s&session=%s",
		c.HostURL,
		api,
		version,
		method,
		c.Auth.Username,
		c.Auth.Password,
		format,
		session),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) SignOut(authToken *string) error {

	api := "SYNO.API.Auth"
	version := "3"
	method := "logout"
	sid := c.SID

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/webapi/auth.cgi?api=%s&version=%s&method=%s&sid=%s",
		c.HostURL,
		api,
		version,
		method,
		sid),
		nil)

	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "{\"success\":true}" {
		return errors.New(string(body))
	}

	return nil
}
