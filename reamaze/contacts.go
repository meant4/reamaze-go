package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
)

// GetContacts retrieves all contacts for the Account
// Note that unlike other resources, contacts are tied to the account, not the individual brand.
func (c *Client) GetContacts() (*GetContactsResponse, error) {
	var response *GetContactsResponse
	urlEndpoint := contactsEndpoint
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

// GetContact getting Contact for identifier
func (c *Client) GetContact(identifier string) (*GetContactResponse, error) {
	var response *GetContactResponse
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("GetContact identifier cannot be empty, please provide identifier as argument")
	}
	urlEndpoint := contactsEndpoint + "/" + url.PathEscape(identifier)

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

// CreateContact creates new Contact  https://www.reamaze.com/api/post_contacts
func (c *Client) CreateContact(req *CreateContactRequest) (*GetContactResponse, error) {
	var response *GetContactResponse
	emptyReq := &CreateContactRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateContact incorrect request, CreateContactRequest is empty")
	}
	// we are checking if mobile is set and is valid
	if !req.Contact.Mobile.Validate() && len(req.Contact.Email) == 0 {
		return nil, errors.New("CreateContact provided phone number is not correct and email is not set, please provide emial or vaild E.164 phone number")
	}
	if !req.Contact.Mobile.Validate() && len(req.Contact.Mobile) > 0 && len(req.Contact.Email) > 0 && len(req.Contact.Name) == 0 {
		return nil, errors.New("CreateContact if phone number is incorrect and email is set, the Name has to be set")
	}
	urlEndpoint := contactsEndpoint
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

// UpdateContact update existing Contact  https://www.reamaze.com/api/put_contacts
// This re:amaze endpoint can only change name, friendly_name, external_avatar_url and data attributes
func (c *Client) UpdateContact(identifier string, req *UpdateContactRequest, identifierType ...ReamazeIdentifier) (*GetContactResponse, error) {
	// We set default identifier type to email
	idType := string(ReamazeIdentifierEmail)
	// If we have explicitly defined identifier type than we set idType to this identifier type
	if len(identifierType) > 0 {
		idType = string(identifierType[0])
	}

	var response *GetContactResponse
	emptyReq := &UpdateContactRequest{}
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("UpdateContact identifier cannot be empty, please provide identifier as argument")
	}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("UpdateContact incorrect request, UpdateContactRequest is empty")
	}

	urlEndpoint := contactsEndpoint + "/" + url.QueryEscape(identifier) + "?identifier_type=" + idType
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

// GetContactIdentities gets Contact's Identities  https://www.reamaze.com/api/get_identities
// Identities types are one of 'email', 'twitter','facebook', 'instagram', 'igsid' (Instagram-scoped ID), or 'mobile'.
func (c *Client) GetContactIdentities(identifier string) (*GetContactIdentitiesResponse, error) {
	var response *GetContactIdentitiesResponse
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("GetContactIdentities identifier cannot be empty, please provide identifier as argument")
	}
	urlEndpoint := contactsEndpoint + "/" + url.PathEscape(identifier) + "/identities"

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

// CreateContactIdentities creates Contact Identities  https://www.reamaze.com/api/post_identities
// Call to identities will allow you to attach an email, mobile number, twitter handle or instagram user to a contact.
func (c *Client) CreateContactIdentities(identifier string, req *CreateContactIdentitiesRequest, identifierType ...ReamazeIdentifier) (*GetContactIdentitiesResponse, error) {
	// We set default identifier type to email
	idType := string(ReamazeIdentifierEmail)
	// If we have explicitly defined identifier type than we set idType to this identifier type
	if len(identifierType) > 0 {
		idType = string(identifierType[0])
	}
	var response *GetContactIdentitiesResponse

	emptyReq := &CreateContactIdentitiesRequest{}
	// checking if identifier is set
	if len(identifier) == 0 {
		return nil, errors.New("CreateContactIdentities identifier cannot be empty, please provide identifier as argument")
	}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateContactIdentities incorrect request, CreateContactIdentities is empty")
	}
	// NOTE there is information in the documentation that facebook id cannot be created so we check if the type in the request is not Facebook identifier
	if req.Identity.Type == ReamazeIdentifierFacebook {
		return nil, errors.New("CreateContactIdentities cannot create Facebook identites")
	}
	// setting up urlEndpoint
	urlEndpoint := contactsEndpoint + "/" + url.QueryEscape(identifier) + "/identities?identifier_type=" + idType
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
