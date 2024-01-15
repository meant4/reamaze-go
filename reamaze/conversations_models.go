package reamaze

import (
	"fmt"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ReamazeStatus int
type ReamazeFilter string
type ReamazeFor string
type ReamazeForID string
type ReamazeSort string
type ReamazeStartDate time.Time
type ReamazeEndDate time.Time
type ReamazePage int
type ReamazeCategory string
type ReamazeTags []string
type ReamazeData map[string]string

const conversationsEndpoint string = "/api/v1/conversations"

const (
	ReamazeFilterArchived   ReamazeFilter = "archived"
	ReamazeFilterOpen       ReamazeFilter = "open"
	ReamazeFilterUnassigned ReamazeFilter = "unassigned"
	ReamazeFilterAll        ReamazeFilter = "all"
	ReamazeSortUpdated      ReamazeSort   = "updated"
	ReamazeSortChanged      ReamazeSort   = "changed"
	ReamazeSortCreatedAt    ReamazeSort   = "create_at"
)

type Option interface {
	Apply(*ReamazeOptions)
}

func (w ReamazeFilter) Apply(o *ReamazeOptions) {
	if len(w) > 0 {
		o.ReamazeFilter = "filter=" + url.QueryEscape(string(w))
	}
}
func (w ReamazeFor) Apply(o *ReamazeOptions) {
	_, err := mail.ParseAddress(string(w))
	if err == nil {
		o.ReamazeFor = "for=" + url.QueryEscape(string(w))
	}
}
func (w ReamazeForID) Apply(o *ReamazeOptions) {
	if len(w) > 0 {
		o.ReamazeForID = "for_id=" + url.QueryEscape(string(w))
	}
}

func (w ReamazeSort) Apply(o *ReamazeOptions) {
	if len(w) > 0 {
		o.ReamazeSort = "sort=" + url.QueryEscape(string(w))
	}
}

func (w ReamazeData) Apply(o *ReamazeOptions) {
	if len(w) > 0 {
		var parts []string
		for k, v := range w {
			partsString := "data[" + url.QueryEscape(k) + "]=" + url.QueryEscape(v)
			parts = append(parts, partsString)
		}
		o.ReamazeData = strings.Join(parts, "&")
	}
}

func (w ReamazeCategory) Apply(o *ReamazeOptions) {
	if len(w) > 0 {
		o.ReamazeCategory = "category=" + url.QueryEscape(string(w))
	}
}

func (w ReamazeEndDate) Apply(o *ReamazeOptions) {
	date := time.Time(w)
	if date.Year() > 1 && date.Month() > 0 && date.Day() > 0 {
		o.ReamazeEndDate = "end_date=" + fmt.Sprintf("%04d", date.Year()) + "-" + fmt.Sprintf("%02d", date.Month()) + "-" + fmt.Sprintf("%02d", date.Day())
	}
}

func (w ReamazeStartDate) Apply(o *ReamazeOptions) {
	date := time.Time(w)
	if date.Year() > 1 && date.Month() > 0 && date.Day() > 0 {
		o.ReamazeStartDate = "start_date=" + fmt.Sprintf("%04d", date.Year()) + "-" + fmt.Sprintf("%02d", date.Month()) + "-" + fmt.Sprintf("%02d", date.Day())
	}
}

func (w ReamazePage) Apply(o *ReamazeOptions) {
	if w > 0 {
		o.ReamazePage = "page=" + strconv.Itoa(int(w))
	}
}

func (w ReamazeTags) Apply(o *ReamazeOptions) {
	if len(w) > 0 {
		o.ReamazeTag = "tag=" + url.QueryEscape(strings.Join(w, ","))
	}
}

func WithFilter(filter ReamazeFilter) ReamazeFilter {
	return filter
}

func WithFor(email string) ReamazeFor {
	return ReamazeFor(email)
}

func WithForID(id string) ReamazeForID {
	return ReamazeForID(id)
}

func WithSort(sort ReamazeSort) ReamazeSort {
	return sort
}

func WithCategory(category string) ReamazeCategory {
	return ReamazeCategory(category)
}

func WithData(data map[string]string) ReamazeData {
	var reamazeDataMap = make(ReamazeData)
	for k, v := range data {
		reamazeDataMap[k] = v
	}
	return reamazeDataMap
}

func WithEndDate(year int, month int, day int) ReamazeEndDate {
	endDate := ReamazeEndDate{}
	if year > 0 && month > 0 && day > 0 {
		endDate = ReamazeEndDate(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	}
	return endDate
}

func WithStartDate(year, month, day int) ReamazeStartDate {
	startDate := ReamazeStartDate{}
	if year > 0 && month > 0 && day > 0 {

		startDate = ReamazeStartDate(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	}
	return startDate
}

func WithPage(page int) ReamazePage {
	return ReamazePage(page)
}

func WithTags(w ...string) ReamazeTags {
	return ReamazeTags(w)
}

func newSettings(opts []Option) (*ReamazeOptions, error) {
	var o ReamazeOptions
	for _, opt := range opts {
		opt.Apply(&o)
	}
	return &o, nil
}

type ReamazeOptions struct {
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

func (r ReamazeOptions) GetQuery() string {
	output := ""
	var queryParams []string
	// checking if we have filter
	if len(r.ReamazeFilter) > 0 {
		queryParams = append(queryParams, r.ReamazeFilter)
	}
	// checking if we want message with specific email address
	if len(r.ReamazeFor) > 0 {
		queryParams = append(queryParams, r.ReamazeFor)
	}
	// checking if for_id has been set
	if len(r.ReamazeForID) > 0 {
		queryParams = append(queryParams, r.ReamazeForID)
	}
	// checking if sort param has been set
	if len(r.ReamazeSort) > 0 {
		queryParams = append(queryParams, r.ReamazeSort)
	}
	// checking if start_date is set
	if len(r.ReamazeStartDate) > 0 {
		queryParams = append(queryParams, r.ReamazeStartDate)
	}
	// checking if end_date is set
	if len(r.ReamazeStartDate) > 0 {
		queryParams = append(queryParams, r.ReamazeEndDate)
	}
	// checking if page is set
	if len(r.ReamazePage) > 0 {
		queryParams = append(queryParams, r.ReamazePage)
	}
	// checking if category is set
	if len(r.ReamazeCategory) > 0 {
		queryParams = append(queryParams, r.ReamazeCategory)
	}
	// checking if we have any tags set
	if len(r.ReamazeTag) > 0 {
		queryParams = append(queryParams, r.ReamazeTag)
	}
	// checking if we have data set
	if len(r.ReamazeData) > 0 {
		queryParams = append(queryParams, r.ReamazeData)
	}
	output = strings.Join(queryParams, "&")
	if len(output) > 0 {
		output = "?" + output
	}
	return output
}

const (
	ReamazeStatusUnresolved ReamazeStatus = iota
	ReamazeStatusPending
	ReamazeStatusResolved
	ReamazeStatusSpam
	ReamazeStatusArchived
	ReamazeStatusOnHold
	ReamazeStatusAutoResolved
	ReamazeStatusChatbotAssigned
	ReamazeStatusChatbotResolved
)

type CreateConversationResponse struct {
	Subject   string    `json:"subject"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Origin    int       `json:"origin"`
	Data      struct {
		Nda       string `json:"NDA"`
		Firstname string `json:"Firstname"`
		Lastname  string `json:"Lastname"`
	} `json:"data"`
	HoldUntil any `json:"hold_until"`
	Author    struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Data         string `json:"data"`
		Email        string `json:"email"`
		Twitter      string `json:"twitter"`
		Facebook     string `json:"facebook"`
		Instagram    string `json:"instagram"`
		Mobile       string `json:"mobile"`
		FriendlyName string `json:"friendly_name"`
		DisplayName  string `json:"display_name"`
	} `json:"author"`
	Assignee       any      `json:"assignee"`
	PermaURL       string   `json:"perma_url"`
	TagList        []string `json:"tag_list"`
	Status         int      `json:"status"`
	DisplaySubject string   `json:"display_subject"`
	Category       struct {
		Name                     string `json:"name"`
		Slug                     string `json:"slug"`
		Email                    string `json:"email"`
		Channel                  int    `json:"channel"`
		SettingsDisplayHTMLEmail string `json:"settings_display_html_email"`
	} `json:"category"`
	LastCustomerMessage struct {
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"last_customer_message"`
	Followers []struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		Data         interface{} `json:"data"`
		Email        string      `json:"email"`
		Twitter      string      `json:"twitter"`
		Facebook     string      `json:"facebook"`
		Instagram    string      `json:"instagram"`
		Mobile       string      `json:"mobile"`
		FriendlyName string      `json:"friendly_name"`
		DisplayName  string      `json:"display_name"`
		Staff        bool        `json:"staff?"`
		Customer     bool        `json:"customer?"`
		Bot          bool        `json:"bot?"`
	} `json:"followers"`
	Message struct {
		Body string `json:"body"`
	} `json:"message"`
	ReadOnly     bool `json:"readOnly"`
	MessageCount int  `json:"message_count"`
}

type CreateConversationRequest struct {
	Conversation struct {
		Subject             string        `json:"subject,omitempty"`
		Category            string        `json:"category,omitempty"`
		TagList             []string      `json:"tag_list,omitempty"`
		Status              ReamazeStatus `json:"status,omitempty"`
		SupressNotification bool          `json:"suppress_notifications,omitempty"` // You can optionally pass in a message[suppress_notifications] boolean attribute with a value of true to prevent Reamaze from sending any email (or integration) notifications related to this message.
		SupressAutoresolve  bool          `json:"suppress_autoresolve,omitempty"`   // You can optionally pass in a message[suppress_autoresolve] boolean attribute with a value of true to prevent Reamaze from marking the conversation as resolved when message[user] is a staff user.
		Data                interface{}   `json:"data,omitempty"`
		Message             struct {
			Body        string   `json:"body,omitempty"`
			Attachment  string   `json:"attachment,omitempty"`
			Attachments []string `json:"attachments,omitempty"`
		} `json:"message,omitempty"`
		User struct {
			Name  string      `json:"name,omitempty"`
			Email string      `json:"email,omitempty"`
			Data  interface{} `json:"data,omitempty"`
		} `json:"user,omitempty"`
	} `json:"conversation,omitempty"`
}
type UpdateConversationRequest struct {
	Conversation struct {
		TagList  []string      `json:"tag_list,omitempty"`
		Status   ReamazeStatus `json:"status,omitempty"`
		Data     interface{}   `json:"data,omitempty"`
		Assignee string        `json:"assignee,omitempty"`
		Category string        `json:"category,omitempty"`
		Brand    string        `json:"brand,omitempty"`
	} `json:"conversation,omitempty"`
}

type GetConversationResponse struct {
	Subject   string    `json:"subject,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Origin    int       `json:"origin,omitempty"`
	Data      struct {
		LastName  string `json:"last_name,omitempty"`
		FirstName string `json:"first_name,omitempty"`
	} `json:"data,omitempty"`
	HoldUntil any `json:"hold_until,omitempty"`
	Author    struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Data struct {
			LastName       string `json:"last_name,omitempty"`
			FirstName      string `json:"first_name,omitempty"`
			JobApplication string `json:"job_application,omitempty"`
		} `json:"data"`
		Email        string `json:"email,omitempty"`
		Twitter      string `json:"twitter,omitempty"`
		Facebook     string `json:"facebook,omitempty"`
		Instagram    string `json:"instagram,omitempty"`
		Mobile       string `json:"mobile,omitempty"`
		FriendlyName string `json:"friendly_name,omitempty"`
		DisplayName  string `json:"display_name,omitempty"`
	} `json:"author"`
	Assignee       any      `json:"assignee,omitempty"`
	PermaURL       string   `json:"perma_url,omitempty"`
	TagList        []string `json:"tag_list,omitempty"`
	Status         int      `json:"status,omitempty"`
	DisplaySubject string   `json:"display_subject,omitempty"`
	Category       struct {
		Name                     string `json:"name,omitempty"`
		Slug                     string `json:"slug,omitempty"`
		Email                    string `json:"email,omitempty"`
		Channel                  int    `json:"channel,omitempty"`
		SettingsDisplayHTMLEmail string `json:"settings_display_html_email,omitempty"`
	} `json:"category,omitempty"`
	LastCustomerMessage struct {
		Body      string    `json:"body,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	} `json:"last_customer_message,omitempty"`
	Followers []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Data struct {
			LastName       string `json:"last_name,omitempty"`
			FirstName      string `json:"first_name,omitempty"`
			JobApplication string `json:"job_application,omitempty"`
		} `json:"data"`
		Email        string `json:"email,omitempty"`
		Twitter      string `json:"twitter,omitempty"`
		Facebook     string `json:"facebook,omitempty"`
		Instagram    string `json:"instagram,omitempty"`
		Mobile       string `json:"mobile,omitempty"`
		FriendlyName string `json:"friendly_name,omitempty"`
		DisplayName  string `json:"display_name,omitempty"`
		Staff        bool   `json:"staff?,omitempty"`
		Customer     bool   `json:"customer?,omitempty"`
		Bot          bool   `json:"bot?,omitempty"`
	} `json:"followers,omitempty"`
	Message struct {
		Body string `json:"body,omitempty"`
	} `json:"message"`
	ReadOnly     bool `json:"readOnly,omitempty"`
	MessageCount int  `json:"message_count,omitempty"`
}
type GetConversationsResponse struct {
	PageSize      int                       `json:"page_size,omitempty"`
	PageCount     int                       `json:"page_count,omitempty"`
	TotalCount    int                       `json:"total_count,omitempty"`
	Conversations []GetConversationResponse `json:"conversations,omitempty"`
}
