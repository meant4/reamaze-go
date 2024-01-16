package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestClient_GetIncidents(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetIncidentsResponse
		wantErr bool
	}{
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
			want: &GetIncidentsResponse{struct {
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
			}{}},
			wantErr: false,
		},
		{
			name: "Testing incorrect JSON response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`[{}`)),
					}
				}),
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`[{}]`)),
					}
				}),
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetIncidents()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetIncidents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetIncidents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIncident(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetIncidentResponse
		wantErr bool
	}{
		{
			name: "Testing empty identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: ""},
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
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy"},
			want:    &GetIncidentResponse{},
			wantErr: false,
		},
		{
			name: "Testing incorrect JSON response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`[{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`[{}]`)),
					}
				}),
			}},
			args:    args{identifier: "dummy"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetIncident(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetIncident() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetIncident() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateIncident(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier string
		req        *UpdateIncidentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateIncidentResponse
		wantErr bool
	}{
		{
			name: "Testing correct request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy", req: &UpdateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    &UpdateIncidentResponse{},
			wantErr: false,
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
			args: args{identifier: "dummy", req: &UpdateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy", req: &UpdateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "", req: &UpdateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty UpdateIncidentRequest request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy", req: &UpdateIncidentRequest{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.UpdateIncident(tt.args.identifier, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateIncident() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UpdateIncident() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateIncident(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		req *CreateIncidentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateIncidentResponse
		wantErr bool
	}{
		{
			name: "Testing correct request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    &CreateIncidentResponse{},
			wantErr: false,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
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
			args: args{req: &CreateIncidentRequest{Incident: struct {
				Title             string "json:\"title,omitempty\""
				UpdatesAttributes []struct {
					Status  ReamazeIncidentUpdateStatus "json:\"status,omitempty\""
					Message string                      "json:\"message,omitempty\""
				} "json:\"updates_attributes,omitempty\""
				IncidentsSystemsAttributes []struct {
					ID       string                      "json:\"id,omitempty\""
					SystemID string                      "json:\"system_id,omitempty\""
					Status   ReamazeIncidentSystemStatus "json:\"status,omitempty\""
				} "json:\"incidents_systems_attributes,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty CreateIncidentRequest request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{req: &CreateIncidentRequest{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.CreateIncident(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateIncident() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateIncident() = %v, want %v", got, tt.want)
			}
		})
	}
}
