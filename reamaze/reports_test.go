package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestClient_GetReportsVolume(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ReportsOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetReportsVolumeResponse
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
			args:    args{},
			want:    &GetReportsVolumeResponse{},
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
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect invalid endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
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
			got, err := c.GetReportsVolume(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetReportsVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetReportsVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetReportsResponseTime(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ReportsOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetReportsResponseTimeRespone
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
			args:    args{},
			want:    &GetReportsResponseTimeRespone{},
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
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect invalid endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
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
			got, err := c.GetReportsResponseTime(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetReportsResponseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetReportsResponseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetReportsStaff(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ReportsOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetReportsStaffResponse
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
			args:    args{},
			want:    &GetReportsStaffResponse{},
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
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect invalid endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
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
			got, err := c.GetReportsStaff(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetReportsStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetReportsStaff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetReportsTags(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ReportsOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetReportsTagsResponse
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
			args:    args{},
			want:    &GetReportsTagsResponse{},
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
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect invalid endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
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
			got, err := c.GetReportsTags(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetReportsTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetReportsTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetReportsChannelSummary(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ReportsOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetReportsChannelSummaryResponse
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
			args:    args{},
			want:    &GetReportsChannelSummaryResponse{},
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
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing incorrect invalid endpoint response",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable Entity",
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
			got, err := c.GetReportsChannelSummary(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetReportsChannelSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetReportsChannelSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithReportsEndDate(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    ReamazeReportsEndDate
		wantErr bool
	}{
		{
			name: "Testing empty date",
			args: args{},
			want: ReamazeReportsEndDate{},
		},
		{
			name: "Testing correct date",
			args: args{year: 2024, month: 1, day: 1},
			want: ReamazeReportsEndDate(time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithReportsEndDate(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithReportsEndDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithReportsStartDate(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    ReamazeReportsStartDate
		wantErr bool
	}{
		{
			name: "Testing empty date",
			args: args{},
			want: ReamazeReportsStartDate{},
		},
		{
			name: "Testing correct date",
			args: args{year: 2024, month: 1, day: 1},
			want: ReamazeReportsStartDate(time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithReportsStartDate(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithReportsStartDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeReportsStartDate_Apply(t *testing.T) {
	type args struct {
		o *ReamazeReportOptions
	}
	tests := []struct {
		name string
		w    ReamazeReportsStartDate
		args args
		want *ReamazeReportOptions
	}{
		{
			name: "Testing if ReamazeStartDate is  being set in ReamazeReportOptions",
			w:    ReamazeReportsStartDate(time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
			args: args{o: &ReamazeReportOptions{}},
			want: &ReamazeReportOptions{ReamazeStartDate: "start_date=2024-01-01"},
		},
		{
			name: "Testing if ReamazeStartDate is not being set in ReamazeReportOptions if date is empty",
			w:    ReamazeReportsStartDate{},
			args: args{o: &ReamazeReportOptions{}},
			want: &ReamazeReportOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeReportsStartDate.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReamazeReportsEndDate_Apply(t *testing.T) {
	type args struct {
		o *ReamazeReportOptions
	}
	tests := []struct {
		name string
		w    ReamazeReportsEndDate
		args args
		want *ReamazeReportOptions
	}{
		{
			name: "Testing if ReamazeReportsEndDate is  being set in ReamazeReportOptions",
			w:    ReamazeReportsEndDate(time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
			args: args{o: &ReamazeReportOptions{}},
			want: &ReamazeReportOptions{ReamazeEndDate: "end_date=2024-01-01"},
		},
		{
			name: "Testing if ReamazeEndDate is not being set in ReamazeReportOptions if date is empty",
			w:    ReamazeReportsEndDate{},
			args: args{o: &ReamazeReportOptions{}},
			want: &ReamazeReportOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeReportsEndDate.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newReportsSettings(t *testing.T) {
	type args struct {
		opts []ReportsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *ReamazeReportOptions
		wantErr bool
	}{
		{
			name:    "Testing newReportsSettings expansion when no arguments provided",
			args:    args{},
			want:    &ReamazeReportOptions{},
			wantErr: false,
		},
		{
			name:    "Testing newReportsSettings WithFilter Option ReamazeFilterAll",
			args:    args{opts: []ReportsOption{WithReportsStartDate(2024, 1, 1)}},
			want:    &ReamazeReportOptions{ReamazeStartDate: string("start_date=2024-01-01")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newReportsSettings(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("newReportsSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newReportsSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeReportOptions_GetQuery(t *testing.T) {
	type fields struct {
		ReamazeStartDate string
		ReamazeEndDate   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Testing ReamazeReportOptions.GetQuery no fields set",
			fields: fields{},
			want:   "",
		},
		{
			name: "Testing ReamazeReportOptions.GetQuery with all fields set to dummy",
			fields: fields{
				ReamazeStartDate: "dummy",
				ReamazeEndDate:   "dummy",
			},
			want: "?dummy&dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReamazeReportOptions{
				ReamazeStartDate: tt.fields.ReamazeStartDate,
				ReamazeEndDate:   tt.fields.ReamazeEndDate,
			}
			if got := r.GetQuery(); got != tt.want {
				t.Errorf("ReamazeReportOptions.GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
