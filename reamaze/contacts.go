package reamaze

import "fmt"

// Getting all Contacts for brand https://www.reamaze.com/api/get_contacts
func (c *Client) GetContacts() (string, error) {
	fmt.Print("GetContacts Not implemented")
	return "", nil
}

// Create new Contact for brand https://www.reamaze.com/api/post_contacts
func (c *Client) CreateContact() (string, error) {
	fmt.Print("CreateContact Not implemented")
	return "", nil
}

// Create new Contact for brand https://www.reamaze.com/api/put_contacts
// This re:amaze endpoint can only change name, friendly_name, external_avatar_url and data attributes
func (c *Client) UpdateContact() (string, error) {
	fmt.Print("UpdateContact Not implemented")
	return "", nil
}

// Get Contact Identities for brand https://www.reamaze.com/api/get_identities
func (c *Client) GetContactIdentities() (string, error) {
	fmt.Print("GetContactIdentities Not implemented")
	return "", nil
}

// Create Contact Identities for brand https://www.reamaze.com/api/post_identities
// Call to identities will allow you to attach an email, mobile number, twitter handle or instagram user to a contact.
func (c *Client) CreateContactIdentities() (string, error) {
	fmt.Print("CreateContactIdentities Not implemented")
	return "", nil
}
