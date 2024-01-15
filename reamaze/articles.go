package reamaze

import (
	"encoding/json"
	"net/http"
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
