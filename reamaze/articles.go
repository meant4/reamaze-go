package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
)

// GetArticles will allow you to retrieve Help Articles for the Brand
// optional parameters WithArticlePage(int), WithArticleStatus(ReamazeArticleStatus),WithArticleQuery(string)
// https://www.reamaze.com/api/get_articles
func (c *Client) GetArticles(o ...ArticlesOption) (*GetArticlesResponse, error) {
	var response *GetArticlesResponse
	settings, _ := newArticlesSettings(o)
	urlEndpoint := articlesEndpoint + settings.GetQuery()

	resp, err := c.reamazeRequest(http.MethodGet, urlEndpoint, []byte{})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetArticle will allow you to retrieve a specific Help Article
// https://www.reamaze.com/api/get_article
func (c *Client) GetArticle(slug string) (*GetArticleResponse, error) {
	var response *GetArticleResponse
	// checking if slug is set
	if len(slug) == 0 {
		return nil, errors.New("GetArticle slug cannot be empty, please provide slug as argument")
	}
	urlEndpoint := articlesEndpoint + "/" + url.PathEscape(slug)
	resp, err := c.reamazeRequest(http.MethodGet, urlEndpoint, []byte{})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateArticle will allow you to update the article.
// https://www.reamaze.com/api/put_article
func (c *Client) UpdateArticle(slug string, req *UpdateArticleRequest) (*UpdateArticleResponse, error) {
	var response *UpdateArticleResponse
	emptyReq := &UpdateArticleRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("UpdateArticle incorrect request, UpdateArticleRequest is empty")
	}
	// checking if slug is set
	if len(slug) == 0 {
		return nil, errors.New("UpdateArticle slug cannot be empty, please provide slug as argument")
	}
	urlEndpoint := articlesEndpoint + "/" + url.PathEscape(slug)
	resp, err := c.reamazeRequest(http.MethodGet, urlEndpoint, []byte{})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// CreateArticle will allow you to update the article.
// https://www.reamaze.com/api/put_article
func (c *Client) CreateArticle(req *CreateArticleRequest) (*CreateArticleResponse, error) {
	var response *CreateArticleResponse
	emptyReq := &CreateArticleRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateArticle incorrect request, CreateArticleRequest is empty")
	}

	urlEndpoint := articlesEndpoint
	data, _ := json.Marshal(req)
	resp, err := c.reamazeRequest(http.MethodPost, urlEndpoint, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
