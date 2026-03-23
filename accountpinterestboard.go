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

// AccountPinterestBoardService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPinterestBoardService] method instead.
type AccountPinterestBoardService struct {
	Options []option.RequestOption
}

// NewAccountPinterestBoardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPinterestBoardService(opts ...option.RequestOption) (r *AccountPinterestBoardService) {
	r = &AccountPinterestBoardService{}
	r.Options = opts
	return
}

// Fetch Pinterest boards for an account
func (r *AccountPinterestBoardService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountPinterestBoardGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/pinterest-boards", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set default Pinterest board
func (r *AccountPinterestBoardService) SetDefault(ctx context.Context, id string, body AccountPinterestBoardSetDefaultParams, opts ...option.RequestOption) (res *AccountPinterestBoardSetDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/pinterest-boards", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AccountPinterestBoardGetResponse struct {
	Data []AccountPinterestBoardGetResponseData `json:"data" api:"required"`
	JSON accountPinterestBoardGetResponseJSON   `json:"-"`
}

// accountPinterestBoardGetResponseJSON contains the JSON metadata for the struct
// [AccountPinterestBoardGetResponse]
type accountPinterestBoardGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPinterestBoardGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPinterestBoardGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPinterestBoardGetResponseData struct {
	ID   string                                   `json:"id" api:"required"`
	Name string                                   `json:"name" api:"required"`
	URL  string                                   `json:"url" api:"required,nullable"`
	JSON accountPinterestBoardGetResponseDataJSON `json:"-"`
}

// accountPinterestBoardGetResponseDataJSON contains the JSON metadata for the
// struct [AccountPinterestBoardGetResponseData]
type accountPinterestBoardGetResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPinterestBoardGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPinterestBoardGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountPinterestBoardSetDefaultResponse struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             AccountPinterestBoardSetDefaultResponseGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                          `json:"metadata" api:"required,nullable"`
	Platform          AccountPinterestBoardSetDefaultResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                                          `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                       `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                          `json:"username" api:"required,nullable"`
	JSON              accountPinterestBoardSetDefaultResponseJSON     `json:"-"`
}

// accountPinterestBoardSetDefaultResponseJSON contains the JSON metadata for the
// struct [AccountPinterestBoardSetDefaultResponse]
type accountPinterestBoardSetDefaultResponseJSON struct {
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

func (r *AccountPinterestBoardSetDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPinterestBoardSetDefaultResponseJSON) RawJSON() string {
	return r.raw
}

// Account group
type AccountPinterestBoardSetDefaultResponseGroup struct {
	ID   string                                           `json:"id" api:"required"`
	Name string                                           `json:"name" api:"required"`
	JSON accountPinterestBoardSetDefaultResponseGroupJSON `json:"-"`
}

// accountPinterestBoardSetDefaultResponseGroupJSON contains the JSON metadata for
// the struct [AccountPinterestBoardSetDefaultResponseGroup]
type accountPinterestBoardSetDefaultResponseGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPinterestBoardSetDefaultResponseGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPinterestBoardSetDefaultResponseGroupJSON) RawJSON() string {
	return r.raw
}

type AccountPinterestBoardSetDefaultResponsePlatform string

const (
	AccountPinterestBoardSetDefaultResponsePlatformTwitter        AccountPinterestBoardSetDefaultResponsePlatform = "twitter"
	AccountPinterestBoardSetDefaultResponsePlatformInstagram      AccountPinterestBoardSetDefaultResponsePlatform = "instagram"
	AccountPinterestBoardSetDefaultResponsePlatformFacebook       AccountPinterestBoardSetDefaultResponsePlatform = "facebook"
	AccountPinterestBoardSetDefaultResponsePlatformLinkedin       AccountPinterestBoardSetDefaultResponsePlatform = "linkedin"
	AccountPinterestBoardSetDefaultResponsePlatformTiktok         AccountPinterestBoardSetDefaultResponsePlatform = "tiktok"
	AccountPinterestBoardSetDefaultResponsePlatformYoutube        AccountPinterestBoardSetDefaultResponsePlatform = "youtube"
	AccountPinterestBoardSetDefaultResponsePlatformPinterest      AccountPinterestBoardSetDefaultResponsePlatform = "pinterest"
	AccountPinterestBoardSetDefaultResponsePlatformReddit         AccountPinterestBoardSetDefaultResponsePlatform = "reddit"
	AccountPinterestBoardSetDefaultResponsePlatformBluesky        AccountPinterestBoardSetDefaultResponsePlatform = "bluesky"
	AccountPinterestBoardSetDefaultResponsePlatformThreads        AccountPinterestBoardSetDefaultResponsePlatform = "threads"
	AccountPinterestBoardSetDefaultResponsePlatformTelegram       AccountPinterestBoardSetDefaultResponsePlatform = "telegram"
	AccountPinterestBoardSetDefaultResponsePlatformSnapchat       AccountPinterestBoardSetDefaultResponsePlatform = "snapchat"
	AccountPinterestBoardSetDefaultResponsePlatformGooglebusiness AccountPinterestBoardSetDefaultResponsePlatform = "googlebusiness"
	AccountPinterestBoardSetDefaultResponsePlatformWhatsapp       AccountPinterestBoardSetDefaultResponsePlatform = "whatsapp"
	AccountPinterestBoardSetDefaultResponsePlatformMastodon       AccountPinterestBoardSetDefaultResponsePlatform = "mastodon"
	AccountPinterestBoardSetDefaultResponsePlatformDiscord        AccountPinterestBoardSetDefaultResponsePlatform = "discord"
	AccountPinterestBoardSetDefaultResponsePlatformSMS            AccountPinterestBoardSetDefaultResponsePlatform = "sms"
)

func (r AccountPinterestBoardSetDefaultResponsePlatform) IsKnown() bool {
	switch r {
	case AccountPinterestBoardSetDefaultResponsePlatformTwitter, AccountPinterestBoardSetDefaultResponsePlatformInstagram, AccountPinterestBoardSetDefaultResponsePlatformFacebook, AccountPinterestBoardSetDefaultResponsePlatformLinkedin, AccountPinterestBoardSetDefaultResponsePlatformTiktok, AccountPinterestBoardSetDefaultResponsePlatformYoutube, AccountPinterestBoardSetDefaultResponsePlatformPinterest, AccountPinterestBoardSetDefaultResponsePlatformReddit, AccountPinterestBoardSetDefaultResponsePlatformBluesky, AccountPinterestBoardSetDefaultResponsePlatformThreads, AccountPinterestBoardSetDefaultResponsePlatformTelegram, AccountPinterestBoardSetDefaultResponsePlatformSnapchat, AccountPinterestBoardSetDefaultResponsePlatformGooglebusiness, AccountPinterestBoardSetDefaultResponsePlatformWhatsapp, AccountPinterestBoardSetDefaultResponsePlatformMastodon, AccountPinterestBoardSetDefaultResponsePlatformDiscord, AccountPinterestBoardSetDefaultResponsePlatformSMS:
		return true
	}
	return false
}

type AccountPinterestBoardSetDefaultParams struct {
	// Pinterest board ID to set as default
	BoardID param.Field[string] `json:"board_id" api:"required"`
}

func (r AccountPinterestBoardSetDefaultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
