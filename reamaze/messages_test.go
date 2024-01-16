package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetMessages(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetMessagesResponse
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

			want:    &GetMessagesResponse{},
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

			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 UnprocessableEntity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
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
			got, err := c.GetMessages()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateMessage(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		slug string
		req  *CreateMessageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateMessageResponse
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
			args: args{slug: "dummy", req: &CreateMessageRequest{Message: struct {
				Body       string            "json:\"body\""
				Visibility ReamazeVisibility "json:\"visibility,omitempty\""
				OriginID   string            "json:\"origin_id,omitempty\""
				User       *struct {
					Name  string "json:\"name,omitempty\""
					Email string "json:\"email,omitempty\""
				} "json:\"user,omitempty\""
				SupressNotification bool     "json:\"suppress_notifications,omitempty\""
				SupressAutoresolve  bool     "json:\"suppress_autoresolve,omitempty\""
				Attachment          string   "json:\"attachment,omitempty\""
				Attachments         []string "json:\"attachments,omitempty\""
			}{
				Body: "dummy",
			}}},
			want:    &CreateMessageResponse{},
			wantErr: false,
		},
		{
			name: "Testing invalid JSON response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args: args{slug: "dummy", req: &CreateMessageRequest{Message: struct {
				Body       string            "json:\"body\""
				Visibility ReamazeVisibility "json:\"visibility,omitempty\""
				OriginID   string            "json:\"origin_id,omitempty\""
				User       *struct {
					Name  string "json:\"name,omitempty\""
					Email string "json:\"email,omitempty\""
				} "json:\"user,omitempty\""
				SupressNotification bool     "json:\"suppress_notifications,omitempty\""
				SupressAutoresolve  bool     "json:\"suppress_autoresolve,omitempty\""
				Attachment          string   "json:\"attachment,omitempty\""
				Attachments         []string "json:\"attachments,omitempty\""
			}{
				Body: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing invalid endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{slug: "dummy", req: &CreateMessageRequest{Message: struct {
				Body       string            "json:\"body\""
				Visibility ReamazeVisibility "json:\"visibility,omitempty\""
				OriginID   string            "json:\"origin_id,omitempty\""
				User       *struct {
					Name  string "json:\"name,omitempty\""
					Email string "json:\"email,omitempty\""
				} "json:\"user,omitempty\""
				SupressNotification bool     "json:\"suppress_notifications,omitempty\""
				SupressAutoresolve  bool     "json:\"suppress_autoresolve,omitempty\""
				Attachment          string   "json:\"attachment,omitempty\""
				Attachments         []string "json:\"attachments,omitempty\""
			}{
				Body: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty CreateMessageRequest",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "dummy", req: &CreateMessageRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty slug",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{slug: "", req: &CreateMessageRequest{Message: struct {
				Body       string            "json:\"body\""
				Visibility ReamazeVisibility "json:\"visibility,omitempty\""
				OriginID   string            "json:\"origin_id,omitempty\""
				User       *struct {
					Name  string "json:\"name,omitempty\""
					Email string "json:\"email,omitempty\""
				} "json:\"user,omitempty\""
				SupressNotification bool     "json:\"suppress_notifications,omitempty\""
				SupressAutoresolve  bool     "json:\"suppress_autoresolve,omitempty\""
				Attachment          string   "json:\"attachment,omitempty\""
				Attachments         []string "json:\"attachments,omitempty\""
			}{
				Body: "dummy",
			}}},
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
			got, err := c.CreateMessage(tt.args.slug, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
