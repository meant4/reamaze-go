package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestClient_GetSystems(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetSystemsResponse
		wantErr bool
	}{
		{
			name: "Testing incorrect endpopint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Improcessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect JSON response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`[{}]`)),
					}
				}),
			}},
			want: &GetSystemsResponse{struct {
				ID              string    "json:\"id,omitempty\""
				Title           string    "json:\"title,omitempty\""
				AccountID       int       "json:\"account_id,omitempty\""
				BrandID         int       "json:\"brand_id,omitempty\""
				CreatedAt       time.Time "json:\"created_at,omitempty\""
				UpdatedAt       time.Time "json:\"updated_at,omitempty\""
				Status          string    "json:\"status,omitempty\""
				ActiveIncidents []struct {
					ID          string    "json:\"id,omitempty\""
					Title       string    "json:\"title,omitempty\""
					AccountID   int       "json:\"account_id,omitempty\""
					BrandID     int       "json:\"brand_id,omitempty\""
					CreatedAt   time.Time "json:\"created_at,omitempty\""
					UpdatedAt   time.Time "json:\"updated_at,omitempty\""
					Status      string    "json:\"status,omitempty\""
					ExternalURL string    "json:\"external_url,omitempty\""
					Updates     []struct {
						ID        string    "json:\"id,omitempty\""
						Status    string    "json:\"status,omitempty\""
						Message   string    "json:\"message,omitempty\""
						CreatedAt time.Time "json:\"created_at,omitempty\""
					} "json:\"updates,omitempty\""
					IncidentsSystems []struct {
						ID       string "json:\"id,omitempty\""
						SystemID string "json:\"system_id,omitempty\""
						Status   string "json:\"status,omitempty\""
						System   struct {
							ID    string "json:\"id,omitempty\""
							Title string "json:\"title,omitempty\""
						} "json:\"system,omitempty\""
					} "json:\"incidents_systems,omitempty\""
				} "json:\"active_incidents,omitempty\""
			}{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetSystems()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetSystems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetSystems() = %v, want %v", got, tt.want)
			}
		})
	}
}
