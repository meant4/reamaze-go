package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

// GetStaff will allow you to retrieve staff users for the Account https://www.reamaze.com/api/get_staff
func (c *Client) GetStaff(o ...StaffOption) (*GetStaffResponse, error) {
	var response *GetStaffResponse
	settings, _ := newStaffSettings(o)
	urlEndpoint := staffEndpoint + settings.GetQuery()

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

// CreateStaff will allow you to create new staff user. Please keep in mind that this may have an impact on your monthly subscription cost https://www.reamaze.com/api/post_staff
// No invite emails are sent to the created user's email. This is to prevent potential abuse.
// The created user will be asked to change their password the first time they log in.
func (c *Client) CreateStaff(req *CreateStaffRequest) (*CreateStaffResponse, error) {
	var response *CreateStaffResponse
	emptyReq := &CreateStaffRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateNote incorrect request, CreateNoteRequest is empty")
	}
	urlEndpoint := staffEndpoint
	data, _ := json.Marshal(req)
	resp, err := c.reamazeRequest(http.MethodPost, urlEndpoint, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
