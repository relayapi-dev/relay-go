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

// ToolValidateService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolValidateService] method instead.
type ToolValidateService struct {
	Options []option.RequestOption
}

// NewToolValidateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolValidateService(opts ...option.RequestOption) (r *ToolValidateService) {
	r = &ToolValidateService{}
	r.Options = opts
	return
}

// Check character counts against platform limits
func (r *ToolValidateService) CheckPostLength(ctx context.Context, body ToolValidateCheckPostLengthParams, opts ...option.RequestOption) (res *ToolValidateCheckPostLengthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/validate/post-length"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Check if a subreddit exists and get its details
func (r *ToolValidateService) GetSubreddit(ctx context.Context, query ToolValidateGetSubredditParams, opts ...option.RequestOption) (res *ToolValidateGetSubredditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/validate/subreddit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Validate a media URL for platform compatibility
func (r *ToolValidateService) ValidateMedia(ctx context.Context, body ToolValidateValidateMediaParams, opts ...option.RequestOption) (res *ToolValidateValidateMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/validate/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Validate a post (dry-run without publishing)
func (r *ToolValidateService) ValidatePost(ctx context.Context, body ToolValidateValidatePostParams, opts ...option.RequestOption) (res *ToolValidateValidatePostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/validate/post"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ToolValidateCheckPostLengthResponse struct {
	// Character count per platform
	Platforms ToolValidateCheckPostLengthResponsePlatforms `json:"platforms" api:"required"`
	JSON      toolValidateCheckPostLengthResponseJSON      `json:"-"`
}

// toolValidateCheckPostLengthResponseJSON contains the JSON metadata for the
// struct [ToolValidateCheckPostLengthResponse]
type toolValidateCheckPostLengthResponseJSON struct {
	Platforms   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponseJSON) RawJSON() string {
	return r.raw
}

// Character count per platform
type ToolValidateCheckPostLengthResponsePlatforms struct {
	Bluesky        ToolValidateCheckPostLengthResponsePlatformsBluesky        `json:"bluesky"`
	Discord        ToolValidateCheckPostLengthResponsePlatformsDiscord        `json:"discord"`
	Facebook       ToolValidateCheckPostLengthResponsePlatformsFacebook       `json:"facebook"`
	Googlebusiness ToolValidateCheckPostLengthResponsePlatformsGooglebusiness `json:"googlebusiness"`
	Instagram      ToolValidateCheckPostLengthResponsePlatformsInstagram      `json:"instagram"`
	Linkedin       ToolValidateCheckPostLengthResponsePlatformsLinkedin       `json:"linkedin"`
	Mastodon       ToolValidateCheckPostLengthResponsePlatformsMastodon       `json:"mastodon"`
	Pinterest      ToolValidateCheckPostLengthResponsePlatformsPinterest      `json:"pinterest"`
	Reddit         ToolValidateCheckPostLengthResponsePlatformsReddit         `json:"reddit"`
	SMS            ToolValidateCheckPostLengthResponsePlatformsSMS            `json:"sms"`
	Snapchat       ToolValidateCheckPostLengthResponsePlatformsSnapchat       `json:"snapchat"`
	Telegram       ToolValidateCheckPostLengthResponsePlatformsTelegram       `json:"telegram"`
	Threads        ToolValidateCheckPostLengthResponsePlatformsThreads        `json:"threads"`
	Tiktok         ToolValidateCheckPostLengthResponsePlatformsTiktok         `json:"tiktok"`
	Twitter        ToolValidateCheckPostLengthResponsePlatformsTwitter        `json:"twitter"`
	Whatsapp       ToolValidateCheckPostLengthResponsePlatformsWhatsapp       `json:"whatsapp"`
	Youtube        ToolValidateCheckPostLengthResponsePlatformsYoutube        `json:"youtube"`
	JSON           toolValidateCheckPostLengthResponsePlatformsJSON           `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsJSON contains the JSON metadata for
// the struct [ToolValidateCheckPostLengthResponsePlatforms]
type toolValidateCheckPostLengthResponsePlatformsJSON struct {
	Bluesky        apijson.Field
	Discord        apijson.Field
	Facebook       apijson.Field
	Googlebusiness apijson.Field
	Instagram      apijson.Field
	Linkedin       apijson.Field
	Mastodon       apijson.Field
	Pinterest      apijson.Field
	Reddit         apijson.Field
	SMS            apijson.Field
	Snapchat       apijson.Field
	Telegram       apijson.Field
	Threads        apijson.Field
	Tiktok         apijson.Field
	Twitter        apijson.Field
	Whatsapp       apijson.Field
	Youtube        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatforms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsBluesky struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                    `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsBlueskyJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsBlueskyJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsBluesky]
type toolValidateCheckPostLengthResponsePlatformsBlueskyJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsBluesky) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsBlueskyJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsDiscord struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                    `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsDiscordJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsDiscordJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsDiscord]
type toolValidateCheckPostLengthResponsePlatformsDiscordJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsDiscord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsDiscordJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsFacebook struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                     `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsFacebookJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsFacebookJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsFacebook]
type toolValidateCheckPostLengthResponsePlatformsFacebookJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsFacebookJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsGooglebusiness struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                           `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsGooglebusinessJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsGooglebusinessJSON contains the JSON
// metadata for the struct
// [ToolValidateCheckPostLengthResponsePlatformsGooglebusiness]
type toolValidateCheckPostLengthResponsePlatformsGooglebusinessJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsGooglebusiness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsGooglebusinessJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsInstagram struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                      `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsInstagramJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsInstagramJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsInstagram]
type toolValidateCheckPostLengthResponsePlatformsInstagramJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsInstagram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsInstagramJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsLinkedin struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                     `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsLinkedinJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsLinkedinJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsLinkedin]
type toolValidateCheckPostLengthResponsePlatformsLinkedinJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsLinkedinJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsMastodon struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                     `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsMastodonJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsMastodonJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsMastodon]
type toolValidateCheckPostLengthResponsePlatformsMastodonJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsMastodon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsMastodonJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsPinterest struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                      `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsPinterestJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsPinterestJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsPinterest]
type toolValidateCheckPostLengthResponsePlatformsPinterestJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsPinterest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsPinterestJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsReddit struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                   `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsRedditJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsRedditJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsReddit]
type toolValidateCheckPostLengthResponsePlatformsRedditJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsReddit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsRedditJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsSMS struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsSMSJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsSMSJSON contains the JSON metadata
// for the struct [ToolValidateCheckPostLengthResponsePlatformsSMS]
type toolValidateCheckPostLengthResponsePlatformsSMSJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsSMS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsSMSJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsSnapchat struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                     `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsSnapchatJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsSnapchatJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsSnapchat]
type toolValidateCheckPostLengthResponsePlatformsSnapchatJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsSnapchat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsSnapchatJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsTelegram struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                     `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsTelegramJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsTelegramJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsTelegram]
type toolValidateCheckPostLengthResponsePlatformsTelegramJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsTelegram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsTelegramJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsThreads struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                    `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsThreadsJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsThreadsJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsThreads]
type toolValidateCheckPostLengthResponsePlatformsThreadsJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsThreads) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsThreadsJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsTiktok struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                   `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsTiktokJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsTiktokJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsTiktok]
type toolValidateCheckPostLengthResponsePlatformsTiktokJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsTiktok) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsTiktokJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsTwitter struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                    `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsTwitterJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsTwitterJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsTwitter]
type toolValidateCheckPostLengthResponsePlatformsTwitterJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsTwitter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsTwitterJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsWhatsapp struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                     `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsWhatsappJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsWhatsappJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsWhatsapp]
type toolValidateCheckPostLengthResponsePlatformsWhatsappJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsWhatsapp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsWhatsappJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthResponsePlatformsYoutube struct {
	// Character count for this platform
	Count float64 `json:"count" api:"required"`
	// Character limit for this platform
	Limit float64 `json:"limit" api:"required"`
	// Whether content is within limit
	WithinLimit bool                                                    `json:"within_limit" api:"required"`
	JSON        toolValidateCheckPostLengthResponsePlatformsYoutubeJSON `json:"-"`
}

