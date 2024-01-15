package reamaze

import "time"

const messagesEndpoint string = "/api/v1/messages"

type ReamazeVisibility int

const (
	ReamazeVisibilityRegular      ReamazeVisibility = 0
	ReamazeVisibilityInternalNote ReamazeVisibility = 1
)

type GetMessagesResponse struct {
	PageSize   int `json:"page_size"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
	Messages   []struct {
		Visibility   int       `json:"visibility"`
		Origin       int       `json:"origin"`
		CreatedAt    time.Time `json:"created_at"`
		Conversation struct {
			Subject   string    `json:"subject"`
			Slug      string    `json:"slug"`
			CreatedAt time.Time `json:"created_at"`
			Category  struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Slug    string `json:"slug"`
				Email   string `json:"email"`
				Channel int    `json:"channel"`
			} `json:"category"`
			Followers []any `json:"followers"`
		} `json:"conversation"`
		Attachments []struct {
			ThumbURL        string `json:"thumb_url"`
			URL             string `json:"url"`
			Image           bool   `json:"image?"`
			FileContentType string `json:"file_content_type"`
			FileFileName    string `json:"file_file_name"`
			FileFileSize    int    `json:"file_file_size"`
		} `json:"attachments"`
		Body             string `json:"body"`
		DirectRecipients []any  `json:"direct_recipients"`
		Recipients       []any  `json:"recipients"`
		User             struct {
			Name   string `json:"name"`
			Email  string `json:"email"`
			Mobile any    `json:"mobile"`
			Staff  bool   `json:"staff?"`
		} `json:"user,omitempty"`
		Meta struct {
			Subject  string `json:"Subject"`
			Language struct {
				Name     string `json:"name"`
				Code     string `json:"code"`
				Reliable bool   `json:"reliable"`
			} `json:"language"`
		} `json:"meta,omitempty"`
	} `json:"messages"`
}

type CreateMessageResponse struct {
	Body       string `json:"body"`
	Visibility int    `json:"visibility"`
	CreatedAt  string `json:"created_at"`
	OriginID   string `json:"origin_id"`
	User       struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
	Conversation struct {
		Subject   string `json:"subject"`
		Slug      string `json:"slug"`
		CreatedAt string `json:"created_at"`
		Category  struct {
			Name    string `json:"name"`
			Slug    string `json:"slug"`
			Email   string `json:"email"`
			Channel int    `json:"channel"`
		} `json:"category"`
	} `json:"conversation"`
}

type CreateMessageRequest struct {
	Message struct {
		Body       string            `json:"body"`
		Visibility ReamazeVisibility `json:"visibility,omitempty"`
		OriginID   string            `json:"origin_id,omitempty"`
		User       struct {
			Name  string `json:"name,omitempty"`
			Email string `json:"email,omitempty"`
		} `json:"user,omitempty"`
		SupressNotification bool     `json:"suppress_notifications,omitempty"` // You can optionally pass in a message[suppress_notifications] boolean attribute with a value of true to prevent Reamaze from sending any email (or integration) notifications related to this message.
		SupressAutoresolve  bool     `json:"suppress_autoresolve,omitempty"`
		Attachment          string   `json:"attachment,omitempty"`
		Attachments         []string `json:"attachments,omitempty"`
	} `json:"message"`
}
