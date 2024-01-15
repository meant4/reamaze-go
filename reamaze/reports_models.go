package reamaze

import (
	"fmt"
	"strings"
	"time"
)

const reportsEndpoint string = "/api/v1/reports"

type ReamazeReportsStartDate time.Time
type ReamazeReportsEndDate time.Time

type ReportsOption interface {
	Apply(*ReamazeReportOptions)
}

func (w ReamazeReportsEndDate) Apply(o *ReamazeReportOptions) {
	date := time.Time(w)
	if date.Year() > 1 && date.Month() > 0 && date.Day() > 0 {
		o.ReamazeEndDate = "end_date=" + fmt.Sprintf("%04d", date.Year()) + "-" + fmt.Sprintf("%02d", date.Month()) + "-" + fmt.Sprintf("%02d", date.Day())
	}
}

func (w ReamazeReportsStartDate) Apply(o *ReamazeReportOptions) {
	date := time.Time(w)
	if date.Year() > 1 && date.Month() > 0 && date.Day() > 0 {
		o.ReamazeStartDate = "start_date=" + fmt.Sprintf("%04d", date.Year()) + "-" + fmt.Sprintf("%02d", date.Month()) + "-" + fmt.Sprintf("%02d", date.Day())
	}
}

func WithReportsEndDate(year int, month int, day int) ReamazeReportsEndDate {
	endDate := ReamazeReportsEndDate{}
	if year > 0 && month > 0 && day > 0 {
		endDate = ReamazeReportsEndDate(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	}
	return endDate
}

func WithReportsStartDate(year, month, day int) ReamazeReportsStartDate {
	startDate := ReamazeReportsStartDate{}
	if year > 0 && month > 0 && day > 0 {

		startDate = ReamazeReportsStartDate(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	}
	return startDate
}
func (r ReamazeReportOptions) GetQuery() string {
	output := ""
	var queryParams []string
	// checking if start_date is set
	if len(r.ReamazeStartDate) > 0 {
		queryParams = append(queryParams, r.ReamazeStartDate)
	}
	// checking if end_date is set
	if len(r.ReamazeStartDate) > 0 {
		queryParams = append(queryParams, r.ReamazeEndDate)
	}

	output = strings.Join(queryParams, "&")
	if len(output) > 0 {
		output = "?" + output
	}
	return output
}
func newReportsSettings(opts []ReportsOption) (*ReamazeReportOptions, error) {
	var o ReamazeReportOptions
	for _, opt := range opts {
		opt.Apply(&o)
	}
	return &o, nil
}

type ReamazeReportOptions struct {
	ReamazeStartDate string
	ReamazeEndDate   string
}

type GetReportsVolumeResponse struct {
	ConversationCounts map[string]int `json:"conversation_counts,omitempty"`
	StartDate          string         `json:"start_date,omitempty"`
	EndDate            string         `json:"end_date,omitempty"`
}

type GetReportsResponseTimeRespone struct {
	ResponseTimes map[string]float64 `json:"response_times,omitempty"`
	Summary       struct {
		Averages struct {
			InRange   float64 `json:"in_range,omitempty"`
			ThisMonth float64 `json:"this_month,omitempty"`
			ThisWeek  float64 `json:"this_week,omitempty"`
		} `json:"averages,omitempty"`
		Trends struct {
			Last30Days struct {
				Average    float64 `json:"average,omitempty"`
				ChangeRate any     `json:"change_rate,omitempty"`
			} `json:"last_30_days,omitempty"`
			Last7Days struct {
				Average    float64 `json:"average,omitempty"`
				ChangeRate string  `json:"change_rate,omitempty"`
			} `json:"last_7_days,omitempty"`
		} `json:"trends,omitempty"`
		Ratio struct {
			Under1Hour float64 `json:"under_1_hour,omitempty"`
			Under1Day  float64 `json:"under_1_day,omitempty"`
			Under1Week float64 `json:"under_1_week,omitempty"`
		} `json:"ratio,omitempty"`
	} `json:"summary,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type GetReportsStaffResponse struct {
	Report struct {
		RelacjeHiLIVE struct {
			ResponseCount       int            `json:"response_count,omitempty"`
			ArchivedCount       int            `json:"archived_count,omitempty"`
			ResolvedCount       int            `json:"resolved_count,omitempty"`
			SatisfactionAverage any            `json:"satisfaction_average,omitempty"`
			AppreciationsCount  any            `json:"appreciations_count,omitempty"`
			ResponsesTrend      map[string]int `json:"responses_trend,omitempty"`
			ResponseTimeSeconds float64        `json:"response_time_seconds,omitempty"`
		} `json:"Relacje HiLIVE,omitempty"`
	} `json:"report,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type GetReportsTagsResponse struct {
	Tags      map[string]int `json:"tags,omitempty"`
	StartDate string         `json:"start_date,omitempty"`
	EndDate   string         `json:"end_date,omitempty"`
}

type GetReportsChannelSummaryResponse struct {
	Channels map[string]struct {
		Category struct {
			ID              int    `json:"id,omitempty"`
			Name            string `json:"name,omitempty"`
			ChannelType     int    `json:"channel_type,omitempty"`
			ChannelTypeName string `json:"channel_type_name,omitempty"`
		} `json:"category,omitempty"`
		Brand struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"brand,omitempty"`
		StaffResponses             int     `json:"staff_responses,omitempty"`
		CustomerResponses          int     `json:"customer_responses,omitempty"`
		AverageResponseTimeSeconds any     `json:"average_response_time_seconds,omitempty"`
		Appreciations              int     `json:"appreciations,omitempty"`
		ActiveConversations        int     `json:"active_conversations,omitempty"`
		ResolvedConversations      int     `json:"resolved_conversations,omitempty"`
		ArchivedConversations      int     `json:"archived_conversations,omitempty"`
		AverageSatisfactionRating  float64 `json:"average_satisfaction_rating,omitempty"`
		AverageThreadSize          float64 `json:"average_thread_size,omitempty"`
	} `json:"channels,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}
