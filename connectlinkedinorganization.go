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

// ConnectLinkedinOrganizationService contains methods and other services that help
// with interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectLinkedinOrganizationService] method instead.
type ConnectLinkedinOrganizationService struct {
	Options []option.RequestOption
}

// NewConnectLinkedinOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectLinkedinOrganizationService(opts ...option.RequestOption) (r *ConnectLinkedinOrganizationService) {
	r = &ConnectLinkedinOrganizationService{}
	r.Options = opts
	return
}

// List LinkedIn organizations after OAuth
func (r *ConnectLinkedinOrganizationService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectLinkedinOrganizationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/linkedin/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Select LinkedIn organization
func (r *ConnectLinkedinOrganizationService) Select(ctx context.Context, body ConnectLinkedinOrganizationSelectParams, opts ...option.RequestOption) (res *ConnectLinkedinOrganizationSelectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/linkedin/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConnectLinkedinOrganizationListResponse struct {
	Organizations []ConnectLinkedinOrganizationListResponseOrganization `json:"organizations" api:"required"`
	// User's personal LinkedIn profile
	PersonalProfile ConnectLinkedinOrganizationListResponsePersonalProfile `json:"personal_profile"`
	JSON            connectLinkedinOrganizationListResponseJSON            `json:"-"`
}

// connectLinkedinOrganizationListResponseJSON contains the JSON metadata for the
// struct [ConnectLinkedinOrganizationListResponse]
type connectLinkedinOrganizationListResponseJSON struct {
	Organizations   apijson.Field
	PersonalProfile apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectLinkedinOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectLinkedinOrganizationListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectLinkedinOrganizationListResponseOrganization struct {
	// Organization name
	Name string `json:"name" api:"required"`
	// LinkedIn organization URN
	Urn string `json:"urn" api:"required"`
	// Organization logo URL
	LogoURL string `json:"logo_url" api:"nullable"`
	// Organization vanity name
	VanityName string                                                  `json:"vanity_name" api:"nullable"`
	JSON       connectLinkedinOrganizationListResponseOrganizationJSON `json:"-"`
}

// connectLinkedinOrganizationListResponseOrganizationJSON contains the JSON
// metadata for the struct [ConnectLinkedinOrganizationListResponseOrganization]
type connectLinkedinOrganizationListResponseOrganizationJSON struct {
	Name        apijson.Field
	Urn         apijson.Field
	LogoURL     apijson.Field
	VanityName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectLinkedinOrganizationListResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectLinkedinOrganizationListResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

// User's personal LinkedIn profile
type ConnectLinkedinOrganizationListResponsePersonalProfile struct {
	Name string                                                     `json:"name" api:"required"`
	Urn  string                                                     `json:"urn" api:"required"`
	JSON connectLinkedinOrganizationListResponsePersonalProfileJSON `json:"-"`
}

// connectLinkedinOrganizationListResponsePersonalProfileJSON contains the JSON
// metadata for the struct [ConnectLinkedinOrganizationListResponsePersonalProfile]
type connectLinkedinOrganizationListResponsePersonalProfileJSON struct {
	Name        apijson.Field
	Urn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectLinkedinOrganizationListResponsePersonalProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectLinkedinOrganizationListResponsePersonalProfileJSON) RawJSON() string {
	return r.raw
}

type ConnectLinkedinOrganizationSelectResponse struct {
	Account ConnectLinkedinOrganizationSelectResponseAccount `json:"account" api:"required"`
	JSON    connectLinkedinOrganizationSelectResponseJSON    `json:"-"`
}

// connectLinkedinOrganizationSelectResponseJSON contains the JSON metadata for the
// struct [ConnectLinkedinOrganizationSelectResponse]
type connectLinkedinOrganizationSelectResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectLinkedinOrganizationSelectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectLinkedinOrganizationSelectResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectLinkedinOrganizationSelectResponseAccount struct {
	// Account ID
	ID          string    `json:"id" api:"required"`
	AvatarURL   string    `json:"avatar_url" api:"required,nullable"`
	ConnectedAt time.Time `json:"connected_at" api:"required" format:"date-time"`
	DisplayName string    `json:"display_name" api:"required,nullable"`
	// Account group
	Group             ConnectLinkedinOrganizationSelectResponseAccountGroup    `json:"group" api:"required,nullable"`
	Metadata          map[string]interface{}                                   `json:"metadata" api:"required,nullable"`
	Platform          ConnectLinkedinOrganizationSelectResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                   `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                                `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                   `json:"username" api:"required,nullable"`
	JSON              connectLinkedinOrganizationSelectResponseAccountJSON     `json:"-"`
}

// connectLinkedinOrganizationSelectResponseAccountJSON contains the JSON metadata
// for the struct [ConnectLinkedinOrganizationSelectResponseAccount]
type connectLinkedinOrganizationSelectResponseAccountJSON struct {
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

func (r *ConnectLinkedinOrganizationSelectResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectLinkedinOrganizationSelectResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Account group
type ConnectLinkedinOrganizationSelectResponseAccountGroup struct {
	ID   string                                                    `json:"id" api:"required"`
	Name string                                                    `json:"name" api:"required"`
	JSON connectLinkedinOrganizationSelectResponseAccountGroupJSON `json:"-"`
}

// connectLinkedinOrganizationSelectResponseAccountGroupJSON contains the JSON
// metadata for the struct [ConnectLinkedinOrganizationSelectResponseAccountGroup]
type connectLinkedinOrganizationSelectResponseAccountGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectLinkedinOrganizationSelectResponseAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectLinkedinOrganizationSelectResponseAccountGroupJSON) RawJSON() string {
	return r.raw
}

type ConnectLinkedinOrganizationSelectResponseAccountPlatform string

const (
	ConnectLinkedinOrganizationSelectResponseAccountPlatformTwitter        ConnectLinkedinOrganizationSelectResponseAccountPlatform = "twitter"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformInstagram      ConnectLinkedinOrganizationSelectResponseAccountPlatform = "instagram"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformFacebook       ConnectLinkedinOrganizationSelectResponseAccountPlatform = "facebook"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformLinkedin       ConnectLinkedinOrganizationSelectResponseAccountPlatform = "linkedin"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformTiktok         ConnectLinkedinOrganizationSelectResponseAccountPlatform = "tiktok"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformYoutube        ConnectLinkedinOrganizationSelectResponseAccountPlatform = "youtube"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformPinterest      ConnectLinkedinOrganizationSelectResponseAccountPlatform = "pinterest"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformReddit         ConnectLinkedinOrganizationSelectResponseAccountPlatform = "reddit"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformBluesky        ConnectLinkedinOrganizationSelectResponseAccountPlatform = "bluesky"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformThreads        ConnectLinkedinOrganizationSelectResponseAccountPlatform = "threads"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformTelegram       ConnectLinkedinOrganizationSelectResponseAccountPlatform = "telegram"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformSnapchat       ConnectLinkedinOrganizationSelectResponseAccountPlatform = "snapchat"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformGooglebusiness ConnectLinkedinOrganizationSelectResponseAccountPlatform = "googlebusiness"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformWhatsapp       ConnectLinkedinOrganizationSelectResponseAccountPlatform = "whatsapp"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformMastodon       ConnectLinkedinOrganizationSelectResponseAccountPlatform = "mastodon"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformDiscord        ConnectLinkedinOrganizationSelectResponseAccountPlatform = "discord"
	ConnectLinkedinOrganizationSelectResponseAccountPlatformSMS            ConnectLinkedinOrganizationSelectResponseAccountPlatform = "sms"
)

func (r ConnectLinkedinOrganizationSelectResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectLinkedinOrganizationSelectResponseAccountPlatformTwitter, ConnectLinkedinOrganizationSelectResponseAccountPlatformInstagram, ConnectLinkedinOrganizationSelectResponseAccountPlatformFacebook, ConnectLinkedinOrganizationSelectResponseAccountPlatformLinkedin, ConnectLinkedinOrganizationSelectResponseAccountPlatformTiktok, ConnectLinkedinOrganizationSelectResponseAccountPlatformYoutube, ConnectLinkedinOrganizationSelectResponseAccountPlatformPinterest, ConnectLinkedinOrganizationSelectResponseAccountPlatformReddit, ConnectLinkedinOrganizationSelectResponseAccountPlatformBluesky, ConnectLinkedinOrganizationSelectResponseAccountPlatformThreads, ConnectLinkedinOrganizationSelectResponseAccountPlatformTelegram, ConnectLinkedinOrganizationSelectResponseAccountPlatformSnapchat, ConnectLinkedinOrganizationSelectResponseAccountPlatformGooglebusiness, ConnectLinkedinOrganizationSelectResponseAccountPlatformWhatsapp, ConnectLinkedinOrganizationSelectResponseAccountPlatformMastodon, ConnectLinkedinOrganizationSelectResponseAccountPlatformDiscord, ConnectLinkedinOrganizationSelectResponseAccountPlatformSMS:
		return true
	}
	return false
}

type ConnectLinkedinOrganizationSelectParams struct {
	// Whether to connect as a personal profile or organization
	AccountType param.Field[ConnectLinkedinOrganizationSelectParamsAccountType] `json:"account_type" api:"required"`
	// Token from pending data or OAuth flow
	ConnectToken param.Field[string] `json:"connect_token" api:"required"`
	// LinkedIn organization URN (required if account_type is organization)
	OrganizationUrn param.Field[string] `json:"organization_urn"`
}

func (r ConnectLinkedinOrganizationSelectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to connect as a personal profile or organization
type ConnectLinkedinOrganizationSelectParamsAccountType string

const (
	ConnectLinkedinOrganizationSelectParamsAccountTypePersonal     ConnectLinkedinOrganizationSelectParamsAccountType = "personal"
	ConnectLinkedinOrganizationSelectParamsAccountTypeOrganization ConnectLinkedinOrganizationSelectParamsAccountType = "organization"
)

func (r ConnectLinkedinOrganizationSelectParamsAccountType) IsKnown() bool {
	switch r {
	case ConnectLinkedinOrganizationSelectParamsAccountTypePersonal, ConnectLinkedinOrganizationSelectParamsAccountTypeOrganization:
		return true
	}
	return false
}
