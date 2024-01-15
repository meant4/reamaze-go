package reamaze

import "time"

const staffEndpoint string = "/api/v1/staff"

type GetStaffResponse struct {
	PageSize   int `json:"page_size,omitempty"`
	PageCount  int `json:"page_count,omitempty"`
	TotalCount int `json:"total_count,omitempty"`
	Staff      []struct {
		Name              string    `json:"name,omitempty"`
		CreatedAt         time.Time `json:"created_at,omitempty"`
		Email             string    `json:"email,omitempty"`
		DisplayName       string    `json:"display_name,omitempty"`
		NotificationEmail string    `json:"notification_email,omitempty"`
		Role              struct {
			ID          any    `json:"id,omitempty"`
			Name        string `json:"name,omitempty"`
			Description any    `json:"description,omitempty"`
			Admin       bool   `json:"admin?,omitempty"`
			Default     bool   `json:"default?,omitempty"`
			Permissions struct {
				ManageStaff                      bool  `json:"manage_staff?,omitempty"`
				ManageStaffRole                  bool  `json:"manage_staff_role?,omitempty"`
				ManageDepartments                bool  `json:"manage_departments?,omitempty"`
				ViewStaff                        bool  `json:"view_staff?,omitempty"`
				ManageSubscriptions              bool  `json:"manage_subscriptions?,omitempty"`
				ManageInvoiceEmail               bool  `json:"manage_invoice_email?,omitempty"`
				ManageKb                         bool  `json:"manage_kb?,omitempty"`
				ManageAccount                    bool  `json:"manage_account?,omitempty"`
				ManageResponseTemplates          bool  `json:"manage_response_templates?,omitempty"`
				ManagePersonalResponseTemplates  bool  `json:"manage_personal_response_templates?,omitempty"`
				ManageWorkflows                  bool  `json:"manage_workflows?,omitempty"`
				ManageChatbots                   bool  `json:"manage_chatbots?,omitempty"`
				ManagePushCampaigns              bool  `json:"manage_push_campaigns?,omitempty"`
				ManageWebsiteIntegrations        bool  `json:"manage_website_integrations?,omitempty"`
				ManageDeveloperSettings          bool  `json:"manage_developer_settings?,omitempty"`
				ManageAssignments                bool  `json:"manage_assignments?,omitempty"`
				ManageIncidents                  bool  `json:"manage_incidents?,omitempty"`
				ManageNotes                      bool  `json:"manage_notes?,omitempty"`
				DeleteConversations              bool  `json:"delete_conversations?,omitempty"`
				AccessVoice                      bool  `json:"access_voice?,omitempty"`
				AccessVideoCall                  bool  `json:"access_video_call?,omitempty"`
				AccessAiFeatures                 bool  `json:"access_ai_features?,omitempty"`
				AccessWebhookSubscriptionsAPI    bool  `json:"access_webhook_subscriptions_api?,omitempty"`
				ManageTags                       bool  `json:"manage_tags?,omitempty"`
				AccessChat                       bool  `json:"access_chat?,omitempty"`
				AccessLiveView                   bool  `json:"access_live_view?,omitempty"`
				AccessReports                    bool  `json:"access_reports?,omitempty"`
				AccessStaffReports               bool  `json:"access_staff_reports?,omitempty"`
				ReplyToCustomers                 bool  `json:"reply_to_customers?,omitempty"`
				EditCustomers                    bool  `json:"edit_customers?,omitempty"`
				RestrictChannels                 bool  `json:"restrict_channels?,omitempty"`
				MoveAcrossRestrictedChannels     bool  `json:"move_across_restricted_channels?,omitempty"`
				AssignAcrossRestrictedChannels   bool  `json:"assign_across_restricted_channels?,omitempty"`
				ViewReportsForRestrictedChannels bool  `json:"view_reports_for_restricted_channels?,omitempty"`
				VisibleChannelIds                []any `json:"visible_channel_ids,omitempty"`
				ViewAllContacts                  bool  `json:"view_all_contacts?,omitempty"`
				ExportContacts                   bool  `json:"export_contacts?,omitempty"`
				MaxChats                         int   `json:"max_chats,omitempty"`
				BigcommerceAccess                struct {
					Access         bool `json:"access,omitempty"`
					ProcessRefunds bool `json:"process_refunds,omitempty"`
				} `json:"bigcommerce_access,omitempty"`
				LoyaltylionAccess struct {
					Access bool `json:"access,omitempty"`
					Edit   bool `json:"edit,omitempty"`
				} `json:"loyaltylion_access,omitempty"`
				PipedriveAccess struct {
					Access      bool `json:"access,omitempty"`
					ManageDeals bool `json:"manage_deals,omitempty"`
				} `json:"pipedrive_access,omitempty"`
				ShopifyAccess struct {
					Access             bool `json:"access,omitempty"`
					EditDetails        bool `json:"edit_details,omitempty"`
					ProcessRefunds     bool `json:"process_refunds,omitempty"`
					ProcessCancels     bool `json:"process_cancels,omitempty"`
					ManageDraftOrders  bool `json:"manage_draft_orders,omitempty"`
					ManageFulfillments bool `json:"manage_fulfillments,omitempty"`
				} `json:"shopify_access,omitempty"`
				StripeAccess struct {
					Access              bool `json:"access,omitempty"`
					ProcessRefunds      bool `json:"process_refunds,omitempty"`
					CancelSubscriptions bool `json:"cancel_subscriptions,omitempty"`
				} `json:"stripe_access,omitempty"`
				WoocommerceAccess struct {
					Access         bool `json:"access,omitempty"`
					ProcessRefunds bool `json:"process_refunds,omitempty"`
				} `json:"woocommerce_access,omitempty"`
				YotpoAccess struct {
					Access bool `json:"access,omitempty"`
				} `json:"yotpo_access,omitempty"`
				GbmAccess struct {
					Access bool `json:"access,omitempty"`
				} `json:"gbm_access,omitempty"`
				WixAccess struct {
					Access bool `json:"access,omitempty"`
				} `json:"wix_access,omitempty"`
			} `json:"permissions,omitempty"`
		} `json:"role,omitempty"`
	} `json:"staff,omitempty"`
}
