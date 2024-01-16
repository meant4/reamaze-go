package reamaze

import (
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestClient_GetArticles(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ArticlesOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetArticlesResponse
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
			want:    &GetArticlesResponse{},
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
			got, err := c.GetArticles(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetArticle(t *testing.T) {
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
		want    *GetArticleResponse
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
			args:    args{slug: "dummy"},
			want:    &GetArticleResponse{},
			wantErr: false,
		},
		{
			name: "Testing empty slug parameter",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetArticle(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateArticle(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		slug string
		req  *UpdateArticleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateArticleResponse
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
			args:    args{slug: "dummy", req: &UpdateArticleRequest{Title: "dummy"}},
			want:    &UpdateArticleResponse{},
			wantErr: false,
		},
		{
			name: "Testing empty UpdateArticleRequest request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "dummy", req: &UpdateArticleRequest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty slug parameter",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{slug: "", req: &UpdateArticleRequest{Title: "dummy"}},
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
			args:    args{slug: "dummy", req: &UpdateArticleRequest{Title: "dummy"}},
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
			args:    args{slug: "dummy", req: &UpdateArticleRequest{Title: "dummy"}},
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
			got, err := c.UpdateArticle(tt.args.slug, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UpdateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateArticle(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		req *CreateArticleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateArticleResponse
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
			args: args{req: &CreateArticleRequest{Article: struct {
				Title   string               "json:\"title,omitempty\""
				Body    string               "json:\"body,omitempty\""
				Status  ReamazeArticleStatus "json:\"status,omitempty\""
				TopicID string               "json:\"topic_id,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    &CreateArticleResponse{},
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
			args: args{req: &CreateArticleRequest{Article: struct {
				Title   string               "json:\"title,omitempty\""
				Body    string               "json:\"body,omitempty\""
				Status  ReamazeArticleStatus "json:\"status,omitempty\""
				TopicID string               "json:\"topic_id,omitempty\""
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
			args: args{req: &CreateArticleRequest{Article: struct {
				Title   string               "json:\"title,omitempty\""
				Body    string               "json:\"body,omitempty\""
				Status  ReamazeArticleStatus "json:\"status,omitempty\""
				TopicID string               "json:\"topic_id,omitempty\""
			}{
				Title: "dummy",
			}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing empty  CreateArticleRequest request",
			fields: fields{baseURL: "https://dummy.reamaze.io", auth: "dummy", httpClient: &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Status:     "200 Status OK",
						Body:       io.NopCloser(strings.NewReader(`{}`)),
					}
				}),
			}},
			args:    args{req: &CreateArticleRequest{}},
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
			got, err := c.CreateArticle(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeArticlesOptions_GetQuery(t *testing.T) {
	type fields struct {
		ReamazeArticlesStatus string
		ReamazeArticlesQuery  string
		ReamazeArticlesPage   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Testing ReamazeArticlesOptions.GetQuery no fields set",
			fields: fields{},
			want:   "",
		},
		{
			name: "Testing ReamazeArticlesOptions.GetQuery with fields set",
			fields: fields{
				ReamazeArticlesQuery:  "q=dummy",
				ReamazeArticlesPage:   "page=1",
				ReamazeArticlesStatus: "status=" + strconv.Itoa(int(ReamazeArticleStatusPublished)),
			},
			want: "?status=" + strconv.Itoa(int(ReamazeArticleStatusPublished)) + "&q=dummy&page=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReamazeArticlesOptions{
				ReamazeArticlesStatus: tt.fields.ReamazeArticlesStatus,
				ReamazeArticlesQuery:  tt.fields.ReamazeArticlesQuery,
				ReamazeArticlesPage:   tt.fields.ReamazeArticlesPage,
			}
			if got := r.GetQuery(); got != tt.want {
				t.Errorf("ReamazeArticlesOptions.GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newArticlesSettings(t *testing.T) {
	type args struct {
		opts []ArticlesOption
	}
	tests := []struct {
		name    string
		args    args
		want    *ReamazeArticlesOptions
		wantErr bool
	}{
		{
			name:    "Testing newArticlesSettings expansion when no arguments provided",
			args:    args{},
			want:    &ReamazeArticlesOptions{},
			wantErr: false,
		},
		{
			name:    "Testing newArticlesSettings WithArticleStatus with ReamazeArticleStatusPublished",
			args:    args{opts: []ArticlesOption{WithArticleStatus(ReamazeArticleStatusDraft)}},
			want:    &ReamazeArticlesOptions{ReamazeArticlesStatus: "status=1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newArticlesSettings(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("newArticlesSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newArticlesSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArticlePage(t *testing.T) {
	type args struct {
		page int
	}
	tests := []struct {
		name string
		args args
		want ReamazeArticlePage
	}{
		{
			name: "Testing ReamazeArticlePage",
			args: args{page: 2},
			want: ReamazeArticlePage(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArticlePage(tt.args.page); got != tt.want {
				t.Errorf("WithArticlePage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArticleQuery(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want ReamazeArticleQuery
	}{
		{
			name: "Testing WithArticleQuery",
			args: args{query: "dummy"},
			want: ReamazeArticleQuery("dummy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArticleQuery(tt.args.query); got != tt.want {
				t.Errorf("WithArticleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeArticlePage_Apply(t *testing.T) {
	type args struct {
		o *ReamazeArticlesOptions
	}
	tests := []struct {
		name string
		w    ReamazeArticlePage
		args args
		want *ReamazeArticlesOptions
	}{
		{
			name: "Testing if ReamazeArticlesPage is not being set when 0",
			w:    ReamazeArticlePage(0),
			args: args{o: &ReamazeArticlesOptions{}},
			want: &ReamazeArticlesOptions{},
		},
		{
			name: "Testing if ReamazeArticlePage is being set",
			w:    ReamazeArticlePage(1),
			args: args{o: &ReamazeArticlesOptions{}},
			want: &ReamazeArticlesOptions{ReamazeArticlesPage: "page=1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeArticlePage.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReamazeArticleQuery_Apply(t *testing.T) {
	type args struct {
		o *ReamazeArticlesOptions
	}
	tests := []struct {
		name string
		w    ReamazeArticleQuery
		args args
		want *ReamazeArticlesOptions
	}{
		{
			name: "Testing if ReamazeArticlesQuery is being set",
			w:    ReamazeArticleQuery("dummy"),
			args: args{o: &ReamazeArticlesOptions{}},
			want: &ReamazeArticlesOptions{ReamazeArticlesQuery: "q=dummy"},
		},
		{
			name: "Testing if ReamazeArticleQuery is not being set when empty",
			w:    ReamazeArticleQuery(""),
			args: args{o: &ReamazeArticlesOptions{}},
			want: &ReamazeArticlesOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Apply(tt.args.o)
			got := tt.args.o
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReamazeArticleQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
