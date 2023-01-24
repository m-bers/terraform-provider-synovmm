package synovmm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetHosts(authToken *string) (*Host, error) {

	api := "SYNO.Virtualization.API.Host"
	method := "list"
	version := "1"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/webapi/entry.cgi?_sid=%s&api=%s&method=%s&version=%s",
		c.HostURL,
		c.SID,
		api,
		method,
		version),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	h := Host{}
	err = json.Unmarshal(body, &h)

	if err != nil {
		return nil, err
	}

	return &h, nil
}
