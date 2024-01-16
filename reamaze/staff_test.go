package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetStaff(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []StaffOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetStaffResponse
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
			want:    &GetStaffResponse{},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetStaff(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStaff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateStaff(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		req *CreateStaffRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateStaffResponse
		wantErr bool
	}{
		{
			name: "Testing empty CreateStaffRequest",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{`)),
					}
				}),
			}},
			args:    args{req: &CreateStaffRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Improcessable Entity",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args: args{req: &CreateStaffRequest{Staff: struct {
				Name     string "json:\"name,omitempty\""
				Email    string "json:\"email,omitempty\""
				Password string "json:\"password,omitempty\""
			}{
				Name: "dummy",
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
			args: args{req: &CreateStaffRequest{Staff: struct {
				Name     string "json:\"name,omitempty\""
				Email    string "json:\"email,omitempty\""
				Password string "json:\"password,omitempty\""
			}{
				Name: "dummy",
			}}},
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
			args: args{req: &CreateStaffRequest{Staff: struct {
				Name     string "json:\"name,omitempty\""
				Email    string "json:\"email,omitempty\""
				Password string "json:\"password,omitempty\""
			}{
				Name: "dummy",
			}}},
			want:    &CreateStaffResponse{},
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
			got, err := c.CreateStaff(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateStaff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newStaffSettings(t *testing.T) {
	type args struct {
		opts []StaffOption
	}
	tests := []struct {
		name    string
		args    args
		want    *ReamazeStaffOptions
		wantErr bool
	}{
		{
			name:    "Testing newStaffSettings expansion when no arguments provided",
			args:    args{},
			want:    &ReamazeStaffOptions{},
			wantErr: false,
		},
		{
			name:    "Testing newSettings WithFilter Option ReamazeFilterAll",
			args:    args{opts: []StaffOption{WithStaffPage(1)}},
			want:    &ReamazeStaffOptions{ReamazeStaffPage: string("page=1")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newStaffSettings(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("newStaffSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStaffSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeStaffOptions_GetQuery(t *testing.T) {
	type fields struct {
		ReamazeStaffPage string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Testing ReamazeStaffOptions_.GetQuery no fields set",
			fields: fields{},
			want:   "",
		},
		{
			name: "Testing ReamazeStaffOptions_.GetQuery with fields set",
			fields: fields{
				ReamazeStaffPage: "page=1",
			},
			want: "?page=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReamazeStaffOptions{
				ReamazeStaffPage: tt.fields.ReamazeStaffPage,
			}
			if got := r.GetQuery(); got != tt.want {
				t.Errorf("ReamazeStaffOptions.GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithStaffPage(t *testing.T) {
	type args struct {
		page int
	}
	tests := []struct {
		name string
		args args
		want ReamazeStaffPage
	}{
		{
			name: "Testing ReamazePage",
			args: args{page: 2},
			want: ReamazeStaffPage(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithStaffPage(tt.args.page); got != tt.want {
				t.Errorf("WithStaffPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeStaffPage_Apply(t *testing.T) {
	type args struct {
		o *ReamazeStaffOptions
	}
	tests := []struct {
		name string
		w    ReamazeStaffPage
		args args
		want *ReamazeStaffOptions
	}{
		{
			name: "Testing if ReamazePage is not being set in ReamazeOptions if page 0",
			w:    ReamazeStaffPage(0),
			args: args{o: &ReamazeStaffOptions{}},
			want: &ReamazeStaffOptions{},
		},
		{
			name: "Testing if ReamazePage is being set in ReamazeOptions",
			w:    ReamazeStaffPage(2),
			args: args{o: &ReamazeStaffOptions{}},
			want: &ReamazeStaffOptions{ReamazeStaffPage: "page=2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazePage.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
