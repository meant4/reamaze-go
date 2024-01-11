package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestClient_CreateConversation(t *testing.T) {
	correctConversation := &CreateConversationRequest{}
	correctConversation.Conversation.Category = "test"
	correctConversation.Conversation.Message.Body = "test"
	correctConversation.Conversation.User.Email = "test"

	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		req *CreateConversationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateConversationResponse
		wantErr bool
	}{
		{
			name: "Testing incorrect params",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Status:     "Bad Request",
						Body:       io.NopCloser(strings.NewReader(`{"error":"Bad Request"}`)),
					}
				}),
			}},
			args:    args{req: &CreateConversationRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct params incorrect request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Status:     "Bad Request",
						Body:       io.NopCloser(strings.NewReader(`{"error":"Bad Request"}`)),
					}
				}),
			}},
			args:    args{req: correctConversation},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct params correct request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "Status OK",
						Body:       io.NopCloser(strings.NewReader(`{"error":"Bad Request"}`)),
					}
				}),
			}},
			args:    args{req: correctConversation},
			want:    &CreateConversationResponse{},
			wantErr: false,
		},
		{
			name: "Testing correct params correct request invalid JSON response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "Status OK",
						Body:       io.NopCloser(strings.NewReader(`{"error":"Bad Request"`)),
					}
				}),
			}},
			args:    args{req: correctConversation},
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
			got, err := c.CreateConversation(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateConversation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateConversation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConversation(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetConversationResponse
		wantErr bool
	}{
		{
			name: "Testing empty slug parameter error",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "404 Status Not Found",
						Body:       io.NopCloser(strings.NewReader(`{"errors":"Not Found"}`)),
					}
				}),
			}},
			args:    args{slug: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing non existing slug request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{"errors":"Not Found"}`)),
					}
				}),
			}},
			args:    args{slug: "dummy"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 OK",
						Body:       io.NopCloser(strings.NewReader(`{"subject":"test"}`)),
					}
				}),
			}},
			args:    args{slug: "dummy"},
			want:    &GetConversationResponse{Subject: "test"},
			wantErr: false,
		},
		{
			name: "Testing incorrect response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 StatusOK",
						Body:       io.NopCloser(strings.NewReader(`{"subject":"test"`)),
					}
				}),
			}},
			args:    args{slug: "dummy"},
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
			got, err := c.GetConversation(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetConversation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetConversation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConversations(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetConversationsResponse
		wantErr bool
	}{
		{
			name: "Testing empty args",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{},
			want:    &GetConversationsResponse{},
			wantErr: false,
		},
		{
			name: "Testing incorrect response from reamaze",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing reamaze error response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotAcceptable,
						Status:     "406 Status Not Acceptable",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{},
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
			got, err := c.GetConversations(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetConversations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetConversations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateConversation(t *testing.T) {
	updateConversationReq := &UpdateConversationRequest{}
	updateConversationReq.Conversation.Status = ReamazeStatusArchived
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		slug string
		req  *UpdateConversationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetConversationResponse
		wantErr bool
	}{
		{
			name: "Testing no empty slug argument passed to UpdateConversation",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "", req: updateConversationReq},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct slug and empty UpdateConversationRequest",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "dummy", req: &UpdateConversationRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct slug and correct UpdateConversationRequest incorrect json response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args:    args{slug: "dummy", req: updateConversationReq},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct slug and correct UpdateConversationRequest correct json response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "dummy", req: updateConversationReq},
			want:    &GetConversationResponse{},
			wantErr: false,
		},
		{
			name: "Testing reamaze client error",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotAcceptable,
						Status:     "406 Status Not Acceptable",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "dummy", req: updateConversationReq},
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
			got, err := c.UpdateConversation(tt.args.slug, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateConversation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UpdateConversation() = %v, want %v", got, tt.want)
			}
		})
	}
}
