package reamaze

import (
	"encoding/json"
	"net/http"
)

// GetStaff will allow you to retrieve staff users for the Account https://www.reamaze.com/api/get_staff
func (c *Client) GetStaff() (*GetStaffResponse, error) {
	var response *GetStaffResponse

	urlEndpoint := staffEndpoint
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
