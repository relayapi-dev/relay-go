// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// ConnectFacebookPageService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectFacebookPageService] method instead.
type ConnectFacebookPageService struct {
	Options []option.RequestOption
}

// NewConnectFacebookPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectFacebookPageService(opts ...option.RequestOption) (r *ConnectFacebookPageService) {
	r = &ConnectFacebookPageService{}
	r.Options = opts
	return
}

// List Facebook Pages after OAuth
func (r *ConnectFacebookPageService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectFacebookPageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/facebook/pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Select Facebook Page to connect
func (r *ConnectFacebookPageService) Select(ctx context.Context, body ConnectFacebookPageSelectParams, opts ...option.RequestOption) (res *ConnectFacebookPageSelectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/facebook/pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConnectFacebookPageListResponse struct {
	Pages []ConnectFacebookPageListResponsePage `json:"pages" api:"required"`
	JSON  connectFacebookPageListResponseJSON   `json:"-"`
}

// connectFacebookPageListResponseJSON contains the JSON metadata for the struct
// [ConnectFacebookPageListResponse]
type connectFacebookPageListResponseJSON struct {
	Pages       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectFacebookPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFacebookPageListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectFacebookPageListResponsePage struct {
	// Facebook page ID
	ID string `json:"id" api:"required"`
	// Page name
	Name string `json:"name" api:"required"`
	// Page category
	Category string `json:"category" api:"nullable"`
	// Page profile picture URL
	PictureURL string                                  `json:"picture_url" api:"nullable"`
	JSON       connectFacebookPageListResponsePageJSON `json:"-"`
}

// connectFacebookPageListResponsePageJSON contains the JSON metadata for the
// struct [ConnectFacebookPageListResponsePage]
type connectFacebookPageListResponsePageJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Category    apijson.Field
	PictureURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectFacebookPageListResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFacebookPageListResponsePageJSON) RawJSON() string {
	return r.raw
}

type ConnectFacebookPageSelectResponse struct {
	Account ConnectFacebookPageSelectResponseAccount `json:"account" api:"required"`
	JSON    connectFacebookPageSelectResponseJSON    `json:"-"`
}

// connectFacebookPageSelectResponseJSON contains the JSON metadata for the struct
// [ConnectFacebookPageSelectResponse]
type connectFacebookPageSelectResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectFacebookPageSelectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFacebookPageSelectResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectFacebookPageSelectResponseAccount struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             ConnectFacebookPageSelectResponseAccountGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                           `json:"metadata" api:"required,nullable"`
	Platform          ConnectFacebookPageSelectResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                           `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                           `json:"username" api:"required,nullable"`
	JSON              connectFacebookPageSelectResponseAccountJSON     `json:"-"`
}

// connectFacebookPageSelectResponseAccountJSON contains the JSON metadata for the
// struct [ConnectFacebookPageSelectResponseAccount]
type connectFacebookPageSelectResponseAccountJSON struct {
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

func (r *ConnectFacebookPageSelectResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFacebookPageSelectResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account group
type ConnectFacebookPageSelectResponseAccountGroup struct {
	ID   string                                            `json:"id" api:"required"`
	Name string                                            `json:"name" api:"required"`
	JSON connectFacebookPageSelectResponseAccountGroupJSON `json:"-"`
}

// connectFacebookPageSelectResponseAccountGroupJSON contains the JSON metadata for
// the struct [ConnectFacebookPageSelectResponseAccountGroup]
type connectFacebookPageSelectResponseAccountGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectFacebookPageSelectResponseAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFacebookPageSelectResponseAccountGroupJSON) RawJSON() string {
	return r.raw
}

type ConnectFacebookPageSelectResponseAccountPlatform string

const (
	ConnectFacebookPageSelectResponseAccountPlatformTwitter        ConnectFacebookPageSelectResponseAccountPlatform = "twitter"
	ConnectFacebookPageSelectResponseAccountPlatformInstagram      ConnectFacebookPageSelectResponseAccountPlatform = "instagram"
	ConnectFacebookPageSelectResponseAccountPlatformFacebook       ConnectFacebookPageSelectResponseAccountPlatform = "facebook"
	ConnectFacebookPageSelectResponseAccountPlatformLinkedin       ConnectFacebookPageSelectResponseAccountPlatform = "linkedin"
	ConnectFacebookPageSelectResponseAccountPlatformTiktok         ConnectFacebookPageSelectResponseAccountPlatform = "tiktok"
	ConnectFacebookPageSelectResponseAccountPlatformYoutube        ConnectFacebookPageSelectResponseAccountPlatform = "youtube"
	ConnectFacebookPageSelectResponseAccountPlatformPinterest      ConnectFacebookPageSelectResponseAccountPlatform = "pinterest"
	ConnectFacebookPageSelectResponseAccountPlatformReddit         ConnectFacebookPageSelectResponseAccountPlatform = "reddit"
	ConnectFacebookPageSelectResponseAccountPlatformBluesky        ConnectFacebookPageSelectResponseAccountPlatform = "bluesky"
	ConnectFacebookPageSelectResponseAccountPlatformThreads        ConnectFacebookPageSelectResponseAccountPlatform = "threads"
	ConnectFacebookPageSelectResponseAccountPlatformTelegram       ConnectFacebookPageSelectResponseAccountPlatform = "telegram"
	ConnectFacebookPageSelectResponseAccountPlatformSnapchat       ConnectFacebookPageSelectResponseAccountPlatform = "snapchat"
	ConnectFacebookPageSelectResponseAccountPlatformGooglebusiness ConnectFacebookPageSelectResponseAccountPlatform = "googlebusiness"
	ConnectFacebookPageSelectResponseAccountPlatformWhatsapp       ConnectFacebookPageSelectResponseAccountPlatform = "whatsapp"
	ConnectFacebookPageSelectResponseAccountPlatformMastodon       ConnectFacebookPageSelectResponseAccountPlatform = "mastodon"
	ConnectFacebookPageSelectResponseAccountPlatformDiscord        ConnectFacebookPageSelectResponseAccountPlatform = "discord"
	ConnectFacebookPageSelectResponseAccountPlatformSMS            ConnectFacebookPageSelectResponseAccountPlatform = "sms"
	ConnectFacebookPageSelectResponseAccountPlatformBeehiiv        ConnectFacebookPageSelectResponseAccountPlatform = "beehiiv"
	ConnectFacebookPageSelectResponseAccountPlatformConvertkit     ConnectFacebookPageSelectResponseAccountPlatform = "convertkit"
	ConnectFacebookPageSelectResponseAccountPlatformMailchimp      ConnectFacebookPageSelectResponseAccountPlatform = "mailchimp"
	ConnectFacebookPageSelectResponseAccountPlatformListmonk       ConnectFacebookPageSelectResponseAccountPlatform = "listmonk"
)

func (r ConnectFacebookPageSelectResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectFacebookPageSelectResponseAccountPlatformTwitter, ConnectFacebookPageSelectResponseAccountPlatformInstagram, ConnectFacebookPageSelectResponseAccountPlatformFacebook, ConnectFacebookPageSelectResponseAccountPlatformLinkedin, ConnectFacebookPageSelectResponseAccountPlatformTiktok, ConnectFacebookPageSelectResponseAccountPlatformYoutube, ConnectFacebookPageSelectResponseAccountPlatformPinterest, ConnectFacebookPageSelectResponseAccountPlatformReddit, ConnectFacebookPageSelectResponseAccountPlatformBluesky, ConnectFacebookPageSelectResponseAccountPlatformThreads, ConnectFacebookPageSelectResponseAccountPlatformTelegram, ConnectFacebookPageSelectResponseAccountPlatformSnapchat, ConnectFacebookPageSelectResponseAccountPlatformGooglebusiness, ConnectFacebookPageSelectResponseAccountPlatformWhatsapp, ConnectFacebookPageSelectResponseAccountPlatformMastodon, ConnectFacebookPageSelectResponseAccountPlatformDiscord, ConnectFacebookPageSelectResponseAccountPlatformSMS, ConnectFacebookPageSelectResponseAccountPlatformBeehiiv, ConnectFacebookPageSelectResponseAccountPlatformConvertkit, ConnectFacebookPageSelectResponseAccountPlatformMailchimp, ConnectFacebookPageSelectResponseAccountPlatformListmonk:
		return true
	}
	return false
}

type ConnectFacebookPageSelectParams struct {
	// Token from pending data or OAuth flow
	ConnectToken param.Field[string] `json:"connect_token" api:"required"`
	// Selected Facebook page ID
	PageID param.Field[string] `json:"page_id" api:"required"`
}

func (r ConnectFacebookPageSelectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
