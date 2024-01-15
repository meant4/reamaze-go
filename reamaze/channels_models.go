package reamaze

import "time"

const channelsEndpoint string = "/api/v1/channels"

type ReamazeChannelType int

const (
	ReamazeChannelEmail             ReamazeChannelType = 1
	ReamazeChannelTwitter           ReamazeChannelType = 2
	ReamazeChannelFacebook          ReamazeChannelType = 3
	ReamazeChannelChat              ReamazeChannelType = 6
	ReamazeChannelInstagram         ReamazeChannelType = 8
	ReamazeChannelSMS               ReamazeChannelType = 9
	ReamazeChannelVoice             ReamazeChannelType = 10
	ReamazeChannelFacebookMessanger ReamazeChannelType = 12
	ReamazeChannelFacebookLead      ReamazeChannelType = 13
	ReamazeChannelInstagramAd       ReamazeChannelType = 14
	ReamazeChannelWhatsApp          ReamazeChannelType = 15
	ReamazeChannelInstagramDM       ReamazeChannelType = 16
)

func (w ReamazeChannelType) Name() string {
	switch int(w) {
	case 1:
		return "ReamazeChannelEmail"
	case 2:
		return "ReamazeChannelTwitter"
	case 3:
		return "ReamazeChannelFacebook"
	case 6:
		return "ReamazeChannelChat"
	case 8:
		return "ReamazeChannelInstagram"
	case 9:
		return "ReamazeChannelSMS"
	case 10:
		return "ReamazeChannelVoice"
	case 12:
		return "ReamazeChannelFacebookMessanger"
	case 13:
		return "ReamazeChannelFacebookLead"
	case 14:
		return "ReamazeChannelInstagramAd"
	case 15:
		return "ReamazeChannelWhatsApp"
	case 16:
		return "ReamazeChannelInstagramDM"
	default:
		return "Undefined"
	}
}

type ReamazeChannelVisibility int

const (
	ReamazeChannelVisibilityPrivate ReamazeChannelVisibility = 0
	ReamazeChannelVisibilityPublic  ReamazeChannelVisibility = 1
)

func (w ReamazeChannelVisibility) Name() string {
	switch int(w) {
	case 0:
		return "ReamazeChannelVisibilityPrivate"
	case 1:
		return "ReamazeChannelVisibilityPublic"
	default:
		return "Undefined"
	}
}

// The settings_reply_from_name value determines the whether the "From" name for replies is set to the channel name ("channel"), the brand name ("brand"), or the responding staff user's name ("staff")
type GetChannelsResponse struct {
	TotalCount int `json:"total_count"`
	Channels   []struct {
		Name                  string                   `json:"name"`
		Slug                  string                   `json:"slug"`
		Email                 string                   `json:"email"`
		CreatedAt             time.Time                `json:"created_at"`
		UpdatedAt             time.Time                `json:"updated_at"`
		Channel               ReamazeChannelType       `json:"channel"`
		Visibility            ReamazeChannelVisibility `json:"visibility"`
		SpamFilterEnabled     bool                     `json:"spam_filter_enabled"`
		ReplyFromOrigin       bool                     `json:"reply_from_origin"`
		Verified              bool                     `json:"verified"`
		VerificationEmail     string                   `json:"verification_email"`
		LastVerified          time.Time                `json:"last_verified"`
		SettingsReplyFromName string                   `json:"settings_reply_from_name"`
		SettingsSignature     string                   `json:"settings_signature"`
		Brand                 struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"brand"`
	} `json:"channels"`
}

// The settings_reply_from_name value determines the whether the "From" name for replies is set to the channel name ("channel"), the brand name ("brand"), or the responding staff user's name ("staff")
type GetChannelResponse struct {
	Name                  string                   `json:"name"`
	Slug                  string                   `json:"slug"`
	Email                 string                   `json:"email"`
	CreatedAt             time.Time                `json:"created_at"`
	UpdatedAt             time.Time                `json:"updated_at"`
	Channel               ReamazeChannelType       `json:"channel"`
	Visibility            ReamazeChannelVisibility `json:"visibility"`
	SpamFilterEnabled     bool                     `json:"spam_filter_enabled"`
	ReplyFromOrigin       bool                     `json:"reply_from_origin"`
	Verified              bool                     `json:"verified"`
	VerificationEmail     string                   `json:"verification_email"`
	LastVerified          time.Time                `json:"last_verified"`
	SettingsReplyFromName string                   `json:"settings_reply_from_name"`
	SettingsSignature     string                   `json:"settings_signature"`
	Brand                 struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"brand"`
}
