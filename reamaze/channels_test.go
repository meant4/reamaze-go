package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetChannels(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetChannelsResponse
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
			want:    &GetChannelsResponse{},
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
						Status:     "422 Unprocessable Entity",
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
			got, err := c.GetChannels()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetChannel(t *testing.T) {
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
		want    *GetChannelResponse
		wantErr bool
	}{
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
			args:    args{slug: "dummy"},
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
			args:    args{slug: "dummy"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing missing slug argument",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "dummy"},
			want:    &GetChannelResponse{},
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
			got, err := c.GetChannel(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeChannelVisibility_Name(t *testing.T) {
	tests := []struct {
		name string
		w    ReamazeChannelVisibility
		want string
	}{
		{
			name: "Case 0",
			w:    ReamazeChannelVisibilityPrivate,
			want: "ReamazeChannelVisibilityPrivate",
		},
		{
			name: "Case 1",
			w:    ReamazeChannelVisibilityPublic,
			want: "ReamazeChannelVisibilityPublic",
		},
		{
			name: "Case 1",
			w:    ReamazeChannelVisibility(10),
			want: "Undefined",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Name(); got != tt.want {
				t.Errorf("ReamazeChannelVisibility.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeChannelType_Name(t *testing.T) {
	tests := []struct {
		name string
		w    ReamazeChannelType
		want string
	}{
		{
			name: "Case 1",
			w:    ReamazeChannelEmail,
			want: "ReamazeChannelEmail",
		},
		{
			name: "Case 2",
			w:    ReamazeChannelTwitter,
			want: "ReamazeChannelTwitter",
		},
		{
			name: "Case 3",
			w:    ReamazeChannelFacebook,
			want: "ReamazeChannelFacebook",
		},
		{
			name: "Case 6",
			w:    ReamazeChannelChat,
			want: "ReamazeChannelChat",
		},
		{
			name: "Case 8",
			w:    ReamazeChannelInstagram,
			want: "ReamazeChannelInstagram",
		},
		{
			name: "Case 9",
			w:    ReamazeChannelSMS,
			want: "ReamazeChannelSMS",
		},
		{
			name: "Case 10",
			w:    ReamazeChannelVoice,
			want: "ReamazeChannelVoice",
		},
		{
			name: "Case 12",
			w:    ReamazeChannelFacebookMessanger,
			want: "ReamazeChannelFacebookMessanger",
		},
		{
			name: "Case 13",
			w:    ReamazeChannelFacebookLead,
			want: "ReamazeChannelFacebookLead",
		},
		{
			name: "Case 14",
			w:    ReamazeChannelInstagramAd,
			want: "ReamazeChannelInstagramAd",
		},
		{
			name: "Case 15",
			w:    ReamazeChannelWhatsApp,
			want: "ReamazeChannelWhatsApp",
		},
		{
			name: "Case 16",
			w:    ReamazeChannelInstagramDM,
			want: "ReamazeChannelInstagramDM",
		},
		{
			name: "Case 13",
			w:    ReamazeChannelType(100),
			want: "Undefined",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Name(); got != tt.want {
				t.Errorf("ReamazeChannelType.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
