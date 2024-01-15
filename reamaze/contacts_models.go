package reamaze

import (
	"regexp"
	"strings"
	"time"
)

// Contacts models
type ReamazePhoneNumber string
type ReamazeIdentifier string

const contactsEndpoint string = "/api/v1/contacts"
const (
	ReamazeIdentifierEmail     ReamazeIdentifier = "email"
	ReamazeIdentifierMobile    ReamazeIdentifier = "mobile"
	ReamazeIdentifierFacebook  ReamazeIdentifier = "facebook"
	ReamazeIdentifierTwitter   ReamazeIdentifier = "twitter"
	ReamazeIdentifierInstagram ReamazeIdentifier = "instagram"
)

func (w ReamazePhoneNumber) Validate() bool {
	phoneNumber := string(w)
	e164RegexString := "^\\+[1-9]?[0-9]{7,14}$"
	re := regexp.MustCompile(e164RegexString)
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	validPhone := re.Find([]byte(phoneNumber))
	return validPhone != nil
}

type GetContactResponse struct {
	Name         string      `json:"name"`
	Data         interface{} `json:"data"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Email        string      `json:"email"`
	Twitter      string      `json:"twitter"`
	Facebook     string      `json:"facebook"`
	Instagram    string      `json:"instagram"`
	Mobile       string      `json:"mobile"`
	FriendlyName string      `json:"friendly_name"`
	ID           string      `json:"id"`
	ID0          int         `json:"_id"`
	Notes        []Note      `json:"notes"`
}

type GetContactsResponse struct {
	PageSize   int `json:"page_size"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
	Contacts   []struct {
		Name         string      `json:"name"`
		Data         interface{} `json:"data"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		Email        string      `json:"email"`
		Twitter      string      `json:"twitter"`
		Facebook     string      `json:"facebook"`
		Instagram    string      `json:"instagram"`
		Mobile       string      `json:"mobile"`
		FriendlyName string      `json:"friendly_name"`
		ID           int         `json:"_id"`
		Notes        []Note      `json:"notes"`
		ID0          string      `json:"id,omitempty"`
	} `json:"contacts"`
}

type CreateContactRequest struct {
	Contact struct {
		Name              string             `json:"name"`
		Email             string             `json:"email"`
		Mobile            ReamazePhoneNumber `json:"mobile"`
		FriendlyName      string             `json:"friendly_name"`
		ID                string             `json:"id"`
		ExternalAvatarURL string             `json:"external_avatar_url"`
		Notes             []string           `json:"notes"`
		Data              interface{}        `json:"data"`
	} `json:"contact"`
}

type UpdateContactRequest struct {
	Contact struct {
		Name              string      `json:"name"`
		FriendlyName      string      `json:"friendly_name"`
		ExternalAvatarURL string      `json:"external_avatar_url"`
		Notes             []string    `json:"notes"`
		Data              interface{} `json:"data"`
	} `json:"contact"`
}

type GetContactIdentitiesResponse struct {
	Identities []struct {
		Type       ReamazeIdentifier `json:"type"`
		Identifier string            `json:"identifier"`
	} `json:"identities"`
}
type CreateContactIdentitiesRequest struct {
	Identity struct {
		Type       ReamazeIdentifier `json:"type"`
		Identifier string            `json:"identifier"`
	} `json:"identity"`
}
