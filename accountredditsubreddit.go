// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// AccountRedditSubredditService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRedditSubredditService] method instead.
type AccountRedditSubredditService struct {
	Options []option.RequestOption
}

// NewAccountRedditSubredditService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRedditSubredditService(opts ...option.RequestOption) (r *AccountRedditSubredditService) {
	r = &AccountRedditSubredditService{}
	r.Options = opts
	return
}

// Fetch Reddit subreddits for an account
func (r *AccountRedditSubredditService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountRedditSubredditGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/reddit-subreddits", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set default Reddit subreddit
func (r *AccountRedditSubredditService) SetDefault(ctx context.Context, id string, body AccountRedditSubredditSetDefaultParams, opts ...option.RequestOption) (res *AccountRedditSubredditSetDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/reddit-subreddits", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AccountRedditSubredditGetResponse struct {
	Data []AccountRedditSubredditGetResponseData `json:"data" api:"required"`
	JSON accountRedditSubredditGetResponseJSON   `json:"-"`
}

// accountRedditSubredditGetResponseJSON contains the JSON metadata for the struct
// [AccountRedditSubredditGetResponse]
type accountRedditSubredditGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRedditSubredditGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRedditSubredditGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRedditSubredditGetResponseData struct {
	DisplayName string                                    `json:"display_name" api:"required"`
	Name        string                                    `json:"name" api:"required"`
	Subscribers float64                                   `json:"subscribers" api:"required,nullable"`
	JSON        accountRedditSubredditGetResponseDataJSON `json:"-"`
}

// accountRedditSubredditGetResponseDataJSON contains the JSON metadata for the
// struct [AccountRedditSubredditGetResponseData]
type accountRedditSubredditGetResponseDataJSON struct {
	DisplayName apijson.Field
	Name        apijson.Field
	Subscribers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRedditSubredditGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRedditSubredditGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountRedditSubredditSetDefaultResponse struct {
	// Account ID
	ID                string                                           `json:"id" api:"required"`
	AvatarURL         string                                           `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                        `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                           `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                           `json:"metadata" api:"required,nullable"`
	Platform          AccountRedditSubredditSetDefaultResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                                           `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                           `json:"username" api:"required,nullable"`
	JSON              accountRedditSubredditSetDefaultResponseJSON     `json:"-"`
}

// accountRedditSubredditSetDefaultResponseJSON contains the JSON metadata for the
// struct [AccountRedditSubredditSetDefaultResponse]
type accountRedditSubredditSetDefaultResponseJSON struct {
	ID                apijson.Field
	AvatarURL         apijson.Field
	ConnectedAt       apijson.Field
	DisplayName       apijson.Field
	Metadata          apijson.Field
	Platform          apijson.Field
	PlatformAccountID apijson.Field
	UpdatedAt         apijson.Field
	Username          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountRedditSubredditSetDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRedditSubredditSetDefaultResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRedditSubredditSetDefaultResponsePlatform string

const (
	AccountRedditSubredditSetDefaultResponsePlatformTwitter        AccountRedditSubredditSetDefaultResponsePlatform = "twitter"
	AccountRedditSubredditSetDefaultResponsePlatformInstagram      AccountRedditSubredditSetDefaultResponsePlatform = "instagram"
	AccountRedditSubredditSetDefaultResponsePlatformFacebook       AccountRedditSubredditSetDefaultResponsePlatform = "facebook"
	AccountRedditSubredditSetDefaultResponsePlatformLinkedin       AccountRedditSubredditSetDefaultResponsePlatform = "linkedin"
	AccountRedditSubredditSetDefaultResponsePlatformTiktok         AccountRedditSubredditSetDefaultResponsePlatform = "tiktok"
	AccountRedditSubredditSetDefaultResponsePlatformYoutube        AccountRedditSubredditSetDefaultResponsePlatform = "youtube"
	AccountRedditSubredditSetDefaultResponsePlatformPinterest      AccountRedditSubredditSetDefaultResponsePlatform = "pinterest"
	AccountRedditSubredditSetDefaultResponsePlatformReddit         AccountRedditSubredditSetDefaultResponsePlatform = "reddit"
	AccountRedditSubredditSetDefaultResponsePlatformBluesky        AccountRedditSubredditSetDefaultResponsePlatform = "bluesky"
	AccountRedditSubredditSetDefaultResponsePlatformThreads        AccountRedditSubredditSetDefaultResponsePlatform = "threads"
	AccountRedditSubredditSetDefaultResponsePlatformTelegram       AccountRedditSubredditSetDefaultResponsePlatform = "telegram"
	AccountRedditSubredditSetDefaultResponsePlatformSnapchat       AccountRedditSubredditSetDefaultResponsePlatform = "snapchat"
	AccountRedditSubredditSetDefaultResponsePlatformGooglebusiness AccountRedditSubredditSetDefaultResponsePlatform = "googlebusiness"
	AccountRedditSubredditSetDefaultResponsePlatformWhatsapp       AccountRedditSubredditSetDefaultResponsePlatform = "whatsapp"
	AccountRedditSubredditSetDefaultResponsePlatformMastodon       AccountRedditSubredditSetDefaultResponsePlatform = "mastodon"
	AccountRedditSubredditSetDefaultResponsePlatformDiscord        AccountRedditSubredditSetDefaultResponsePlatform = "discord"
	AccountRedditSubredditSetDefaultResponsePlatformSMS            AccountRedditSubredditSetDefaultResponsePlatform = "sms"
)

func (r AccountRedditSubredditSetDefaultResponsePlatform) IsKnown() bool {
	switch r {
	case AccountRedditSubredditSetDefaultResponsePlatformTwitter, AccountRedditSubredditSetDefaultResponsePlatformInstagram, AccountRedditSubredditSetDefaultResponsePlatformFacebook, AccountRedditSubredditSetDefaultResponsePlatformLinkedin, AccountRedditSubredditSetDefaultResponsePlatformTiktok, AccountRedditSubredditSetDefaultResponsePlatformYoutube, AccountRedditSubredditSetDefaultResponsePlatformPinterest, AccountRedditSubredditSetDefaultResponsePlatformReddit, AccountRedditSubredditSetDefaultResponsePlatformBluesky, AccountRedditSubredditSetDefaultResponsePlatformThreads, AccountRedditSubredditSetDefaultResponsePlatformTelegram, AccountRedditSubredditSetDefaultResponsePlatformSnapchat, AccountRedditSubredditSetDefaultResponsePlatformGooglebusiness, AccountRedditSubredditSetDefaultResponsePlatformWhatsapp, AccountRedditSubredditSetDefaultResponsePlatformMastodon, AccountRedditSubredditSetDefaultResponsePlatformDiscord, AccountRedditSubredditSetDefaultResponsePlatformSMS:
		return true
	}
	return false
}

type AccountRedditSubredditSetDefaultParams struct {
	// Subreddit name to set as default
	Subreddit param.Field[string] `json:"subreddit" api:"required"`
}

func (r AccountRedditSubredditSetDefaultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
