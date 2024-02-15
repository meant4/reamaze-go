package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetNotes(t *testing.T) {

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
		want    *GetNotesResponse
		wantErr bool
	}{
		{
			name: "Testing lack of identifier",
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
			name: "Testing incorrect response from the REST endpoint",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect JSON response from the REST endpoint",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args:    args{identifier: "dummy"},
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
						Body:       io.NopCloser(strings.NewReader(`[{}]`)),
					}
				}),
			}},
			args:    args{identifier: "dummy"},
			want:    &GetNotesResponse{Note{}},
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
			got, err := c.GetNotes(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateNote(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier string
		req        *CreateNoteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateNoteResponse
		wantErr bool
	}{
		{
			name: "Testing incorrect empty identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "", req: &CreateNoteRequest{Body: "dummy"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty CraeteNoteRequest",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy", req: &CreateNoteRequest{}},
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
			args:    args{identifier: "dummy", req: &CreateNoteRequest{Body: "dummy"}},
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
			args:    args{identifier: "dummy", req: &CreateNoteRequest{Body: "dummy"}},
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
			args:    args{identifier: "dummy", req: &CreateNoteRequest{Body: "dummy"}},
			want:    &CreateNoteResponse{},
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
			got, err := c.CreateNote(tt.args.identifier, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateNote(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier string
		noteID     string
		req        *UpdateNoteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateNoteResponse
		wantErr bool
	}{
		{
			name: "Testing incorrect empty identifier",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "", noteID: "", req: &UpdateNoteRequest{Body: "dummy"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty UpdateNoteRequest",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy", noteID: "", req: &UpdateNoteRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty noteID",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy", noteID: "", req: &UpdateNoteRequest{Body: "dummy"}},
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
			args:    args{identifier: "dummy", noteID: "dummy", req: &UpdateNoteRequest{Body: "dummy"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy", noteID: "dummy", req: &UpdateNoteRequest{Body: "dummy"}},
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
			args:    args{identifier: "dummy", noteID: "dummy", req: &UpdateNoteRequest{Body: "dummy"}},
			want:    &UpdateNoteResponse{},
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
			got, err := c.UpdateNote(tt.args.identifier, tt.args.noteID, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UpdateNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteNote(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		identifier string
		noteID     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DeleteNoteResponse
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
			args:    args{identifier: "dummy", noteID: "dummy"},
			want:    &DeleteNoteResponse{},
			wantErr: false,
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
			args:    args{identifier: "", noteID: "dummy"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty noteID",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{identifier: "dummy", noteID: ""},
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
			args:    args{identifier: "dummy", noteID: "dummy"},
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
			args:    args{identifier: "dummy", noteID: "dummy"},
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
			got, err := c.DeleteNote(tt.args.identifier, tt.args.noteID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DeleteNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
