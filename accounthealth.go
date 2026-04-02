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

// AccountHealthService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountHealthService] method instead.
type AccountHealthService struct {
	Options []option.RequestOption
}

// NewAccountHealthService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountHealthService(opts ...option.RequestOption) (r *AccountHealthService) {
	r = &AccountHealthService{}
	r.Options = opts
	return
}

// Check health of a single connected account
func (r *AccountHealthService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountHealthGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/health", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Check health of all connected accounts
func (r *AccountHealthService) List(ctx context.Context, query AccountHealthListParams, opts ...option.RequestOption) (res *AccountHealthListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AccountHealthGetResponse struct {
	ID             string                        `json:"id" api:"required"`
	AvatarURL      string                        `json:"avatar_url" api:"required,nullable"`
	DisplayName    string                        `json:"display_name" api:"required,nullable"`
	Healthy        bool                          `json:"healthy" api:"required"`
	Platform       string                        `json:"platform" api:"required"`
	Scopes         []string                      `json:"scopes" api:"required"`
	TokenExpiresAt string                        `json:"token_expires_at" api:"required,nullable"`
	Username       string                        `json:"username" api:"required,nullable"`
	Error          AccountHealthGetResponseError `json:"error"`
	JSON           accountHealthGetResponseJSON  `json:"-"`
}

// accountHealthGetResponseJSON contains the JSON metadata for the struct
// [AccountHealthGetResponse]
type accountHealthGetResponseJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	DisplayName    apijson.Field
	Healthy        apijson.Field
	Platform       apijson.Field
	Scopes         apijson.Field
	TokenExpiresAt apijson.Field
	Username       apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountHealthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHealthGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHealthGetResponseError struct {
	Code    string                            `json:"code" api:"required"`
	Message string                            `json:"message" api:"required"`
	JSON    accountHealthGetResponseErrorJSON `json:"-"`
}

// accountHealthGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountHealthGetResponseError]
type accountHealthGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHealthGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHealthGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountHealthListResponse struct {
	Data []AccountHealthListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                        `json:"next_cursor" api:"required,nullable"`
	JSON       accountHealthListResponseJSON `json:"-"`
}

// accountHealthListResponseJSON contains the JSON metadata for the struct
// [AccountHealthListResponse]
type accountHealthListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHealthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHealthListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHealthListResponseData struct {
	ID             string                                `json:"id" api:"required"`
	Healthy        bool                                  `json:"healthy" api:"required"`
	Platform       AccountHealthListResponseDataPlatform `json:"platform" api:"required"`
	TokenExpiresAt time.Time                             `json:"token_expires_at" api:"required,nullable" format:"date-time"`
	Username       string                                `json:"username" api:"required,nullable"`
	Error          AccountHealthListResponseDataError    `json:"error"`
	JSON           accountHealthListResponseDataJSON     `json:"-"`
}

// accountHealthListResponseDataJSON contains the JSON metadata for the struct
// [AccountHealthListResponseData]
type accountHealthListResponseDataJSON struct {
	ID             apijson.Field
	Healthy        apijson.Field
	Platform       apijson.Field
	TokenExpiresAt apijson.Field
	Username       apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountHealthListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHealthListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountHealthListResponseDataPlatform string

const (
	AccountHealthListResponseDataPlatformTwitter        AccountHealthListResponseDataPlatform = "twitter"
	AccountHealthListResponseDataPlatformInstagram      AccountHealthListResponseDataPlatform = "instagram"
	AccountHealthListResponseDataPlatformFacebook       AccountHealthListResponseDataPlatform = "facebook"
	AccountHealthListResponseDataPlatformLinkedin       AccountHealthListResponseDataPlatform = "linkedin"
	AccountHealthListResponseDataPlatformTiktok         AccountHealthListResponseDataPlatform = "tiktok"
	AccountHealthListResponseDataPlatformYoutube        AccountHealthListResponseDataPlatform = "youtube"
	AccountHealthListResponseDataPlatformPinterest      AccountHealthListResponseDataPlatform = "pinterest"
	AccountHealthListResponseDataPlatformReddit         AccountHealthListResponseDataPlatform = "reddit"
	AccountHealthListResponseDataPlatformBluesky        AccountHealthListResponseDataPlatform = "bluesky"
	AccountHealthListResponseDataPlatformThreads        AccountHealthListResponseDataPlatform = "threads"
	AccountHealthListResponseDataPlatformTelegram       AccountHealthListResponseDataPlatform = "telegram"
	AccountHealthListResponseDataPlatformSnapchat       AccountHealthListResponseDataPlatform = "snapchat"
	AccountHealthListResponseDataPlatformGooglebusiness AccountHealthListResponseDataPlatform = "googlebusiness"
	AccountHealthListResponseDataPlatformWhatsapp       AccountHealthListResponseDataPlatform = "whatsapp"
	AccountHealthListResponseDataPlatformMastodon       AccountHealthListResponseDataPlatform = "mastodon"
	AccountHealthListResponseDataPlatformDiscord        AccountHealthListResponseDataPlatform = "discord"
	AccountHealthListResponseDataPlatformSMS            AccountHealthListResponseDataPlatform = "sms"
)

func (r AccountHealthListResponseDataPlatform) IsKnown() bool {
	switch r {
	case AccountHealthListResponseDataPlatformTwitter, AccountHealthListResponseDataPlatformInstagram, AccountHealthListResponseDataPlatformFacebook, AccountHealthListResponseDataPlatformLinkedin, AccountHealthListResponseDataPlatformTiktok, AccountHealthListResponseDataPlatformYoutube, AccountHealthListResponseDataPlatformPinterest, AccountHealthListResponseDataPlatformReddit, AccountHealthListResponseDataPlatformBluesky, AccountHealthListResponseDataPlatformThreads, AccountHealthListResponseDataPlatformTelegram, AccountHealthListResponseDataPlatformSnapchat, AccountHealthListResponseDataPlatformGooglebusiness, AccountHealthListResponseDataPlatformWhatsapp, AccountHealthListResponseDataPlatformMastodon, AccountHealthListResponseDataPlatformDiscord, AccountHealthListResponseDataPlatformSMS:
		return true
	}
	return false
}

type AccountHealthListResponseDataError struct {
	Code    string                                 `json:"code" api:"required"`
	Message string                                 `json:"message" api:"required"`
	JSON    accountHealthListResponseDataErrorJSON `json:"-"`
}

// accountHealthListResponseDataErrorJSON contains the JSON metadata for the struct
// [AccountHealthListResponseDataError]
type accountHealthListResponseDataErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHealthListResponseDataError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHealthListResponseDataErrorJSON) RawJSON() string {
	return r.raw
}

type AccountHealthListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AccountHealthListParams]'s query parameters as
// `url.Values`.
func (r AccountHealthListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
