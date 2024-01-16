package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

// GetResponseTemplates will allow you to retrieve Response Templates for the Brand. This will also return personal Response Templates depending on the user role https://www.reamaze.com/api/get_response_templates
// TODO add the q and page handlers
// q with any string will search over response templates by keywords.
// page with any number will allow you to paginate through results
func (c *Client) GetResponseTemplates() (*GetResponseTemplatesResponse, error) {
	var response *GetResponseTemplatesResponse
	urlEndpoint := responseTemplatesEndpoint

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

// GetResponseTemplate will allow you to retrieve a specific Response Template https://www.reamaze.com/api/get_response_template
func (c *Client) GetResponseTemplate(identifier string) (*GetResponseTemplateResponse, error) {
	var response *GetResponseTemplateResponse
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("GetResponseTemplate identifier cannot be empty, please provide identifier as argument")
	}
	urlEndpoint := responseTemplatesEndpoint + "/" + identifier
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

// CreateResponseTemplate will allow you to create a new Response Template https://www.reamaze.com/api/post_response_template
func (c *Client) CreateResponseTemplate(req *CreateResponseTemplateRequest) (*CreateResponseTemplateResponse, error) {
	var response *CreateResponseTemplateResponse
	emptyReq := &CreateResponseTemplateRequest{}

	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateResponseTemplate incorrect request, CreateResponseTemplateRequest is empty")
	}
	urlEndpoint := responseTemplatesEndpoint

	// preparing request
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

// UpdateResponseTemplate will allow you to update the response template https://www.reamaze.com/api/put_response_template
func (c *Client) UpdateResponseTemplate(identifier string, req *UpdateResponseTemplateRequest) (*UpdateResponseTemplateResponse, error) {
	var response *UpdateResponseTemplateResponse
	emptyReq := &UpdateResponseTemplateRequest{}

	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("UpdateResponseTemplate incorrect request, UpdateResponseTemplateRequest is empty")
	}
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("UpdateResponseTemplate identifier cannot be empty, please provide identifier as argument")
	}

	urlEndpoint := responseTemplatesEndpoint + "/" + identifier
	// preparing request
	data, _ := json.Marshal(req)

	resp, err := c.reamazeRequest(http.MethodPut, urlEndpoint, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
