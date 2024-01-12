package reamaze

import "time"

// Contacts models

type GetContactResponse struct {
	Name string `json:"name"`
	Data struct {
		LastName       string `json:"last_name"`
		FirstName      string `json:"first_name"`
		JobApplication string `json:"job_application"`
	} `json:"data"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Twitter      string    `json:"twitter"`
	Facebook     string    `json:"facebook"`
	Instagram    string    `json:"instagram"`
	Mobile       string    `json:"mobile"`
	FriendlyName string    `json:"friendly_name"`
	ID           string    `json:"id"`
	ID0          int       `json:"_id"`
	Notes        []any     `json:"notes"`
}
type CreateContactRequest struct {
	Contact struct {
		Name              string      `json:"name"`
		Email             string      `json:"email"`
		ID                string      `json:"id"`
		ExternalAvatarURL string      `json:"external_avatar_url"`
		Notes             []string    `json:"notes"`
		Data              interface{} `json:"data"`
	} `json:"contact"`
}

type GetContactsResponse struct {
	PageSize   any `json:"page_size"`
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
