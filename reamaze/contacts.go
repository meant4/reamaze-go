package reamaze

import (
	"encoding/json"
	"errors"
	"fmt"
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

// GetContactIdentities fgt Contact's Identities  https://www.reamaze.com/api/get_identities
func (c *Client) GetContactIdentities() (string, error) {
	fmt.Print("GetContactIdentities Not implemented")
	return "", nil
}

// CreateContactIdentities creates Contact Identities  https://www.reamaze.com/api/post_identities
// Call to identities will allow you to attach an email, mobile number, twitter handle or instagram user to a contact.
func (c *Client) CreateContactIdentities() (string, error) {
	fmt.Print("CreateContactIdentities Not implemented")
	return "", nil
}
