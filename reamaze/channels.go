package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// GetChannels will allow you to retrieve channels for the Brand https://www.reamaze.com/api/get_channels
func (c *Client) GetChannels() (*GetChannelsResponse, error) {
	var response *GetChannelsResponse

	urlEndpoint := channelsEndpoint
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

// GetChannel will allow you to retrieve a specific channel https://www.reamaze.com/api/get_channel
func (c *Client) GetChannel(slug string) (*GetChannelResponse, error) {
	var response *GetChannelResponse
	// checking if slug is set
	if len(slug) == 0 {
		return nil, errors.New("GetChannel slug cannot be empty, please provide slug as argument")
	}
	urlEndpoint := channelsEndpoint + "/" + url.PathEscape(slug)

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
