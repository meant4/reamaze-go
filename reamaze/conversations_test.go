package reamaze

import (
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestClient_CreateConversation(t *testing.T) {
	correctConversation := &CreateConversationRequest{}
	correctConversation.Conversation.Category = "test"
	correctConversation.Conversation.Message.Body = "test"
	correctConversation.Conversation.User.Email = "test"
	correctConversation.Conversation.Data = struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	}{
		FirstName: "dummy",
		LastName:  "dummy",
	}
	correctConversation.Conversation.User.Data = struct {
		JobApplication bool   `json:"job_application,omitempty"`
		FirstName      string `json:"first_name,omitempty"`
		LastName       string `json:"last_name,omitempty"`
	}{
		JobApplication: true,
		FirstName:      "dummy",
		LastName:       "dummy",
	}

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
		o []ConversationsOption
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

func TestReamazeFilter_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeFilter
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if reamaze filter all is being set",
			w:    ReamazeFilterAll,
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeFilter: string("filter=" + ReamazeFilterAll)},
		},
		{
			name: "Testing if reamaze filter is not being set when it's empty",
			w:    ReamazeFilter(""),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeFilter.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeFor_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeFor
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if reamaze for filter is being added to ReamazeOptions with dummy email",
			w:    ReamazeFor("dummy@example.com"),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeFor: "for=" + url.QueryEscape(string(ReamazeFor("dummy@example.com")))},
		},
		{
			name: "Testing if reamaze for filter is not being added to ReamazeOptions with invalid email",
			w:    ReamazeFor("dummy"),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeFor.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeForID_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeForID
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeForID is being set in ReamazeOptions with id",
			w:    ReamazeForID("dummy"),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeForID: "for_id=" + url.QueryEscape("dummy")},
		},
		{
			name: "Testing if ReamazeForID is not being set in ReamazeOptions with id when empty",
			w:    ReamazeForID(""),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeForID.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeSort_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeSort
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeSort is being set in ReamazeOptions with ReamazeSort value",
			w:    ReamazeSortChanged,
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeSort: "sort=" + string(ReamazeSortChanged)},
		},
		{
			name: "Testing if ReamazeSort is not being set in ReamazeOptions with when empty",
			w:    ReamazeSortChanged,
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeSort: "sort=" + string(ReamazeSortChanged)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeSort.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeData_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeData
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeSort is being set in ReamazeOptions with ReamazeSort value",
			w:    ReamazeData{"dummy": "dummy"},
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeData: "data[dummy]=dummy"},
		},
		{
			name: "Testing if ReamazeSort is not being set in ReamazeOptions with when empty",
			w:    ReamazeData{},
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeData.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeCategory_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeCategory
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeCategory is being set in ReamazeOptions with ReamazeCategory value",
			w:    ReamazeCategory("dummy"),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeCategory: "category=dummy"},
		},
		{
			name: "Testing if ReamazeCategory is not being set in ReamazeOptions when empty",
			w:    ReamazeCategory("dummy"),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeCategory: "category=dummy"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeCategory.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeEndDate_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeEndDate
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeEndDate is  being set in ReamazeOptions",
			w:    ReamazeEndDate(time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeEndDate: "end_date=2024-01-01"},
		},
		{
			name: "Testing if ReamazeEndDate is not being set in ReamazeOptions if date is empty",
			w:    ReamazeEndDate{},
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeEndDate.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeStartDate_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeStartDate
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeStartDate is  being set in ReamazeOptions",
			w:    ReamazeStartDate(time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeStartDate: "start_date=2024-01-01"},
		},
		{
			name: "Testing if ReamazeStartDate is not being set in ReamazeOptions if date is empty",
			w:    ReamazeStartDate{},
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeStartDate.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeTags_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazeTags
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazeTags is not being set in ReamazeOptions if we have empty tags",
			w:    ReamazeTags{},
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
		{
			name: "Testing if ReamazeTags is being set in ReamazeOptions",
			w:    ReamazeTags{"dummy", "dummy2"},
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazeTag: "tag=" + url.QueryEscape("dummy,dummy2")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeTags.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazePage_Apply(t *testing.T) {
	type args struct {
		o *ReamazeOptions
	}
	tests := []struct {
		name string
		w    ReamazePage
		args args
		want *ReamazeOptions
	}{
		{
			name: "Testing if ReamazePage is not being set in ReamazeOptions if page 0",
			w:    ReamazePage(0),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{},
		},
		{
			name: "Testing if ReamazePage is being set in ReamazeOptions",
			w:    ReamazePage(2),
			args: args{o: &ReamazeOptions{}},
			want: &ReamazeOptions{ReamazePage: "page=2"},
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

func TestWithFilter(t *testing.T) {
	type args struct {
		filter ReamazeFilter
	}
	tests := []struct {
		name string
		args args
		want ReamazeFilter
	}{
		{
			name: "Testing WithFilter",
			args: args{filter: ReamazeFilterAll},
			want: ReamazeFilterAll,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithFilter(tt.args.filter); got != tt.want {
				t.Errorf("WithFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithFor(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want ReamazeFor
	}{
		{
			name: "Testing WithFor",
			args: args{email: "dummy@example.com"},
			want: ReamazeFor("dummy@example.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithFor(tt.args.email); got != tt.want {
				t.Errorf("WithFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithForID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want ReamazeForID
	}{
		{
			name: "Testing WithForID",
			args: args{id: "dummy"},
			want: ReamazeForID("dummy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithForID(tt.args.id); got != tt.want {
				t.Errorf("WithForID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSort(t *testing.T) {
	type args struct {
		sort ReamazeSort
	}
	tests := []struct {
		name string
		args args
		want ReamazeSort
	}{
		{
			name: "Testing WithSort",
			args: args{sort: ReamazeSortUpdated},
			want: ReamazeSortUpdated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithSort(tt.args.sort); got != tt.want {
				t.Errorf("WithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCategory(t *testing.T) {
	type args struct {
		category string
	}
	tests := []struct {
		name string
		args args
		want ReamazeCategory
	}{
		{
			name: "Testing WithCategory",
			args: args{category: "dummy"},
			want: ReamazeCategory("dummy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCategory(tt.args.category); got != tt.want {
				t.Errorf("WithCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithData(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name string
		args args
		want ReamazeData
	}{
		{
			name: "Testing WithData",
			args: args{data: map[string]string{"dummy": "dummy"}},
			want: ReamazeData{"dummy": "dummy"},
		},
		{
			name: "Testing WithData with empty map",
			args: args{data: map[string]string{}},
			want: ReamazeData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEndDate(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want ReamazeEndDate
	}{
		{
			name: "Testing WithEndDate",
			args: args{year: 2024, month: 1, day: 1},
			want: ReamazeEndDate(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "Testing WithEndDate when one of the args is 0",
			args: args{year: 2024, month: 0, day: 0},
			want: ReamazeEndDate{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEndDate(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEndDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithStartDate(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want ReamazeStartDate
	}{
		{
			name: "Testing ReamazeStartDate",
			args: args{year: 2024, month: 1, day: 1},
			want: ReamazeStartDate(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "Testing ReamazeStartDate when one of the args is 0",
			args: args{year: 2024, month: 0, day: 0},
			want: ReamazeStartDate{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithStartDate(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithStartDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPage(t *testing.T) {
	type args struct {
		page int
	}
	tests := []struct {
		name string
		args args
		want ReamazePage
	}{
		{
			name: "Testing ReamazePage",
			args: args{page: 2},
			want: ReamazePage(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPage(tt.args.page); got != tt.want {
				t.Errorf("WithPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithTags(t *testing.T) {
	type args struct {
		w []string
	}
	tests := []struct {
		name string
		args args
		want ReamazeTags
	}{
		{
			name: "Testing ReamazeTags",
			args: args{w: []string{"dummy", "dummy2"}},
			want: ReamazeTags{"dummy", "dummy2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithTags(tt.args.w...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSettings(t *testing.T) {
	type args struct {
		opts []ConversationsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *ReamazeOptions
		wantErr bool
	}{
		{
			name:    "Testing newSettings expansion when no arguments provided",
			args:    args{},
			want:    &ReamazeOptions{},
			wantErr: false,
		},
		{
			name:    "Testing newSettings WithFilter Option ReamazeFilterAll",
			args:    args{opts: []ConversationsOption{WithFilter(ReamazeFilterAll)}},
			want:    &ReamazeOptions{ReamazeFilter: string("filter=" + ReamazeFilterAll)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newSettings(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeOptions_GetQuery(t *testing.T) {
	type fields struct {
		ReamazeFilter    string
		ReamazeFor       string
		ReamazeForID     string
		ReamazeSort      string
		ReamazeTag       string
		ReamazeCategory  string
		ReamazeData      string
		ReamazePage      string
		ReamazeStartDate string
		ReamazeEndDate   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Testing ReamazeOptions.GetQuery no fields set",
			fields: fields{},
			want:   "",
		},
		{
			name: "Testing ReamazeOptions.GetQuery with all fields set to dummy",
			fields: fields{
				ReamazeFilter:    "dummy",
				ReamazeFor:       "dummy",
				ReamazeForID:     "dummy",
				ReamazeSort:      "dummy",
				ReamazeTag:       "dummy",
				ReamazeCategory:  "dummy",
				ReamazeData:      "dummy",
				ReamazePage:      "dummy",
				ReamazeStartDate: "dummy",
				ReamazeEndDate:   "dummy",
			},
			want: "?dummy&dummy&dummy&dummy&dummy&dummy&dummy&dummy&dummy&dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReamazeOptions{
				ReamazeFilter:    tt.fields.ReamazeFilter,
				ReamazeFor:       tt.fields.ReamazeFor,
				ReamazeForID:     tt.fields.ReamazeForID,
				ReamazeSort:      tt.fields.ReamazeSort,
				ReamazeTag:       tt.fields.ReamazeTag,
				ReamazeCategory:  tt.fields.ReamazeCategory,
				ReamazeData:      tt.fields.ReamazeData,
				ReamazePage:      tt.fields.ReamazePage,
				ReamazeStartDate: tt.fields.ReamazeStartDate,
				ReamazeEndDate:   tt.fields.ReamazeEndDate,
			}
			if got := r.GetQuery(); got != tt.want {
				t.Errorf("ReamazeOptions.GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
