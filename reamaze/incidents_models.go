package reamaze

import "time"

const incidentsEndpoint string = "/api/v1/incidents"

type ReamazeIncidentSystemStatus string
type ReamazeIncidentUpdateStatus string

const (
	ReamazeIncidentUpdateSatusInvestigating ReamazeIncidentUpdateStatus = "investigating"
	ReamazeIncidentUpdateStatusIdentified   ReamazeIncidentUpdateStatus = "identified"
	ReamazeIncidentUpdateStatusMonitoring   ReamazeIncidentUpdateStatus = "monitoring"
	ReamazeIncidentUpdateStatusResolved     ReamazeIncidentUpdateStatus = "resolved"
)

const (
	ReamazeIncidentSystemStatusOperational         ReamazeIncidentSystemStatus = "operational"
	ReamazeIncidentSystemStatusDegradedPerformance ReamazeIncidentSystemStatus = "degraded_performance"
	ReamazeIncidentSystemStatusPartialOutage       ReamazeIncidentSystemStatus = "partial_outage"
	ReamazeIncidentSystemStatusMajorOutage         ReamazeIncidentSystemStatus = "major_outage"
	ReamazeIncidentSystemStatusUnderMaintenance    ReamazeIncidentSystemStatus = "under_maintenance"
)

type GetIncidentsResponse []struct {
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
}

type GetIncidentResponse struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Status    string `json:"status,omitempty"`
	Updates   []struct {
		ID      string `json:"id,omitempty"`
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
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
}

type UpdateIncidentRequest struct {
	Incident struct {
		Title             string `json:"title,omitempty"`
		UpdatesAttributes []struct {
			Status  ReamazeIncidentUpdateStatus `json:"status,omitempty"`
			Message string                      `json:"message,omitempty"`
		} `json:"updates_attributes,omitempty"`
		IncidentsSystemsAttributes []struct {
			ID       string                      `json:"id,omitempty"`
			SystemID string                      `json:"system_id,omitempty"`
			Status   ReamazeIncidentSystemStatus `json:"status,omitempty"`
		} `json:"incidents_systems_attributes,omitempty"`
	} `json:"incident,omitempty"`
}

type UpdateIncidentResponse GetIncidentResponse
type CreateIncidentRequest UpdateIncidentRequest
type CreateIncidentResponse GetIncidentResponse
