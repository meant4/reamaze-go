package reamaze

import (
	"encoding/json"
	"net/http"
)

// GetSystems will allow you to retrieve the systems for the Brand
// https://www.reamaze.com/api/get_systems
func (c *Client) GetSystems() (*GetSystemsResponse, error) {
	var response *GetSystemsResponse
	urlEndpoint := systemsEndpoint
	resp, err := c.reamazeRequest(http.MethodGet, urlEndpoint, []byte{})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
