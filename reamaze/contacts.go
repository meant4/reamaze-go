package reamaze

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// GetContacts retrieves all contacts for the Account
// Note that unlike other resources, contacts are tied to the account, not the individual brand.
func (c *Client) GetContacts() (*GetContactsResponse, error) {
	var response *GetContactsResponse
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
func (c *Client) CreateContact() (string, error) {
	fmt.Print("CreateContact Not implemented")
	return "", nil
}

// UpdateContact update existing Contact  https://www.reamaze.com/api/put_contacts
// This re:amaze endpoint can only change name, friendly_name, external_avatar_url and data attributes
func (c *Client) UpdateContact() (string, error) {
	fmt.Print("UpdateContact Not implemented")
	return "", nil
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
