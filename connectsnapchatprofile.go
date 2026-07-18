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

// ConnectSnapchatProfileService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectSnapchatProfileService] method instead.
type ConnectSnapchatProfileService struct {
	Options []option.RequestOption
}

// NewConnectSnapchatProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectSnapchatProfileService(opts ...option.RequestOption) (r *ConnectSnapchatProfileService) {
	r = &ConnectSnapchatProfileService{}
	r.Options = opts
	return
}

// List Snapchat Public Profiles after OAuth
func (r *ConnectSnapchatProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectSnapchatProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/snapchat/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Select Snapchat Public Profile
func (r *ConnectSnapchatProfileService) Select(ctx context.Context, body ConnectSnapchatProfileSelectParams, opts ...option.RequestOption) (res *ConnectSnapchatProfileSelectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/snapchat/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConnectSnapchatProfileListResponse struct {
	Profiles []ConnectSnapchatProfileListResponseProfile `json:"profiles" api:"required"`
	JSON     connectSnapchatProfileListResponseJSON      `json:"-"`
}

// connectSnapchatProfileListResponseJSON contains the JSON metadata for the struct
// [ConnectSnapchatProfileListResponse]
type connectSnapchatProfileListResponseJSON struct {
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectSnapchatProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSnapchatProfileListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectSnapchatProfileListResponseProfile struct {
	// Snapchat profile ID
	ID string `json:"id" api:"required"`
	// Display name
	DisplayName string `json:"display_name" api:"required"`
	// Snapchat username
	Username string `json:"username" api:"required"`
	// Profile image URL
	ProfileImageURL string `json:"profile_image_url" api:"nullable"`
	// Number of subscribers
	SubscriberCount int64                                         `json:"subscriber_count"`
	JSON            connectSnapchatProfileListResponseProfileJSON `json:"-"`
}

// connectSnapchatProfileListResponseProfileJSON contains the JSON metadata for the
// struct [ConnectSnapchatProfileListResponseProfile]
type connectSnapchatProfileListResponseProfileJSON struct {
	ID              apijson.Field
	DisplayName     apijson.Field
	Username        apijson.Field
	ProfileImageURL apijson.Field
	SubscriberCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectSnapchatProfileListResponseProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSnapchatProfileListResponseProfileJSON) RawJSON() string {
	return r.raw
}

type ConnectSnapchatProfileSelectResponse struct {
	Account ConnectSnapchatProfileSelectResponseAccount `json:"account" api:"required"`
	JSON    connectSnapchatProfileSelectResponseJSON    `json:"-"`
}

// connectSnapchatProfileSelectResponseJSON contains the JSON metadata for the
// struct [ConnectSnapchatProfileSelectResponse]
type connectSnapchatProfileSelectResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectSnapchatProfileSelectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSnapchatProfileSelectResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectSnapchatProfileSelectResponseAccount struct {
	// Account ID
	ID                string                                              `json:"id" api:"required"`
	AvatarURL         string                                              `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                           `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                              `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                              `json:"metadata" api:"required,nullable"`
	Platform          ConnectSnapchatProfileSelectResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                              `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                              `json:"username" api:"required,nullable"`
	// Account workspace
	Workspace ConnectSnapchatProfileSelectResponseAccountWorkspace `json:"workspace" api:"required,nullable"`
	JSON      connectSnapchatProfileSelectResponseAccountJSON      `json:"-"`
}

// connectSnapchatProfileSelectResponseAccountJSON contains the JSON metadata for
// the struct [ConnectSnapchatProfileSelectResponseAccount]
type connectSnapchatProfileSelectResponseAccountJSON struct {
	ID                apijson.Field
	AvatarURL         apijson.Field
	ConnectedAt       apijson.Field
	DisplayName       apijson.Field
	Metadata          apijson.Field
	Platform          apijson.Field
	PlatformAccountID apijson.Field
	UpdatedAt         apijson.Field
	Username          apijson.Field
	Workspace         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConnectSnapchatProfileSelectResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSnapchatProfileSelectResponseAccountJSON) RawJSON() string {
	return r.raw
}

type ConnectSnapchatProfileSelectResponseAccountPlatform string

const (
	ConnectSnapchatProfileSelectResponseAccountPlatformTwitter        ConnectSnapchatProfileSelectResponseAccountPlatform = "twitter"
	ConnectSnapchatProfileSelectResponseAccountPlatformInstagram      ConnectSnapchatProfileSelectResponseAccountPlatform = "instagram"
	ConnectSnapchatProfileSelectResponseAccountPlatformFacebook       ConnectSnapchatProfileSelectResponseAccountPlatform = "facebook"
	ConnectSnapchatProfileSelectResponseAccountPlatformLinkedin       ConnectSnapchatProfileSelectResponseAccountPlatform = "linkedin"
	ConnectSnapchatProfileSelectResponseAccountPlatformTiktok         ConnectSnapchatProfileSelectResponseAccountPlatform = "tiktok"
	ConnectSnapchatProfileSelectResponseAccountPlatformYoutube        ConnectSnapchatProfileSelectResponseAccountPlatform = "youtube"
	ConnectSnapchatProfileSelectResponseAccountPlatformPinterest      ConnectSnapchatProfileSelectResponseAccountPlatform = "pinterest"
	ConnectSnapchatProfileSelectResponseAccountPlatformReddit         ConnectSnapchatProfileSelectResponseAccountPlatform = "reddit"
	ConnectSnapchatProfileSelectResponseAccountPlatformBluesky        ConnectSnapchatProfileSelectResponseAccountPlatform = "bluesky"
	ConnectSnapchatProfileSelectResponseAccountPlatformThreads        ConnectSnapchatProfileSelectResponseAccountPlatform = "threads"
	ConnectSnapchatProfileSelectResponseAccountPlatformTelegram       ConnectSnapchatProfileSelectResponseAccountPlatform = "telegram"
	ConnectSnapchatProfileSelectResponseAccountPlatformSnapchat       ConnectSnapchatProfileSelectResponseAccountPlatform = "snapchat"
	ConnectSnapchatProfileSelectResponseAccountPlatformGooglebusiness ConnectSnapchatProfileSelectResponseAccountPlatform = "googlebusiness"
	ConnectSnapchatProfileSelectResponseAccountPlatformWhatsapp       ConnectSnapchatProfileSelectResponseAccountPlatform = "whatsapp"
	ConnectSnapchatProfileSelectResponseAccountPlatformMastodon       ConnectSnapchatProfileSelectResponseAccountPlatform = "mastodon"
	ConnectSnapchatProfileSelectResponseAccountPlatformDiscord        ConnectSnapchatProfileSelectResponseAccountPlatform = "discord"
	ConnectSnapchatProfileSelectResponseAccountPlatformSMS            ConnectSnapchatProfileSelectResponseAccountPlatform = "sms"
	ConnectSnapchatProfileSelectResponseAccountPlatformBeehiiv        ConnectSnapchatProfileSelectResponseAccountPlatform = "beehiiv"
	ConnectSnapchatProfileSelectResponseAccountPlatformConvertkit     ConnectSnapchatProfileSelectResponseAccountPlatform = "convertkit"
	ConnectSnapchatProfileSelectResponseAccountPlatformMailchimp      ConnectSnapchatProfileSelectResponseAccountPlatform = "mailchimp"
	ConnectSnapchatProfileSelectResponseAccountPlatformListmonk       ConnectSnapchatProfileSelectResponseAccountPlatform = "listmonk"
)

func (r ConnectSnapchatProfileSelectResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectSnapchatProfileSelectResponseAccountPlatformTwitter, ConnectSnapchatProfileSelectResponseAccountPlatformInstagram, ConnectSnapchatProfileSelectResponseAccountPlatformFacebook, ConnectSnapchatProfileSelectResponseAccountPlatformLinkedin, ConnectSnapchatProfileSelectResponseAccountPlatformTiktok, ConnectSnapchatProfileSelectResponseAccountPlatformYoutube, ConnectSnapchatProfileSelectResponseAccountPlatformPinterest, ConnectSnapchatProfileSelectResponseAccountPlatformReddit, ConnectSnapchatProfileSelectResponseAccountPlatformBluesky, ConnectSnapchatProfileSelectResponseAccountPlatformThreads, ConnectSnapchatProfileSelectResponseAccountPlatformTelegram, ConnectSnapchatProfileSelectResponseAccountPlatformSnapchat, ConnectSnapchatProfileSelectResponseAccountPlatformGooglebusiness, ConnectSnapchatProfileSelectResponseAccountPlatformWhatsapp, ConnectSnapchatProfileSelectResponseAccountPlatformMastodon, ConnectSnapchatProfileSelectResponseAccountPlatformDiscord, ConnectSnapchatProfileSelectResponseAccountPlatformSMS, ConnectSnapchatProfileSelectResponseAccountPlatformBeehiiv, ConnectSnapchatProfileSelectResponseAccountPlatformConvertkit, ConnectSnapchatProfileSelectResponseAccountPlatformMailchimp, ConnectSnapchatProfileSelectResponseAccountPlatformListmonk:
		return true
	}
	return false
}

// Account workspace
type ConnectSnapchatProfileSelectResponseAccountWorkspace struct {
	ID   string                                                   `json:"id" api:"required"`
	Name string                                                   `json:"name" api:"required"`
	JSON connectSnapchatProfileSelectResponseAccountWorkspaceJSON `json:"-"`
}

// connectSnapchatProfileSelectResponseAccountWorkspaceJSON contains the JSON
// metadata for the struct [ConnectSnapchatProfileSelectResponseAccountWorkspace]
type connectSnapchatProfileSelectResponseAccountWorkspaceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectSnapchatProfileSelectResponseAccountWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSnapchatProfileSelectResponseAccountWorkspaceJSON) RawJSON() string {
	return r.raw
}

type ConnectSnapchatProfileSelectParams struct {
	// Token from pending data or OAuth flow
	ConnectToken param.Field[string] `json:"connect_token" api:"required"`
	// Selected Snapchat profile ID
	ProfileID param.Field[string] `json:"profile_id" api:"required"`
}

func (r ConnectSnapchatProfileSelectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
