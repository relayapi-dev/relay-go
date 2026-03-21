// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// AnalyticsYoutubeService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsYoutubeService] method instead.
type AnalyticsYoutubeService struct {
	Options []option.RequestOption
}

// NewAnalyticsYoutubeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsYoutubeService(opts ...option.RequestOption) (r *AnalyticsYoutubeService) {
	r = &AnalyticsYoutubeService{}
	r.Options = opts
	return
}

// Get YouTube daily views and watch time
func (r *AnalyticsYoutubeService) GetDailyViews(ctx context.Context, query AnalyticsYoutubeGetDailyViewsParams, opts ...option.RequestOption) (res *AnalyticsYoutubeGetDailyViewsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics/youtube/daily-views"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AnalyticsYoutubeGetDailyViewsResponse struct {
	Data []AnalyticsYoutubeGetDailyViewsResponseData `json:"data" api:"required"`
	JSON analyticsYoutubeGetDailyViewsResponseJSON   `json:"-"`
}

// analyticsYoutubeGetDailyViewsResponseJSON contains the JSON metadata for the
// struct [AnalyticsYoutubeGetDailyViewsResponse]
type analyticsYoutubeGetDailyViewsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsYoutubeGetDailyViewsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsYoutubeGetDailyViewsResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsYoutubeGetDailyViewsResponseData struct {
	// Date (YYYY-MM-DD)
	Date string `json:"date" api:"required"`
	// Net subscribers gained
	SubscribersGained float64 `json:"subscribers_gained" api:"required"`
	// Total views
	Views float64 `json:"views" api:"required"`
	// Watch time in minutes
	WatchTimeMinutes float64                                       `json:"watch_time_minutes" api:"required"`
	JSON             analyticsYoutubeGetDailyViewsResponseDataJSON `json:"-"`
}

// analyticsYoutubeGetDailyViewsResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsYoutubeGetDailyViewsResponseData]
type analyticsYoutubeGetDailyViewsResponseDataJSON struct {
	Date              apijson.Field
	SubscribersGained apijson.Field
	Views             apijson.Field
	WatchTimeMinutes  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AnalyticsYoutubeGetDailyViewsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsYoutubeGetDailyViewsResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsYoutubeGetDailyViewsParams struct {
	// YouTube account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
	// Start date (ISO 8601)
	FromDate param.Field[string] `query:"from_date"`
	// End date (ISO 8601)
	ToDate param.Field[string] `query:"to_date"`
}

// URLQuery serializes [AnalyticsYoutubeGetDailyViewsParams]'s query parameters as
// `url.Values`.
func (r AnalyticsYoutubeGetDailyViewsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
