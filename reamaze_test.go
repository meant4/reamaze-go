package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestNewClient(t *testing.T) {
	type args struct {
		email    string
		apiToken string
		brand    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name:    "Testing invalid email address",
			args:    args{email: "invalid", apiToken: "token", brand: "brand"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Testing valid baseURL and auth creation",
			args:    args{email: "test@example.com", apiToken: "something", brand: "brand"},
			want:    &Client{baseURL: "https://brand.reamaze.io", auth: "dGVzdEBleGFtcGxlLmNvbTpzb21ldGhpbmc=", httpClient: &http.Client{}},
			wantErr: false,
		},
		{
			name:    "Testing invalid apiToken",
			args:    args{email: "test@example.com", apiToken: "", brand: "brand"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Testing invalid brand",
			args:    args{email: "test@example.com", apiToken: "something", brand: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Testing empty email",
			args:    args{email: "", apiToken: "something", brand: "brand"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.email, tt.args.apiToken, tt.args.brand)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_reamazeRequest(t *testing.T) {

	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		method   string
		endpoint string
		payload  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Testing 4xx error",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Status:     "Bad Request",
						Body:       io.NopCloser(strings.NewReader(`{"error":"Bad Request"}`)),
					}
				}),
			}},
			args:    args{method: http.MethodPost, endpoint: "/api/v1/conversations", payload: []byte{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing Status OK",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "Status OK",
						Body:       io.NopCloser(strings.NewReader(`{"status":"ok"}`)),
					}
				}),
			}},
			args:    args{method: http.MethodPost, endpoint: "/api/v1/conversations", payload: []byte{}},
			want:    []byte{123, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 111, 107, 34, 125},
			wantErr: false,
		},
		{
			name: "Testing http.NewRequest passed incorrect url",
			fields: fields{baseURL: ".reamaze.io:", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "Status OK",
						Body:       io.NopCloser(strings.NewReader(`{"status":"ok"}`)),
					}
				}),
			}},
			args:    args{method: http.MethodPost, endpoint: "/api/v1/conversations", payload: []byte{}},
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
			got, err := c.reamazeRequest(tt.args.method, tt.args.endpoint, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.reamazeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.reamazeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