// toolValidateCheckPostLengthResponsePlatformsYoutubeJSON contains the JSON
// metadata for the struct [ToolValidateCheckPostLengthResponsePlatformsYoutube]
type toolValidateCheckPostLengthResponsePlatformsYoutubeJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	WithinLimit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateCheckPostLengthResponsePlatformsYoutube) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateCheckPostLengthResponsePlatformsYoutubeJSON) RawJSON() string {
	return r.raw
}

type ToolValidateGetSubredditResponse struct {
	// Whether the subreddit exists
	Exists bool `json:"exists" api:"required"`
	// Canonical subreddit name
	Name string `json:"name" api:"nullable"`
	// Whether NSFW
	Nsfw bool `json:"nsfw" api:"nullable"`
	// Allowed post types
	PostTypes ToolValidateGetSubredditResponsePostTypes `json:"post_types"`
	// Subscriber count
	Subscribers float64 `json:"subscribers" api:"nullable"`
	// Subreddit title
	Title string                               `json:"title" api:"nullable"`
	JSON  toolValidateGetSubredditResponseJSON `json:"-"`
}

// toolValidateGetSubredditResponseJSON contains the JSON metadata for the struct
// [ToolValidateGetSubredditResponse]
type toolValidateGetSubredditResponseJSON struct {
	Exists      apijson.Field
	Name        apijson.Field
	Nsfw        apijson.Field
	PostTypes   apijson.Field
	Subscribers apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateGetSubredditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateGetSubredditResponseJSON) RawJSON() string {
	return r.raw
}

