package reamaze

import (
	"encoding/json"
	"net/http"
)

// GetReportsVolume returns a daily volume count
// The start and end dates of the report will default to the last 30 days. Time frames can be no smaller than 1 day and no larger than 1 year.
// https://www.reamaze.com/api/get_reports_volume
func (c *Client) GetReportsVolume(o ...ReportsOption) (*GetReportsVolumeResponse, error) {
	var response *GetReportsVolumeResponse
	settings, _ := newReportsSettings(o)
	urlEndpoint := reportsEndpoint + "/volume" + settings.GetQuery()

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

// GetReportsResponseTime Returns a daily response time metric and a response times summary object. Response times are reported in seconds.
// https://www.reamaze.com/api/get_reports_response_time
func (c *Client) GetReportsResponseTime(o ...ReportsOption) (*GetReportsResponseTimeRespone, error) {
	var response *GetReportsResponseTimeRespone
	settings, _ := newReportsSettings(o)
	urlEndpoint := reportsEndpoint + "/response_time" + settings.GetQuery()
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

// GetReportsStaffResponse returns a staff report summarizing staff metrics
// The start and end dates of the report will default to the last 30 days. Time frames can be no smaller than 1 day and no larger than 1 year.
// https://www.reamaze.com/api/get_reports_staff
func (c *Client) GetReportsStaff(o ...ReportsOption) (*GetReportsStaffResponse, error) {
	var response *GetReportsStaffResponse
	settings, _ := newReportsSettings(o)
	urlEndpoint := reportsEndpoint + "/staff" + settings.GetQuery()
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

// GetReportsTags returns a tag report summarizing tag usage
// The start and end dates of the report will default to the last 30 days. Time frames can be no smaller than 1 day and no larger than 1 year.
// https://www.reamaze.com/api/get_reports_tags
func (c *Client) GetReportsTags(o ...ReportsOption) (*GetReportsTagsResponse, error) {
	var response *GetReportsTagsResponse
	settings, _ := newReportsSettings(o)
	urlEndpoint := reportsEndpoint + "/tags" + settings.GetQuery()
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

// GetReportsChannelSummary returns a channel report summarizing channel metrics
// The start and end dates of the report will default to the last 30 days. Time frames can be no smaller than 1 day and no larger than 1 year.
// https://www.reamaze.com/api/get_reports_channel_summary
func (c *Client) GetReportsChannelSummary(o ...ReportsOption) (*GetReportsChannelSummaryResponse, error) {
	var response *GetReportsChannelSummaryResponse
	settings, _ := newReportsSettings(o)

	urlEndpoint := reportsEndpoint + "/channel_summary" + settings.GetQuery()
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
