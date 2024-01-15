package reamaze

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
)

// call to messages will allow you to retrieve individual messages for all conversations in the Brand
// https://www.reamaze.com/api/get_messages
func (c *Client) GetMessages() (*GetMessagesResponse, error) {
	var response *GetMessagesResponse
	urlEndpoint := messagesEndpoint
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

// CreateMessage will allow you to create a new message under a specific conversation https://www.reamaze.com/api/post_messages
func (c *Client) CreateMessage(slug string, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	var response *CreateMessageResponse
	emptyReq := &CreateMessageRequest{}
	// checking if we don't have empty request
	if reflect.DeepEqual(req, emptyReq) {
		return nil, errors.New("CreateMessage incorrect request, CreateMessageRequest is empty")
	}
	// checking if slug is set
	if len(slug) == 0 {
		return nil, errors.New("CreateMessage slug cannot be empty, please provide slug as argument")
	}

	urlEndpoint := conversationsEndpoint + "/" + url.PathEscape(slug) + "/messages"

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
