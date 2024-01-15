package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

// GetNotes gets all the Notes for provided identifier
func (c *Client) GetNotes(identifier string) (*GetNotesResponse, error) {
	var response *GetNotesResponse
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("GetNotes identifier cannot be empty, please provide identifier as argument")
	}
	urlEndpoint := contactsEndpoint + "/" + url.PathEscape(identifier) + "/notes"
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

// CreateNote will allow you to attach an note to a contact
func (c *Client) CreateNote(identifier string, req *CreateNoteRequest) (*CreateNoteResponse, error) {
	var response *CreateNoteResponse
	emptyReq := &CreateContactRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateNote incorrect request, CreateNoteRequest is empty")
	}
	if len(identifier) == 0 {
		return nil, errors.New("CreateNote identifier cannot be empty, please provide identifier as argument")
	}
	urlEndpoint := contactsEndpoint + "/" + url.PathEscape(identifier) + "/notes"
	// Documentations says it's optional but during the test we are having empty date 0001-01-01T00:00:00.000Z so we pass current time
	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now()
	}
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

// UpdateNote will allow you to update a note with the given id
func (c *Client) UpdateNote(identifier string, noteID string, req *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	var response *UpdateNoteResponse
	emptyReq := &CreateContactRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("UpdateNote incorrect request, UpdateNoteRequest is empty")
	}
	if len(identifier) == 0 {
		return nil, errors.New("UpdateNote identifier cannot be empty, please provide identifier as argument")
	}
	if len(noteID) == 0 {
		return nil, errors.New("UpdateNote noteID cannot be empty, please provide noteID as argument")
	}
	urlEndpoint := contactsEndpoint + "/" + url.PathEscape(identifier) + "/notes/" + url.PathEscape(noteID)
	// Documentations says it's optional but during the test we are having empty date 0001-01-01T00:00:00.000Z so we pass current time
	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now()
	}
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

// DeleteNote will delete a note with the given id from the contact
func (c *Client) DeleteNote(identifier string, noteID string) (*DeleteNoteResponse, error) {
	var response *DeleteNoteResponse

	if len(identifier) == 0 {
		return nil, errors.New("DeleteNote identifier cannot be empty, please provide identifier as argument")
	}
	if len(noteID) == 0 {
		return nil, errors.New("DeleteNote noteID cannot be empty, please provide noteID as argument")
	}
	urlEndpoint := contactsEndpoint + "/" + url.PathEscape(identifier) + "/notes/" + url.PathEscape(noteID)
	resp, err := c.reamazeRequest(http.MethodDelete, urlEndpoint, []byte{})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
