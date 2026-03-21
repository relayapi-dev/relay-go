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

// ConnectWhatsappService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectWhatsappService] method instead.
type ConnectWhatsappService struct {
	Options []option.RequestOption
}

// NewConnectWhatsappService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectWhatsappService(opts ...option.RequestOption) (r *ConnectWhatsappService) {
	r = &ConnectWhatsappService{}
	r.Options = opts
	return
}

// Complete WhatsApp Embedded Signup
func (r *ConnectWhatsappService) CompleteEmbeddedSignup(ctx context.Context, body ConnectWhatsappCompleteEmbeddedSignupParams, opts ...option.RequestOption) (res *ConnectWhatsappCompleteEmbeddedSignupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/whatsapp/embedded-signup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Connect WhatsApp via System User credentials
func (r *ConnectWhatsappService) ConnectViaCredentials(ctx context.Context, body ConnectWhatsappConnectViaCredentialsParams, opts ...option.RequestOption) (res *ConnectWhatsappConnectViaCredentialsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/whatsapp/credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get WhatsApp Embedded Signup SDK config
func (r *ConnectWhatsappService) GetSDKConfig(ctx context.Context, opts ...option.RequestOption) (res *ConnectWhatsappGetSDKConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/whatsapp/sdk-config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ConnectWhatsappCompleteEmbeddedSignupResponse struct {
	Account ConnectWhatsappCompleteEmbeddedSignupResponseAccount `json:"account" api:"required"`
	JSON    connectWhatsappCompleteEmbeddedSignupResponseJSON    `json:"-"`
}

// connectWhatsappCompleteEmbeddedSignupResponseJSON contains the JSON metadata for
// the struct [ConnectWhatsappCompleteEmbeddedSignupResponse]
type connectWhatsappCompleteEmbeddedSignupResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectWhatsappCompleteEmbeddedSignupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectWhatsappCompleteEmbeddedSignupResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectWhatsappCompleteEmbeddedSignupResponseAccount struct {
	// Account ID
	ID                string                                                       `json:"id" api:"required"`
	AvatarURL         string                                                       `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                                    `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                                       `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                                       `json:"metadata" api:"required,nullable"`
	Platform          ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                       `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                                    `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                       `json:"username" api:"required,nullable"`
	JSON              connectWhatsappCompleteEmbeddedSignupResponseAccountJSON     `json:"-"`
}

// connectWhatsappCompleteEmbeddedSignupResponseAccountJSON contains the JSON
// metadata for the struct [ConnectWhatsappCompleteEmbeddedSignupResponseAccount]
type connectWhatsappCompleteEmbeddedSignupResponseAccountJSON struct {
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

func (r *ConnectWhatsappCompleteEmbeddedSignupResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectWhatsappCompleteEmbeddedSignupResponseAccountJSON) RawJSON() string {
	return r.raw
}

type ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform string

const (
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformTwitter        ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "twitter"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformInstagram      ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "instagram"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformFacebook       ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "facebook"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformLinkedin       ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "linkedin"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformTiktok         ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "tiktok"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformYoutube        ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "youtube"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformPinterest      ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "pinterest"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformReddit         ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "reddit"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformBluesky        ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "bluesky"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformThreads        ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "threads"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformTelegram       ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "telegram"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformSnapchat       ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "snapchat"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformGooglebusiness ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "googlebusiness"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformWhatsapp       ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "whatsapp"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformMastodon       ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "mastodon"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformDiscord        ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "discord"
	ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformSMS            ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform = "sms"
)

func (r ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformTwitter, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformInstagram, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformFacebook, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformLinkedin, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformTiktok, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformYoutube, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformPinterest, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformReddit, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformBluesky, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformThreads, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformTelegram, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformSnapchat, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformGooglebusiness, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformWhatsapp, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformMastodon, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformDiscord, ConnectWhatsappCompleteEmbeddedSignupResponseAccountPlatformSMS:
		return true
	}
	return false
}

type ConnectWhatsappConnectViaCredentialsResponse struct {
	Account ConnectWhatsappConnectViaCredentialsResponseAccount `json:"account" api:"required"`
	JSON    connectWhatsappConnectViaCredentialsResponseJSON    `json:"-"`
}

// connectWhatsappConnectViaCredentialsResponseJSON contains the JSON metadata for
// the struct [ConnectWhatsappConnectViaCredentialsResponse]
type connectWhatsappConnectViaCredentialsResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectWhatsappConnectViaCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectWhatsappConnectViaCredentialsResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectWhatsappConnectViaCredentialsResponseAccount struct {
	// Account ID
	ID                string                                                      `json:"id" api:"required"`
	AvatarURL         string                                                      `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                                   `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                                      `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                                      `json:"metadata" api:"required,nullable"`
	Platform          ConnectWhatsappConnectViaCredentialsResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                      `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                                   `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                      `json:"username" api:"required,nullable"`
	JSON              connectWhatsappConnectViaCredentialsResponseAccountJSON     `json:"-"`
}

// connectWhatsappConnectViaCredentialsResponseAccountJSON contains the JSON
// metadata for the struct [ConnectWhatsappConnectViaCredentialsResponseAccount]
type connectWhatsappConnectViaCredentialsResponseAccountJSON struct {
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

func (r *ConnectWhatsappConnectViaCredentialsResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectWhatsappConnectViaCredentialsResponseAccountJSON) RawJSON() string {
	return r.raw
}

type ConnectWhatsappConnectViaCredentialsResponseAccountPlatform string

const (
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformTwitter        ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "twitter"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformInstagram      ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "instagram"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformFacebook       ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "facebook"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformLinkedin       ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "linkedin"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformTiktok         ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "tiktok"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformYoutube        ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "youtube"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformPinterest      ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "pinterest"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformReddit         ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "reddit"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformBluesky        ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "bluesky"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformThreads        ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "threads"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformTelegram       ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "telegram"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformSnapchat       ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "snapchat"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformGooglebusiness ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "googlebusiness"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformWhatsapp       ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "whatsapp"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformMastodon       ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "mastodon"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformDiscord        ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "discord"
	ConnectWhatsappConnectViaCredentialsResponseAccountPlatformSMS            ConnectWhatsappConnectViaCredentialsResponseAccountPlatform = "sms"
)

func (r ConnectWhatsappConnectViaCredentialsResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectWhatsappConnectViaCredentialsResponseAccountPlatformTwitter, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformInstagram, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformFacebook, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformLinkedin, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformTiktok, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformYoutube, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformPinterest, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformReddit, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformBluesky, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformThreads, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformTelegram, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformSnapchat, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformGooglebusiness, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformWhatsapp, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformMastodon, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformDiscord, ConnectWhatsappConnectViaCredentialsResponseAccountPlatformSMS:
		return true
	}
	return false
}

type ConnectWhatsappGetSDKConfigResponse struct {
	// Facebook App ID for WhatsApp embedded signup
	AppID string `json:"app_id" api:"required"`
	// WhatsApp configuration ID
	ConfigID string                                  `json:"config_id" api:"required"`
	JSON     connectWhatsappGetSDKConfigResponseJSON `json:"-"`
}

// connectWhatsappGetSDKConfigResponseJSON contains the JSON metadata for the
// struct [ConnectWhatsappGetSDKConfigResponse]
type connectWhatsappGetSDKConfigResponseJSON struct {
	AppID       apijson.Field
	ConfigID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectWhatsappGetSDKConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectWhatsappGetSDKConfigResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectWhatsappCompleteEmbeddedSignupParams struct {
	// Code from WhatsApp embedded signup flow
	Code param.Field[string] `json:"code" api:"required"`
}

func (r ConnectWhatsappCompleteEmbeddedSignupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectWhatsappConnectViaCredentialsParams struct {
	// WhatsApp Business API access token
	AccessToken param.Field[string] `json:"access_token" api:"required"`
	// WhatsApp phone number ID
	PhoneNumberID param.Field[string] `json:"phone_number_id" api:"required"`
	// WhatsApp Business Account ID
	WabaID param.Field[string] `json:"waba_id" api:"required"`
}

func (r ConnectWhatsappConnectViaCredentialsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
