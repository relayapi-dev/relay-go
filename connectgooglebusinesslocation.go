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

// ConnectGooglebusinessLocationService contains methods and other services that
// help with interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectGooglebusinessLocationService] method instead.
type ConnectGooglebusinessLocationService struct {
	Options []option.RequestOption
}

// NewConnectGooglebusinessLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectGooglebusinessLocationService(opts ...option.RequestOption) (r *ConnectGooglebusinessLocationService) {
	r = &ConnectGooglebusinessLocationService{}
	r.Options = opts
	return
}

// List Google Business locations after OAuth
func (r *ConnectGooglebusinessLocationService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectGooglebusinessLocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/googlebusiness/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Select Google Business location
func (r *ConnectGooglebusinessLocationService) Select(ctx context.Context, body ConnectGooglebusinessLocationSelectParams, opts ...option.RequestOption) (res *ConnectGooglebusinessLocationSelectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/googlebusiness/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConnectGooglebusinessLocationListResponse struct {
	Locations []ConnectGooglebusinessLocationListResponseLocation `json:"locations" api:"required"`
	JSON      connectGooglebusinessLocationListResponseJSON       `json:"-"`
}

// connectGooglebusinessLocationListResponseJSON contains the JSON metadata for the
// struct [ConnectGooglebusinessLocationListResponse]
type connectGooglebusinessLocationListResponseJSON struct {
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectGooglebusinessLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectGooglebusinessLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectGooglebusinessLocationListResponseLocation struct {
	// Google Business location ID
	ID string `json:"id" api:"required"`
	// Business name
	Name string `json:"name" api:"required"`
	// Business address
	Address string `json:"address" api:"nullable"`
	// Business phone number
	Phone string                                                `json:"phone" api:"nullable"`
	JSON  connectGooglebusinessLocationListResponseLocationJSON `json:"-"`
}

// connectGooglebusinessLocationListResponseLocationJSON contains the JSON metadata
// for the struct [ConnectGooglebusinessLocationListResponseLocation]
type connectGooglebusinessLocationListResponseLocationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Address     apijson.Field
	Phone       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectGooglebusinessLocationListResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectGooglebusinessLocationListResponseLocationJSON) RawJSON() string {
	return r.raw
}

type ConnectGooglebusinessLocationSelectResponse struct {
	Account ConnectGooglebusinessLocationSelectResponseAccount `json:"account" api:"required"`
	JSON    connectGooglebusinessLocationSelectResponseJSON    `json:"-"`
}

// connectGooglebusinessLocationSelectResponseJSON contains the JSON metadata for
// the struct [ConnectGooglebusinessLocationSelectResponse]
type connectGooglebusinessLocationSelectResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectGooglebusinessLocationSelectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectGooglebusinessLocationSelectResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectGooglebusinessLocationSelectResponseAccount struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             ConnectGooglebusinessLocationSelectResponseAccountGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                                     `json:"metadata" api:"required,nullable"`
	Platform          ConnectGooglebusinessLocationSelectResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                     `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                                  `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                     `json:"username" api:"required,nullable"`
	JSON              connectGooglebusinessLocationSelectResponseAccountJSON     `json:"-"`
}

// connectGooglebusinessLocationSelectResponseAccountJSON contains the JSON
// metadata for the struct [ConnectGooglebusinessLocationSelectResponseAccount]
type connectGooglebusinessLocationSelectResponseAccountJSON struct {
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

func (r *ConnectGooglebusinessLocationSelectResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectGooglebusinessLocationSelectResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account group
type ConnectGooglebusinessLocationSelectResponseAccountGroup struct {
	ID   string                                                      `json:"id" api:"required"`
	Name string                                                      `json:"name" api:"required"`
	JSON connectGooglebusinessLocationSelectResponseAccountGroupJSON `json:"-"`
}

// connectGooglebusinessLocationSelectResponseAccountGroupJSON contains the JSON
// metadata for the struct
// [ConnectGooglebusinessLocationSelectResponseAccountGroup]
type connectGooglebusinessLocationSelectResponseAccountGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectGooglebusinessLocationSelectResponseAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectGooglebusinessLocationSelectResponseAccountGroupJSON) RawJSON() string {
	return r.raw
}

type ConnectGooglebusinessLocationSelectResponseAccountPlatform string

const (
	ConnectGooglebusinessLocationSelectResponseAccountPlatformTwitter        ConnectGooglebusinessLocationSelectResponseAccountPlatform = "twitter"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformInstagram      ConnectGooglebusinessLocationSelectResponseAccountPlatform = "instagram"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformFacebook       ConnectGooglebusinessLocationSelectResponseAccountPlatform = "facebook"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformLinkedin       ConnectGooglebusinessLocationSelectResponseAccountPlatform = "linkedin"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformTiktok         ConnectGooglebusinessLocationSelectResponseAccountPlatform = "tiktok"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformYoutube        ConnectGooglebusinessLocationSelectResponseAccountPlatform = "youtube"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformPinterest      ConnectGooglebusinessLocationSelectResponseAccountPlatform = "pinterest"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformReddit         ConnectGooglebusinessLocationSelectResponseAccountPlatform = "reddit"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformBluesky        ConnectGooglebusinessLocationSelectResponseAccountPlatform = "bluesky"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformThreads        ConnectGooglebusinessLocationSelectResponseAccountPlatform = "threads"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformTelegram       ConnectGooglebusinessLocationSelectResponseAccountPlatform = "telegram"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformSnapchat       ConnectGooglebusinessLocationSelectResponseAccountPlatform = "snapchat"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformGooglebusiness ConnectGooglebusinessLocationSelectResponseAccountPlatform = "googlebusiness"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformWhatsapp       ConnectGooglebusinessLocationSelectResponseAccountPlatform = "whatsapp"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformMastodon       ConnectGooglebusinessLocationSelectResponseAccountPlatform = "mastodon"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformDiscord        ConnectGooglebusinessLocationSelectResponseAccountPlatform = "discord"
	ConnectGooglebusinessLocationSelectResponseAccountPlatformSMS            ConnectGooglebusinessLocationSelectResponseAccountPlatform = "sms"
)

func (r ConnectGooglebusinessLocationSelectResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectGooglebusinessLocationSelectResponseAccountPlatformTwitter, ConnectGooglebusinessLocationSelectResponseAccountPlatformInstagram, ConnectGooglebusinessLocationSelectResponseAccountPlatformFacebook, ConnectGooglebusinessLocationSelectResponseAccountPlatformLinkedin, ConnectGooglebusinessLocationSelectResponseAccountPlatformTiktok, ConnectGooglebusinessLocationSelectResponseAccountPlatformYoutube, ConnectGooglebusinessLocationSelectResponseAccountPlatformPinterest, ConnectGooglebusinessLocationSelectResponseAccountPlatformReddit, ConnectGooglebusinessLocationSelectResponseAccountPlatformBluesky, ConnectGooglebusinessLocationSelectResponseAccountPlatformThreads, ConnectGooglebusinessLocationSelectResponseAccountPlatformTelegram, ConnectGooglebusinessLocationSelectResponseAccountPlatformSnapchat, ConnectGooglebusinessLocationSelectResponseAccountPlatformGooglebusiness, ConnectGooglebusinessLocationSelectResponseAccountPlatformWhatsapp, ConnectGooglebusinessLocationSelectResponseAccountPlatformMastodon, ConnectGooglebusinessLocationSelectResponseAccountPlatformDiscord, ConnectGooglebusinessLocationSelectResponseAccountPlatformSMS:
		return true
	}
	return false
}

type ConnectGooglebusinessLocationSelectParams struct {
	// Token from pending data or OAuth flow
	ConnectToken param.Field[string] `json:"connect_token" api:"required"`
	// Selected Google Business location ID
	LocationID param.Field[string] `json:"location_id" api:"required"`
}

func (r ConnectGooglebusinessLocationSelectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
