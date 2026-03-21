// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// AnalyticsService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options []option.RequestOption
	Youtube *AnalyticsYoutubeService
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r *AnalyticsService) {
	r = &AnalyticsService{}
	r.Options = opts
	r.Youtube = NewAnalyticsYoutubeService(opts...)
	return
}

// Get post analytics
func (r *AnalyticsService) Get(ctx context.Context, query AnalyticsGetParams, opts ...option.RequestOption) (res *AnalyticsGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get best posting times based on engagement
func (r *AnalyticsService) GetBestTime(ctx context.Context, query AnalyticsGetBestTimeParams, opts ...option.RequestOption) (res *AnalyticsGetBestTimeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics/best-time"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get engagement decay curve for a post
func (r *AnalyticsService) GetContentDecay(ctx context.Context, query AnalyticsGetContentDecayParams, opts ...option.RequestOption) (res *AnalyticsGetContentDecayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics/content-decay"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get per-post daily timeline of metrics
func (r *AnalyticsService) GetPostTimeline(ctx context.Context, query AnalyticsGetPostTimelineParams, opts ...option.RequestOption) (res *AnalyticsGetPostTimelineResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics/post-timeline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get posting frequency vs engagement analysis
func (r *AnalyticsService) GetPostingFrequency(ctx context.Context, query AnalyticsGetPostingFrequencyParams, opts ...option.RequestOption) (res *AnalyticsGetPostingFrequencyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics/posting-frequency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get daily aggregated metrics
func (r *AnalyticsService) ListDailyMetrics(ctx context.Context, query AnalyticsListDailyMetricsParams, opts ...option.RequestOption) (res *AnalyticsListDailyMetricsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/analytics/daily-metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AnalyticsGetResponse struct {
	Data     []AnalyticsGetResponseData   `json:"data" api:"required"`
	Overview AnalyticsGetResponseOverview `json:"overview"`
	JSON     analyticsGetResponseJSON     `json:"-"`
}

// analyticsGetResponseJSON contains the JSON metadata for the struct
// [AnalyticsGetResponse]
type analyticsGetResponseJSON struct {
	Data        apijson.Field
	Overview    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetResponseData struct {
	Platform AnalyticsGetResponseDataPlatform `json:"platform" api:"required"`
	// Post ID
	PostID string `json:"post_id" api:"required"`
	// Published timestamp
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Total clicks
	Clicks float64 `json:"clicks" api:"nullable"`
	// Total comments
	Comments float64 `json:"comments" api:"nullable"`
	// Total impressions
	Impressions float64 `json:"impressions" api:"nullable"`
	// Total likes
	Likes float64 `json:"likes" api:"nullable"`
	// Total reach
	Reach float64 `json:"reach" api:"nullable"`
	// Total saves
	Saves float64 `json:"saves" api:"nullable"`
	// Total shares
	Shares float64 `json:"shares" api:"nullable"`
	// Total views
	Views float64                      `json:"views" api:"nullable"`
	JSON  analyticsGetResponseDataJSON `json:"-"`
}

// analyticsGetResponseDataJSON contains the JSON metadata for the struct
// [AnalyticsGetResponseData]
type analyticsGetResponseDataJSON struct {
	Platform    apijson.Field
	PostID      apijson.Field
	PublishedAt apijson.Field
	Clicks      apijson.Field
	Comments    apijson.Field
	Impressions apijson.Field
	Likes       apijson.Field
	Reach       apijson.Field
	Saves       apijson.Field
	Shares      apijson.Field
	Views       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetResponseDataPlatform string

const (
	AnalyticsGetResponseDataPlatformTwitter        AnalyticsGetResponseDataPlatform = "twitter"
	AnalyticsGetResponseDataPlatformInstagram      AnalyticsGetResponseDataPlatform = "instagram"
	AnalyticsGetResponseDataPlatformFacebook       AnalyticsGetResponseDataPlatform = "facebook"
	AnalyticsGetResponseDataPlatformLinkedin       AnalyticsGetResponseDataPlatform = "linkedin"
	AnalyticsGetResponseDataPlatformTiktok         AnalyticsGetResponseDataPlatform = "tiktok"
	AnalyticsGetResponseDataPlatformYoutube        AnalyticsGetResponseDataPlatform = "youtube"
	AnalyticsGetResponseDataPlatformPinterest      AnalyticsGetResponseDataPlatform = "pinterest"
	AnalyticsGetResponseDataPlatformReddit         AnalyticsGetResponseDataPlatform = "reddit"
	AnalyticsGetResponseDataPlatformBluesky        AnalyticsGetResponseDataPlatform = "bluesky"
	AnalyticsGetResponseDataPlatformThreads        AnalyticsGetResponseDataPlatform = "threads"
	AnalyticsGetResponseDataPlatformTelegram       AnalyticsGetResponseDataPlatform = "telegram"
	AnalyticsGetResponseDataPlatformSnapchat       AnalyticsGetResponseDataPlatform = "snapchat"
	AnalyticsGetResponseDataPlatformGooglebusiness AnalyticsGetResponseDataPlatform = "googlebusiness"
	AnalyticsGetResponseDataPlatformWhatsapp       AnalyticsGetResponseDataPlatform = "whatsapp"
	AnalyticsGetResponseDataPlatformMastodon       AnalyticsGetResponseDataPlatform = "mastodon"
	AnalyticsGetResponseDataPlatformDiscord        AnalyticsGetResponseDataPlatform = "discord"
	AnalyticsGetResponseDataPlatformSMS            AnalyticsGetResponseDataPlatform = "sms"
)

func (r AnalyticsGetResponseDataPlatform) IsKnown() bool {
	switch r {
	case AnalyticsGetResponseDataPlatformTwitter, AnalyticsGetResponseDataPlatformInstagram, AnalyticsGetResponseDataPlatformFacebook, AnalyticsGetResponseDataPlatformLinkedin, AnalyticsGetResponseDataPlatformTiktok, AnalyticsGetResponseDataPlatformYoutube, AnalyticsGetResponseDataPlatformPinterest, AnalyticsGetResponseDataPlatformReddit, AnalyticsGetResponseDataPlatformBluesky, AnalyticsGetResponseDataPlatformThreads, AnalyticsGetResponseDataPlatformTelegram, AnalyticsGetResponseDataPlatformSnapchat, AnalyticsGetResponseDataPlatformGooglebusiness, AnalyticsGetResponseDataPlatformWhatsapp, AnalyticsGetResponseDataPlatformMastodon, AnalyticsGetResponseDataPlatformDiscord, AnalyticsGetResponseDataPlatformSMS:
		return true
	}
	return false
}

type AnalyticsGetResponseOverview struct {
	// Total clicks across posts
	TotalClicks float64 `json:"total_clicks" api:"required"`
	// Total comments across posts
	TotalComments float64 `json:"total_comments" api:"required"`
	// Total impressions across posts
	TotalImpressions float64 `json:"total_impressions" api:"required"`
	// Total likes across posts
	TotalLikes float64 `json:"total_likes" api:"required"`
	// Total number of posts
	TotalPosts float64 `json:"total_posts" api:"required"`
	// Total shares across posts
	TotalShares float64 `json:"total_shares" api:"required"`
	// Total views across posts
	TotalViews float64                          `json:"total_views" api:"required"`
	JSON       analyticsGetResponseOverviewJSON `json:"-"`
}

// analyticsGetResponseOverviewJSON contains the JSON metadata for the struct
// [AnalyticsGetResponseOverview]
type analyticsGetResponseOverviewJSON struct {
	TotalClicks      apijson.Field
	TotalComments    apijson.Field
	TotalImpressions apijson.Field
	TotalLikes       apijson.Field
	TotalPosts       apijson.Field
	TotalShares      apijson.Field
	TotalViews       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsGetResponseOverview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetResponseOverviewJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetBestTimeResponse struct {
	Data []AnalyticsGetBestTimeResponseData `json:"data" api:"required"`
	JSON analyticsGetBestTimeResponseJSON   `json:"-"`
}

// analyticsGetBestTimeResponseJSON contains the JSON metadata for the struct
// [AnalyticsGetBestTimeResponse]
type analyticsGetBestTimeResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsGetBestTimeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetBestTimeResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetBestTimeResponseData struct {
	// Average engagement score
	AvgEngagement float64 `json:"avg_engagement" api:"required"`
	// Day of week (0=Sunday)
	DayOfWeek int64 `json:"day_of_week" api:"required"`
	// Hour in UTC
	HourUtc int64 `json:"hour_utc" api:"required"`
	// Number of posts analyzed
	PostCount float64                              `json:"post_count" api:"required"`
	JSON      analyticsGetBestTimeResponseDataJSON `json:"-"`
}

// analyticsGetBestTimeResponseDataJSON contains the JSON metadata for the struct
// [AnalyticsGetBestTimeResponseData]
type analyticsGetBestTimeResponseDataJSON struct {
	AvgEngagement apijson.Field
	DayOfWeek     apijson.Field
	HourUtc       apijson.Field
	PostCount     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AnalyticsGetBestTimeResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetBestTimeResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetContentDecayResponse struct {
	Data []AnalyticsGetContentDecayResponseData `json:"data" api:"required"`
	// Days until engagement halved
	HalfLifeDays float64                                  `json:"half_life_days" api:"required,nullable"`
	Platform     AnalyticsGetContentDecayResponsePlatform `json:"platform" api:"required"`
	PostID       string                                   `json:"post_id" api:"required"`
	JSON         analyticsGetContentDecayResponseJSON     `json:"-"`
}

// analyticsGetContentDecayResponseJSON contains the JSON metadata for the struct
// [AnalyticsGetContentDecayResponse]
type analyticsGetContentDecayResponseJSON struct {
	Data         apijson.Field
	HalfLifeDays apijson.Field
	Platform     apijson.Field
	PostID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsGetContentDecayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetContentDecayResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetContentDecayResponseData struct {
	// Cumulative engagement
	CumulativeEngagement float64 `json:"cumulative_engagement" api:"required"`
	// Cumulative impressions
	CumulativeImpressions float64 `json:"cumulative_impressions" api:"required"`
	// Days since publication
	Day float64 `json:"day" api:"required"`
	// Engagement on this day
	Engagement float64 `json:"engagement" api:"required"`
	// Impressions on this day
	Impressions float64                                  `json:"impressions" api:"required"`
	JSON        analyticsGetContentDecayResponseDataJSON `json:"-"`
}

// analyticsGetContentDecayResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsGetContentDecayResponseData]
type analyticsGetContentDecayResponseDataJSON struct {
	CumulativeEngagement  apijson.Field
	CumulativeImpressions apijson.Field
	Day                   apijson.Field
	Engagement            apijson.Field
	Impressions           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AnalyticsGetContentDecayResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetContentDecayResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetContentDecayResponsePlatform string

const (
	AnalyticsGetContentDecayResponsePlatformTwitter        AnalyticsGetContentDecayResponsePlatform = "twitter"
	AnalyticsGetContentDecayResponsePlatformInstagram      AnalyticsGetContentDecayResponsePlatform = "instagram"
	AnalyticsGetContentDecayResponsePlatformFacebook       AnalyticsGetContentDecayResponsePlatform = "facebook"
	AnalyticsGetContentDecayResponsePlatformLinkedin       AnalyticsGetContentDecayResponsePlatform = "linkedin"
	AnalyticsGetContentDecayResponsePlatformTiktok         AnalyticsGetContentDecayResponsePlatform = "tiktok"
	AnalyticsGetContentDecayResponsePlatformYoutube        AnalyticsGetContentDecayResponsePlatform = "youtube"
	AnalyticsGetContentDecayResponsePlatformPinterest      AnalyticsGetContentDecayResponsePlatform = "pinterest"
	AnalyticsGetContentDecayResponsePlatformReddit         AnalyticsGetContentDecayResponsePlatform = "reddit"
	AnalyticsGetContentDecayResponsePlatformBluesky        AnalyticsGetContentDecayResponsePlatform = "bluesky"
	AnalyticsGetContentDecayResponsePlatformThreads        AnalyticsGetContentDecayResponsePlatform = "threads"
	AnalyticsGetContentDecayResponsePlatformTelegram       AnalyticsGetContentDecayResponsePlatform = "telegram"
	AnalyticsGetContentDecayResponsePlatformSnapchat       AnalyticsGetContentDecayResponsePlatform = "snapchat"
	AnalyticsGetContentDecayResponsePlatformGooglebusiness AnalyticsGetContentDecayResponsePlatform = "googlebusiness"
	AnalyticsGetContentDecayResponsePlatformWhatsapp       AnalyticsGetContentDecayResponsePlatform = "whatsapp"
	AnalyticsGetContentDecayResponsePlatformMastodon       AnalyticsGetContentDecayResponsePlatform = "mastodon"
	AnalyticsGetContentDecayResponsePlatformDiscord        AnalyticsGetContentDecayResponsePlatform = "discord"
	AnalyticsGetContentDecayResponsePlatformSMS            AnalyticsGetContentDecayResponsePlatform = "sms"
)

func (r AnalyticsGetContentDecayResponsePlatform) IsKnown() bool {
	switch r {
	case AnalyticsGetContentDecayResponsePlatformTwitter, AnalyticsGetContentDecayResponsePlatformInstagram, AnalyticsGetContentDecayResponsePlatformFacebook, AnalyticsGetContentDecayResponsePlatformLinkedin, AnalyticsGetContentDecayResponsePlatformTiktok, AnalyticsGetContentDecayResponsePlatformYoutube, AnalyticsGetContentDecayResponsePlatformPinterest, AnalyticsGetContentDecayResponsePlatformReddit, AnalyticsGetContentDecayResponsePlatformBluesky, AnalyticsGetContentDecayResponsePlatformThreads, AnalyticsGetContentDecayResponsePlatformTelegram, AnalyticsGetContentDecayResponsePlatformSnapchat, AnalyticsGetContentDecayResponsePlatformGooglebusiness, AnalyticsGetContentDecayResponsePlatformWhatsapp, AnalyticsGetContentDecayResponsePlatformMastodon, AnalyticsGetContentDecayResponsePlatformDiscord, AnalyticsGetContentDecayResponsePlatformSMS:
		return true
	}
	return false
}

type AnalyticsGetPostTimelineResponse struct {
	Data   []AnalyticsGetPostTimelineResponseData `json:"data" api:"required"`
	PostID string                                 `json:"post_id" api:"required"`
	JSON   analyticsGetPostTimelineResponseJSON   `json:"-"`
}

// analyticsGetPostTimelineResponseJSON contains the JSON metadata for the struct
// [AnalyticsGetPostTimelineResponse]
type analyticsGetPostTimelineResponseJSON struct {
	Data        apijson.Field
	PostID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsGetPostTimelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetPostTimelineResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetPostTimelineResponseData struct {
	Clicks   float64 `json:"clicks" api:"required"`
	Comments float64 `json:"comments" api:"required"`
	// Date (YYYY-MM-DD)
	Date        string                                   `json:"date" api:"required"`
	Impressions float64                                  `json:"impressions" api:"required"`
	Likes       float64                                  `json:"likes" api:"required"`
	Shares      float64                                  `json:"shares" api:"required"`
	Views       float64                                  `json:"views" api:"required"`
	JSON        analyticsGetPostTimelineResponseDataJSON `json:"-"`
}

// analyticsGetPostTimelineResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsGetPostTimelineResponseData]
type analyticsGetPostTimelineResponseDataJSON struct {
	Clicks      apijson.Field
	Comments    apijson.Field
	Date        apijson.Field
	Impressions apijson.Field
	Likes       apijson.Field
	Shares      apijson.Field
	Views       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsGetPostTimelineResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetPostTimelineResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetPostingFrequencyResponse struct {
	Data []AnalyticsGetPostingFrequencyResponseData `json:"data" api:"required"`
	// Recommended posts per week
	OptimalFrequency float64                                  `json:"optimal_frequency" api:"required,nullable"`
	JSON             analyticsGetPostingFrequencyResponseJSON `json:"-"`
}

// analyticsGetPostingFrequencyResponseJSON contains the JSON metadata for the
// struct [AnalyticsGetPostingFrequencyResponse]
type analyticsGetPostingFrequencyResponseJSON struct {
	Data             apijson.Field
	OptimalFrequency apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsGetPostingFrequencyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetPostingFrequencyResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetPostingFrequencyResponseData struct {
	// Average engagement
	AvgEngagement float64 `json:"avg_engagement" api:"required"`
	// Average impressions
	AvgImpressions float64 `json:"avg_impressions" api:"required"`
	// Average posts per week in bucket
	PostsPerWeek float64 `json:"posts_per_week" api:"required"`
	// Number of weeks in sample
	SampleWeeks float64                                      `json:"sample_weeks" api:"required"`
	JSON        analyticsGetPostingFrequencyResponseDataJSON `json:"-"`
}

// analyticsGetPostingFrequencyResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsGetPostingFrequencyResponseData]
type analyticsGetPostingFrequencyResponseDataJSON struct {
	AvgEngagement  apijson.Field
	AvgImpressions apijson.Field
	PostsPerWeek   apijson.Field
	SampleWeeks    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AnalyticsGetPostingFrequencyResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsGetPostingFrequencyResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsListDailyMetricsResponse struct {
	Data []AnalyticsListDailyMetricsResponseData `json:"data" api:"required"`
	JSON analyticsListDailyMetricsResponseJSON   `json:"-"`
}

// analyticsListDailyMetricsResponseJSON contains the JSON metadata for the struct
// [AnalyticsListDailyMetricsResponse]
type analyticsListDailyMetricsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsListDailyMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsListDailyMetricsResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsListDailyMetricsResponseData struct {
	// Total clicks
	Clicks float64 `json:"clicks" api:"required"`
	// Total comments
	Comments float64 `json:"comments" api:"required"`
	// Date (YYYY-MM-DD)
	Date string `json:"date" api:"required"`
	// Total impressions
	Impressions float64 `json:"impressions" api:"required"`
	// Total likes
	Likes float64 `json:"likes" api:"required"`
	// Post count per platform
	Platforms map[string]float64 `json:"platforms" api:"required"`
	// Posts published on this date
	PostCount float64 `json:"post_count" api:"required"`
	// Total shares
	Shares float64 `json:"shares" api:"required"`
	// Total views
	Views float64                                   `json:"views" api:"required"`
	JSON  analyticsListDailyMetricsResponseDataJSON `json:"-"`
}

// analyticsListDailyMetricsResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsListDailyMetricsResponseData]
type analyticsListDailyMetricsResponseDataJSON struct {
	Clicks      apijson.Field
	Comments    apijson.Field
	Date        apijson.Field
	Impressions apijson.Field
	Likes       apijson.Field
	Platforms   apijson.Field
	PostCount   apijson.Field
	Shares      apijson.Field
	Views       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsListDailyMetricsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsListDailyMetricsResponseDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsGetParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Start date (ISO 8601 date string)
	FromDate param.Field[string] `query:"from_date"`
	// Number of items
	Limit param.Field[int64] `query:"limit"`
	// Offset
	Offset param.Field[int64] `query:"offset"`
	// Filter by platform
	Platform param.Field[AnalyticsGetParamsPlatform] `query:"platform"`
	// Filter by post ID
	PostID param.Field[string] `query:"post_id"`
	// End date (ISO 8601 date string)
	ToDate param.Field[string] `query:"to_date"`
}

// URLQuery serializes [AnalyticsGetParams]'s query parameters as `url.Values`.
func (r AnalyticsGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type AnalyticsGetParamsPlatform string

const (
	AnalyticsGetParamsPlatformTwitter        AnalyticsGetParamsPlatform = "twitter"
	AnalyticsGetParamsPlatformInstagram      AnalyticsGetParamsPlatform = "instagram"
	AnalyticsGetParamsPlatformFacebook       AnalyticsGetParamsPlatform = "facebook"
	AnalyticsGetParamsPlatformLinkedin       AnalyticsGetParamsPlatform = "linkedin"
	AnalyticsGetParamsPlatformTiktok         AnalyticsGetParamsPlatform = "tiktok"
	AnalyticsGetParamsPlatformYoutube        AnalyticsGetParamsPlatform = "youtube"
	AnalyticsGetParamsPlatformPinterest      AnalyticsGetParamsPlatform = "pinterest"
	AnalyticsGetParamsPlatformReddit         AnalyticsGetParamsPlatform = "reddit"
	AnalyticsGetParamsPlatformBluesky        AnalyticsGetParamsPlatform = "bluesky"
	AnalyticsGetParamsPlatformThreads        AnalyticsGetParamsPlatform = "threads"
	AnalyticsGetParamsPlatformTelegram       AnalyticsGetParamsPlatform = "telegram"
	AnalyticsGetParamsPlatformSnapchat       AnalyticsGetParamsPlatform = "snapchat"
	AnalyticsGetParamsPlatformGooglebusiness AnalyticsGetParamsPlatform = "googlebusiness"
	AnalyticsGetParamsPlatformWhatsapp       AnalyticsGetParamsPlatform = "whatsapp"
	AnalyticsGetParamsPlatformMastodon       AnalyticsGetParamsPlatform = "mastodon"
	AnalyticsGetParamsPlatformDiscord        AnalyticsGetParamsPlatform = "discord"
	AnalyticsGetParamsPlatformSMS            AnalyticsGetParamsPlatform = "sms"
)

func (r AnalyticsGetParamsPlatform) IsKnown() bool {
	switch r {
	case AnalyticsGetParamsPlatformTwitter, AnalyticsGetParamsPlatformInstagram, AnalyticsGetParamsPlatformFacebook, AnalyticsGetParamsPlatformLinkedin, AnalyticsGetParamsPlatformTiktok, AnalyticsGetParamsPlatformYoutube, AnalyticsGetParamsPlatformPinterest, AnalyticsGetParamsPlatformReddit, AnalyticsGetParamsPlatformBluesky, AnalyticsGetParamsPlatformThreads, AnalyticsGetParamsPlatformTelegram, AnalyticsGetParamsPlatformSnapchat, AnalyticsGetParamsPlatformGooglebusiness, AnalyticsGetParamsPlatformWhatsapp, AnalyticsGetParamsPlatformMastodon, AnalyticsGetParamsPlatformDiscord, AnalyticsGetParamsPlatformSMS:
		return true
	}
	return false
}

type AnalyticsGetBestTimeParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Start date (ISO 8601)
	FromDate param.Field[string] `query:"from_date"`
	// Filter by platform
	Platform param.Field[AnalyticsGetBestTimeParamsPlatform] `query:"platform"`
	// End date (ISO 8601)
	ToDate param.Field[string] `query:"to_date"`
}

// URLQuery serializes [AnalyticsGetBestTimeParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetBestTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type AnalyticsGetBestTimeParamsPlatform string

const (
	AnalyticsGetBestTimeParamsPlatformTwitter        AnalyticsGetBestTimeParamsPlatform = "twitter"
	AnalyticsGetBestTimeParamsPlatformInstagram      AnalyticsGetBestTimeParamsPlatform = "instagram"
	AnalyticsGetBestTimeParamsPlatformFacebook       AnalyticsGetBestTimeParamsPlatform = "facebook"
	AnalyticsGetBestTimeParamsPlatformLinkedin       AnalyticsGetBestTimeParamsPlatform = "linkedin"
	AnalyticsGetBestTimeParamsPlatformTiktok         AnalyticsGetBestTimeParamsPlatform = "tiktok"
	AnalyticsGetBestTimeParamsPlatformYoutube        AnalyticsGetBestTimeParamsPlatform = "youtube"
	AnalyticsGetBestTimeParamsPlatformPinterest      AnalyticsGetBestTimeParamsPlatform = "pinterest"
	AnalyticsGetBestTimeParamsPlatformReddit         AnalyticsGetBestTimeParamsPlatform = "reddit"
	AnalyticsGetBestTimeParamsPlatformBluesky        AnalyticsGetBestTimeParamsPlatform = "bluesky"
	AnalyticsGetBestTimeParamsPlatformThreads        AnalyticsGetBestTimeParamsPlatform = "threads"
	AnalyticsGetBestTimeParamsPlatformTelegram       AnalyticsGetBestTimeParamsPlatform = "telegram"
	AnalyticsGetBestTimeParamsPlatformSnapchat       AnalyticsGetBestTimeParamsPlatform = "snapchat"
	AnalyticsGetBestTimeParamsPlatformGooglebusiness AnalyticsGetBestTimeParamsPlatform = "googlebusiness"
	AnalyticsGetBestTimeParamsPlatformWhatsapp       AnalyticsGetBestTimeParamsPlatform = "whatsapp"
	AnalyticsGetBestTimeParamsPlatformMastodon       AnalyticsGetBestTimeParamsPlatform = "mastodon"
	AnalyticsGetBestTimeParamsPlatformDiscord        AnalyticsGetBestTimeParamsPlatform = "discord"
	AnalyticsGetBestTimeParamsPlatformSMS            AnalyticsGetBestTimeParamsPlatform = "sms"
)

func (r AnalyticsGetBestTimeParamsPlatform) IsKnown() bool {
	switch r {
	case AnalyticsGetBestTimeParamsPlatformTwitter, AnalyticsGetBestTimeParamsPlatformInstagram, AnalyticsGetBestTimeParamsPlatformFacebook, AnalyticsGetBestTimeParamsPlatformLinkedin, AnalyticsGetBestTimeParamsPlatformTiktok, AnalyticsGetBestTimeParamsPlatformYoutube, AnalyticsGetBestTimeParamsPlatformPinterest, AnalyticsGetBestTimeParamsPlatformReddit, AnalyticsGetBestTimeParamsPlatformBluesky, AnalyticsGetBestTimeParamsPlatformThreads, AnalyticsGetBestTimeParamsPlatformTelegram, AnalyticsGetBestTimeParamsPlatformSnapchat, AnalyticsGetBestTimeParamsPlatformGooglebusiness, AnalyticsGetBestTimeParamsPlatformWhatsapp, AnalyticsGetBestTimeParamsPlatformMastodon, AnalyticsGetBestTimeParamsPlatformDiscord, AnalyticsGetBestTimeParamsPlatformSMS:
		return true
	}
	return false
}

type AnalyticsGetContentDecayParams struct {
	// Post ID to analyze decay for
	PostID param.Field[string] `query:"post_id" api:"required"`
	// Number of days to analyze
	Days param.Field[int64] `query:"days"`
}

// URLQuery serializes [AnalyticsGetContentDecayParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetContentDecayParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticsGetPostTimelineParams struct {
	// Post ID
	PostID param.Field[string] `query:"post_id" api:"required"`
	// Start date (ISO 8601)
	FromDate param.Field[string] `query:"from_date"`
	// End date (ISO 8601)
	ToDate param.Field[string] `query:"to_date"`
}

// URLQuery serializes [AnalyticsGetPostTimelineParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetPostTimelineParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticsGetPostingFrequencyParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Start date (ISO 8601)
	FromDate param.Field[string] `query:"from_date"`
	// Filter by platform
	Platform param.Field[AnalyticsGetPostingFrequencyParamsPlatform] `query:"platform"`
	// End date (ISO 8601)
	ToDate param.Field[string] `query:"to_date"`
}

// URLQuery serializes [AnalyticsGetPostingFrequencyParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetPostingFrequencyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type AnalyticsGetPostingFrequencyParamsPlatform string

const (
	AnalyticsGetPostingFrequencyParamsPlatformTwitter        AnalyticsGetPostingFrequencyParamsPlatform = "twitter"
	AnalyticsGetPostingFrequencyParamsPlatformInstagram      AnalyticsGetPostingFrequencyParamsPlatform = "instagram"
	AnalyticsGetPostingFrequencyParamsPlatformFacebook       AnalyticsGetPostingFrequencyParamsPlatform = "facebook"
	AnalyticsGetPostingFrequencyParamsPlatformLinkedin       AnalyticsGetPostingFrequencyParamsPlatform = "linkedin"
	AnalyticsGetPostingFrequencyParamsPlatformTiktok         AnalyticsGetPostingFrequencyParamsPlatform = "tiktok"
	AnalyticsGetPostingFrequencyParamsPlatformYoutube        AnalyticsGetPostingFrequencyParamsPlatform = "youtube"
	AnalyticsGetPostingFrequencyParamsPlatformPinterest      AnalyticsGetPostingFrequencyParamsPlatform = "pinterest"
	AnalyticsGetPostingFrequencyParamsPlatformReddit         AnalyticsGetPostingFrequencyParamsPlatform = "reddit"
	AnalyticsGetPostingFrequencyParamsPlatformBluesky        AnalyticsGetPostingFrequencyParamsPlatform = "bluesky"
	AnalyticsGetPostingFrequencyParamsPlatformThreads        AnalyticsGetPostingFrequencyParamsPlatform = "threads"
	AnalyticsGetPostingFrequencyParamsPlatformTelegram       AnalyticsGetPostingFrequencyParamsPlatform = "telegram"
	AnalyticsGetPostingFrequencyParamsPlatformSnapchat       AnalyticsGetPostingFrequencyParamsPlatform = "snapchat"
	AnalyticsGetPostingFrequencyParamsPlatformGooglebusiness AnalyticsGetPostingFrequencyParamsPlatform = "googlebusiness"
	AnalyticsGetPostingFrequencyParamsPlatformWhatsapp       AnalyticsGetPostingFrequencyParamsPlatform = "whatsapp"
	AnalyticsGetPostingFrequencyParamsPlatformMastodon       AnalyticsGetPostingFrequencyParamsPlatform = "mastodon"
	AnalyticsGetPostingFrequencyParamsPlatformDiscord        AnalyticsGetPostingFrequencyParamsPlatform = "discord"
	AnalyticsGetPostingFrequencyParamsPlatformSMS            AnalyticsGetPostingFrequencyParamsPlatform = "sms"
)

func (r AnalyticsGetPostingFrequencyParamsPlatform) IsKnown() bool {
	switch r {
	case AnalyticsGetPostingFrequencyParamsPlatformTwitter, AnalyticsGetPostingFrequencyParamsPlatformInstagram, AnalyticsGetPostingFrequencyParamsPlatformFacebook, AnalyticsGetPostingFrequencyParamsPlatformLinkedin, AnalyticsGetPostingFrequencyParamsPlatformTiktok, AnalyticsGetPostingFrequencyParamsPlatformYoutube, AnalyticsGetPostingFrequencyParamsPlatformPinterest, AnalyticsGetPostingFrequencyParamsPlatformReddit, AnalyticsGetPostingFrequencyParamsPlatformBluesky, AnalyticsGetPostingFrequencyParamsPlatformThreads, AnalyticsGetPostingFrequencyParamsPlatformTelegram, AnalyticsGetPostingFrequencyParamsPlatformSnapchat, AnalyticsGetPostingFrequencyParamsPlatformGooglebusiness, AnalyticsGetPostingFrequencyParamsPlatformWhatsapp, AnalyticsGetPostingFrequencyParamsPlatformMastodon, AnalyticsGetPostingFrequencyParamsPlatformDiscord, AnalyticsGetPostingFrequencyParamsPlatformSMS:
		return true
	}
	return false
}

type AnalyticsListDailyMetricsParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Start date (ISO 8601)
	FromDate param.Field[string] `query:"from_date"`
	// Filter by platform
	Platform param.Field[AnalyticsListDailyMetricsParamsPlatform] `query:"platform"`
	// End date (ISO 8601)
	ToDate param.Field[string] `query:"to_date"`
}

// URLQuery serializes [AnalyticsListDailyMetricsParams]'s query parameters as
// `url.Values`.
func (r AnalyticsListDailyMetricsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type AnalyticsListDailyMetricsParamsPlatform string

const (
	AnalyticsListDailyMetricsParamsPlatformTwitter        AnalyticsListDailyMetricsParamsPlatform = "twitter"
	AnalyticsListDailyMetricsParamsPlatformInstagram      AnalyticsListDailyMetricsParamsPlatform = "instagram"
	AnalyticsListDailyMetricsParamsPlatformFacebook       AnalyticsListDailyMetricsParamsPlatform = "facebook"
	AnalyticsListDailyMetricsParamsPlatformLinkedin       AnalyticsListDailyMetricsParamsPlatform = "linkedin"
	AnalyticsListDailyMetricsParamsPlatformTiktok         AnalyticsListDailyMetricsParamsPlatform = "tiktok"
	AnalyticsListDailyMetricsParamsPlatformYoutube        AnalyticsListDailyMetricsParamsPlatform = "youtube"
	AnalyticsListDailyMetricsParamsPlatformPinterest      AnalyticsListDailyMetricsParamsPlatform = "pinterest"
	AnalyticsListDailyMetricsParamsPlatformReddit         AnalyticsListDailyMetricsParamsPlatform = "reddit"
	AnalyticsListDailyMetricsParamsPlatformBluesky        AnalyticsListDailyMetricsParamsPlatform = "bluesky"
	AnalyticsListDailyMetricsParamsPlatformThreads        AnalyticsListDailyMetricsParamsPlatform = "threads"
	AnalyticsListDailyMetricsParamsPlatformTelegram       AnalyticsListDailyMetricsParamsPlatform = "telegram"
	AnalyticsListDailyMetricsParamsPlatformSnapchat       AnalyticsListDailyMetricsParamsPlatform = "snapchat"
	AnalyticsListDailyMetricsParamsPlatformGooglebusiness AnalyticsListDailyMetricsParamsPlatform = "googlebusiness"
	AnalyticsListDailyMetricsParamsPlatformWhatsapp       AnalyticsListDailyMetricsParamsPlatform = "whatsapp"
	AnalyticsListDailyMetricsParamsPlatformMastodon       AnalyticsListDailyMetricsParamsPlatform = "mastodon"
	AnalyticsListDailyMetricsParamsPlatformDiscord        AnalyticsListDailyMetricsParamsPlatform = "discord"
	AnalyticsListDailyMetricsParamsPlatformSMS            AnalyticsListDailyMetricsParamsPlatform = "sms"
)

func (r AnalyticsListDailyMetricsParamsPlatform) IsKnown() bool {
	switch r {
	case AnalyticsListDailyMetricsParamsPlatformTwitter, AnalyticsListDailyMetricsParamsPlatformInstagram, AnalyticsListDailyMetricsParamsPlatformFacebook, AnalyticsListDailyMetricsParamsPlatformLinkedin, AnalyticsListDailyMetricsParamsPlatformTiktok, AnalyticsListDailyMetricsParamsPlatformYoutube, AnalyticsListDailyMetricsParamsPlatformPinterest, AnalyticsListDailyMetricsParamsPlatformReddit, AnalyticsListDailyMetricsParamsPlatformBluesky, AnalyticsListDailyMetricsParamsPlatformThreads, AnalyticsListDailyMetricsParamsPlatformTelegram, AnalyticsListDailyMetricsParamsPlatformSnapchat, AnalyticsListDailyMetricsParamsPlatformGooglebusiness, AnalyticsListDailyMetricsParamsPlatformWhatsapp, AnalyticsListDailyMetricsParamsPlatformMastodon, AnalyticsListDailyMetricsParamsPlatformDiscord, AnalyticsListDailyMetricsParamsPlatformSMS:
		return true
	}
	return false
}
