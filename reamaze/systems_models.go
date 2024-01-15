package reamaze

import "time"

const systemsEndpoint string = "/api/v1/systems"

type GetSystemsResponse []struct {
	ID              string    `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	AccountID       int       `json:"account_id,omitempty"`
	BrandID         int       `json:"brand_id,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	Status          string    `json:"status,omitempty"`
	ActiveIncidents []struct {
		ID          string    `json:"id,omitempty"`
		Title       string    `json:"title,omitempty"`
		AccountID   int       `json:"account_id,omitempty"`
		BrandID     int       `json:"brand_id,omitempty"`
		CreatedAt   time.Time `json:"created_at,omitempty"`
		UpdatedAt   time.Time `json:"updated_at,omitempty"`
		Status      string    `json:"status,omitempty"`
		ExternalURL string    `json:"external_url,omitempty"`
		Updates     []struct {
			ID        string    `json:"id,omitempty"`
			Status    string    `json:"status,omitempty"`
			Message   string    `json:"message,omitempty"`
			CreatedAt time.Time `json:"created_at,omitempty"`
		} `json:"updates,omitempty"`
		IncidentsSystems []struct {
			ID       string `json:"id,omitempty"`
			SystemID string `json:"system_id,omitempty"`
			Status   string `json:"status,omitempty"`
			System   struct {
				ID    string `json:"id,omitempty"`
				Title string `json:"title,omitempty"`
			} `json:"system,omitempty"`
		} `json:"incidents_systems,omitempty"`
	} `json:"active_incidents,omitempty"`
}
