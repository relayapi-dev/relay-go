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

// ConnectPinterestBoardService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectPinterestBoardService] method instead.
type ConnectPinterestBoardService struct {
	Options []option.RequestOption
}

// NewConnectPinterestBoardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectPinterestBoardService(opts ...option.RequestOption) (r *ConnectPinterestBoardService) {
	r = &ConnectPinterestBoardService{}
	r.Options = opts
	return
}

// List Pinterest boards after OAuth
func (r *ConnectPinterestBoardService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectPinterestBoardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/pinterest/boards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Select Pinterest board
func (r *ConnectPinterestBoardService) Select(ctx context.Context, body ConnectPinterestBoardSelectParams, opts ...option.RequestOption) (res *ConnectPinterestBoardSelectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/pinterest/boards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConnectPinterestBoardListResponse struct {
	Boards []ConnectPinterestBoardListResponseBoard `json:"boards" api:"required"`
	JSON   connectPinterestBoardListResponseJSON    `json:"-"`
}

// connectPinterestBoardListResponseJSON contains the JSON metadata for the struct
// [ConnectPinterestBoardListResponse]
type connectPinterestBoardListResponseJSON struct {
	Boards      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectPinterestBoardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectPinterestBoardListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectPinterestBoardListResponseBoard struct {
	// Pinterest board ID
	ID string `json:"id" api:"required"`
	// Board name
	Name string `json:"name" api:"required"`
	// Board description
	Description string `json:"description" api:"nullable"`
	// Number of pins on the board
	PinCount int64                                      `json:"pin_count"`
	JSON     connectPinterestBoardListResponseBoardJSON `json:"-"`
}

// connectPinterestBoardListResponseBoardJSON contains the JSON metadata for the
// struct [ConnectPinterestBoardListResponseBoard]
type connectPinterestBoardListResponseBoardJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Description apijson.Field
	PinCount    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectPinterestBoardListResponseBoard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectPinterestBoardListResponseBoardJSON) RawJSON() string {
	return r.raw
}

type ConnectPinterestBoardSelectResponse struct {
	Account ConnectPinterestBoardSelectResponseAccount `json:"account" api:"required"`
	JSON    connectPinterestBoardSelectResponseJSON    `json:"-"`
}

// connectPinterestBoardSelectResponseJSON contains the JSON metadata for the
// struct [ConnectPinterestBoardSelectResponse]
type connectPinterestBoardSelectResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectPinterestBoardSelectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectPinterestBoardSelectResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectPinterestBoardSelectResponseAccount struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             ConnectPinterestBoardSelectResponseAccountGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                             `json:"metadata" api:"required,nullable"`
	Platform          ConnectPinterestBoardSelectResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                             `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                          `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                             `json:"username" api:"required,nullable"`
	JSON              connectPinterestBoardSelectResponseAccountJSON     `json:"-"`
}

// connectPinterestBoardSelectResponseAccountJSON contains the JSON metadata for
// the struct [ConnectPinterestBoardSelectResponseAccount]
type connectPinterestBoardSelectResponseAccountJSON struct {
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

func (r *ConnectPinterestBoardSelectResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectPinterestBoardSelectResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account group
type ConnectPinterestBoardSelectResponseAccountGroup struct {
	ID   string                                              `json:"id" api:"required"`
	Name string                                              `json:"name" api:"required"`
	JSON connectPinterestBoardSelectResponseAccountGroupJSON `json:"-"`
}

// connectPinterestBoardSelectResponseAccountGroupJSON contains the JSON metadata
// for the struct [ConnectPinterestBoardSelectResponseAccountGroup]
type connectPinterestBoardSelectResponseAccountGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectPinterestBoardSelectResponseAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectPinterestBoardSelectResponseAccountGroupJSON) RawJSON() string {
	return r.raw
}

type ConnectPinterestBoardSelectResponseAccountPlatform string

const (
	ConnectPinterestBoardSelectResponseAccountPlatformTwitter        ConnectPinterestBoardSelectResponseAccountPlatform = "twitter"
	ConnectPinterestBoardSelectResponseAccountPlatformInstagram      ConnectPinterestBoardSelectResponseAccountPlatform = "instagram"
	ConnectPinterestBoardSelectResponseAccountPlatformFacebook       ConnectPinterestBoardSelectResponseAccountPlatform = "facebook"
	ConnectPinterestBoardSelectResponseAccountPlatformLinkedin       ConnectPinterestBoardSelectResponseAccountPlatform = "linkedin"
	ConnectPinterestBoardSelectResponseAccountPlatformTiktok         ConnectPinterestBoardSelectResponseAccountPlatform = "tiktok"
	ConnectPinterestBoardSelectResponseAccountPlatformYoutube        ConnectPinterestBoardSelectResponseAccountPlatform = "youtube"
	ConnectPinterestBoardSelectResponseAccountPlatformPinterest      ConnectPinterestBoardSelectResponseAccountPlatform = "pinterest"
	ConnectPinterestBoardSelectResponseAccountPlatformReddit         ConnectPinterestBoardSelectResponseAccountPlatform = "reddit"
	ConnectPinterestBoardSelectResponseAccountPlatformBluesky        ConnectPinterestBoardSelectResponseAccountPlatform = "bluesky"
	ConnectPinterestBoardSelectResponseAccountPlatformThreads        ConnectPinterestBoardSelectResponseAccountPlatform = "threads"
	ConnectPinterestBoardSelectResponseAccountPlatformTelegram       ConnectPinterestBoardSelectResponseAccountPlatform = "telegram"
	ConnectPinterestBoardSelectResponseAccountPlatformSnapchat       ConnectPinterestBoardSelectResponseAccountPlatform = "snapchat"
	ConnectPinterestBoardSelectResponseAccountPlatformGooglebusiness ConnectPinterestBoardSelectResponseAccountPlatform = "googlebusiness"
	ConnectPinterestBoardSelectResponseAccountPlatformWhatsapp       ConnectPinterestBoardSelectResponseAccountPlatform = "whatsapp"
	ConnectPinterestBoardSelectResponseAccountPlatformMastodon       ConnectPinterestBoardSelectResponseAccountPlatform = "mastodon"
	ConnectPinterestBoardSelectResponseAccountPlatformDiscord        ConnectPinterestBoardSelectResponseAccountPlatform = "discord"
	ConnectPinterestBoardSelectResponseAccountPlatformSMS            ConnectPinterestBoardSelectResponseAccountPlatform = "sms"
	ConnectPinterestBoardSelectResponseAccountPlatformBeehiiv        ConnectPinterestBoardSelectResponseAccountPlatform = "beehiiv"
	ConnectPinterestBoardSelectResponseAccountPlatformConvertkit     ConnectPinterestBoardSelectResponseAccountPlatform = "convertkit"
	ConnectPinterestBoardSelectResponseAccountPlatformMailchimp      ConnectPinterestBoardSelectResponseAccountPlatform = "mailchimp"
	ConnectPinterestBoardSelectResponseAccountPlatformListmonk       ConnectPinterestBoardSelectResponseAccountPlatform = "listmonk"
)

func (r ConnectPinterestBoardSelectResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectPinterestBoardSelectResponseAccountPlatformTwitter, ConnectPinterestBoardSelectResponseAccountPlatformInstagram, ConnectPinterestBoardSelectResponseAccountPlatformFacebook, ConnectPinterestBoardSelectResponseAccountPlatformLinkedin, ConnectPinterestBoardSelectResponseAccountPlatformTiktok, ConnectPinterestBoardSelectResponseAccountPlatformYoutube, ConnectPinterestBoardSelectResponseAccountPlatformPinterest, ConnectPinterestBoardSelectResponseAccountPlatformReddit, ConnectPinterestBoardSelectResponseAccountPlatformBluesky, ConnectPinterestBoardSelectResponseAccountPlatformThreads, ConnectPinterestBoardSelectResponseAccountPlatformTelegram, ConnectPinterestBoardSelectResponseAccountPlatformSnapchat, ConnectPinterestBoardSelectResponseAccountPlatformGooglebusiness, ConnectPinterestBoardSelectResponseAccountPlatformWhatsapp, ConnectPinterestBoardSelectResponseAccountPlatformMastodon, ConnectPinterestBoardSelectResponseAccountPlatformDiscord, ConnectPinterestBoardSelectResponseAccountPlatformSMS, ConnectPinterestBoardSelectResponseAccountPlatformBeehiiv, ConnectPinterestBoardSelectResponseAccountPlatformConvertkit, ConnectPinterestBoardSelectResponseAccountPlatformMailchimp, ConnectPinterestBoardSelectResponseAccountPlatformListmonk:
		return true
	}
	return false
}

type ConnectPinterestBoardSelectParams struct {
	// Selected Pinterest board ID
	BoardID param.Field[string] `json:"board_id" api:"required"`
	// Token from pending data or OAuth flow
	ConnectToken param.Field[string] `json:"connect_token" api:"required"`
}

func (r ConnectPinterestBoardSelectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
