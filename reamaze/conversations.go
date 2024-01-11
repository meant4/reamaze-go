package reamaze

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// CreateConversation is creating new conversation in reamaze
func (c *Client) CreateConversation(req *CreateConversationRequest) (*CreateConversationResponse, error) {
	var response *CreateConversationResponse
	// Checking if we have all required fields the rest is optional
	if len(req.Conversation.Category) == 0 || len(req.Conversation.Message.Body) == 0 || len(req.Conversation.User.Email) == 0 {
		return nil, errors.New("missing one of the required field: conversation.category, conversation.message.body or conversation.user.email")
	}

	data, _ := json.Marshal(req)
	resp, err := c.reamazeRequest(http.MethodPost, conversationsEndpoint, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {

		return nil, err
	}
	return response, nil
}

// UpdateConversation updates a conversation
func (c *Client) UpdateConversation() {
	fmt.Printf("Method UpdateConversation not implemented")
}

// GetConversations retrieve conversations for the Brand
func (c *Client) GetConversations(o ...Option) (*GetConversationsResponse, error) {
	var response *GetConversationsResponse
	settings, _ := newSettings(o)
	urlEndpoint := conversationsEndpoint + settings.GetQuery()
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

// GetConversation retrieve a specific conversation
func (c *Client) GetConversation(slug string) (*GetConversationResponse, error) {
	var response *GetConversationResponse
	if len(slug) == 0 {
		return nil, errors.New("slug parameter cannot be empty")
	}
	urlEndpoint := conversationsEndpoint + "/" + url.QueryEscape(slug)
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
