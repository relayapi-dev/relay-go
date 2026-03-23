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

// AccountFacebookPageService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountFacebookPageService] method instead.
type AccountFacebookPageService struct {
	Options []option.RequestOption
}

// NewAccountFacebookPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountFacebookPageService(opts ...option.RequestOption) (r *AccountFacebookPageService) {
	r = &AccountFacebookPageService{}
	r.Options = opts
	return
}

// Fetch Facebook pages for an account
func (r *AccountFacebookPageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountFacebookPageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/facebook-pages", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set default Facebook page
func (r *AccountFacebookPageService) SetDefault(ctx context.Context, id string, body AccountFacebookPageSetDefaultParams, opts ...option.RequestOption) (res *AccountFacebookPageSetDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/facebook-pages", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AccountFacebookPageGetResponse struct {
	Data []AccountFacebookPageGetResponseData `json:"data" api:"required"`
	JSON accountFacebookPageGetResponseJSON   `json:"-"`
}

// accountFacebookPageGetResponseJSON contains the JSON metadata for the struct
// [AccountFacebookPageGetResponse]
type accountFacebookPageGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFacebookPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFacebookPageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountFacebookPageGetResponseData struct {
	ID          string                                 `json:"id" api:"required"`
	Name        string                                 `json:"name" api:"required"`
	AccessToken string                                 `json:"access_token"`
	JSON        accountFacebookPageGetResponseDataJSON `json:"-"`
}

// accountFacebookPageGetResponseDataJSON contains the JSON metadata for the struct
// [AccountFacebookPageGetResponseData]
type accountFacebookPageGetResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFacebookPageGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFacebookPageGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountFacebookPageSetDefaultResponse struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             AccountFacebookPageSetDefaultResponseGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                        `json:"metadata" api:"required,nullable"`
	Platform          AccountFacebookPageSetDefaultResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                                        `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                     `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                        `json:"username" api:"required,nullable"`
	JSON              accountFacebookPageSetDefaultResponseJSON     `json:"-"`
}

// accountFacebookPageSetDefaultResponseJSON contains the JSON metadata for the
// struct [AccountFacebookPageSetDefaultResponse]
type accountFacebookPageSetDefaultResponseJSON struct {
	ID                apijson.Field
	AvatarURL         apijson.Field
	ConnectedAt       apijson.Field
	DisplayName       apijson.Field
	Group             apijson.Field
	Metadata          apijson.Field
	Platform          apijson.Field
	PlatformAccountID apijson.Field
	UpdatedAt         apijson.Field
	Username          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountFacebookPageSetDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFacebookPageSetDefaultResponseJSON) RawJSON() string {
	return r.raw
}

// Account group
type AccountFacebookPageSetDefaultResponseGroup struct {
	ID   string                                         `json:"id" api:"required"`
	Name string                                         `json:"name" api:"required"`
	JSON accountFacebookPageSetDefaultResponseGroupJSON `json:"-"`
}

// accountFacebookPageSetDefaultResponseGroupJSON contains the JSON metadata for
// the struct [AccountFacebookPageSetDefaultResponseGroup]
type accountFacebookPageSetDefaultResponseGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFacebookPageSetDefaultResponseGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFacebookPageSetDefaultResponseGroupJSON) RawJSON() string {
	return r.raw
}

type AccountFacebookPageSetDefaultResponsePlatform string

const (
	AccountFacebookPageSetDefaultResponsePlatformTwitter        AccountFacebookPageSetDefaultResponsePlatform = "twitter"
	AccountFacebookPageSetDefaultResponsePlatformInstagram      AccountFacebookPageSetDefaultResponsePlatform = "instagram"
	AccountFacebookPageSetDefaultResponsePlatformFacebook       AccountFacebookPageSetDefaultResponsePlatform = "facebook"
	AccountFacebookPageSetDefaultResponsePlatformLinkedin       AccountFacebookPageSetDefaultResponsePlatform = "linkedin"
	AccountFacebookPageSetDefaultResponsePlatformTiktok         AccountFacebookPageSetDefaultResponsePlatform = "tiktok"
	AccountFacebookPageSetDefaultResponsePlatformYoutube        AccountFacebookPageSetDefaultResponsePlatform = "youtube"
	AccountFacebookPageSetDefaultResponsePlatformPinterest      AccountFacebookPageSetDefaultResponsePlatform = "pinterest"
	AccountFacebookPageSetDefaultResponsePlatformReddit         AccountFacebookPageSetDefaultResponsePlatform = "reddit"
	AccountFacebookPageSetDefaultResponsePlatformBluesky        AccountFacebookPageSetDefaultResponsePlatform = "bluesky"
	AccountFacebookPageSetDefaultResponsePlatformThreads        AccountFacebookPageSetDefaultResponsePlatform = "threads"
	AccountFacebookPageSetDefaultResponsePlatformTelegram       AccountFacebookPageSetDefaultResponsePlatform = "telegram"
	AccountFacebookPageSetDefaultResponsePlatformSnapchat       AccountFacebookPageSetDefaultResponsePlatform = "snapchat"
	AccountFacebookPageSetDefaultResponsePlatformGooglebusiness AccountFacebookPageSetDefaultResponsePlatform = "googlebusiness"
	AccountFacebookPageSetDefaultResponsePlatformWhatsapp       AccountFacebookPageSetDefaultResponsePlatform = "whatsapp"
	AccountFacebookPageSetDefaultResponsePlatformMastodon       AccountFacebookPageSetDefaultResponsePlatform = "mastodon"
	AccountFacebookPageSetDefaultResponsePlatformDiscord        AccountFacebookPageSetDefaultResponsePlatform = "discord"
	AccountFacebookPageSetDefaultResponsePlatformSMS            AccountFacebookPageSetDefaultResponsePlatform = "sms"
)

func (r AccountFacebookPageSetDefaultResponsePlatform) IsKnown() bool {
	switch r {
	case AccountFacebookPageSetDefaultResponsePlatformTwitter, AccountFacebookPageSetDefaultResponsePlatformInstagram, AccountFacebookPageSetDefaultResponsePlatformFacebook, AccountFacebookPageSetDefaultResponsePlatformLinkedin, AccountFacebookPageSetDefaultResponsePlatformTiktok, AccountFacebookPageSetDefaultResponsePlatformYoutube, AccountFacebookPageSetDefaultResponsePlatformPinterest, AccountFacebookPageSetDefaultResponsePlatformReddit, AccountFacebookPageSetDefaultResponsePlatformBluesky, AccountFacebookPageSetDefaultResponsePlatformThreads, AccountFacebookPageSetDefaultResponsePlatformTelegram, AccountFacebookPageSetDefaultResponsePlatformSnapchat, AccountFacebookPageSetDefaultResponsePlatformGooglebusiness, AccountFacebookPageSetDefaultResponsePlatformWhatsapp, AccountFacebookPageSetDefaultResponsePlatformMastodon, AccountFacebookPageSetDefaultResponsePlatformDiscord, AccountFacebookPageSetDefaultResponsePlatformSMS:
		return true
	}
	return false
}

type AccountFacebookPageSetDefaultParams struct {
	// Facebook page ID to set as default
	PageID param.Field[string] `json:"page_id" api:"required"`
}

func (r AccountFacebookPageSetDefaultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
