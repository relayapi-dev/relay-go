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

// ConnectTelegramService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectTelegramService] method instead.
type ConnectTelegramService struct {
	Options []option.RequestOption
}

// NewConnectTelegramService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectTelegramService(opts ...option.RequestOption) (r *ConnectTelegramService) {
	r = &ConnectTelegramService{}
	r.Options = opts
	return
}

// Connect Telegram directly with chat ID
func (r *ConnectTelegramService) ConnectDirectly(ctx context.Context, body ConnectTelegramConnectDirectlyParams, opts ...option.RequestOption) (res *ConnectTelegramConnectDirectlyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/telegram/direct"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generates a 6-character access code (valid 15 minutes).
func (r *ConnectTelegramService) InitiateConnection(ctx context.Context, opts ...option.RequestOption) (res *ConnectTelegramInitiateConnectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/telegram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Poll Telegram connection status
func (r *ConnectTelegramService) PollConnectionStatus(ctx context.Context, query ConnectTelegramPollConnectionStatusParams, opts ...option.RequestOption) (res *ConnectTelegramPollConnectionStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/telegram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConnectTelegramConnectDirectlyResponse struct {
	Account ConnectTelegramConnectDirectlyResponseAccount `json:"account" api:"required"`
	JSON    connectTelegramConnectDirectlyResponseJSON    `json:"-"`
}

// connectTelegramConnectDirectlyResponseJSON contains the JSON metadata for the
// struct [ConnectTelegramConnectDirectlyResponse]
type connectTelegramConnectDirectlyResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectTelegramConnectDirectlyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramConnectDirectlyResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectTelegramConnectDirectlyResponseAccount struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             ConnectTelegramConnectDirectlyResponseAccountGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                                `json:"metadata" api:"required,nullable"`
	Platform          ConnectTelegramConnectDirectlyResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                             `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                `json:"username" api:"required,nullable"`
	JSON              connectTelegramConnectDirectlyResponseAccountJSON     `json:"-"`
}

// connectTelegramConnectDirectlyResponseAccountJSON contains the JSON metadata for
// the struct [ConnectTelegramConnectDirectlyResponseAccount]
type connectTelegramConnectDirectlyResponseAccountJSON struct {
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

func (r *ConnectTelegramConnectDirectlyResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramConnectDirectlyResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account group
type ConnectTelegramConnectDirectlyResponseAccountGroup struct {
	ID   string                                                 `json:"id" api:"required"`
	Name string                                                 `json:"name" api:"required"`
	JSON connectTelegramConnectDirectlyResponseAccountGroupJSON `json:"-"`
}

// connectTelegramConnectDirectlyResponseAccountGroupJSON contains the JSON
// metadata for the struct [ConnectTelegramConnectDirectlyResponseAccountGroup]
type connectTelegramConnectDirectlyResponseAccountGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectTelegramConnectDirectlyResponseAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramConnectDirectlyResponseAccountGroupJSON) RawJSON() string {
	return r.raw
}

type ConnectTelegramConnectDirectlyResponseAccountPlatform string

const (
	ConnectTelegramConnectDirectlyResponseAccountPlatformTwitter        ConnectTelegramConnectDirectlyResponseAccountPlatform = "twitter"
	ConnectTelegramConnectDirectlyResponseAccountPlatformInstagram      ConnectTelegramConnectDirectlyResponseAccountPlatform = "instagram"
	ConnectTelegramConnectDirectlyResponseAccountPlatformFacebook       ConnectTelegramConnectDirectlyResponseAccountPlatform = "facebook"
	ConnectTelegramConnectDirectlyResponseAccountPlatformLinkedin       ConnectTelegramConnectDirectlyResponseAccountPlatform = "linkedin"
	ConnectTelegramConnectDirectlyResponseAccountPlatformTiktok         ConnectTelegramConnectDirectlyResponseAccountPlatform = "tiktok"
	ConnectTelegramConnectDirectlyResponseAccountPlatformYoutube        ConnectTelegramConnectDirectlyResponseAccountPlatform = "youtube"
	ConnectTelegramConnectDirectlyResponseAccountPlatformPinterest      ConnectTelegramConnectDirectlyResponseAccountPlatform = "pinterest"
	ConnectTelegramConnectDirectlyResponseAccountPlatformReddit         ConnectTelegramConnectDirectlyResponseAccountPlatform = "reddit"
	ConnectTelegramConnectDirectlyResponseAccountPlatformBluesky        ConnectTelegramConnectDirectlyResponseAccountPlatform = "bluesky"
	ConnectTelegramConnectDirectlyResponseAccountPlatformThreads        ConnectTelegramConnectDirectlyResponseAccountPlatform = "threads"
	ConnectTelegramConnectDirectlyResponseAccountPlatformTelegram       ConnectTelegramConnectDirectlyResponseAccountPlatform = "telegram"
	ConnectTelegramConnectDirectlyResponseAccountPlatformSnapchat       ConnectTelegramConnectDirectlyResponseAccountPlatform = "snapchat"
	ConnectTelegramConnectDirectlyResponseAccountPlatformGooglebusiness ConnectTelegramConnectDirectlyResponseAccountPlatform = "googlebusiness"
	ConnectTelegramConnectDirectlyResponseAccountPlatformWhatsapp       ConnectTelegramConnectDirectlyResponseAccountPlatform = "whatsapp"
	ConnectTelegramConnectDirectlyResponseAccountPlatformMastodon       ConnectTelegramConnectDirectlyResponseAccountPlatform = "mastodon"
	ConnectTelegramConnectDirectlyResponseAccountPlatformDiscord        ConnectTelegramConnectDirectlyResponseAccountPlatform = "discord"
	ConnectTelegramConnectDirectlyResponseAccountPlatformSMS            ConnectTelegramConnectDirectlyResponseAccountPlatform = "sms"
)

func (r ConnectTelegramConnectDirectlyResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectTelegramConnectDirectlyResponseAccountPlatformTwitter, ConnectTelegramConnectDirectlyResponseAccountPlatformInstagram, ConnectTelegramConnectDirectlyResponseAccountPlatformFacebook, ConnectTelegramConnectDirectlyResponseAccountPlatformLinkedin, ConnectTelegramConnectDirectlyResponseAccountPlatformTiktok, ConnectTelegramConnectDirectlyResponseAccountPlatformYoutube, ConnectTelegramConnectDirectlyResponseAccountPlatformPinterest, ConnectTelegramConnectDirectlyResponseAccountPlatformReddit, ConnectTelegramConnectDirectlyResponseAccountPlatformBluesky, ConnectTelegramConnectDirectlyResponseAccountPlatformThreads, ConnectTelegramConnectDirectlyResponseAccountPlatformTelegram, ConnectTelegramConnectDirectlyResponseAccountPlatformSnapchat, ConnectTelegramConnectDirectlyResponseAccountPlatformGooglebusiness, ConnectTelegramConnectDirectlyResponseAccountPlatformWhatsapp, ConnectTelegramConnectDirectlyResponseAccountPlatformMastodon, ConnectTelegramConnectDirectlyResponseAccountPlatformDiscord, ConnectTelegramConnectDirectlyResponseAccountPlatformSMS:
		return true
	}
	return false
}

type ConnectTelegramInitiateConnectionResponse struct {
	// Telegram bot username to message
	BotUsername string `json:"bot_username" api:"required"`
	// 6-character access code
	Code string `json:"code" api:"required"`
	// ISO 8601 expiry timestamp
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Seconds until code expires
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// Step-by-step instructions for the user
	Instructions []string                                      `json:"instructions" api:"required"`
	JSON         connectTelegramInitiateConnectionResponseJSON `json:"-"`
}

// connectTelegramInitiateConnectionResponseJSON contains the JSON metadata for the
// struct [ConnectTelegramInitiateConnectionResponse]
type connectTelegramInitiateConnectionResponseJSON struct {
	BotUsername  apijson.Field
	Code         apijson.Field
	ExpiresAt    apijson.Field
	ExpiresIn    apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectTelegramInitiateConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramInitiateConnectionResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectTelegramPollConnectionStatusResponse struct {
	// Current connection status
	Status ConnectTelegramPollConnectionStatusResponseStatus `json:"status" api:"required"`
	// Connected account details
	Account ConnectTelegramPollConnectionStatusResponseAccount `json:"account"`
	// Telegram chat ID once connected
	ChatID string `json:"chat_id"`
	// Chat or channel title
	ChatTitle string `json:"chat_title"`
	// Chat type (private, group, supergroup, channel)
	ChatType string `json:"chat_type"`
	// Code expiry timestamp
	ExpiresAt time.Time                                       `json:"expires_at" format:"date-time"`
	JSON      connectTelegramPollConnectionStatusResponseJSON `json:"-"`
}

// connectTelegramPollConnectionStatusResponseJSON contains the JSON metadata for
// the struct [ConnectTelegramPollConnectionStatusResponse]
type connectTelegramPollConnectionStatusResponseJSON struct {
	Status      apijson.Field
	Account     apijson.Field
	ChatID      apijson.Field
	ChatTitle   apijson.Field
	ChatType    apijson.Field
	ExpiresAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectTelegramPollConnectionStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramPollConnectionStatusResponseJSON) RawJSON() string {
	return r.raw
}

// Current connection status
type ConnectTelegramPollConnectionStatusResponseStatus string

const (
	ConnectTelegramPollConnectionStatusResponseStatusPending   ConnectTelegramPollConnectionStatusResponseStatus = "pending"
	ConnectTelegramPollConnectionStatusResponseStatusConnected ConnectTelegramPollConnectionStatusResponseStatus = "connected"
	ConnectTelegramPollConnectionStatusResponseStatusExpired   ConnectTelegramPollConnectionStatusResponseStatus = "expired"
)

func (r ConnectTelegramPollConnectionStatusResponseStatus) IsKnown() bool {
	switch r {
	case ConnectTelegramPollConnectionStatusResponseStatusPending, ConnectTelegramPollConnectionStatusResponseStatusConnected, ConnectTelegramPollConnectionStatusResponseStatusExpired:
		return true
	}
	return false
}

// Connected account details
type ConnectTelegramPollConnectionStatusResponseAccount struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             ConnectTelegramPollConnectionStatusResponseAccountGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                                     `json:"metadata" api:"required,nullable"`
	Platform          ConnectTelegramPollConnectionStatusResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                     `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                                  `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                     `json:"username" api:"required,nullable"`
	JSON              connectTelegramPollConnectionStatusResponseAccountJSON     `json:"-"`
}

// connectTelegramPollConnectionStatusResponseAccountJSON contains the JSON
// metadata for the struct [ConnectTelegramPollConnectionStatusResponseAccount]
type connectTelegramPollConnectionStatusResponseAccountJSON struct {
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

func (r *ConnectTelegramPollConnectionStatusResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramPollConnectionStatusResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account group
type ConnectTelegramPollConnectionStatusResponseAccountGroup struct {
	ID   string                                                      `json:"id" api:"required"`
	Name string                                                      `json:"name" api:"required"`
	JSON connectTelegramPollConnectionStatusResponseAccountGroupJSON `json:"-"`
}

// connectTelegramPollConnectionStatusResponseAccountGroupJSON contains the JSON
// metadata for the struct
// [ConnectTelegramPollConnectionStatusResponseAccountGroup]
type connectTelegramPollConnectionStatusResponseAccountGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectTelegramPollConnectionStatusResponseAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectTelegramPollConnectionStatusResponseAccountGroupJSON) RawJSON() string {
	return r.raw
}

type ConnectTelegramPollConnectionStatusResponseAccountPlatform string

const (
	ConnectTelegramPollConnectionStatusResponseAccountPlatformTwitter        ConnectTelegramPollConnectionStatusResponseAccountPlatform = "twitter"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformInstagram      ConnectTelegramPollConnectionStatusResponseAccountPlatform = "instagram"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformFacebook       ConnectTelegramPollConnectionStatusResponseAccountPlatform = "facebook"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformLinkedin       ConnectTelegramPollConnectionStatusResponseAccountPlatform = "linkedin"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformTiktok         ConnectTelegramPollConnectionStatusResponseAccountPlatform = "tiktok"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformYoutube        ConnectTelegramPollConnectionStatusResponseAccountPlatform = "youtube"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformPinterest      ConnectTelegramPollConnectionStatusResponseAccountPlatform = "pinterest"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformReddit         ConnectTelegramPollConnectionStatusResponseAccountPlatform = "reddit"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformBluesky        ConnectTelegramPollConnectionStatusResponseAccountPlatform = "bluesky"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformThreads        ConnectTelegramPollConnectionStatusResponseAccountPlatform = "threads"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformTelegram       ConnectTelegramPollConnectionStatusResponseAccountPlatform = "telegram"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformSnapchat       ConnectTelegramPollConnectionStatusResponseAccountPlatform = "snapchat"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformGooglebusiness ConnectTelegramPollConnectionStatusResponseAccountPlatform = "googlebusiness"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformWhatsapp       ConnectTelegramPollConnectionStatusResponseAccountPlatform = "whatsapp"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformMastodon       ConnectTelegramPollConnectionStatusResponseAccountPlatform = "mastodon"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformDiscord        ConnectTelegramPollConnectionStatusResponseAccountPlatform = "discord"
	ConnectTelegramPollConnectionStatusResponseAccountPlatformSMS            ConnectTelegramPollConnectionStatusResponseAccountPlatform = "sms"
)

func (r ConnectTelegramPollConnectionStatusResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectTelegramPollConnectionStatusResponseAccountPlatformTwitter, ConnectTelegramPollConnectionStatusResponseAccountPlatformInstagram, ConnectTelegramPollConnectionStatusResponseAccountPlatformFacebook, ConnectTelegramPollConnectionStatusResponseAccountPlatformLinkedin, ConnectTelegramPollConnectionStatusResponseAccountPlatformTiktok, ConnectTelegramPollConnectionStatusResponseAccountPlatformYoutube, ConnectTelegramPollConnectionStatusResponseAccountPlatformPinterest, ConnectTelegramPollConnectionStatusResponseAccountPlatformReddit, ConnectTelegramPollConnectionStatusResponseAccountPlatformBluesky, ConnectTelegramPollConnectionStatusResponseAccountPlatformThreads, ConnectTelegramPollConnectionStatusResponseAccountPlatformTelegram, ConnectTelegramPollConnectionStatusResponseAccountPlatformSnapchat, ConnectTelegramPollConnectionStatusResponseAccountPlatformGooglebusiness, ConnectTelegramPollConnectionStatusResponseAccountPlatformWhatsapp, ConnectTelegramPollConnectionStatusResponseAccountPlatformMastodon, ConnectTelegramPollConnectionStatusResponseAccountPlatformDiscord, ConnectTelegramPollConnectionStatusResponseAccountPlatformSMS:
		return true
	}
	return false
}

type ConnectTelegramConnectDirectlyParams struct {
	// Telegram chat or channel ID
	ChatID param.Field[string] `json:"chat_id" api:"required"`
}

func (r ConnectTelegramConnectDirectlyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectTelegramPollConnectionStatusParams struct {
	// The 6-character access code to check
	Code param.Field[string] `query:"code" api:"required"`
}

// URLQuery serializes [ConnectTelegramPollConnectionStatusParams]'s query
// parameters as `url.Values`.
func (r ConnectTelegramPollConnectionStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