// Allowed post types
type ToolValidateGetSubredditResponsePostTypes struct {
	// Allows image posts
	Image bool `json:"image" api:"required"`
	// Allows link posts
	Link bool `json:"link" api:"required"`
	// Allows text posts
	Self bool `json:"self" api:"required"`
	// Allows video posts
	Video bool                                          `json:"video"`
	JSON  toolValidateGetSubredditResponsePostTypesJSON `json:"-"`
}

// toolValidateGetSubredditResponsePostTypesJSON contains the JSON metadata for the
// struct [ToolValidateGetSubredditResponsePostTypes]
type toolValidateGetSubredditResponsePostTypesJSON struct {
	Image       apijson.Field
	Link        apijson.Field
	Self        apijson.Field
	Video       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateGetSubredditResponsePostTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateGetSubredditResponsePostTypesJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponse struct {
	// Whether the URL is accessible
	Accessible bool `json:"accessible" api:"required"`
	// Per-platform size limits
	PlatformLimits ToolValidateValidateMediaResponsePlatformLimits `json:"platform_limits" api:"required"`
	// MIME type
	ContentType string `json:"content_type" api:"nullable"`
	// File size in bytes
	Size float64                               `json:"size" api:"nullable"`
	JSON toolValidateValidateMediaResponseJSON `json:"-"`
}

// toolValidateValidateMediaResponseJSON contains the JSON metadata for the struct
// [ToolValidateValidateMediaResponse]
type toolValidateValidateMediaResponseJSON struct {
	Accessible     apijson.Field
	PlatformLimits apijson.Field
	ContentType    apijson.Field
	Size           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponseJSON) RawJSON() string {
	return r.raw
}

// Per-platform size limits
type ToolValidateValidateMediaResponsePlatformLimits struct {
	Bluesky        ToolValidateValidateMediaResponsePlatformLimitsBluesky        `json:"bluesky"`
	Discord        ToolValidateValidateMediaResponsePlatformLimitsDiscord        `json:"discord"`
	Facebook       ToolValidateValidateMediaResponsePlatformLimitsFacebook       `json:"facebook"`
	Googlebusiness ToolValidateValidateMediaResponsePlatformLimitsGooglebusiness `json:"googlebusiness"`
	Instagram      ToolValidateValidateMediaResponsePlatformLimitsInstagram      `json:"instagram"`
	Linkedin       ToolValidateValidateMediaResponsePlatformLimitsLinkedin       `json:"linkedin"`
	Mastodon       ToolValidateValidateMediaResponsePlatformLimitsMastodon       `json:"mastodon"`
	Pinterest      ToolValidateValidateMediaResponsePlatformLimitsPinterest      `json:"pinterest"`
	Reddit         ToolValidateValidateMediaResponsePlatformLimitsReddit         `json:"reddit"`
	SMS            ToolValidateValidateMediaResponsePlatformLimitsSMS            `json:"sms"`
	Snapchat       ToolValidateValidateMediaResponsePlatformLimitsSnapchat       `json:"snapchat"`
	Telegram       ToolValidateValidateMediaResponsePlatformLimitsTelegram       `json:"telegram"`
	Threads        ToolValidateValidateMediaResponsePlatformLimitsThreads        `json:"threads"`
	Tiktok         ToolValidateValidateMediaResponsePlatformLimitsTiktok         `json:"tiktok"`
	Twitter        ToolValidateValidateMediaResponsePlatformLimitsTwitter        `json:"twitter"`
	Whatsapp       ToolValidateValidateMediaResponsePlatformLimitsWhatsapp       `json:"whatsapp"`
	Youtube        ToolValidateValidateMediaResponsePlatformLimitsYoutube        `json:"youtube"`
	JSON           toolValidateValidateMediaResponsePlatformLimitsJSON           `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsJSON contains the JSON metadata
// for the struct [ToolValidateValidateMediaResponsePlatformLimits]
type toolValidateValidateMediaResponsePlatformLimitsJSON struct {
	Bluesky        apijson.Field
	Discord        apijson.Field
	Facebook       apijson.Field
	Googlebusiness apijson.Field
	Instagram      apijson.Field
	Linkedin       apijson.Field
	Mastodon       apijson.Field
	Pinterest      apijson.Field
	Reddit         apijson.Field
	SMS            apijson.Field
	Snapchat       apijson.Field
	Telegram       apijson.Field
	Threads        apijson.Field
	Tiktok         apijson.Field
	Twitter        apijson.Field
	Whatsapp       apijson.Field
	Youtube        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsBluesky struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                       `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsBlueskyJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsBlueskyJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsBluesky]
type toolValidateValidateMediaResponsePlatformLimitsBlueskyJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsBluesky) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsBlueskyJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsDiscord struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                       `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsDiscordJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsDiscordJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsDiscord]
type toolValidateValidateMediaResponsePlatformLimitsDiscordJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsDiscord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsDiscordJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsFacebook struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                        `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsFacebookJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsFacebookJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsFacebook]
type toolValidateValidateMediaResponsePlatformLimitsFacebookJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsFacebookJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsGooglebusiness struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                              `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsGooglebusinessJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsGooglebusinessJSON contains the
// JSON metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsGooglebusiness]
type toolValidateValidateMediaResponsePlatformLimitsGooglebusinessJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsGooglebusiness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsGooglebusinessJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsInstagram struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                         `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsInstagramJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsInstagramJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsInstagram]
type toolValidateValidateMediaResponsePlatformLimitsInstagramJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsInstagram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsInstagramJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsLinkedin struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                        `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsLinkedinJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsLinkedinJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsLinkedin]
type toolValidateValidateMediaResponsePlatformLimitsLinkedinJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsLinkedinJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsMastodon struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                        `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsMastodonJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsMastodonJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsMastodon]
type toolValidateValidateMediaResponsePlatformLimitsMastodonJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsMastodon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsMastodonJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsPinterest struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                         `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsPinterestJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsPinterestJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsPinterest]
type toolValidateValidateMediaResponsePlatformLimitsPinterestJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsPinterest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsPinterestJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsReddit struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                      `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsRedditJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsRedditJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsReddit]
type toolValidateValidateMediaResponsePlatformLimitsRedditJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsReddit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsRedditJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsSMS struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                   `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsSMSJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsSMSJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsSMS]
type toolValidateValidateMediaResponsePlatformLimitsSMSJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsSMS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsSMSJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsSnapchat struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                        `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsSnapchatJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsSnapchatJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsSnapchat]
type toolValidateValidateMediaResponsePlatformLimitsSnapchatJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsSnapchat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsSnapchatJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsTelegram struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                        `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsTelegramJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsTelegramJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsTelegram]
type toolValidateValidateMediaResponsePlatformLimitsTelegramJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsTelegram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsTelegramJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsThreads struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                       `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsThreadsJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsThreadsJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsThreads]
type toolValidateValidateMediaResponsePlatformLimitsThreadsJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsThreads) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsThreadsJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsTiktok struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                      `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsTiktokJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsTiktokJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsTiktok]
type toolValidateValidateMediaResponsePlatformLimitsTiktokJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsTiktok) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsTiktokJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsTwitter struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                       `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsTwitterJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsTwitterJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsTwitter]
type toolValidateValidateMediaResponsePlatformLimitsTwitterJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsTwitter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsTwitterJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsWhatsapp struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                        `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsWhatsappJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsWhatsappJSON contains the JSON
// metadata for the struct
// [ToolValidateValidateMediaResponsePlatformLimitsWhatsapp]
type toolValidateValidateMediaResponsePlatformLimitsWhatsappJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsWhatsapp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsWhatsappJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidateMediaResponsePlatformLimitsYoutube struct {
	// Maximum file size in bytes
	MaxSize float64 `json:"max_size" api:"required"`
	// Whether file size is within limit
	WithinLimit bool `json:"within_limit" api:"required"`
	// Whether the MIME type is supported by this platform
	MimeTypeSupported bool                                                       `json:"mime_type_supported"`
	JSON              toolValidateValidateMediaResponsePlatformLimitsYoutubeJSON `json:"-"`
}

// toolValidateValidateMediaResponsePlatformLimitsYoutubeJSON contains the JSON
// metadata for the struct [ToolValidateValidateMediaResponsePlatformLimitsYoutube]
type toolValidateValidateMediaResponsePlatformLimitsYoutubeJSON struct {
	MaxSize           apijson.Field
	WithinLimit       apijson.Field
	MimeTypeSupported apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ToolValidateValidateMediaResponsePlatformLimitsYoutube) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidateMediaResponsePlatformLimitsYoutubeJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidatePostResponse struct {
	// Blocking errors
	Errors []ToolValidateValidatePostResponseError `json:"errors" api:"required"`
	// Whether the post is valid for all targets
	Valid bool `json:"valid" api:"required"`
	// Non-blocking warnings
	Warnings []ToolValidateValidatePostResponseWarning `json:"warnings" api:"required"`
	JSON     toolValidateValidatePostResponseJSON      `json:"-"`
}

// toolValidateValidatePostResponseJSON contains the JSON metadata for the struct
// [ToolValidateValidatePostResponse]
type toolValidateValidatePostResponseJSON struct {
	Errors      apijson.Field
	Valid       apijson.Field
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateValidatePostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidatePostResponseJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidatePostResponseError struct {
	// Error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Target identifier (account ID, platform, or field name)
	Target string                                    `json:"target" api:"required"`
	JSON   toolValidateValidatePostResponseErrorJSON `json:"-"`
}

// toolValidateValidatePostResponseErrorJSON contains the JSON metadata for the
// struct [ToolValidateValidatePostResponseError]
type toolValidateValidatePostResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateValidatePostResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidatePostResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ToolValidateValidatePostResponseWarning struct {
	// Error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Target identifier (account ID, platform, or field name)
	Target string                                      `json:"target" api:"required"`
	JSON   toolValidateValidatePostResponseWarningJSON `json:"-"`
}

// toolValidateValidatePostResponseWarningJSON contains the JSON metadata for the
// struct [ToolValidateValidatePostResponseWarning]
type toolValidateValidatePostResponseWarningJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolValidateValidatePostResponseWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolValidateValidatePostResponseWarningJSON) RawJSON() string {
	return r.raw
}

type ToolValidateCheckPostLengthParams struct {
	// Post content to check
	Content param.Field[string] `json:"content" api:"required"`
}

func (r ToolValidateCheckPostLengthParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolValidateGetSubredditParams struct {
	// Subreddit name (without r/ prefix)
	Name param.Field[string] `query:"name" api:"required"`
}

// URLQuery serializes [ToolValidateGetSubredditParams]'s query parameters as
// `url.Values`.
func (r ToolValidateGetSubredditParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolValidateValidateMediaParams struct {
	// Media URL to validate
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
}

func (r ToolValidateValidateMediaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolValidateValidatePostParams struct {
	// Publish intent. Use "now" to publish immediately, "draft" to save as draft, or
	// an ISO 8601 timestamp to schedule.
	ScheduledAt param.Field[string] `json:"scheduled_at" api:"required"`
	// Account IDs, platform names, or workspace IDs to publish to
	Targets param.Field[[]string] `json:"targets" api:"required"`
	// Post text. Optional if target_options provide per-target content.
	Content param.Field[string] `json:"content"`
	// Cross-post actions to execute after publishing (e.g., repost from another
	// account, comment from another account)
	CrossPostActions param.Field[[]ToolValidateValidatePostParamsCrossPostAction] `json:"cross_post_actions"`
	// Media attachments
	Media param.Field[[]ToolValidateValidatePostParamsMedia] `json:"media"`
	// Recycling configuration for evergreen content (Pro plan only)
	Recycling param.Field[ToolValidateValidatePostParamsRecycling] `json:"recycling"`
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

func (r ToolValidateValidatePostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolValidateValidatePostParamsCrossPostAction struct {
	// Type of cross-post action
	ActionType param.Field[ToolValidateValidatePostParamsCrossPostActionsActionType] `json:"action_type" api:"required"`
	// Account to perform the action from
	TargetAccountID param.Field[string] `json:"target_account_id" api:"required"`
	// Text content for comment/quote actions (required for comment and quote)
	Content param.Field[string] `json:"content"`
	// Delay in minutes after publishing
	DelayMinutes param.Field[int64] `json:"delay_minutes"`
}

func (r ToolValidateValidatePostParamsCrossPostAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of cross-post action
type ToolValidateValidatePostParamsCrossPostActionsActionType string

const (
	ToolValidateValidatePostParamsCrossPostActionsActionTypeRepost  ToolValidateValidatePostParamsCrossPostActionsActionType = "repost"
	ToolValidateValidatePostParamsCrossPostActionsActionTypeComment ToolValidateValidatePostParamsCrossPostActionsActionType = "comment"
	ToolValidateValidatePostParamsCrossPostActionsActionTypeQuote   ToolValidateValidatePostParamsCrossPostActionsActionType = "quote"
)

func (r ToolValidateValidatePostParamsCrossPostActionsActionType) IsKnown() bool {
	switch r {
	case ToolValidateValidatePostParamsCrossPostActionsActionTypeRepost, ToolValidateValidatePostParamsCrossPostActionsActionTypeComment, ToolValidateValidatePostParamsCrossPostActionsActionTypeQuote:
		return true
	}
	return false
}

type ToolValidateValidatePostParamsMedia struct {
	// Public URL of the media file
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Media type. Inferred from URL extension if omitted.
	Type param.Field[ToolValidateValidatePostParamsMediaType] `json:"type"`
}

func (r ToolValidateValidatePostParamsMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media type. Inferred from URL extension if omitted.
type ToolValidateValidatePostParamsMediaType string

const (
	ToolValidateValidatePostParamsMediaTypeImage    ToolValidateValidatePostParamsMediaType = "image"
	ToolValidateValidatePostParamsMediaTypeVideo    ToolValidateValidatePostParamsMediaType = "video"
	ToolValidateValidatePostParamsMediaTypeGif      ToolValidateValidatePostParamsMediaType = "gif"
	ToolValidateValidatePostParamsMediaTypeDocument ToolValidateValidatePostParamsMediaType = "document"
)

func (r ToolValidateValidatePostParamsMediaType) IsKnown() bool {
	switch r {
	case ToolValidateValidatePostParamsMediaTypeImage, ToolValidateValidatePostParamsMediaTypeVideo, ToolValidateValidatePostParamsMediaTypeGif, ToolValidateValidatePostParamsMediaTypeDocument:
		return true
	}
	return false
}

// Recycling configuration for evergreen content (Pro plan only)
type ToolValidateValidatePostParamsRecycling struct {
	// Interval value
	Gap param.Field[int64] `json:"gap" api:"required"`
	// Interval unit
	GapFreq param.Field[ToolValidateValidatePostParamsRecyclingGapFreq] `json:"gap_freq" api:"required"`
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

func (r ToolValidateValidatePostParamsRecycling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Interval unit
type ToolValidateValidatePostParamsRecyclingGapFreq string

const (
	ToolValidateValidatePostParamsRecyclingGapFreqDay   ToolValidateValidatePostParamsRecyclingGapFreq = "day"
	ToolValidateValidatePostParamsRecyclingGapFreqWeek  ToolValidateValidatePostParamsRecyclingGapFreq = "week"
	ToolValidateValidatePostParamsRecyclingGapFreqMonth ToolValidateValidatePostParamsRecyclingGapFreq = "month"
)

func (r ToolValidateValidatePostParamsRecyclingGapFreq) IsKnown() bool {
	switch r {
	case ToolValidateValidatePostParamsRecyclingGapFreqDay, ToolValidateValidatePostParamsRecyclingGapFreqWeek, ToolValidateValidatePostParamsRecyclingGapFreqMonth:
		return true
	}
	return false
}
