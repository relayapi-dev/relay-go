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

// AccountService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options               []option.RequestOption
	Health                *AccountHealthService
	RedditFlairs          *AccountRedditFlairService
	FacebookPages         *AccountFacebookPageService
	LinkedinOrganizations *AccountLinkedinOrganizationService
	PinterestBoards       *AccountPinterestBoardService
	RedditSubreddits      *AccountRedditSubredditService
	GmbLocations          *AccountGmbLocationService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Health = NewAccountHealthService(opts...)
	r.RedditFlairs = NewAccountRedditFlairService(opts...)
	r.FacebookPages = NewAccountFacebookPageService(opts...)
	r.LinkedinOrganizations = NewAccountLinkedinOrganizationService(opts...)
	r.PinterestBoards = NewAccountPinterestBoardService(opts...)
	r.RedditSubreddits = NewAccountRedditSubredditService(opts...)
	r.GmbLocations = NewAccountGmbLocationService(opts...)
	return
}

// Get a connected account
func (r *AccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update account metadata
func (r *AccountService) Update(ctx context.Context, id string, body AccountUpdateParams, opts ...option.RequestOption) (res *AccountUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List connected accounts
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Disconnect a social account
func (r *AccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AccountGetResponse struct {
	// Account ID
	ID                string                     `json:"id" api:"required"`
	AvatarURL         string                     `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                  `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                     `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}     `json:"metadata" api:"required,nullable"`
	Platform          AccountGetResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                     `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	Username          string                     `json:"username" api:"required,nullable"`
	JSON              accountGetResponseJSON     `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
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

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponsePlatform string

const (
	AccountGetResponsePlatformTwitter        AccountGetResponsePlatform = "twitter"
	AccountGetResponsePlatformInstagram      AccountGetResponsePlatform = "instagram"
	AccountGetResponsePlatformFacebook       AccountGetResponsePlatform = "facebook"
	AccountGetResponsePlatformLinkedin       AccountGetResponsePlatform = "linkedin"
	AccountGetResponsePlatformTiktok         AccountGetResponsePlatform = "tiktok"
	AccountGetResponsePlatformYoutube        AccountGetResponsePlatform = "youtube"
	AccountGetResponsePlatformPinterest      AccountGetResponsePlatform = "pinterest"
	AccountGetResponsePlatformReddit         AccountGetResponsePlatform = "reddit"
	AccountGetResponsePlatformBluesky        AccountGetResponsePlatform = "bluesky"
	AccountGetResponsePlatformThreads        AccountGetResponsePlatform = "threads"
	AccountGetResponsePlatformTelegram       AccountGetResponsePlatform = "telegram"
	AccountGetResponsePlatformSnapchat       AccountGetResponsePlatform = "snapchat"
	AccountGetResponsePlatformGooglebusiness AccountGetResponsePlatform = "googlebusiness"
	AccountGetResponsePlatformWhatsapp       AccountGetResponsePlatform = "whatsapp"
	AccountGetResponsePlatformMastodon       AccountGetResponsePlatform = "mastodon"
	AccountGetResponsePlatformDiscord        AccountGetResponsePlatform = "discord"
	AccountGetResponsePlatformSMS            AccountGetResponsePlatform = "sms"
)

func (r AccountGetResponsePlatform) IsKnown() bool {
	switch r {
	case AccountGetResponsePlatformTwitter, AccountGetResponsePlatformInstagram, AccountGetResponsePlatformFacebook, AccountGetResponsePlatformLinkedin, AccountGetResponsePlatformTiktok, AccountGetResponsePlatformYoutube, AccountGetResponsePlatformPinterest, AccountGetResponsePlatformReddit, AccountGetResponsePlatformBluesky, AccountGetResponsePlatformThreads, AccountGetResponsePlatformTelegram, AccountGetResponsePlatformSnapchat, AccountGetResponsePlatformGooglebusiness, AccountGetResponsePlatformWhatsapp, AccountGetResponsePlatformMastodon, AccountGetResponsePlatformDiscord, AccountGetResponsePlatformSMS:
		return true
	}
	return false
}

type AccountUpdateResponse struct {
	// Account ID
	ID                string                        `json:"id" api:"required"`
	AvatarURL         string                        `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                     `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                        `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}        `json:"metadata" api:"required,nullable"`
	Platform          AccountUpdateResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                        `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                     `json:"updated_at" api:"required" format:"date-time"`
	Username          string                        `json:"username" api:"required,nullable"`
	JSON              accountUpdateResponseJSON     `json:"-"`
}

// accountUpdateResponseJSON contains the JSON metadata for the struct
// [AccountUpdateResponse]
type accountUpdateResponseJSON struct {
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

func (r *AccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUpdateResponsePlatform string

const (
	AccountUpdateResponsePlatformTwitter        AccountUpdateResponsePlatform = "twitter"
	AccountUpdateResponsePlatformInstagram      AccountUpdateResponsePlatform = "instagram"
	AccountUpdateResponsePlatformFacebook       AccountUpdateResponsePlatform = "facebook"
	AccountUpdateResponsePlatformLinkedin       AccountUpdateResponsePlatform = "linkedin"
	AccountUpdateResponsePlatformTiktok         AccountUpdateResponsePlatform = "tiktok"
	AccountUpdateResponsePlatformYoutube        AccountUpdateResponsePlatform = "youtube"
	AccountUpdateResponsePlatformPinterest      AccountUpdateResponsePlatform = "pinterest"
	AccountUpdateResponsePlatformReddit         AccountUpdateResponsePlatform = "reddit"
	AccountUpdateResponsePlatformBluesky        AccountUpdateResponsePlatform = "bluesky"
	AccountUpdateResponsePlatformThreads        AccountUpdateResponsePlatform = "threads"
	AccountUpdateResponsePlatformTelegram       AccountUpdateResponsePlatform = "telegram"
	AccountUpdateResponsePlatformSnapchat       AccountUpdateResponsePlatform = "snapchat"
	AccountUpdateResponsePlatformGooglebusiness AccountUpdateResponsePlatform = "googlebusiness"
	AccountUpdateResponsePlatformWhatsapp       AccountUpdateResponsePlatform = "whatsapp"
	AccountUpdateResponsePlatformMastodon       AccountUpdateResponsePlatform = "mastodon"
	AccountUpdateResponsePlatformDiscord        AccountUpdateResponsePlatform = "discord"
	AccountUpdateResponsePlatformSMS            AccountUpdateResponsePlatform = "sms"
)

func (r AccountUpdateResponsePlatform) IsKnown() bool {
	switch r {
	case AccountUpdateResponsePlatformTwitter, AccountUpdateResponsePlatformInstagram, AccountUpdateResponsePlatformFacebook, AccountUpdateResponsePlatformLinkedin, AccountUpdateResponsePlatformTiktok, AccountUpdateResponsePlatformYoutube, AccountUpdateResponsePlatformPinterest, AccountUpdateResponsePlatformReddit, AccountUpdateResponsePlatformBluesky, AccountUpdateResponsePlatformThreads, AccountUpdateResponsePlatformTelegram, AccountUpdateResponsePlatformSnapchat, AccountUpdateResponsePlatformGooglebusiness, AccountUpdateResponsePlatformWhatsapp, AccountUpdateResponsePlatformMastodon, AccountUpdateResponsePlatformDiscord, AccountUpdateResponsePlatformSMS:
		return true
	}
	return false
}

type AccountListResponse struct {
	Data []AccountListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                  `json:"next_cursor" api:"required,nullable"`
	JSON       accountListResponseJSON `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListResponseData struct {
	// Account ID
	ID                string                          `json:"id" api:"required"`
	AvatarURL         string                          `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                       `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                          `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}          `json:"metadata" api:"required,nullable"`
	Platform          AccountListResponseDataPlatform `json:"platform" api:"required"`
	PlatformAccountID string                          `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                       `json:"updated_at" api:"required" format:"date-time"`
	Username          string                          `json:"username" api:"required,nullable"`
	JSON              accountListResponseDataJSON     `json:"-"`
}

// accountListResponseDataJSON contains the JSON metadata for the struct
// [AccountListResponseData]
type accountListResponseDataJSON struct {
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

func (r *AccountListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountListResponseDataPlatform string

const (
	AccountListResponseDataPlatformTwitter        AccountListResponseDataPlatform = "twitter"
	AccountListResponseDataPlatformInstagram      AccountListResponseDataPlatform = "instagram"
	AccountListResponseDataPlatformFacebook       AccountListResponseDataPlatform = "facebook"
	AccountListResponseDataPlatformLinkedin       AccountListResponseDataPlatform = "linkedin"
	AccountListResponseDataPlatformTiktok         AccountListResponseDataPlatform = "tiktok"
	AccountListResponseDataPlatformYoutube        AccountListResponseDataPlatform = "youtube"
	AccountListResponseDataPlatformPinterest      AccountListResponseDataPlatform = "pinterest"
	AccountListResponseDataPlatformReddit         AccountListResponseDataPlatform = "reddit"
	AccountListResponseDataPlatformBluesky        AccountListResponseDataPlatform = "bluesky"
	AccountListResponseDataPlatformThreads        AccountListResponseDataPlatform = "threads"
	AccountListResponseDataPlatformTelegram       AccountListResponseDataPlatform = "telegram"
	AccountListResponseDataPlatformSnapchat       AccountListResponseDataPlatform = "snapchat"
	AccountListResponseDataPlatformGooglebusiness AccountListResponseDataPlatform = "googlebusiness"
	AccountListResponseDataPlatformWhatsapp       AccountListResponseDataPlatform = "whatsapp"
	AccountListResponseDataPlatformMastodon       AccountListResponseDataPlatform = "mastodon"
	AccountListResponseDataPlatformDiscord        AccountListResponseDataPlatform = "discord"
	AccountListResponseDataPlatformSMS            AccountListResponseDataPlatform = "sms"
)

func (r AccountListResponseDataPlatform) IsKnown() bool {
	switch r {
	case AccountListResponseDataPlatformTwitter, AccountListResponseDataPlatformInstagram, AccountListResponseDataPlatformFacebook, AccountListResponseDataPlatformLinkedin, AccountListResponseDataPlatformTiktok, AccountListResponseDataPlatformYoutube, AccountListResponseDataPlatformPinterest, AccountListResponseDataPlatformReddit, AccountListResponseDataPlatformBluesky, AccountListResponseDataPlatformThreads, AccountListResponseDataPlatformTelegram, AccountListResponseDataPlatformSnapchat, AccountListResponseDataPlatformGooglebusiness, AccountListResponseDataPlatformWhatsapp, AccountListResponseDataPlatformMastodon, AccountListResponseDataPlatformDiscord, AccountListResponseDataPlatformSMS:
		return true
	}
	return false
}

type AccountUpdateParams struct {
	DisplayName param.Field[string]                 `json:"display_name"`
	Metadata    param.Field[map[string]interface{}] `json:"metadata"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
