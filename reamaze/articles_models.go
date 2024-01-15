package reamaze

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const articlesEndpoint string = "/api/v1/articles"

type ReamazeArticleStatus int
type ReamazeArticleQuery string
type ReamazeArticlePage int

const (
	ReamazeArticleStatusPublished ReamazeArticleStatus = 0
	ReamazeArticleStatusDraft     ReamazeArticleStatus = 1
	ReamazeArticleStatusInternal  ReamazeArticleStatus = 2
)

type ArticlesOption interface {
	Apply(*ReamazeArticlesOptions)
}

func (w ReamazeArticleStatus) Apply(o *ReamazeArticlesOptions) {
	if w > 0 {
		o.ReamazeArticlesStatus = "status=" + strconv.Itoa(int(w))
	}
}

func (w ReamazeArticlePage) Apply(o *ReamazeArticlesOptions) {
	if w > 0 {
		o.ReamazeArticlesPage = "page=" + strconv.Itoa(int(w))
	}
}

func (w ReamazeArticleQuery) Apply(o *ReamazeArticlesOptions) {
	if len(w) > 0 {
		o.ReamazeArticlesQuery = "q=" + url.PathEscape(string(w))
	}
}

func WithArticleQuery(query string) ReamazeArticleQuery {
	return ReamazeArticleQuery(query)
}

func WithArticleStatus(status ReamazeArticleStatus) ReamazeArticleStatus {
	return status
}

func WithArticlePage(page int) ReamazeArticlePage {
	return ReamazeArticlePage(page)
}

func newArticlesSettings(opts []ArticlesOption) (*ReamazeArticlesOptions, error) {
	var o ReamazeArticlesOptions
	for _, opt := range opts {
		opt.Apply(&o)
	}
	return &o, nil
}

func (r ReamazeArticlesOptions) GetQuery() string {
	output := ""
	var queryParams []string

	// checking if status is set
	if len(r.ReamazeArticlesStatus) > 0 {
		queryParams = append(queryParams, r.ReamazeArticlesStatus)
	}
	// checking if query is set
	if len(r.ReamazeArticlesQuery) > 0 {
		queryParams = append(queryParams, r.ReamazeArticlesQuery)
	}
	// checking if page is set
	if len(r.ReamazeArticlesPage) > 0 {
		queryParams = append(queryParams, r.ReamazeArticlesPage)
	}

	output = strings.Join(queryParams, "&")
	if len(output) > 0 {
		output = "?" + output
	}
	return output
}

type ReamazeArticlesOptions struct {
	ReamazeArticlesStatus string
	ReamazeArticlesQuery  string
	ReamazeArticlesPage   string
}
type GetArticlesResponse struct {
	PageSize   int `json:"page_size,omitempty"`
	PageCount  int `json:"page_count,omitempty"`
	TotalCount int `json:"total_count,omitempty"`
	Articles   []struct {
		Title     string    `json:"title,omitempty"`
		Body      string    `json:"body,omitempty"`
		Slug      string    `json:"slug,omitempty"`
		Status    int       `json:"status,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
		URL       string    `json:"url,omitempty"`
		Author    struct {
			ID           int    `json:"id,omitempty"`
			Name         string `json:"name,omitempty"`
			Data         any    `json:"data,omitempty"`
			Email        string `json:"email,omitempty"`
			Twitter      string `json:"twitter,omitempty"`
			Facebook     string `json:"facebook,omitempty"`
			Instagram    string `json:"instagram,omitempty"`
			Mobile       string `json:"mobile,omitempty"`
			FriendlyName string `json:"friendly_name,omitempty"`
			DisplayName  string `json:"display_name,omitempty"`
		} `json:"author,omitempty"`
		EmbeddedURL string `json:"embedded_url,omitempty"`
		Topic       struct {
			Name string `json:"name,omitempty"`
			Slug string `json:"slug,omitempty"`
		} `json:"topic,omitempty"`
	} `json:"articles,omitempty"`
}
type ReamazeArticle struct {
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	Status    int       `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	URL       string    `json:"url,omitempty"`
	Author    struct {
		ID           int    `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		Data         any    `json:"data,omitempty"`
		Email        string `json:"email,omitempty"`
		Twitter      string `json:"twitter,omitempty"`
		Facebook     string `json:"facebook,omitempty"`
		Instagram    string `json:"instagram,omitempty"`
		Mobile       string `json:"mobile,omitempty"`
		FriendlyName string `json:"friendly_name,omitempty"`
		DisplayName  string `json:"display_name,omitempty"`
	} `json:"author,omitempty"`
	EmbeddedURL string `json:"embedded_url,omitempty"`
	Topic       struct {
		Name string `json:"name,omitempty"`
		Slug string `json:"slug,omitempty"`
	} `json:"topic,omitempty"`
}

type GetArticleResponse ReamazeArticle
type UpdateArticleResponse ReamazeArticle
type UpdateArticleRequest ReamazeArticle
type CreateArticleResponse ReamazeArticle

type CreateArticleRequest struct {
	Article struct {
		Title   string               `json:"title,omitempty"`
		Body    string               `json:"body,omitempty"`
		Status  ReamazeArticleStatus `json:"status,omitempty"`
		TopicID string               `json:"topic_id,omitempty"`
	} `json:"article,omitempty"`
}
