package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
)

// GetIncidents will allow you to retrieve the incidents for the Brand
// https://www.reamaze.com/api/get_incidents
func (c *Client) GetIncidents() (*GetIncidentsResponse, error) {
	var response *GetIncidentsResponse
	urlEndpoint := incidentsEndpoint
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

// GetIncidents will allow you to retrieve the incidents for the Brand
// https://www.reamaze.com/api/get_incident
func (c *Client) GetIncident(identifier string) (*GetIncidentResponse, error) {
	var response *GetIncidentResponse

	if len(identifier) == 0 {
		return nil, errors.New("GetIncident identifier cannot be empty, please provide identifier as argument")
	}
	urlEndpoint := incidentsEndpoint + "/" + url.PathEscape(identifier)

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

// UpdateIncident will allow you to retrieve the incidents for the Brand
// https://www.reamaze.com/api/put_incident
func (c *Client) UpdateIncident(identifier string, req *UpdateIncidentRequest) (*UpdateIncidentResponse, error) {
	var response *UpdateIncidentResponse
	emptyReq := &UpdateIncidentRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("UpdateIncident incorrect request, UpdateIncidentRequest is empty")
	}
	if len(identifier) == 0 {
		return nil, errors.New("UpdateIncident identifier cannot be empty, please provide identifier as argument")
	}

	urlEndpoint := incidentsEndpoint + "/" + url.PathEscape(identifier)
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

// CreateIncident  will allow you to create a new incident
// https://www.reamaze.com/api/post_incident
// The incidents_systems_attributes attribute is optional.
// The status attribute in updates_attributes can be one of the following: [investigating, identified, monitoring, resolved]
// The status attribute in incidents_systems_attributes can be one of the following: [operational, degraded_performance, partial_outage, major_outage, under_maintenance]
func (c *Client) CreateIncident(req *CreateIncidentRequest) (*CreateIncidentResponse, error) {
	var response *CreateIncidentResponse
	emptyReq := &CreateIncidentRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateIncident incorrect request, CreateIncidentRequest is empty")
	}

	urlEndpoint := incidentsEndpoint
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
