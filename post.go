// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
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

// PostService contains methods and other services that help with interacting with
// the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostService] method instead.
type PostService struct {
	Options []option.RequestOption
	Logs    *PostLogService
}

// NewPostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPostService(opts ...option.RequestOption) (r *PostService) {
	r = &PostService{}
	r.Options = opts
	r.Logs = NewPostLogService(opts...)
	return
}

// Create a post. Use scheduled_at: "now" to publish immediately, "draft" to save
// as draft, or an ISO timestamp to schedule.
func (r *PostService) New(ctx context.Context, body PostNewParams, opts ...option.RequestOption) (res *PostNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a post
func (r *PostService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PostGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/posts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a draft or scheduled post.
func (r *PostService) Update(ctx context.Context, id string, body PostUpdateParams, opts ...option.RequestOption) (res *PostUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/posts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List posts
func (r *PostService) List(ctx context.Context, query PostListParams, opts ...option.RequestOption) (res *PostListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a post.
func (r *PostService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/posts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Create multiple posts in a single request. Each item follows the same schema as
// single post creation.
func (r *PostService) BulkNew(ctx context.Context, body PostBulkNewParams, opts ...option.RequestOption) (res *PostBulkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/posts/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retry publishing for failed targets on a post.
func (r *PostService) Retry(ctx context.Context, id string, opts ...option.RequestOption) (res *PostRetryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/posts/%s/retry", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Attempt to delete the post from each platform and set the post status to
// cancelled.
func (r *PostService) Unpublish(ctx context.Context, id string, body PostUnpublishParams, opts ...option.RequestOption) (res *PostUnpublishResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/posts/%s/unpublish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PostNewResponse struct {
	// Post ID
	ID        string                 `json:"id" api:"required"`
	Content   string                 `json:"content" api:"required,nullable"`
	CreatedAt time.Time              `json:"created_at" api:"required" format:"date-time"`
	Media     []PostNewResponseMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostNewResponseRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                   `json:"scheduled_at" api:"required,nullable"`
	Status      PostNewResponseStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostNewResponseTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostNewResponseMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string              `json:"timezone" api:"nullable"`
	JSON     postNewResponseJSON `json:"-"`
}

// postNewResponseJSON contains the JSON metadata for the struct [PostNewResponse]
type postNewResponseJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseJSON) RawJSON() string {
	return r.raw
}

type PostNewResponseMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostNewResponseMediaType `json:"type"`
	JSON postNewResponseMediaJSON `json:"-"`
}

// postNewResponseMediaJSON contains the JSON metadata for the struct
// [PostNewResponseMedia]
type postNewResponseMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostNewResponseMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostNewResponseMediaType string

const (
	PostNewResponseMediaTypeImage    PostNewResponseMediaType = "image"
	PostNewResponseMediaTypeVideo    PostNewResponseMediaType = "video"
	PostNewResponseMediaTypeGif      PostNewResponseMediaType = "gif"
	PostNewResponseMediaTypeDocument PostNewResponseMediaType = "document"
)

func (r PostNewResponseMediaType) IsKnown() bool {
	switch r {
	case PostNewResponseMediaTypeImage, PostNewResponseMediaTypeVideo, PostNewResponseMediaTypeGif, PostNewResponseMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostNewResponseRecycling struct {
	ID                    string                          `json:"id" api:"required"`
	ContentVariationIndex float64                         `json:"content_variation_index" api:"required"`
	ContentVariations     []string                        `json:"content_variations" api:"required"`
	CreatedAt             time.Time                       `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                            `json:"enabled" api:"required"`
	ExpireCount           float64                         `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                       `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                         `json:"gap" api:"required"`
	GapFreq               PostNewResponseRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                       `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                       `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                         `json:"recycle_count" api:"required"`
	StartDate             time.Time                       `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                       `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postNewResponseRecyclingJSON    `json:"-"`
}

// postNewResponseRecyclingJSON contains the JSON metadata for the struct
// [PostNewResponseRecycling]
type postNewResponseRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostNewResponseRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostNewResponseRecyclingGapFreq string

const (
	PostNewResponseRecyclingGapFreqDay   PostNewResponseRecyclingGapFreq = "day"
	PostNewResponseRecyclingGapFreqWeek  PostNewResponseRecyclingGapFreq = "week"
	PostNewResponseRecyclingGapFreqMonth PostNewResponseRecyclingGapFreq = "month"
)

func (r PostNewResponseRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostNewResponseRecyclingGapFreqDay, PostNewResponseRecyclingGapFreqWeek, PostNewResponseRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostNewResponseStatus string

const (
	PostNewResponseStatusDraft      PostNewResponseStatus = "draft"
	PostNewResponseStatusScheduled  PostNewResponseStatus = "scheduled"
	PostNewResponseStatusPublishing PostNewResponseStatus = "publishing"
	PostNewResponseStatusPublished  PostNewResponseStatus = "published"
	PostNewResponseStatusFailed     PostNewResponseStatus = "failed"
	PostNewResponseStatusPartial    PostNewResponseStatus = "partial"
)

func (r PostNewResponseStatus) IsKnown() bool {
	switch r {
	case PostNewResponseStatusDraft, PostNewResponseStatusScheduled, PostNewResponseStatusPublishing, PostNewResponseStatusPublished, PostNewResponseStatusFailed, PostNewResponseStatusPartial:
		return true
	}
	return false
}

type PostNewResponseTarget struct {
	Platform PostNewResponseTargetsPlatform  `json:"platform" api:"required"`
	Status   PostNewResponseTargetsStatus    `json:"status" api:"required"`
	Accounts []PostNewResponseTargetsAccount `json:"accounts"`
	Error    PostNewResponseTargetsError     `json:"error"`
	JSON     postNewResponseTargetJSON       `json:"-"`
}

// postNewResponseTargetJSON contains the JSON metadata for the struct
// [PostNewResponseTarget]
type postNewResponseTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostNewResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseTargetJSON) RawJSON() string {
	return r.raw
}

type PostNewResponseTargetsPlatform string

const (
	PostNewResponseTargetsPlatformTwitter        PostNewResponseTargetsPlatform = "twitter"
	PostNewResponseTargetsPlatformInstagram      PostNewResponseTargetsPlatform = "instagram"
	PostNewResponseTargetsPlatformFacebook       PostNewResponseTargetsPlatform = "facebook"
	PostNewResponseTargetsPlatformLinkedin       PostNewResponseTargetsPlatform = "linkedin"
	PostNewResponseTargetsPlatformTiktok         PostNewResponseTargetsPlatform = "tiktok"
	PostNewResponseTargetsPlatformYoutube        PostNewResponseTargetsPlatform = "youtube"
	PostNewResponseTargetsPlatformPinterest      PostNewResponseTargetsPlatform = "pinterest"
	PostNewResponseTargetsPlatformReddit         PostNewResponseTargetsPlatform = "reddit"
	PostNewResponseTargetsPlatformBluesky        PostNewResponseTargetsPlatform = "bluesky"
	PostNewResponseTargetsPlatformThreads        PostNewResponseTargetsPlatform = "threads"
	PostNewResponseTargetsPlatformTelegram       PostNewResponseTargetsPlatform = "telegram"
	PostNewResponseTargetsPlatformSnapchat       PostNewResponseTargetsPlatform = "snapchat"
	PostNewResponseTargetsPlatformGooglebusiness PostNewResponseTargetsPlatform = "googlebusiness"
	PostNewResponseTargetsPlatformWhatsapp       PostNewResponseTargetsPlatform = "whatsapp"
	PostNewResponseTargetsPlatformMastodon       PostNewResponseTargetsPlatform = "mastodon"
	PostNewResponseTargetsPlatformDiscord        PostNewResponseTargetsPlatform = "discord"
	PostNewResponseTargetsPlatformSMS            PostNewResponseTargetsPlatform = "sms"
	PostNewResponseTargetsPlatformBeehiiv        PostNewResponseTargetsPlatform = "beehiiv"
	PostNewResponseTargetsPlatformConvertkit     PostNewResponseTargetsPlatform = "convertkit"
	PostNewResponseTargetsPlatformMailchimp      PostNewResponseTargetsPlatform = "mailchimp"
	PostNewResponseTargetsPlatformListmonk       PostNewResponseTargetsPlatform = "listmonk"
)

func (r PostNewResponseTargetsPlatform) IsKnown() bool {
	switch r {
	case PostNewResponseTargetsPlatformTwitter, PostNewResponseTargetsPlatformInstagram, PostNewResponseTargetsPlatformFacebook, PostNewResponseTargetsPlatformLinkedin, PostNewResponseTargetsPlatformTiktok, PostNewResponseTargetsPlatformYoutube, PostNewResponseTargetsPlatformPinterest, PostNewResponseTargetsPlatformReddit, PostNewResponseTargetsPlatformBluesky, PostNewResponseTargetsPlatformThreads, PostNewResponseTargetsPlatformTelegram, PostNewResponseTargetsPlatformSnapchat, PostNewResponseTargetsPlatformGooglebusiness, PostNewResponseTargetsPlatformWhatsapp, PostNewResponseTargetsPlatformMastodon, PostNewResponseTargetsPlatformDiscord, PostNewResponseTargetsPlatformSMS, PostNewResponseTargetsPlatformBeehiiv, PostNewResponseTargetsPlatformConvertkit, PostNewResponseTargetsPlatformMailchimp, PostNewResponseTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostNewResponseTargetsStatus string

const (
	PostNewResponseTargetsStatusDraft      PostNewResponseTargetsStatus = "draft"
	PostNewResponseTargetsStatusScheduled  PostNewResponseTargetsStatus = "scheduled"
	PostNewResponseTargetsStatusPublishing PostNewResponseTargetsStatus = "publishing"
	PostNewResponseTargetsStatusPublished  PostNewResponseTargetsStatus = "published"
	PostNewResponseTargetsStatusFailed     PostNewResponseTargetsStatus = "failed"
	PostNewResponseTargetsStatusPartial    PostNewResponseTargetsStatus = "partial"
)

func (r PostNewResponseTargetsStatus) IsKnown() bool {
	switch r {
	case PostNewResponseTargetsStatusDraft, PostNewResponseTargetsStatusScheduled, PostNewResponseTargetsStatusPublishing, PostNewResponseTargetsStatusPublished, PostNewResponseTargetsStatusFailed, PostNewResponseTargetsStatusPartial:
		return true
	}
	return false
}

type PostNewResponseTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                            `json:"url" api:"required,nullable"`
	Username string                            `json:"username" api:"required,nullable"`
	JSON     postNewResponseTargetsAccountJSON `json:"-"`
}

// postNewResponseTargetsAccountJSON contains the JSON metadata for the struct
// [PostNewResponseTargetsAccount]
type postNewResponseTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostNewResponseTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostNewResponseTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                          `json:"detail"`
	JSON   postNewResponseTargetsErrorJSON `json:"-"`
}

// postNewResponseTargetsErrorJSON contains the JSON metadata for the struct
// [PostNewResponseTargetsError]
type postNewResponseTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostNewResponseTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostNewResponseMetrics struct {
	Clicks         float64                    `json:"clicks"`
	Comments       float64                    `json:"comments"`
	EngagementRate float64                    `json:"engagement_rate"`
	Impressions    float64                    `json:"impressions"`
	Likes          float64                    `json:"likes"`
	Reach          float64                    `json:"reach"`
	Saves          float64                    `json:"saves"`
	Shares         float64                    `json:"shares"`
	Views          float64                    `json:"views"`
	JSON           postNewResponseMetricsJSON `json:"-"`
}

// postNewResponseMetricsJSON contains the JSON metadata for the struct
// [PostNewResponseMetrics]
type postNewResponseMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostNewResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postNewResponseMetricsJSON) RawJSON() string {
	return r.raw
}

type PostGetResponse struct {
	// Post ID
	ID        string                 `json:"id" api:"required"`
	Content   string                 `json:"content" api:"required,nullable"`
	CreatedAt time.Time              `json:"created_at" api:"required" format:"date-time"`
	Media     []PostGetResponseMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostGetResponseRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                   `json:"scheduled_at" api:"required,nullable"`
	Status      PostGetResponseStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostGetResponseTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostGetResponseMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string              `json:"timezone" api:"nullable"`
	JSON     postGetResponseJSON `json:"-"`
}

// postGetResponseJSON contains the JSON metadata for the struct [PostGetResponse]
type postGetResponseJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseJSON) RawJSON() string {
	return r.raw
}

type PostGetResponseMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostGetResponseMediaType `json:"type"`
	JSON postGetResponseMediaJSON `json:"-"`
}

// postGetResponseMediaJSON contains the JSON metadata for the struct
// [PostGetResponseMedia]
type postGetResponseMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostGetResponseMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostGetResponseMediaType string

const (
	PostGetResponseMediaTypeImage    PostGetResponseMediaType = "image"
	PostGetResponseMediaTypeVideo    PostGetResponseMediaType = "video"
	PostGetResponseMediaTypeGif      PostGetResponseMediaType = "gif"
	PostGetResponseMediaTypeDocument PostGetResponseMediaType = "document"
)

func (r PostGetResponseMediaType) IsKnown() bool {
	switch r {
	case PostGetResponseMediaTypeImage, PostGetResponseMediaTypeVideo, PostGetResponseMediaTypeGif, PostGetResponseMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostGetResponseRecycling struct {
	ID                    string                          `json:"id" api:"required"`
	ContentVariationIndex float64                         `json:"content_variation_index" api:"required"`
	ContentVariations     []string                        `json:"content_variations" api:"required"`
	CreatedAt             time.Time                       `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                            `json:"enabled" api:"required"`
	ExpireCount           float64                         `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                       `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                         `json:"gap" api:"required"`
	GapFreq               PostGetResponseRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                       `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                       `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                         `json:"recycle_count" api:"required"`
	StartDate             time.Time                       `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                       `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postGetResponseRecyclingJSON    `json:"-"`
}

// postGetResponseRecyclingJSON contains the JSON metadata for the struct
// [PostGetResponseRecycling]
type postGetResponseRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostGetResponseRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostGetResponseRecyclingGapFreq string

const (
	PostGetResponseRecyclingGapFreqDay   PostGetResponseRecyclingGapFreq = "day"
	PostGetResponseRecyclingGapFreqWeek  PostGetResponseRecyclingGapFreq = "week"
	PostGetResponseRecyclingGapFreqMonth PostGetResponseRecyclingGapFreq = "month"
)

func (r PostGetResponseRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostGetResponseRecyclingGapFreqDay, PostGetResponseRecyclingGapFreqWeek, PostGetResponseRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostGetResponseStatus string

const (
	PostGetResponseStatusDraft      PostGetResponseStatus = "draft"
	PostGetResponseStatusScheduled  PostGetResponseStatus = "scheduled"
	PostGetResponseStatusPublishing PostGetResponseStatus = "publishing"
	PostGetResponseStatusPublished  PostGetResponseStatus = "published"
	PostGetResponseStatusFailed     PostGetResponseStatus = "failed"
	PostGetResponseStatusPartial    PostGetResponseStatus = "partial"
)

func (r PostGetResponseStatus) IsKnown() bool {
	switch r {
	case PostGetResponseStatusDraft, PostGetResponseStatusScheduled, PostGetResponseStatusPublishing, PostGetResponseStatusPublished, PostGetResponseStatusFailed, PostGetResponseStatusPartial:
		return true
	}
	return false
}

type PostGetResponseTarget struct {
	Platform PostGetResponseTargetsPlatform  `json:"platform" api:"required"`
	Status   PostGetResponseTargetsStatus    `json:"status" api:"required"`
	Accounts []PostGetResponseTargetsAccount `json:"accounts"`
	Error    PostGetResponseTargetsError     `json:"error"`
	JSON     postGetResponseTargetJSON       `json:"-"`
}

// postGetResponseTargetJSON contains the JSON metadata for the struct
// [PostGetResponseTarget]
type postGetResponseTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostGetResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseTargetJSON) RawJSON() string {
	return r.raw
}

type PostGetResponseTargetsPlatform string

const (
	PostGetResponseTargetsPlatformTwitter        PostGetResponseTargetsPlatform = "twitter"
	PostGetResponseTargetsPlatformInstagram      PostGetResponseTargetsPlatform = "instagram"
	PostGetResponseTargetsPlatformFacebook       PostGetResponseTargetsPlatform = "facebook"
	PostGetResponseTargetsPlatformLinkedin       PostGetResponseTargetsPlatform = "linkedin"
	PostGetResponseTargetsPlatformTiktok         PostGetResponseTargetsPlatform = "tiktok"
	PostGetResponseTargetsPlatformYoutube        PostGetResponseTargetsPlatform = "youtube"
	PostGetResponseTargetsPlatformPinterest      PostGetResponseTargetsPlatform = "pinterest"
	PostGetResponseTargetsPlatformReddit         PostGetResponseTargetsPlatform = "reddit"
	PostGetResponseTargetsPlatformBluesky        PostGetResponseTargetsPlatform = "bluesky"
	PostGetResponseTargetsPlatformThreads        PostGetResponseTargetsPlatform = "threads"
	PostGetResponseTargetsPlatformTelegram       PostGetResponseTargetsPlatform = "telegram"
	PostGetResponseTargetsPlatformSnapchat       PostGetResponseTargetsPlatform = "snapchat"
	PostGetResponseTargetsPlatformGooglebusiness PostGetResponseTargetsPlatform = "googlebusiness"
	PostGetResponseTargetsPlatformWhatsapp       PostGetResponseTargetsPlatform = "whatsapp"
	PostGetResponseTargetsPlatformMastodon       PostGetResponseTargetsPlatform = "mastodon"
	PostGetResponseTargetsPlatformDiscord        PostGetResponseTargetsPlatform = "discord"
	PostGetResponseTargetsPlatformSMS            PostGetResponseTargetsPlatform = "sms"
	PostGetResponseTargetsPlatformBeehiiv        PostGetResponseTargetsPlatform = "beehiiv"
	PostGetResponseTargetsPlatformConvertkit     PostGetResponseTargetsPlatform = "convertkit"
	PostGetResponseTargetsPlatformMailchimp      PostGetResponseTargetsPlatform = "mailchimp"
	PostGetResponseTargetsPlatformListmonk       PostGetResponseTargetsPlatform = "listmonk"
)

func (r PostGetResponseTargetsPlatform) IsKnown() bool {
	switch r {
	case PostGetResponseTargetsPlatformTwitter, PostGetResponseTargetsPlatformInstagram, PostGetResponseTargetsPlatformFacebook, PostGetResponseTargetsPlatformLinkedin, PostGetResponseTargetsPlatformTiktok, PostGetResponseTargetsPlatformYoutube, PostGetResponseTargetsPlatformPinterest, PostGetResponseTargetsPlatformReddit, PostGetResponseTargetsPlatformBluesky, PostGetResponseTargetsPlatformThreads, PostGetResponseTargetsPlatformTelegram, PostGetResponseTargetsPlatformSnapchat, PostGetResponseTargetsPlatformGooglebusiness, PostGetResponseTargetsPlatformWhatsapp, PostGetResponseTargetsPlatformMastodon, PostGetResponseTargetsPlatformDiscord, PostGetResponseTargetsPlatformSMS, PostGetResponseTargetsPlatformBeehiiv, PostGetResponseTargetsPlatformConvertkit, PostGetResponseTargetsPlatformMailchimp, PostGetResponseTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostGetResponseTargetsStatus string

const (
	PostGetResponseTargetsStatusDraft      PostGetResponseTargetsStatus = "draft"
	PostGetResponseTargetsStatusScheduled  PostGetResponseTargetsStatus = "scheduled"
	PostGetResponseTargetsStatusPublishing PostGetResponseTargetsStatus = "publishing"
	PostGetResponseTargetsStatusPublished  PostGetResponseTargetsStatus = "published"
	PostGetResponseTargetsStatusFailed     PostGetResponseTargetsStatus = "failed"
	PostGetResponseTargetsStatusPartial    PostGetResponseTargetsStatus = "partial"
)

func (r PostGetResponseTargetsStatus) IsKnown() bool {
	switch r {
	case PostGetResponseTargetsStatusDraft, PostGetResponseTargetsStatusScheduled, PostGetResponseTargetsStatusPublishing, PostGetResponseTargetsStatusPublished, PostGetResponseTargetsStatusFailed, PostGetResponseTargetsStatusPartial:
		return true
	}
	return false
}

type PostGetResponseTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                            `json:"url" api:"required,nullable"`
	Username string                            `json:"username" api:"required,nullable"`
	JSON     postGetResponseTargetsAccountJSON `json:"-"`
}

// postGetResponseTargetsAccountJSON contains the JSON metadata for the struct
// [PostGetResponseTargetsAccount]
type postGetResponseTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostGetResponseTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostGetResponseTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                          `json:"detail"`
	JSON   postGetResponseTargetsErrorJSON `json:"-"`
}

// postGetResponseTargetsErrorJSON contains the JSON metadata for the struct
// [PostGetResponseTargetsError]
type postGetResponseTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostGetResponseTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostGetResponseMetrics struct {
	Clicks         float64                    `json:"clicks"`
	Comments       float64                    `json:"comments"`
	EngagementRate float64                    `json:"engagement_rate"`
	Impressions    float64                    `json:"impressions"`
	Likes          float64                    `json:"likes"`
	Reach          float64                    `json:"reach"`
	Saves          float64                    `json:"saves"`
	Shares         float64                    `json:"shares"`
	Views          float64                    `json:"views"`
	JSON           postGetResponseMetricsJSON `json:"-"`
}

// postGetResponseMetricsJSON contains the JSON metadata for the struct
// [PostGetResponseMetrics]
type postGetResponseMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostGetResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postGetResponseMetricsJSON) RawJSON() string {
	return r.raw
}

type PostUpdateResponse struct {
	// Post ID
	ID        string                    `json:"id" api:"required"`
	Content   string                    `json:"content" api:"required,nullable"`
	CreatedAt time.Time                 `json:"created_at" api:"required" format:"date-time"`
	Media     []PostUpdateResponseMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostUpdateResponseRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                      `json:"scheduled_at" api:"required,nullable"`
	Status      PostUpdateResponseStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostUpdateResponseTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                           `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostUpdateResponseMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string                 `json:"timezone" api:"nullable"`
	JSON     postUpdateResponseJSON `json:"-"`
}

// postUpdateResponseJSON contains the JSON metadata for the struct
// [PostUpdateResponse]
type postUpdateResponseJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PostUpdateResponseMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostUpdateResponseMediaType `json:"type"`
	JSON postUpdateResponseMediaJSON `json:"-"`
}

// postUpdateResponseMediaJSON contains the JSON metadata for the struct
// [PostUpdateResponseMedia]
type postUpdateResponseMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostUpdateResponseMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostUpdateResponseMediaType string

const (
	PostUpdateResponseMediaTypeImage    PostUpdateResponseMediaType = "image"
	PostUpdateResponseMediaTypeVideo    PostUpdateResponseMediaType = "video"
	PostUpdateResponseMediaTypeGif      PostUpdateResponseMediaType = "gif"
	PostUpdateResponseMediaTypeDocument PostUpdateResponseMediaType = "document"
)

func (r PostUpdateResponseMediaType) IsKnown() bool {
	switch r {
	case PostUpdateResponseMediaTypeImage, PostUpdateResponseMediaTypeVideo, PostUpdateResponseMediaTypeGif, PostUpdateResponseMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostUpdateResponseRecycling struct {
	ID                    string                             `json:"id" api:"required"`
	ContentVariationIndex float64                            `json:"content_variation_index" api:"required"`
	ContentVariations     []string                           `json:"content_variations" api:"required"`
	CreatedAt             time.Time                          `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                               `json:"enabled" api:"required"`
	ExpireCount           float64                            `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                          `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                            `json:"gap" api:"required"`
	GapFreq               PostUpdateResponseRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                          `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                          `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                            `json:"recycle_count" api:"required"`
	StartDate             time.Time                          `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                          `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postUpdateResponseRecyclingJSON    `json:"-"`
}

// postUpdateResponseRecyclingJSON contains the JSON metadata for the struct
// [PostUpdateResponseRecycling]
type postUpdateResponseRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostUpdateResponseRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostUpdateResponseRecyclingGapFreq string

const (
	PostUpdateResponseRecyclingGapFreqDay   PostUpdateResponseRecyclingGapFreq = "day"
	PostUpdateResponseRecyclingGapFreqWeek  PostUpdateResponseRecyclingGapFreq = "week"
	PostUpdateResponseRecyclingGapFreqMonth PostUpdateResponseRecyclingGapFreq = "month"
)

func (r PostUpdateResponseRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostUpdateResponseRecyclingGapFreqDay, PostUpdateResponseRecyclingGapFreqWeek, PostUpdateResponseRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostUpdateResponseStatus string

const (
	PostUpdateResponseStatusDraft      PostUpdateResponseStatus = "draft"
	PostUpdateResponseStatusScheduled  PostUpdateResponseStatus = "scheduled"
	PostUpdateResponseStatusPublishing PostUpdateResponseStatus = "publishing"
	PostUpdateResponseStatusPublished  PostUpdateResponseStatus = "published"
	PostUpdateResponseStatusFailed     PostUpdateResponseStatus = "failed"
	PostUpdateResponseStatusPartial    PostUpdateResponseStatus = "partial"
)

func (r PostUpdateResponseStatus) IsKnown() bool {
	switch r {
	case PostUpdateResponseStatusDraft, PostUpdateResponseStatusScheduled, PostUpdateResponseStatusPublishing, PostUpdateResponseStatusPublished, PostUpdateResponseStatusFailed, PostUpdateResponseStatusPartial:
		return true
	}
	return false
}

type PostUpdateResponseTarget struct {
	Platform PostUpdateResponseTargetsPlatform  `json:"platform" api:"required"`
	Status   PostUpdateResponseTargetsStatus    `json:"status" api:"required"`
	Accounts []PostUpdateResponseTargetsAccount `json:"accounts"`
	Error    PostUpdateResponseTargetsError     `json:"error"`
	JSON     postUpdateResponseTargetJSON       `json:"-"`
}

// postUpdateResponseTargetJSON contains the JSON metadata for the struct
// [PostUpdateResponseTarget]
type postUpdateResponseTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostUpdateResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseTargetJSON) RawJSON() string {
	return r.raw
}

type PostUpdateResponseTargetsPlatform string

const (
	PostUpdateResponseTargetsPlatformTwitter        PostUpdateResponseTargetsPlatform = "twitter"
	PostUpdateResponseTargetsPlatformInstagram      PostUpdateResponseTargetsPlatform = "instagram"
	PostUpdateResponseTargetsPlatformFacebook       PostUpdateResponseTargetsPlatform = "facebook"
	PostUpdateResponseTargetsPlatformLinkedin       PostUpdateResponseTargetsPlatform = "linkedin"
	PostUpdateResponseTargetsPlatformTiktok         PostUpdateResponseTargetsPlatform = "tiktok"
	PostUpdateResponseTargetsPlatformYoutube        PostUpdateResponseTargetsPlatform = "youtube"
	PostUpdateResponseTargetsPlatformPinterest      PostUpdateResponseTargetsPlatform = "pinterest"
	PostUpdateResponseTargetsPlatformReddit         PostUpdateResponseTargetsPlatform = "reddit"
	PostUpdateResponseTargetsPlatformBluesky        PostUpdateResponseTargetsPlatform = "bluesky"
	PostUpdateResponseTargetsPlatformThreads        PostUpdateResponseTargetsPlatform = "threads"
	PostUpdateResponseTargetsPlatformTelegram       PostUpdateResponseTargetsPlatform = "telegram"
	PostUpdateResponseTargetsPlatformSnapchat       PostUpdateResponseTargetsPlatform = "snapchat"
	PostUpdateResponseTargetsPlatformGooglebusiness PostUpdateResponseTargetsPlatform = "googlebusiness"
	PostUpdateResponseTargetsPlatformWhatsapp       PostUpdateResponseTargetsPlatform = "whatsapp"
	PostUpdateResponseTargetsPlatformMastodon       PostUpdateResponseTargetsPlatform = "mastodon"
	PostUpdateResponseTargetsPlatformDiscord        PostUpdateResponseTargetsPlatform = "discord"
	PostUpdateResponseTargetsPlatformSMS            PostUpdateResponseTargetsPlatform = "sms"
	PostUpdateResponseTargetsPlatformBeehiiv        PostUpdateResponseTargetsPlatform = "beehiiv"
	PostUpdateResponseTargetsPlatformConvertkit     PostUpdateResponseTargetsPlatform = "convertkit"
	PostUpdateResponseTargetsPlatformMailchimp      PostUpdateResponseTargetsPlatform = "mailchimp"
	PostUpdateResponseTargetsPlatformListmonk       PostUpdateResponseTargetsPlatform = "listmonk"
)

func (r PostUpdateResponseTargetsPlatform) IsKnown() bool {
	switch r {
	case PostUpdateResponseTargetsPlatformTwitter, PostUpdateResponseTargetsPlatformInstagram, PostUpdateResponseTargetsPlatformFacebook, PostUpdateResponseTargetsPlatformLinkedin, PostUpdateResponseTargetsPlatformTiktok, PostUpdateResponseTargetsPlatformYoutube, PostUpdateResponseTargetsPlatformPinterest, PostUpdateResponseTargetsPlatformReddit, PostUpdateResponseTargetsPlatformBluesky, PostUpdateResponseTargetsPlatformThreads, PostUpdateResponseTargetsPlatformTelegram, PostUpdateResponseTargetsPlatformSnapchat, PostUpdateResponseTargetsPlatformGooglebusiness, PostUpdateResponseTargetsPlatformWhatsapp, PostUpdateResponseTargetsPlatformMastodon, PostUpdateResponseTargetsPlatformDiscord, PostUpdateResponseTargetsPlatformSMS, PostUpdateResponseTargetsPlatformBeehiiv, PostUpdateResponseTargetsPlatformConvertkit, PostUpdateResponseTargetsPlatformMailchimp, PostUpdateResponseTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostUpdateResponseTargetsStatus string

const (
	PostUpdateResponseTargetsStatusDraft      PostUpdateResponseTargetsStatus = "draft"
	PostUpdateResponseTargetsStatusScheduled  PostUpdateResponseTargetsStatus = "scheduled"
	PostUpdateResponseTargetsStatusPublishing PostUpdateResponseTargetsStatus = "publishing"
	PostUpdateResponseTargetsStatusPublished  PostUpdateResponseTargetsStatus = "published"
	PostUpdateResponseTargetsStatusFailed     PostUpdateResponseTargetsStatus = "failed"
	PostUpdateResponseTargetsStatusPartial    PostUpdateResponseTargetsStatus = "partial"
)

func (r PostUpdateResponseTargetsStatus) IsKnown() bool {
	switch r {
	case PostUpdateResponseTargetsStatusDraft, PostUpdateResponseTargetsStatusScheduled, PostUpdateResponseTargetsStatusPublishing, PostUpdateResponseTargetsStatusPublished, PostUpdateResponseTargetsStatusFailed, PostUpdateResponseTargetsStatusPartial:
		return true
	}
	return false
}

type PostUpdateResponseTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                               `json:"url" api:"required,nullable"`
	Username string                               `json:"username" api:"required,nullable"`
	JSON     postUpdateResponseTargetsAccountJSON `json:"-"`
}

// postUpdateResponseTargetsAccountJSON contains the JSON metadata for the struct
// [PostUpdateResponseTargetsAccount]
type postUpdateResponseTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostUpdateResponseTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostUpdateResponseTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                             `json:"detail"`
	JSON   postUpdateResponseTargetsErrorJSON `json:"-"`
}

// postUpdateResponseTargetsErrorJSON contains the JSON metadata for the struct
// [PostUpdateResponseTargetsError]
type postUpdateResponseTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostUpdateResponseTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostUpdateResponseMetrics struct {
	Clicks         float64                       `json:"clicks"`
	Comments       float64                       `json:"comments"`
	EngagementRate float64                       `json:"engagement_rate"`
	Impressions    float64                       `json:"impressions"`
	Likes          float64                       `json:"likes"`
	Reach          float64                       `json:"reach"`
	Saves          float64                       `json:"saves"`
	Shares         float64                       `json:"shares"`
	Views          float64                       `json:"views"`
	JSON           postUpdateResponseMetricsJSON `json:"-"`
}

// postUpdateResponseMetricsJSON contains the JSON metadata for the struct
// [PostUpdateResponseMetrics]
type postUpdateResponseMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostUpdateResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUpdateResponseMetricsJSON) RawJSON() string {
	return r.raw
}

type PostListResponse struct {
	Data []PostListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string               `json:"next_cursor" api:"required,nullable"`
	JSON       postListResponseJSON `json:"-"`
}

// postListResponseJSON contains the JSON metadata for the struct
// [PostListResponse]
type postListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseJSON) RawJSON() string {
	return r.raw
}

type PostListResponseData struct {
	// Post ID
	ID        string                      `json:"id" api:"required"`
	Content   string                      `json:"content" api:"required,nullable"`
	CreatedAt time.Time                   `json:"created_at" api:"required" format:"date-time"`
	Media     []PostListResponseDataMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostListResponseDataRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                        `json:"scheduled_at" api:"required,nullable"`
	Status      PostListResponseDataStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostListResponseDataTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostListResponseDataMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string                   `json:"timezone" api:"nullable"`
	JSON     postListResponseDataJSON `json:"-"`
}

// postListResponseDataJSON contains the JSON metadata for the struct
// [PostListResponseData]
type postListResponseDataJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataJSON) RawJSON() string {
	return r.raw
}

type PostListResponseDataMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostListResponseDataMediaType `json:"type"`
	JSON postListResponseDataMediaJSON `json:"-"`
}

// postListResponseDataMediaJSON contains the JSON metadata for the struct
// [PostListResponseDataMedia]
type postListResponseDataMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostListResponseDataMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostListResponseDataMediaType string

const (
	PostListResponseDataMediaTypeImage    PostListResponseDataMediaType = "image"
	PostListResponseDataMediaTypeVideo    PostListResponseDataMediaType = "video"
	PostListResponseDataMediaTypeGif      PostListResponseDataMediaType = "gif"
	PostListResponseDataMediaTypeDocument PostListResponseDataMediaType = "document"
)

func (r PostListResponseDataMediaType) IsKnown() bool {
	switch r {
	case PostListResponseDataMediaTypeImage, PostListResponseDataMediaTypeVideo, PostListResponseDataMediaTypeGif, PostListResponseDataMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostListResponseDataRecycling struct {
	ID                    string                               `json:"id" api:"required"`
	ContentVariationIndex float64                              `json:"content_variation_index" api:"required"`
	ContentVariations     []string                             `json:"content_variations" api:"required"`
	CreatedAt             time.Time                            `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                                 `json:"enabled" api:"required"`
	ExpireCount           float64                              `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                            `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                              `json:"gap" api:"required"`
	GapFreq               PostListResponseDataRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                            `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                            `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                              `json:"recycle_count" api:"required"`
	StartDate             time.Time                            `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                            `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postListResponseDataRecyclingJSON    `json:"-"`
}

// postListResponseDataRecyclingJSON contains the JSON metadata for the struct
// [PostListResponseDataRecycling]
type postListResponseDataRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostListResponseDataRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostListResponseDataRecyclingGapFreq string

const (
	PostListResponseDataRecyclingGapFreqDay   PostListResponseDataRecyclingGapFreq = "day"
	PostListResponseDataRecyclingGapFreqWeek  PostListResponseDataRecyclingGapFreq = "week"
	PostListResponseDataRecyclingGapFreqMonth PostListResponseDataRecyclingGapFreq = "month"
)

func (r PostListResponseDataRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostListResponseDataRecyclingGapFreqDay, PostListResponseDataRecyclingGapFreqWeek, PostListResponseDataRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostListResponseDataStatus string

const (
	PostListResponseDataStatusDraft      PostListResponseDataStatus = "draft"
	PostListResponseDataStatusScheduled  PostListResponseDataStatus = "scheduled"
	PostListResponseDataStatusPublishing PostListResponseDataStatus = "publishing"
	PostListResponseDataStatusPublished  PostListResponseDataStatus = "published"
	PostListResponseDataStatusFailed     PostListResponseDataStatus = "failed"
	PostListResponseDataStatusPartial    PostListResponseDataStatus = "partial"
)

func (r PostListResponseDataStatus) IsKnown() bool {
	switch r {
	case PostListResponseDataStatusDraft, PostListResponseDataStatusScheduled, PostListResponseDataStatusPublishing, PostListResponseDataStatusPublished, PostListResponseDataStatusFailed, PostListResponseDataStatusPartial:
		return true
	}
	return false
}

type PostListResponseDataTarget struct {
	Platform PostListResponseDataTargetsPlatform  `json:"platform" api:"required"`
	Status   PostListResponseDataTargetsStatus    `json:"status" api:"required"`
	Accounts []PostListResponseDataTargetsAccount `json:"accounts"`
	Error    PostListResponseDataTargetsError     `json:"error"`
	JSON     postListResponseDataTargetJSON       `json:"-"`
}

// postListResponseDataTargetJSON contains the JSON metadata for the struct
// [PostListResponseDataTarget]
type postListResponseDataTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostListResponseDataTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataTargetJSON) RawJSON() string {
	return r.raw
}

type PostListResponseDataTargetsPlatform string

const (
	PostListResponseDataTargetsPlatformTwitter        PostListResponseDataTargetsPlatform = "twitter"
	PostListResponseDataTargetsPlatformInstagram      PostListResponseDataTargetsPlatform = "instagram"
	PostListResponseDataTargetsPlatformFacebook       PostListResponseDataTargetsPlatform = "facebook"
	PostListResponseDataTargetsPlatformLinkedin       PostListResponseDataTargetsPlatform = "linkedin"
	PostListResponseDataTargetsPlatformTiktok         PostListResponseDataTargetsPlatform = "tiktok"
	PostListResponseDataTargetsPlatformYoutube        PostListResponseDataTargetsPlatform = "youtube"
	PostListResponseDataTargetsPlatformPinterest      PostListResponseDataTargetsPlatform = "pinterest"
	PostListResponseDataTargetsPlatformReddit         PostListResponseDataTargetsPlatform = "reddit"
	PostListResponseDataTargetsPlatformBluesky        PostListResponseDataTargetsPlatform = "bluesky"
	PostListResponseDataTargetsPlatformThreads        PostListResponseDataTargetsPlatform = "threads"
	PostListResponseDataTargetsPlatformTelegram       PostListResponseDataTargetsPlatform = "telegram"
	PostListResponseDataTargetsPlatformSnapchat       PostListResponseDataTargetsPlatform = "snapchat"
	PostListResponseDataTargetsPlatformGooglebusiness PostListResponseDataTargetsPlatform = "googlebusiness"
	PostListResponseDataTargetsPlatformWhatsapp       PostListResponseDataTargetsPlatform = "whatsapp"
	PostListResponseDataTargetsPlatformMastodon       PostListResponseDataTargetsPlatform = "mastodon"
	PostListResponseDataTargetsPlatformDiscord        PostListResponseDataTargetsPlatform = "discord"
	PostListResponseDataTargetsPlatformSMS            PostListResponseDataTargetsPlatform = "sms"
	PostListResponseDataTargetsPlatformBeehiiv        PostListResponseDataTargetsPlatform = "beehiiv"
	PostListResponseDataTargetsPlatformConvertkit     PostListResponseDataTargetsPlatform = "convertkit"
	PostListResponseDataTargetsPlatformMailchimp      PostListResponseDataTargetsPlatform = "mailchimp"
	PostListResponseDataTargetsPlatformListmonk       PostListResponseDataTargetsPlatform = "listmonk"
)

func (r PostListResponseDataTargetsPlatform) IsKnown() bool {
	switch r {
	case PostListResponseDataTargetsPlatformTwitter, PostListResponseDataTargetsPlatformInstagram, PostListResponseDataTargetsPlatformFacebook, PostListResponseDataTargetsPlatformLinkedin, PostListResponseDataTargetsPlatformTiktok, PostListResponseDataTargetsPlatformYoutube, PostListResponseDataTargetsPlatformPinterest, PostListResponseDataTargetsPlatformReddit, PostListResponseDataTargetsPlatformBluesky, PostListResponseDataTargetsPlatformThreads, PostListResponseDataTargetsPlatformTelegram, PostListResponseDataTargetsPlatformSnapchat, PostListResponseDataTargetsPlatformGooglebusiness, PostListResponseDataTargetsPlatformWhatsapp, PostListResponseDataTargetsPlatformMastodon, PostListResponseDataTargetsPlatformDiscord, PostListResponseDataTargetsPlatformSMS, PostListResponseDataTargetsPlatformBeehiiv, PostListResponseDataTargetsPlatformConvertkit, PostListResponseDataTargetsPlatformMailchimp, PostListResponseDataTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostListResponseDataTargetsStatus string

const (
	PostListResponseDataTargetsStatusDraft      PostListResponseDataTargetsStatus = "draft"
	PostListResponseDataTargetsStatusScheduled  PostListResponseDataTargetsStatus = "scheduled"
	PostListResponseDataTargetsStatusPublishing PostListResponseDataTargetsStatus = "publishing"
	PostListResponseDataTargetsStatusPublished  PostListResponseDataTargetsStatus = "published"
	PostListResponseDataTargetsStatusFailed     PostListResponseDataTargetsStatus = "failed"
	PostListResponseDataTargetsStatusPartial    PostListResponseDataTargetsStatus = "partial"
)

func (r PostListResponseDataTargetsStatus) IsKnown() bool {
	switch r {
	case PostListResponseDataTargetsStatusDraft, PostListResponseDataTargetsStatusScheduled, PostListResponseDataTargetsStatusPublishing, PostListResponseDataTargetsStatusPublished, PostListResponseDataTargetsStatusFailed, PostListResponseDataTargetsStatusPartial:
		return true
	}
	return false
}

type PostListResponseDataTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                                 `json:"url" api:"required,nullable"`
	Username string                                 `json:"username" api:"required,nullable"`
	JSON     postListResponseDataTargetsAccountJSON `json:"-"`
}

// postListResponseDataTargetsAccountJSON contains the JSON metadata for the struct
// [PostListResponseDataTargetsAccount]
type postListResponseDataTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostListResponseDataTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostListResponseDataTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                               `json:"detail"`
	JSON   postListResponseDataTargetsErrorJSON `json:"-"`
}

// postListResponseDataTargetsErrorJSON contains the JSON metadata for the struct
// [PostListResponseDataTargetsError]
type postListResponseDataTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostListResponseDataTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostListResponseDataMetrics struct {
	Clicks         float64                         `json:"clicks"`
	Comments       float64                         `json:"comments"`
	EngagementRate float64                         `json:"engagement_rate"`
	Impressions    float64                         `json:"impressions"`
	Likes          float64                         `json:"likes"`
	Reach          float64                         `json:"reach"`
	Saves          float64                         `json:"saves"`
	Shares         float64                         `json:"shares"`
	Views          float64                         `json:"views"`
	JSON           postListResponseDataMetricsJSON `json:"-"`
}

// postListResponseDataMetricsJSON contains the JSON metadata for the struct
// [PostListResponseDataMetrics]
type postListResponseDataMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostListResponseDataMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postListResponseDataMetricsJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponse struct {
	Data    []PostBulkNewResponseData  `json:"data" api:"required"`
	Summary PostBulkNewResponseSummary `json:"summary" api:"required"`
	JSON    postBulkNewResponseJSON    `json:"-"`
}

// postBulkNewResponseJSON contains the JSON metadata for the struct
// [PostBulkNewResponse]
type postBulkNewResponseJSON struct {
	Data        apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponseData struct {
	// Post ID
	ID        string                         `json:"id" api:"required"`
	Content   string                         `json:"content" api:"required,nullable"`
	CreatedAt time.Time                      `json:"created_at" api:"required" format:"date-time"`
	Media     []PostBulkNewResponseDataMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostBulkNewResponseDataRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                           `json:"scheduled_at" api:"required,nullable"`
	Status      PostBulkNewResponseDataStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostBulkNewResponseDataTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostBulkNewResponseDataMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string                      `json:"timezone" api:"nullable"`
	JSON     postBulkNewResponseDataJSON `json:"-"`
}

// postBulkNewResponseDataJSON contains the JSON metadata for the struct
// [PostBulkNewResponseData]
type postBulkNewResponseDataJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostBulkNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponseDataMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostBulkNewResponseDataMediaType `json:"type"`
	JSON postBulkNewResponseDataMediaJSON `json:"-"`
}

// postBulkNewResponseDataMediaJSON contains the JSON metadata for the struct
// [PostBulkNewResponseDataMedia]
type postBulkNewResponseDataMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostBulkNewResponseDataMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostBulkNewResponseDataMediaType string

const (
	PostBulkNewResponseDataMediaTypeImage    PostBulkNewResponseDataMediaType = "image"
	PostBulkNewResponseDataMediaTypeVideo    PostBulkNewResponseDataMediaType = "video"
	PostBulkNewResponseDataMediaTypeGif      PostBulkNewResponseDataMediaType = "gif"
	PostBulkNewResponseDataMediaTypeDocument PostBulkNewResponseDataMediaType = "document"
)

func (r PostBulkNewResponseDataMediaType) IsKnown() bool {
	switch r {
	case PostBulkNewResponseDataMediaTypeImage, PostBulkNewResponseDataMediaTypeVideo, PostBulkNewResponseDataMediaTypeGif, PostBulkNewResponseDataMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostBulkNewResponseDataRecycling struct {
	ID                    string                                  `json:"id" api:"required"`
	ContentVariationIndex float64                                 `json:"content_variation_index" api:"required"`
	ContentVariations     []string                                `json:"content_variations" api:"required"`
	CreatedAt             time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                                    `json:"enabled" api:"required"`
	ExpireCount           float64                                 `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                               `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                                 `json:"gap" api:"required"`
	GapFreq               PostBulkNewResponseDataRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                               `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                               `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                                 `json:"recycle_count" api:"required"`
	StartDate             time.Time                               `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postBulkNewResponseDataRecyclingJSON    `json:"-"`
}

// postBulkNewResponseDataRecyclingJSON contains the JSON metadata for the struct
// [PostBulkNewResponseDataRecycling]
type postBulkNewResponseDataRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostBulkNewResponseDataRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponseDataRecyclingGapFreq string

const (
	PostBulkNewResponseDataRecyclingGapFreqDay   PostBulkNewResponseDataRecyclingGapFreq = "day"
	PostBulkNewResponseDataRecyclingGapFreqWeek  PostBulkNewResponseDataRecyclingGapFreq = "week"
	PostBulkNewResponseDataRecyclingGapFreqMonth PostBulkNewResponseDataRecyclingGapFreq = "month"
)

func (r PostBulkNewResponseDataRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostBulkNewResponseDataRecyclingGapFreqDay, PostBulkNewResponseDataRecyclingGapFreqWeek, PostBulkNewResponseDataRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostBulkNewResponseDataStatus string

const (
	PostBulkNewResponseDataStatusDraft      PostBulkNewResponseDataStatus = "draft"
	PostBulkNewResponseDataStatusScheduled  PostBulkNewResponseDataStatus = "scheduled"
	PostBulkNewResponseDataStatusPublishing PostBulkNewResponseDataStatus = "publishing"
	PostBulkNewResponseDataStatusPublished  PostBulkNewResponseDataStatus = "published"
	PostBulkNewResponseDataStatusFailed     PostBulkNewResponseDataStatus = "failed"
	PostBulkNewResponseDataStatusPartial    PostBulkNewResponseDataStatus = "partial"
)

func (r PostBulkNewResponseDataStatus) IsKnown() bool {
	switch r {
	case PostBulkNewResponseDataStatusDraft, PostBulkNewResponseDataStatusScheduled, PostBulkNewResponseDataStatusPublishing, PostBulkNewResponseDataStatusPublished, PostBulkNewResponseDataStatusFailed, PostBulkNewResponseDataStatusPartial:
		return true
	}
	return false
}

type PostBulkNewResponseDataTarget struct {
	Platform PostBulkNewResponseDataTargetsPlatform  `json:"platform" api:"required"`
	Status   PostBulkNewResponseDataTargetsStatus    `json:"status" api:"required"`
	Accounts []PostBulkNewResponseDataTargetsAccount `json:"accounts"`
	Error    PostBulkNewResponseDataTargetsError     `json:"error"`
	JSON     postBulkNewResponseDataTargetJSON       `json:"-"`
}

// postBulkNewResponseDataTargetJSON contains the JSON metadata for the struct
// [PostBulkNewResponseDataTarget]
type postBulkNewResponseDataTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostBulkNewResponseDataTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataTargetJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponseDataTargetsPlatform string

const (
	PostBulkNewResponseDataTargetsPlatformTwitter        PostBulkNewResponseDataTargetsPlatform = "twitter"
	PostBulkNewResponseDataTargetsPlatformInstagram      PostBulkNewResponseDataTargetsPlatform = "instagram"
	PostBulkNewResponseDataTargetsPlatformFacebook       PostBulkNewResponseDataTargetsPlatform = "facebook"
	PostBulkNewResponseDataTargetsPlatformLinkedin       PostBulkNewResponseDataTargetsPlatform = "linkedin"
	PostBulkNewResponseDataTargetsPlatformTiktok         PostBulkNewResponseDataTargetsPlatform = "tiktok"
	PostBulkNewResponseDataTargetsPlatformYoutube        PostBulkNewResponseDataTargetsPlatform = "youtube"
	PostBulkNewResponseDataTargetsPlatformPinterest      PostBulkNewResponseDataTargetsPlatform = "pinterest"
	PostBulkNewResponseDataTargetsPlatformReddit         PostBulkNewResponseDataTargetsPlatform = "reddit"
	PostBulkNewResponseDataTargetsPlatformBluesky        PostBulkNewResponseDataTargetsPlatform = "bluesky"
	PostBulkNewResponseDataTargetsPlatformThreads        PostBulkNewResponseDataTargetsPlatform = "threads"
	PostBulkNewResponseDataTargetsPlatformTelegram       PostBulkNewResponseDataTargetsPlatform = "telegram"
	PostBulkNewResponseDataTargetsPlatformSnapchat       PostBulkNewResponseDataTargetsPlatform = "snapchat"
	PostBulkNewResponseDataTargetsPlatformGooglebusiness PostBulkNewResponseDataTargetsPlatform = "googlebusiness"
	PostBulkNewResponseDataTargetsPlatformWhatsapp       PostBulkNewResponseDataTargetsPlatform = "whatsapp"
	PostBulkNewResponseDataTargetsPlatformMastodon       PostBulkNewResponseDataTargetsPlatform = "mastodon"
	PostBulkNewResponseDataTargetsPlatformDiscord        PostBulkNewResponseDataTargetsPlatform = "discord"
	PostBulkNewResponseDataTargetsPlatformSMS            PostBulkNewResponseDataTargetsPlatform = "sms"
	PostBulkNewResponseDataTargetsPlatformBeehiiv        PostBulkNewResponseDataTargetsPlatform = "beehiiv"
	PostBulkNewResponseDataTargetsPlatformConvertkit     PostBulkNewResponseDataTargetsPlatform = "convertkit"
	PostBulkNewResponseDataTargetsPlatformMailchimp      PostBulkNewResponseDataTargetsPlatform = "mailchimp"
	PostBulkNewResponseDataTargetsPlatformListmonk       PostBulkNewResponseDataTargetsPlatform = "listmonk"
)

func (r PostBulkNewResponseDataTargetsPlatform) IsKnown() bool {
	switch r {
	case PostBulkNewResponseDataTargetsPlatformTwitter, PostBulkNewResponseDataTargetsPlatformInstagram, PostBulkNewResponseDataTargetsPlatformFacebook, PostBulkNewResponseDataTargetsPlatformLinkedin, PostBulkNewResponseDataTargetsPlatformTiktok, PostBulkNewResponseDataTargetsPlatformYoutube, PostBulkNewResponseDataTargetsPlatformPinterest, PostBulkNewResponseDataTargetsPlatformReddit, PostBulkNewResponseDataTargetsPlatformBluesky, PostBulkNewResponseDataTargetsPlatformThreads, PostBulkNewResponseDataTargetsPlatformTelegram, PostBulkNewResponseDataTargetsPlatformSnapchat, PostBulkNewResponseDataTargetsPlatformGooglebusiness, PostBulkNewResponseDataTargetsPlatformWhatsapp, PostBulkNewResponseDataTargetsPlatformMastodon, PostBulkNewResponseDataTargetsPlatformDiscord, PostBulkNewResponseDataTargetsPlatformSMS, PostBulkNewResponseDataTargetsPlatformBeehiiv, PostBulkNewResponseDataTargetsPlatformConvertkit, PostBulkNewResponseDataTargetsPlatformMailchimp, PostBulkNewResponseDataTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostBulkNewResponseDataTargetsStatus string

const (
	PostBulkNewResponseDataTargetsStatusDraft      PostBulkNewResponseDataTargetsStatus = "draft"
	PostBulkNewResponseDataTargetsStatusScheduled  PostBulkNewResponseDataTargetsStatus = "scheduled"
	PostBulkNewResponseDataTargetsStatusPublishing PostBulkNewResponseDataTargetsStatus = "publishing"
	PostBulkNewResponseDataTargetsStatusPublished  PostBulkNewResponseDataTargetsStatus = "published"
	PostBulkNewResponseDataTargetsStatusFailed     PostBulkNewResponseDataTargetsStatus = "failed"
	PostBulkNewResponseDataTargetsStatusPartial    PostBulkNewResponseDataTargetsStatus = "partial"
)

func (r PostBulkNewResponseDataTargetsStatus) IsKnown() bool {
	switch r {
	case PostBulkNewResponseDataTargetsStatusDraft, PostBulkNewResponseDataTargetsStatusScheduled, PostBulkNewResponseDataTargetsStatusPublishing, PostBulkNewResponseDataTargetsStatusPublished, PostBulkNewResponseDataTargetsStatusFailed, PostBulkNewResponseDataTargetsStatusPartial:
		return true
	}
	return false
}

type PostBulkNewResponseDataTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                                    `json:"url" api:"required,nullable"`
	Username string                                    `json:"username" api:"required,nullable"`
	JSON     postBulkNewResponseDataTargetsAccountJSON `json:"-"`
}

// postBulkNewResponseDataTargetsAccountJSON contains the JSON metadata for the
// struct [PostBulkNewResponseDataTargetsAccount]
type postBulkNewResponseDataTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostBulkNewResponseDataTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponseDataTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                                  `json:"detail"`
	JSON   postBulkNewResponseDataTargetsErrorJSON `json:"-"`
}

// postBulkNewResponseDataTargetsErrorJSON contains the JSON metadata for the
// struct [PostBulkNewResponseDataTargetsError]
type postBulkNewResponseDataTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostBulkNewResponseDataTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostBulkNewResponseDataMetrics struct {
	Clicks         float64                            `json:"clicks"`
	Comments       float64                            `json:"comments"`
	EngagementRate float64                            `json:"engagement_rate"`
	Impressions    float64                            `json:"impressions"`
	Likes          float64                            `json:"likes"`
	Reach          float64                            `json:"reach"`
	Saves          float64                            `json:"saves"`
	Shares         float64                            `json:"shares"`
	Views          float64                            `json:"views"`
	JSON           postBulkNewResponseDataMetricsJSON `json:"-"`
}

// postBulkNewResponseDataMetricsJSON contains the JSON metadata for the struct
// [PostBulkNewResponseDataMetrics]
type postBulkNewResponseDataMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostBulkNewResponseDataMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseDataMetricsJSON) RawJSON() string {
	return r.raw
}

type PostBulkNewResponseSummary struct {
	Failed    float64                        `json:"failed" api:"required"`
	Succeeded float64                        `json:"succeeded" api:"required"`
	Total     float64                        `json:"total" api:"required"`
	JSON      postBulkNewResponseSummaryJSON `json:"-"`
}

// postBulkNewResponseSummaryJSON contains the JSON metadata for the struct
// [PostBulkNewResponseSummary]
type postBulkNewResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PostRetryResponse struct {
	// Post ID
	ID        string                   `json:"id" api:"required"`
	Content   string                   `json:"content" api:"required,nullable"`
	CreatedAt time.Time                `json:"created_at" api:"required" format:"date-time"`
	Media     []PostRetryResponseMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostRetryResponseRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                     `json:"scheduled_at" api:"required,nullable"`
	Status      PostRetryResponseStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostRetryResponseTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                          `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostRetryResponseMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string                `json:"timezone" api:"nullable"`
	JSON     postRetryResponseJSON `json:"-"`
}

// postRetryResponseJSON contains the JSON metadata for the struct
// [PostRetryResponse]
type postRetryResponseJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostRetryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseJSON) RawJSON() string {
	return r.raw
}

type PostRetryResponseMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostRetryResponseMediaType `json:"type"`
	JSON postRetryResponseMediaJSON `json:"-"`
}

// postRetryResponseMediaJSON contains the JSON metadata for the struct
// [PostRetryResponseMedia]
type postRetryResponseMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostRetryResponseMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostRetryResponseMediaType string

const (
	PostRetryResponseMediaTypeImage    PostRetryResponseMediaType = "image"
	PostRetryResponseMediaTypeVideo    PostRetryResponseMediaType = "video"
	PostRetryResponseMediaTypeGif      PostRetryResponseMediaType = "gif"
	PostRetryResponseMediaTypeDocument PostRetryResponseMediaType = "document"
)

func (r PostRetryResponseMediaType) IsKnown() bool {
	switch r {
	case PostRetryResponseMediaTypeImage, PostRetryResponseMediaTypeVideo, PostRetryResponseMediaTypeGif, PostRetryResponseMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostRetryResponseRecycling struct {
	ID                    string                            `json:"id" api:"required"`
	ContentVariationIndex float64                           `json:"content_variation_index" api:"required"`
	ContentVariations     []string                          `json:"content_variations" api:"required"`
	CreatedAt             time.Time                         `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                              `json:"enabled" api:"required"`
	ExpireCount           float64                           `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                         `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                           `json:"gap" api:"required"`
	GapFreq               PostRetryResponseRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                         `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                         `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                           `json:"recycle_count" api:"required"`
	StartDate             time.Time                         `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                         `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postRetryResponseRecyclingJSON    `json:"-"`
}

// postRetryResponseRecyclingJSON contains the JSON metadata for the struct
// [PostRetryResponseRecycling]
type postRetryResponseRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostRetryResponseRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostRetryResponseRecyclingGapFreq string

const (
	PostRetryResponseRecyclingGapFreqDay   PostRetryResponseRecyclingGapFreq = "day"
	PostRetryResponseRecyclingGapFreqWeek  PostRetryResponseRecyclingGapFreq = "week"
	PostRetryResponseRecyclingGapFreqMonth PostRetryResponseRecyclingGapFreq = "month"
)

func (r PostRetryResponseRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostRetryResponseRecyclingGapFreqDay, PostRetryResponseRecyclingGapFreqWeek, PostRetryResponseRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostRetryResponseStatus string

const (
	PostRetryResponseStatusDraft      PostRetryResponseStatus = "draft"
	PostRetryResponseStatusScheduled  PostRetryResponseStatus = "scheduled"
	PostRetryResponseStatusPublishing PostRetryResponseStatus = "publishing"
	PostRetryResponseStatusPublished  PostRetryResponseStatus = "published"
	PostRetryResponseStatusFailed     PostRetryResponseStatus = "failed"
	PostRetryResponseStatusPartial    PostRetryResponseStatus = "partial"
)

func (r PostRetryResponseStatus) IsKnown() bool {
	switch r {
	case PostRetryResponseStatusDraft, PostRetryResponseStatusScheduled, PostRetryResponseStatusPublishing, PostRetryResponseStatusPublished, PostRetryResponseStatusFailed, PostRetryResponseStatusPartial:
		return true
	}
	return false
}

type PostRetryResponseTarget struct {
	Platform PostRetryResponseTargetsPlatform  `json:"platform" api:"required"`
	Status   PostRetryResponseTargetsStatus    `json:"status" api:"required"`
	Accounts []PostRetryResponseTargetsAccount `json:"accounts"`
	Error    PostRetryResponseTargetsError     `json:"error"`
	JSON     postRetryResponseTargetJSON       `json:"-"`
}

// postRetryResponseTargetJSON contains the JSON metadata for the struct
// [PostRetryResponseTarget]
type postRetryResponseTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostRetryResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseTargetJSON) RawJSON() string {
	return r.raw
}

type PostRetryResponseTargetsPlatform string

const (
	PostRetryResponseTargetsPlatformTwitter        PostRetryResponseTargetsPlatform = "twitter"
	PostRetryResponseTargetsPlatformInstagram      PostRetryResponseTargetsPlatform = "instagram"
	PostRetryResponseTargetsPlatformFacebook       PostRetryResponseTargetsPlatform = "facebook"
	PostRetryResponseTargetsPlatformLinkedin       PostRetryResponseTargetsPlatform = "linkedin"
	PostRetryResponseTargetsPlatformTiktok         PostRetryResponseTargetsPlatform = "tiktok"
	PostRetryResponseTargetsPlatformYoutube        PostRetryResponseTargetsPlatform = "youtube"
	PostRetryResponseTargetsPlatformPinterest      PostRetryResponseTargetsPlatform = "pinterest"
	PostRetryResponseTargetsPlatformReddit         PostRetryResponseTargetsPlatform = "reddit"
	PostRetryResponseTargetsPlatformBluesky        PostRetryResponseTargetsPlatform = "bluesky"
	PostRetryResponseTargetsPlatformThreads        PostRetryResponseTargetsPlatform = "threads"
	PostRetryResponseTargetsPlatformTelegram       PostRetryResponseTargetsPlatform = "telegram"
	PostRetryResponseTargetsPlatformSnapchat       PostRetryResponseTargetsPlatform = "snapchat"
	PostRetryResponseTargetsPlatformGooglebusiness PostRetryResponseTargetsPlatform = "googlebusiness"
	PostRetryResponseTargetsPlatformWhatsapp       PostRetryResponseTargetsPlatform = "whatsapp"
	PostRetryResponseTargetsPlatformMastodon       PostRetryResponseTargetsPlatform = "mastodon"
	PostRetryResponseTargetsPlatformDiscord        PostRetryResponseTargetsPlatform = "discord"
	PostRetryResponseTargetsPlatformSMS            PostRetryResponseTargetsPlatform = "sms"
	PostRetryResponseTargetsPlatformBeehiiv        PostRetryResponseTargetsPlatform = "beehiiv"
	PostRetryResponseTargetsPlatformConvertkit     PostRetryResponseTargetsPlatform = "convertkit"
	PostRetryResponseTargetsPlatformMailchimp      PostRetryResponseTargetsPlatform = "mailchimp"
	PostRetryResponseTargetsPlatformListmonk       PostRetryResponseTargetsPlatform = "listmonk"
)

func (r PostRetryResponseTargetsPlatform) IsKnown() bool {
	switch r {
	case PostRetryResponseTargetsPlatformTwitter, PostRetryResponseTargetsPlatformInstagram, PostRetryResponseTargetsPlatformFacebook, PostRetryResponseTargetsPlatformLinkedin, PostRetryResponseTargetsPlatformTiktok, PostRetryResponseTargetsPlatformYoutube, PostRetryResponseTargetsPlatformPinterest, PostRetryResponseTargetsPlatformReddit, PostRetryResponseTargetsPlatformBluesky, PostRetryResponseTargetsPlatformThreads, PostRetryResponseTargetsPlatformTelegram, PostRetryResponseTargetsPlatformSnapchat, PostRetryResponseTargetsPlatformGooglebusiness, PostRetryResponseTargetsPlatformWhatsapp, PostRetryResponseTargetsPlatformMastodon, PostRetryResponseTargetsPlatformDiscord, PostRetryResponseTargetsPlatformSMS, PostRetryResponseTargetsPlatformBeehiiv, PostRetryResponseTargetsPlatformConvertkit, PostRetryResponseTargetsPlatformMailchimp, PostRetryResponseTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostRetryResponseTargetsStatus string

const (
	PostRetryResponseTargetsStatusDraft      PostRetryResponseTargetsStatus = "draft"
	PostRetryResponseTargetsStatusScheduled  PostRetryResponseTargetsStatus = "scheduled"
	PostRetryResponseTargetsStatusPublishing PostRetryResponseTargetsStatus = "publishing"
	PostRetryResponseTargetsStatusPublished  PostRetryResponseTargetsStatus = "published"
	PostRetryResponseTargetsStatusFailed     PostRetryResponseTargetsStatus = "failed"
	PostRetryResponseTargetsStatusPartial    PostRetryResponseTargetsStatus = "partial"
)

func (r PostRetryResponseTargetsStatus) IsKnown() bool {
	switch r {
	case PostRetryResponseTargetsStatusDraft, PostRetryResponseTargetsStatusScheduled, PostRetryResponseTargetsStatusPublishing, PostRetryResponseTargetsStatusPublished, PostRetryResponseTargetsStatusFailed, PostRetryResponseTargetsStatusPartial:
		return true
	}
	return false
}

type PostRetryResponseTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                              `json:"url" api:"required,nullable"`
	Username string                              `json:"username" api:"required,nullable"`
	JSON     postRetryResponseTargetsAccountJSON `json:"-"`
}

// postRetryResponseTargetsAccountJSON contains the JSON metadata for the struct
// [PostRetryResponseTargetsAccount]
type postRetryResponseTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostRetryResponseTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostRetryResponseTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                            `json:"detail"`
	JSON   postRetryResponseTargetsErrorJSON `json:"-"`
}

// postRetryResponseTargetsErrorJSON contains the JSON metadata for the struct
// [PostRetryResponseTargetsError]
type postRetryResponseTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostRetryResponseTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostRetryResponseMetrics struct {
	Clicks         float64                      `json:"clicks"`
	Comments       float64                      `json:"comments"`
	EngagementRate float64                      `json:"engagement_rate"`
	Impressions    float64                      `json:"impressions"`
	Likes          float64                      `json:"likes"`
	Reach          float64                      `json:"reach"`
	Saves          float64                      `json:"saves"`
	Shares         float64                      `json:"shares"`
	Views          float64                      `json:"views"`
	JSON           postRetryResponseMetricsJSON `json:"-"`
}

// postRetryResponseMetricsJSON contains the JSON metadata for the struct
// [PostRetryResponseMetrics]
type postRetryResponseMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostRetryResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postRetryResponseMetricsJSON) RawJSON() string {
	return r.raw
}

type PostUnpublishResponse struct {
	// Post ID
	ID        string                       `json:"id" api:"required"`
	Content   string                       `json:"content" api:"required,nullable"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	Media     []PostUnpublishResponseMedia `json:"media" api:"required,nullable"`
	// When the post was published
	PublishedAt string `json:"published_at" api:"required,nullable"`
	// Source post ID if this is a recycled copy
	RecycledFromID string `json:"recycled_from_id" api:"required,nullable"`
	// Recycling configuration, if any
	Recycling   PostUnpublishResponseRecycling `json:"recycling" api:"required,nullable"`
	ScheduledAt string                         `json:"scheduled_at" api:"required,nullable"`
	Status      PostUnpublishResponseStatus    `json:"status" api:"required"`
	// Per-target results
	Targets   map[string]PostUnpublishResponseTarget `json:"targets" api:"required"`
	UpdatedAt time.Time                              `json:"updated_at" api:"required" format:"date-time"`
	// Engagement metrics (reactions, comments, views, etc.)
	Metrics PostUnpublishResponseMetrics `json:"metrics"`
	// Per-target customizations
	TargetOptions map[string]map[string]interface{} `json:"target_options" api:"nullable"`
	// Thread group ID (non-null if part of a thread)
	ThreadGroupID string `json:"thread_group_id" api:"nullable"`
	// Position within thread (0 = root)
	ThreadPosition float64 `json:"thread_position" api:"nullable"`
	// IANA timezone
	Timezone string                    `json:"timezone" api:"nullable"`
	JSON     postUnpublishResponseJSON `json:"-"`
}

// postUnpublishResponseJSON contains the JSON metadata for the struct
// [PostUnpublishResponse]
type postUnpublishResponseJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Media          apijson.Field
	PublishedAt    apijson.Field
	RecycledFromID apijson.Field
	Recycling      apijson.Field
	ScheduledAt    apijson.Field
	Status         apijson.Field
	Targets        apijson.Field
	UpdatedAt      apijson.Field
	Metrics        apijson.Field
	TargetOptions  apijson.Field
	ThreadGroupID  apijson.Field
	ThreadPosition apijson.Field
	Timezone       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostUnpublishResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseJSON) RawJSON() string {
	return r.raw
}

type PostUnpublishResponseMedia struct {
	// Public URL of the media file
	URL string `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail string `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type PostUnpublishResponseMediaType `json:"type"`
	JSON postUnpublishResponseMediaJSON `json:"-"`
}

// postUnpublishResponseMediaJSON contains the JSON metadata for the struct
// [PostUnpublishResponseMedia]
type postUnpublishResponseMediaJSON struct {
	URL         apijson.Field
	Thumbnail   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostUnpublishResponseMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseMediaJSON) RawJSON() string {
	return r.raw
}

// Media type. Inferred from URL extension if omitted.
type PostUnpublishResponseMediaType string

const (
	PostUnpublishResponseMediaTypeImage    PostUnpublishResponseMediaType = "image"
	PostUnpublishResponseMediaTypeVideo    PostUnpublishResponseMediaType = "video"
	PostUnpublishResponseMediaTypeGif      PostUnpublishResponseMediaType = "gif"
	PostUnpublishResponseMediaTypeDocument PostUnpublishResponseMediaType = "document"
)

func (r PostUnpublishResponseMediaType) IsKnown() bool {
	switch r {
	case PostUnpublishResponseMediaTypeImage, PostUnpublishResponseMediaTypeVideo, PostUnpublishResponseMediaTypeGif, PostUnpublishResponseMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration, if any
type PostUnpublishResponseRecycling struct {
	ID                    string                                `json:"id" api:"required"`
	ContentVariationIndex float64                               `json:"content_variation_index" api:"required"`
	ContentVariations     []string                              `json:"content_variations" api:"required"`
	CreatedAt             time.Time                             `json:"created_at" api:"required" format:"date-time"`
	Enabled               bool                                  `json:"enabled" api:"required"`
	ExpireCount           float64                               `json:"expire_count" api:"required,nullable"`
	ExpireDate            time.Time                             `json:"expire_date" api:"required,nullable" format:"date-time"`
	Gap                   float64                               `json:"gap" api:"required"`
	GapFreq               PostUnpublishResponseRecyclingGapFreq `json:"gap_freq" api:"required"`
	LastRecycledAt        time.Time                             `json:"last_recycled_at" api:"required,nullable" format:"date-time"`
	NextRecycleAt         time.Time                             `json:"next_recycle_at" api:"required,nullable" format:"date-time"`
	RecycleCount          float64                               `json:"recycle_count" api:"required"`
	StartDate             time.Time                             `json:"start_date" api:"required" format:"date-time"`
	UpdatedAt             time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	JSON                  postUnpublishResponseRecyclingJSON    `json:"-"`
}

// postUnpublishResponseRecyclingJSON contains the JSON metadata for the struct
// [PostUnpublishResponseRecycling]
type postUnpublishResponseRecyclingJSON struct {
	ID                    apijson.Field
	ContentVariationIndex apijson.Field
	ContentVariations     apijson.Field
	CreatedAt             apijson.Field
	Enabled               apijson.Field
	ExpireCount           apijson.Field
	ExpireDate            apijson.Field
	Gap                   apijson.Field
	GapFreq               apijson.Field
	LastRecycledAt        apijson.Field
	NextRecycleAt         apijson.Field
	RecycleCount          apijson.Field
	StartDate             apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PostUnpublishResponseRecycling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseRecyclingJSON) RawJSON() string {
	return r.raw
}

type PostUnpublishResponseRecyclingGapFreq string

const (
	PostUnpublishResponseRecyclingGapFreqDay   PostUnpublishResponseRecyclingGapFreq = "day"
	PostUnpublishResponseRecyclingGapFreqWeek  PostUnpublishResponseRecyclingGapFreq = "week"
	PostUnpublishResponseRecyclingGapFreqMonth PostUnpublishResponseRecyclingGapFreq = "month"
)

func (r PostUnpublishResponseRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostUnpublishResponseRecyclingGapFreqDay, PostUnpublishResponseRecyclingGapFreqWeek, PostUnpublishResponseRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostUnpublishResponseStatus string

const (
	PostUnpublishResponseStatusDraft      PostUnpublishResponseStatus = "draft"
	PostUnpublishResponseStatusScheduled  PostUnpublishResponseStatus = "scheduled"
	PostUnpublishResponseStatusPublishing PostUnpublishResponseStatus = "publishing"
	PostUnpublishResponseStatusPublished  PostUnpublishResponseStatus = "published"
	PostUnpublishResponseStatusFailed     PostUnpublishResponseStatus = "failed"
	PostUnpublishResponseStatusPartial    PostUnpublishResponseStatus = "partial"
)

func (r PostUnpublishResponseStatus) IsKnown() bool {
	switch r {
	case PostUnpublishResponseStatusDraft, PostUnpublishResponseStatusScheduled, PostUnpublishResponseStatusPublishing, PostUnpublishResponseStatusPublished, PostUnpublishResponseStatusFailed, PostUnpublishResponseStatusPartial:
		return true
	}
	return false
}

type PostUnpublishResponseTarget struct {
	Platform PostUnpublishResponseTargetsPlatform  `json:"platform" api:"required"`
	Status   PostUnpublishResponseTargetsStatus    `json:"status" api:"required"`
	Accounts []PostUnpublishResponseTargetsAccount `json:"accounts"`
	Error    PostUnpublishResponseTargetsError     `json:"error"`
	JSON     postUnpublishResponseTargetJSON       `json:"-"`
}

// postUnpublishResponseTargetJSON contains the JSON metadata for the struct
// [PostUnpublishResponseTarget]
type postUnpublishResponseTargetJSON struct {
	Platform    apijson.Field
	Status      apijson.Field
	Accounts    apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostUnpublishResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseTargetJSON) RawJSON() string {
	return r.raw
}

type PostUnpublishResponseTargetsPlatform string

const (
	PostUnpublishResponseTargetsPlatformTwitter        PostUnpublishResponseTargetsPlatform = "twitter"
	PostUnpublishResponseTargetsPlatformInstagram      PostUnpublishResponseTargetsPlatform = "instagram"
	PostUnpublishResponseTargetsPlatformFacebook       PostUnpublishResponseTargetsPlatform = "facebook"
	PostUnpublishResponseTargetsPlatformLinkedin       PostUnpublishResponseTargetsPlatform = "linkedin"
	PostUnpublishResponseTargetsPlatformTiktok         PostUnpublishResponseTargetsPlatform = "tiktok"
	PostUnpublishResponseTargetsPlatformYoutube        PostUnpublishResponseTargetsPlatform = "youtube"
	PostUnpublishResponseTargetsPlatformPinterest      PostUnpublishResponseTargetsPlatform = "pinterest"
	PostUnpublishResponseTargetsPlatformReddit         PostUnpublishResponseTargetsPlatform = "reddit"
	PostUnpublishResponseTargetsPlatformBluesky        PostUnpublishResponseTargetsPlatform = "bluesky"
	PostUnpublishResponseTargetsPlatformThreads        PostUnpublishResponseTargetsPlatform = "threads"
	PostUnpublishResponseTargetsPlatformTelegram       PostUnpublishResponseTargetsPlatform = "telegram"
	PostUnpublishResponseTargetsPlatformSnapchat       PostUnpublishResponseTargetsPlatform = "snapchat"
	PostUnpublishResponseTargetsPlatformGooglebusiness PostUnpublishResponseTargetsPlatform = "googlebusiness"
	PostUnpublishResponseTargetsPlatformWhatsapp       PostUnpublishResponseTargetsPlatform = "whatsapp"
	PostUnpublishResponseTargetsPlatformMastodon       PostUnpublishResponseTargetsPlatform = "mastodon"
	PostUnpublishResponseTargetsPlatformDiscord        PostUnpublishResponseTargetsPlatform = "discord"
	PostUnpublishResponseTargetsPlatformSMS            PostUnpublishResponseTargetsPlatform = "sms"
	PostUnpublishResponseTargetsPlatformBeehiiv        PostUnpublishResponseTargetsPlatform = "beehiiv"
	PostUnpublishResponseTargetsPlatformConvertkit     PostUnpublishResponseTargetsPlatform = "convertkit"
	PostUnpublishResponseTargetsPlatformMailchimp      PostUnpublishResponseTargetsPlatform = "mailchimp"
	PostUnpublishResponseTargetsPlatformListmonk       PostUnpublishResponseTargetsPlatform = "listmonk"
)

func (r PostUnpublishResponseTargetsPlatform) IsKnown() bool {
	switch r {
	case PostUnpublishResponseTargetsPlatformTwitter, PostUnpublishResponseTargetsPlatformInstagram, PostUnpublishResponseTargetsPlatformFacebook, PostUnpublishResponseTargetsPlatformLinkedin, PostUnpublishResponseTargetsPlatformTiktok, PostUnpublishResponseTargetsPlatformYoutube, PostUnpublishResponseTargetsPlatformPinterest, PostUnpublishResponseTargetsPlatformReddit, PostUnpublishResponseTargetsPlatformBluesky, PostUnpublishResponseTargetsPlatformThreads, PostUnpublishResponseTargetsPlatformTelegram, PostUnpublishResponseTargetsPlatformSnapchat, PostUnpublishResponseTargetsPlatformGooglebusiness, PostUnpublishResponseTargetsPlatformWhatsapp, PostUnpublishResponseTargetsPlatformMastodon, PostUnpublishResponseTargetsPlatformDiscord, PostUnpublishResponseTargetsPlatformSMS, PostUnpublishResponseTargetsPlatformBeehiiv, PostUnpublishResponseTargetsPlatformConvertkit, PostUnpublishResponseTargetsPlatformMailchimp, PostUnpublishResponseTargetsPlatformListmonk:
		return true
	}
	return false
}

type PostUnpublishResponseTargetsStatus string

const (
	PostUnpublishResponseTargetsStatusDraft      PostUnpublishResponseTargetsStatus = "draft"
	PostUnpublishResponseTargetsStatusScheduled  PostUnpublishResponseTargetsStatus = "scheduled"
	PostUnpublishResponseTargetsStatusPublishing PostUnpublishResponseTargetsStatus = "publishing"
	PostUnpublishResponseTargetsStatusPublished  PostUnpublishResponseTargetsStatus = "published"
	PostUnpublishResponseTargetsStatusFailed     PostUnpublishResponseTargetsStatus = "failed"
	PostUnpublishResponseTargetsStatusPartial    PostUnpublishResponseTargetsStatus = "partial"
)

func (r PostUnpublishResponseTargetsStatus) IsKnown() bool {
	switch r {
	case PostUnpublishResponseTargetsStatusDraft, PostUnpublishResponseTargetsStatusScheduled, PostUnpublishResponseTargetsStatusPublishing, PostUnpublishResponseTargetsStatusPublished, PostUnpublishResponseTargetsStatusFailed, PostUnpublishResponseTargetsStatusPartial:
		return true
	}
	return false
}

type PostUnpublishResponseTargetsAccount struct {
	ID string `json:"id" api:"required"`
	// Account avatar URL
	AvatarURL string `json:"avatar_url" api:"required,nullable"`
	// Account display name
	DisplayName string `json:"display_name" api:"required,nullable"`
	// Platform-native post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Post target ID (pt\_) — pass to /v1/ads/boost as post_target_id
	TargetID string `json:"target_id" api:"required,nullable"`
	// Published post URL on the platform
	URL      string                                  `json:"url" api:"required,nullable"`
	Username string                                  `json:"username" api:"required,nullable"`
	JSON     postUnpublishResponseTargetsAccountJSON `json:"-"`
}

// postUnpublishResponseTargetsAccountJSON contains the JSON metadata for the
// struct [PostUnpublishResponseTargetsAccount]
type postUnpublishResponseTargetsAccountJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	PlatformPostID apijson.Field
	TargetID       apijson.Field
	URL            apijson.Field
	Username       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostUnpublishResponseTargetsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseTargetsAccountJSON) RawJSON() string {
	return r.raw
}

type PostUnpublishResponseTargetsError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Raw platform error (HTTP status + response body), sanitized and truncated
	Detail string                                `json:"detail"`
	JSON   postUnpublishResponseTargetsErrorJSON `json:"-"`
}

// postUnpublishResponseTargetsErrorJSON contains the JSON metadata for the struct
// [PostUnpublishResponseTargetsError]
type postUnpublishResponseTargetsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostUnpublishResponseTargetsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseTargetsErrorJSON) RawJSON() string {
	return r.raw
}

// Engagement metrics (reactions, comments, views, etc.)
type PostUnpublishResponseMetrics struct {
	Clicks         float64                          `json:"clicks"`
	Comments       float64                          `json:"comments"`
	EngagementRate float64                          `json:"engagement_rate"`
	Impressions    float64                          `json:"impressions"`
	Likes          float64                          `json:"likes"`
	Reach          float64                          `json:"reach"`
	Saves          float64                          `json:"saves"`
	Shares         float64                          `json:"shares"`
	Views          float64                          `json:"views"`
	JSON           postUnpublishResponseMetricsJSON `json:"-"`
}

// postUnpublishResponseMetricsJSON contains the JSON metadata for the struct
// [PostUnpublishResponseMetrics]
type postUnpublishResponseMetricsJSON struct {
	Clicks         apijson.Field
	Comments       apijson.Field
	EngagementRate apijson.Field
	Impressions    apijson.Field
	Likes          apijson.Field
	Reach          apijson.Field
	Saves          apijson.Field
	Shares         apijson.Field
	Views          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostUnpublishResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postUnpublishResponseMetricsJSON) RawJSON() string {
	return r.raw
}

type PostNewParams struct {
	// Publish intent. Use "now" to publish immediately, "draft" to save as draft,
	// "auto" to auto-schedule to the best available slot, or an ISO 8601 timestamp to
	// schedule (max 30 days ahead).
	ScheduledAt param.Field[string] `json:"scheduled_at" api:"required"`
	// Account IDs, platform names, or workspace IDs to publish to
	Targets param.Field[[]string] `json:"targets" api:"required"`
	// Post text. Optional if target_options provide per-target content.
	Content param.Field[string] `json:"content"`
	// Cross-post actions to execute after publishing (e.g., repost from another
	// account, comment from another account)
	CrossPostActions param.Field[[]PostNewParamsCrossPostAction] `json:"cross_post_actions"`
	// Create post from an idea. Pre-fills content from the idea. Explicit 'content'
	// field takes precedence.
	IdeaID param.Field[string] `json:"idea_id"`
	// Media attachments
	Media param.Field[[]PostNewParamsMedia] `json:"media"`
	// Recycling configuration for evergreen content (Pro plan only)
	Recycling param.Field[PostNewParamsRecycling] `json:"recycling"`
	// Shorten URLs in post content. Only relevant when short link mode is 'ask'.
	// Ignored when mode is 'always' or 'never'. (Pro plan only)
	ShortenURLs param.Field[bool] `json:"shorten_urls"`
	// When true, the default signature is not auto-appended even if one is configured.
	SkipSignature param.Field[bool] `json:"skip_signature"`
	// Per-target customizations keyed by target value (account ID or platform name).
	// Supports platform-specific features such as Twitter polls (poll.options,
	// poll.duration_minutes), threads, reply_to, and reply_settings.
	TargetOptions param.Field[map[string]map[string]interface{}] `json:"target_options"`
	// Content template ID. When provided, the template content is used as the base for
	// the post. Explicit 'content' field takes precedence.
	TemplateID param.Field[string] `json:"template_id"`
	// Variables to interpolate in the template (e.g., { "promo_code": "SUMMER25" }).
	// Built-in variables: {{date}}, {{account_name}}.
	TemplateVariables param.Field[map[string]string] `json:"template_variables"`
	// IANA timezone for scheduling
	Timezone param.Field[string] `json:"timezone"`
	// Workspace ID to scope this post to
	WorkspaceID param.Field[string] `json:"workspace_id"`
}

func (r PostNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PostNewParamsCrossPostAction struct {
	// Type of cross-post action
	ActionType param.Field[PostNewParamsCrossPostActionsActionType] `json:"action_type" api:"required"`
	// Account to perform the action from
	TargetAccountID param.Field[string] `json:"target_account_id" api:"required"`
	// Text content for comment/quote actions (required for comment and quote)
	Content param.Field[string] `json:"content"`
	// Delay in minutes after publishing
	DelayMinutes param.Field[int64] `json:"delay_minutes"`
}

func (r PostNewParamsCrossPostAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of cross-post action
type PostNewParamsCrossPostActionsActionType string

const (
	PostNewParamsCrossPostActionsActionTypeRepost  PostNewParamsCrossPostActionsActionType = "repost"
	PostNewParamsCrossPostActionsActionTypeComment PostNewParamsCrossPostActionsActionType = "comment"
	PostNewParamsCrossPostActionsActionTypeQuote   PostNewParamsCrossPostActionsActionType = "quote"
)

func (r PostNewParamsCrossPostActionsActionType) IsKnown() bool {
	switch r {
	case PostNewParamsCrossPostActionsActionTypeRepost, PostNewParamsCrossPostActionsActionTypeComment, PostNewParamsCrossPostActionsActionTypeQuote:
		return true
	}
	return false
}

type PostNewParamsMedia struct {
	// Public URL of the media file
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail param.Field[string] `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type param.Field[PostNewParamsMediaType] `json:"type"`
}

func (r PostNewParamsMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media type. Inferred from URL extension if omitted.
type PostNewParamsMediaType string

const (
	PostNewParamsMediaTypeImage    PostNewParamsMediaType = "image"
	PostNewParamsMediaTypeVideo    PostNewParamsMediaType = "video"
	PostNewParamsMediaTypeGif      PostNewParamsMediaType = "gif"
	PostNewParamsMediaTypeDocument PostNewParamsMediaType = "document"
)

func (r PostNewParamsMediaType) IsKnown() bool {
	switch r {
	case PostNewParamsMediaTypeImage, PostNewParamsMediaTypeVideo, PostNewParamsMediaTypeGif, PostNewParamsMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration for evergreen content (Pro plan only)
type PostNewParamsRecycling struct {
	// Interval value
	Gap param.Field[int64] `json:"gap" api:"required"`
	// Interval unit
	GapFreq param.Field[PostNewParamsRecyclingGapFreq] `json:"gap_freq" api:"required"`
	// When to start recycling
	StartDate param.Field[time.Time] `json:"start_date" api:"required" format:"date-time"`
	// Alternate content texts (round-robin)
	ContentVariations param.Field[[]string] `json:"content_variations"`
	// Whether recycling is active
	Enabled param.Field[bool] `json:"enabled"`
	// Stop after this many recycles
	ExpireCount param.Field[int64] `json:"expire_count"`
	// Stop after this date
	ExpireDate param.Field[time.Time] `json:"expire_date" format:"date-time"`
}

func (r PostNewParamsRecycling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Interval unit
type PostNewParamsRecyclingGapFreq string

const (
	PostNewParamsRecyclingGapFreqDay   PostNewParamsRecyclingGapFreq = "day"
	PostNewParamsRecyclingGapFreqWeek  PostNewParamsRecyclingGapFreq = "week"
	PostNewParamsRecyclingGapFreqMonth PostNewParamsRecyclingGapFreq = "month"
)

func (r PostNewParamsRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostNewParamsRecyclingGapFreqDay, PostNewParamsRecyclingGapFreqWeek, PostNewParamsRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostUpdateParams struct {
	// Post text
	Content param.Field[string] `json:"content"`
	// Updated media
	Media param.Field[[]PostUpdateParamsMedia] `json:"media"`
	// Internal notes for this post
	Notes param.Field[string] `json:"notes"`
	// Recycling configuration (Pro plan only)
	Recycling param.Field[PostUpdateParamsRecycling] `json:"recycling"`
	// Publish intent. Use "now" to publish immediately, "draft" to save as draft,
	// "auto" to auto-schedule to the best available slot, or an ISO 8601 timestamp to
	// schedule (max 30 days ahead).
	ScheduledAt   param.Field[string]                            `json:"scheduled_at"`
	TargetOptions param.Field[map[string]map[string]interface{}] `json:"target_options"`
	// Updated targets
	Targets  param.Field[[]string] `json:"targets"`
	Timezone param.Field[string]   `json:"timezone"`
}

func (r PostUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PostUpdateParamsMedia struct {
	// Public URL of the media file
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail param.Field[string] `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type param.Field[PostUpdateParamsMediaType] `json:"type"`
}

func (r PostUpdateParamsMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media type. Inferred from URL extension if omitted.
type PostUpdateParamsMediaType string

const (
	PostUpdateParamsMediaTypeImage    PostUpdateParamsMediaType = "image"
	PostUpdateParamsMediaTypeVideo    PostUpdateParamsMediaType = "video"
	PostUpdateParamsMediaTypeGif      PostUpdateParamsMediaType = "gif"
	PostUpdateParamsMediaTypeDocument PostUpdateParamsMediaType = "document"
)

func (r PostUpdateParamsMediaType) IsKnown() bool {
	switch r {
	case PostUpdateParamsMediaTypeImage, PostUpdateParamsMediaTypeVideo, PostUpdateParamsMediaTypeGif, PostUpdateParamsMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration (Pro plan only)
type PostUpdateParamsRecycling struct {
	// Interval value
	Gap param.Field[int64] `json:"gap" api:"required"`
	// Interval unit
	GapFreq param.Field[PostUpdateParamsRecyclingGapFreq] `json:"gap_freq" api:"required"`
	// When to start recycling
	StartDate param.Field[time.Time] `json:"start_date" api:"required" format:"date-time"`
	// Alternate content texts (round-robin)
	ContentVariations param.Field[[]string] `json:"content_variations"`
	// Whether recycling is active
	Enabled param.Field[bool] `json:"enabled"`
	// Stop after this many recycles
	ExpireCount param.Field[int64] `json:"expire_count"`
	// Stop after this date
	ExpireDate param.Field[time.Time] `json:"expire_date" format:"date-time"`
}

func (r PostUpdateParamsRecycling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Interval unit
type PostUpdateParamsRecyclingGapFreq string

const (
	PostUpdateParamsRecyclingGapFreqDay   PostUpdateParamsRecyclingGapFreq = "day"
	PostUpdateParamsRecyclingGapFreqWeek  PostUpdateParamsRecyclingGapFreq = "week"
	PostUpdateParamsRecyclingGapFreqMonth PostUpdateParamsRecyclingGapFreq = "month"
)

func (r PostUpdateParamsRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostUpdateParamsRecyclingGapFreqDay, PostUpdateParamsRecyclingGapFreqWeek, PostUpdateParamsRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostListParams struct {
	// Filter by specific account ID
	AccountID param.Field[string] `query:"account_id"`
	// Filter by any of several account IDs (comma-separated). Takes precedence over
	// account_id.
	AccountIDs param.Field[string] `query:"account_ids"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Filter: start date (ISO 8601)
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Comma-separated list of fields to include in the response (e.g. 'targets,media')
	Include param.Field[string] `query:"include"`
	// When true, also return external posts merged by published_at (works with
	// status=published or no status filter)
	IncludeExternal param.Field[PostListParamsIncludeExternal] `query:"include_external"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Filter by post status
	Status param.Field[PostListParamsStatus] `query:"status"`
	// Filter: end date (ISO 8601)
	To param.Field[time.Time] `query:"to" format:"date-time"`
	// Filter by workspace ID
	WorkspaceID param.Field[string] `query:"workspace_id"`
}

// URLQuery serializes [PostListParams]'s query parameters as `url.Values`.
func (r PostListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When true, also return external posts merged by published_at (works with
// status=published or no status filter)
type PostListParamsIncludeExternal string

const (
	PostListParamsIncludeExternalTrue  PostListParamsIncludeExternal = "true"
	PostListParamsIncludeExternalFalse PostListParamsIncludeExternal = "false"
)

func (r PostListParamsIncludeExternal) IsKnown() bool {
	switch r {
	case PostListParamsIncludeExternalTrue, PostListParamsIncludeExternalFalse:
		return true
	}
	return false
}

// Filter by post status
type PostListParamsStatus string

const (
	PostListParamsStatusDraft      PostListParamsStatus = "draft"
	PostListParamsStatusScheduled  PostListParamsStatus = "scheduled"
	PostListParamsStatusPublishing PostListParamsStatus = "publishing"
	PostListParamsStatusPublished  PostListParamsStatus = "published"
	PostListParamsStatusFailed     PostListParamsStatus = "failed"
)

func (r PostListParamsStatus) IsKnown() bool {
	switch r {
	case PostListParamsStatusDraft, PostListParamsStatusScheduled, PostListParamsStatusPublishing, PostListParamsStatusPublished, PostListParamsStatusFailed:
		return true
	}
	return false
}

type PostBulkNewParams struct {
	// Array of posts to create (max 50)
	Posts param.Field[[]PostBulkNewParamsPost] `json:"posts" api:"required"`
}

func (r PostBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PostBulkNewParamsPost struct {
	// Publish intent. Use "now" to publish immediately, "draft" to save as draft,
	// "auto" to auto-schedule to the best available slot, or an ISO 8601 timestamp to
	// schedule (max 30 days ahead).
	ScheduledAt param.Field[string] `json:"scheduled_at" api:"required"`
	// Account IDs, platform names, or workspace IDs to publish to
	Targets param.Field[[]string] `json:"targets" api:"required"`
	// Post text. Optional if target_options provide per-target content.
	Content param.Field[string] `json:"content"`
	// Cross-post actions to execute after publishing (e.g., repost from another
	// account, comment from another account)
	CrossPostActions param.Field[[]PostBulkNewParamsPostsCrossPostAction] `json:"cross_post_actions"`
	// Create post from an idea. Pre-fills content from the idea. Explicit 'content'
	// field takes precedence.
	IdeaID param.Field[string] `json:"idea_id"`
	// Media attachments
	Media param.Field[[]PostBulkNewParamsPostsMedia] `json:"media"`
	// Recycling configuration for evergreen content (Pro plan only)
	Recycling param.Field[PostBulkNewParamsPostsRecycling] `json:"recycling"`
	// Shorten URLs in post content. Only relevant when short link mode is 'ask'.
	// Ignored when mode is 'always' or 'never'. (Pro plan only)
	ShortenURLs param.Field[bool] `json:"shorten_urls"`
	// When true, the default signature is not auto-appended even if one is configured.
	SkipSignature param.Field[bool] `json:"skip_signature"`
	// Per-target customizations keyed by target value (account ID or platform name).
	// Supports platform-specific features such as Twitter polls (poll.options,
	// poll.duration_minutes), threads, reply_to, and reply_settings.
	TargetOptions param.Field[map[string]map[string]interface{}] `json:"target_options"`
	// Content template ID. When provided, the template content is used as the base for
	// the post. Explicit 'content' field takes precedence.
	TemplateID param.Field[string] `json:"template_id"`
	// Variables to interpolate in the template (e.g., { "promo_code": "SUMMER25" }).
	// Built-in variables: {{date}}, {{account_name}}.
	TemplateVariables param.Field[map[string]string] `json:"template_variables"`
	// IANA timezone for scheduling
	Timezone param.Field[string] `json:"timezone"`
	// Workspace ID to scope this post to
	WorkspaceID param.Field[string] `json:"workspace_id"`
}

func (r PostBulkNewParamsPost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PostBulkNewParamsPostsCrossPostAction struct {
	// Type of cross-post action
	ActionType param.Field[PostBulkNewParamsPostsCrossPostActionsActionType] `json:"action_type" api:"required"`
	// Account to perform the action from
	TargetAccountID param.Field[string] `json:"target_account_id" api:"required"`
	// Text content for comment/quote actions (required for comment and quote)
	Content param.Field[string] `json:"content"`
	// Delay in minutes after publishing
	DelayMinutes param.Field[int64] `json:"delay_minutes"`
}

func (r PostBulkNewParamsPostsCrossPostAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of cross-post action
type PostBulkNewParamsPostsCrossPostActionsActionType string

const (
	PostBulkNewParamsPostsCrossPostActionsActionTypeRepost  PostBulkNewParamsPostsCrossPostActionsActionType = "repost"
	PostBulkNewParamsPostsCrossPostActionsActionTypeComment PostBulkNewParamsPostsCrossPostActionsActionType = "comment"
	PostBulkNewParamsPostsCrossPostActionsActionTypeQuote   PostBulkNewParamsPostsCrossPostActionsActionType = "quote"
)

func (r PostBulkNewParamsPostsCrossPostActionsActionType) IsKnown() bool {
	switch r {
	case PostBulkNewParamsPostsCrossPostActionsActionTypeRepost, PostBulkNewParamsPostsCrossPostActionsActionTypeComment, PostBulkNewParamsPostsCrossPostActionsActionTypeQuote:
		return true
	}
	return false
}

type PostBulkNewParamsPostsMedia struct {
	// Public URL of the media file
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Read-only. Stable, hyper-optimized preview URL that persists after the full-res
	// original expires. Ignored on write.
	Thumbnail param.Field[string] `json:"thumbnail"`
	// Media type. Inferred from URL extension if omitted.
	Type param.Field[PostBulkNewParamsPostsMediaType] `json:"type"`
}

func (r PostBulkNewParamsPostsMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media type. Inferred from URL extension if omitted.
type PostBulkNewParamsPostsMediaType string

const (
	PostBulkNewParamsPostsMediaTypeImage    PostBulkNewParamsPostsMediaType = "image"
	PostBulkNewParamsPostsMediaTypeVideo    PostBulkNewParamsPostsMediaType = "video"
	PostBulkNewParamsPostsMediaTypeGif      PostBulkNewParamsPostsMediaType = "gif"
	PostBulkNewParamsPostsMediaTypeDocument PostBulkNewParamsPostsMediaType = "document"
)

func (r PostBulkNewParamsPostsMediaType) IsKnown() bool {
	switch r {
	case PostBulkNewParamsPostsMediaTypeImage, PostBulkNewParamsPostsMediaTypeVideo, PostBulkNewParamsPostsMediaTypeGif, PostBulkNewParamsPostsMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration for evergreen content (Pro plan only)
type PostBulkNewParamsPostsRecycling struct {
	// Interval value
	Gap param.Field[int64] `json:"gap" api:"required"`
	// Interval unit
	GapFreq param.Field[PostBulkNewParamsPostsRecyclingGapFreq] `json:"gap_freq" api:"required"`
	// When to start recycling
	StartDate param.Field[time.Time] `json:"start_date" api:"required" format:"date-time"`
	// Alternate content texts (round-robin)
	ContentVariations param.Field[[]string] `json:"content_variations"`
	// Whether recycling is active
	Enabled param.Field[bool] `json:"enabled"`
	// Stop after this many recycles
	ExpireCount param.Field[int64] `json:"expire_count"`
	// Stop after this date
	ExpireDate param.Field[time.Time] `json:"expire_date" format:"date-time"`
}

func (r PostBulkNewParamsPostsRecycling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Interval unit
type PostBulkNewParamsPostsRecyclingGapFreq string

const (
	PostBulkNewParamsPostsRecyclingGapFreqDay   PostBulkNewParamsPostsRecyclingGapFreq = "day"
	PostBulkNewParamsPostsRecyclingGapFreqWeek  PostBulkNewParamsPostsRecyclingGapFreq = "week"
	PostBulkNewParamsPostsRecyclingGapFreqMonth PostBulkNewParamsPostsRecyclingGapFreq = "month"
)

func (r PostBulkNewParamsPostsRecyclingGapFreq) IsKnown() bool {
	switch r {
	case PostBulkNewParamsPostsRecyclingGapFreqDay, PostBulkNewParamsPostsRecyclingGapFreqWeek, PostBulkNewParamsPostsRecyclingGapFreqMonth:
		return true
	}
	return false
}

type PostUnpublishParams struct {
	// Platforms to unpublish from. If omitted, unpublishes from all.
	Platforms param.Field[[]string] `json:"platforms"`
}

func (r PostUnpublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
