package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetContact(t *testing.T) {
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
		want    *GetContactResponse
		wantErr bool
	}{
		{
			name: "Testing correct requset",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy@example.com"},
			want:    &GetContactResponse{},
			wantErr: false,
		},
		{
			name: "Testing no identifier",
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
			name: "Testing bad response code",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing bad response json",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args:    args{identifier: "dummy@example.com"},
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
			got, err := c.GetContact(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetContacts(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetContactsResponse
		wantErr bool
	}{
		{
			name: "Testing bad response code",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing bad response json",
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
			name: "Testing correct response json",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			want:    &GetContactsResponse{},
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
			got, err := c.GetContacts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetContacts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateContact(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		req *CreateContactRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetContactResponse
		wantErr bool
	}{
		{
			name: "Testing empty request failure",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{req: &CreateContactRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect phone number",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateContactRequest{Contact: struct {
				Name              string             "json:\"name\""
				Email             string             "json:\"email\""
				Mobile            ReamazePhoneNumber "json:\"mobile\""
				FriendlyName      string             "json:\"friendly_name\""
				ID                string             "json:\"id\""
				ExternalAvatarURL string             "json:\"external_avatar_url\""
				Notes             []string           "json:\"notes\""
				Data              interface{}        "json:\"data\""
			}{
				Email:  "",
				Mobile: "123",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect phone number and valid email",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateContactRequest{Contact: struct {
				Name              string             "json:\"name\""
				Email             string             "json:\"email\""
				Mobile            ReamazePhoneNumber "json:\"mobile\""
				FriendlyName      string             "json:\"friendly_name\""
				ID                string             "json:\"id\""
				ExternalAvatarURL string             "json:\"external_avatar_url\""
				Notes             []string           "json:\"notes\""
				Data              interface{}        "json:\"data\""
			}{
				Email:  "dummy@example.com",
				Mobile: "123",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing valid email with valid response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateContactRequest{Contact: struct {
				Name              string             "json:\"name\""
				Email             string             "json:\"email\""
				Mobile            ReamazePhoneNumber "json:\"mobile\""
				FriendlyName      string             "json:\"friendly_name\""
				ID                string             "json:\"id\""
				ExternalAvatarURL string             "json:\"external_avatar_url\""
				Notes             []string           "json:\"notes\""
				Data              interface{}        "json:\"data\""
			}{
				Email: "dummy@example.com",
			}}},
			want:    &GetContactResponse{},
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
			args: args{req: &CreateContactRequest{Contact: struct {
				Name              string             "json:\"name\""
				Email             string             "json:\"email\""
				Mobile            ReamazePhoneNumber "json:\"mobile\""
				FriendlyName      string             "json:\"friendly_name\""
				ID                string             "json:\"id\""
				ExternalAvatarURL string             "json:\"external_avatar_url\""
				Notes             []string           "json:\"notes\""
				Data              interface{}        "json:\"data\""
			}{
				Email: "dummy@example.com",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing invalid response code",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateContactRequest{Contact: struct {
				Name              string             "json:\"name\""
				Email             string             "json:\"email\""
				Mobile            ReamazePhoneNumber "json:\"mobile\""
				FriendlyName      string             "json:\"friendly_name\""
				ID                string             "json:\"id\""
				ExternalAvatarURL string             "json:\"external_avatar_url\""
				Notes             []string           "json:\"notes\""
				Data              interface{}        "json:\"data\""
			}{
				Email: "dummy@example.com",
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
			got, err := c.CreateContact(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateContact(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier     string
		req            *UpdateContactRequest
		identifierType []ReamazeIdentifier
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetContactResponse
		wantErr bool
	}{
		{
			name: "Testing empty UpdateContactRequest request with identifier set",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &UpdateContactRequest{Contact: struct {
				Name              string      "json:\"name\""
				FriendlyName      string      "json:\"friendly_name\""
				ExternalAvatarURL string      "json:\"external_avatar_url\""
				Notes             []string    "json:\"notes\""
				Data              interface{} "json:\"data\""
			}{}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty identifier and UpdateContactRequest properly set",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "", req: &UpdateContactRequest{Contact: struct {
				Name              string      "json:\"name\""
				FriendlyName      string      "json:\"friendly_name\""
				ExternalAvatarURL string      "json:\"external_avatar_url\""
				Notes             []string    "json:\"notes\""
				Data              interface{} "json:\"data\""
			}{
				FriendlyName: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty identifier and UpdateContactRequest properly set",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &UpdateContactRequest{Contact: struct {
				Name              string      "json:\"name\""
				FriendlyName      string      "json:\"friendly_name\""
				ExternalAvatarURL string      "json:\"external_avatar_url\""
				Notes             []string    "json:\"notes\""
				Data              interface{} "json:\"data\""
			}{
				FriendlyName: "dummy",
			}},
				identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile},
			},
			want:    &GetContactResponse{},
			wantErr: false,
		},
		{
			name: "Testing invalid JSON respond",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &UpdateContactRequest{Contact: struct {
				Name              string      "json:\"name\""
				FriendlyName      string      "json:\"friendly_name\""
				ExternalAvatarURL string      "json:\"external_avatar_url\""
				Notes             []string    "json:\"notes\""
				Data              interface{} "json:\"data\""
			}{
				FriendlyName: "dummy",
			}},
				identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing invalid response status code",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Status Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &UpdateContactRequest{Contact: struct {
				Name              string      "json:\"name\""
				FriendlyName      string      "json:\"friendly_name\""
				ExternalAvatarURL string      "json:\"external_avatar_url\""
				Notes             []string    "json:\"notes\""
				Data              interface{} "json:\"data\""
			}{
				FriendlyName: "dummy",
			}},
				identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile},
			},
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
			got, err := c.UpdateContact(tt.args.identifier, tt.args.req, tt.args.identifierType...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UpdateContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetContactIdentities(t *testing.T) {
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
		want    *GetContactIdentitiesResponse
		wantErr bool
	}{
		{
			name: "Testing empty identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Status Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy@example.com"},
			want:    &GetContactIdentitiesResponse{},
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
			args:    args{identifier: "dummy@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing invalid response status code",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy@example.com"},
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
			got, err := c.GetContactIdentities(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetContactIdentities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetContactIdentities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateContactIdentities(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier     string
		req            *CreateContactIdentitiesRequest
		identifierType []ReamazeIdentifier
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetContactIdentitiesResponse
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
			name: "Testing correct identifier and empty CreateContactIdentitiesRequest",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy@example.com", req: &CreateContactIdentitiesRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing correct identifier, CreateContactIdentitiesRequest and ReamazeIdentifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &CreateContactIdentitiesRequest{Identity: struct {
				Type       ReamazeIdentifier "json:\"type\""
				Identifier string            "json:\"identifier\""
			}{
				Type:       ReamazeIdentifierEmail,
				Identifier: "dummy@example.com",
			}}, identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile}},
			want:    &GetContactIdentitiesResponse{},
			wantErr: false,
		},
		{
			name: "Testing failure of adding Facebook identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &CreateContactIdentitiesRequest{Identity: struct {
				Type       ReamazeIdentifier "json:\"type\""
				Identifier string            "json:\"identifier\""
			}{
				Type:       ReamazeIdentifierFacebook,
				Identifier: "dummy@example.com",
			}}, identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile}},
			want:    nil,
			wantErr: true,
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
			args: args{identifier: "dummy@example.com", req: &CreateContactIdentitiesRequest{Identity: struct {
				Type       ReamazeIdentifier "json:\"type\""
				Identifier string            "json:\"identifier\""
			}{
				Type:       ReamazeIdentifierEmail,
				Identifier: "dummy@example.com",
			}}, identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing invalid request response code",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{identifier: "dummy@example.com", req: &CreateContactIdentitiesRequest{Identity: struct {
				Type       ReamazeIdentifier "json:\"type\""
				Identifier string            "json:\"identifier\""
			}{
				Type:       ReamazeIdentifierEmail,
				Identifier: "dummy@example.com",
			}}, identifierType: []ReamazeIdentifier{ReamazeIdentifierMobile}},
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
			got, err := c.CreateContactIdentities(tt.args.identifier, tt.args.req, tt.args.identifierType...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateContactIdentities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateContactIdentities() = %v, want %v", got, tt.want)
			}
		})
	}
}
