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

// AccountLinkedinOrganizationService contains methods and other services that help
// with interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLinkedinOrganizationService] method instead.
type AccountLinkedinOrganizationService struct {
	Options []option.RequestOption
}

// NewAccountLinkedinOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLinkedinOrganizationService(opts ...option.RequestOption) (r *AccountLinkedinOrganizationService) {
	r = &AccountLinkedinOrganizationService{}
	r.Options = opts
	return
}

// Fetch LinkedIn organizations for an account
func (r *AccountLinkedinOrganizationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountLinkedinOrganizationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/linkedin-organizations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Switch LinkedIn account type
func (r *AccountLinkedinOrganizationService) SwitchType(ctx context.Context, id string, body AccountLinkedinOrganizationSwitchTypeParams, opts ...option.RequestOption) (res *AccountLinkedinOrganizationSwitchTypeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/linkedin-organizations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AccountLinkedinOrganizationGetResponse struct {
	Data []AccountLinkedinOrganizationGetResponseData `json:"data" api:"required"`
	JSON accountLinkedinOrganizationGetResponseJSON   `json:"-"`
}

// accountLinkedinOrganizationGetResponseJSON contains the JSON metadata for the
// struct [AccountLinkedinOrganizationGetResponse]
type accountLinkedinOrganizationGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLinkedinOrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkedinOrganizationGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountLinkedinOrganizationGetResponseData struct {
	ID         string                                         `json:"id" api:"required"`
	Name       string                                         `json:"name" api:"required"`
	VanityName string                                         `json:"vanity_name" api:"required,nullable"`
	JSON       accountLinkedinOrganizationGetResponseDataJSON `json:"-"`
}

// accountLinkedinOrganizationGetResponseDataJSON contains the JSON metadata for
// the struct [AccountLinkedinOrganizationGetResponseData]
type accountLinkedinOrganizationGetResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	VanityName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLinkedinOrganizationGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkedinOrganizationGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountLinkedinOrganizationSwitchTypeResponse struct {
	// Account ID
	ID                string                                                `json:"id" api:"required"`
	AvatarURL         string                                                `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                             `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                                `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                                `json:"metadata" api:"required,nullable"`
	Platform          AccountLinkedinOrganizationSwitchTypeResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                                                `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                             `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                                `json:"username" api:"required,nullable"`
	JSON              accountLinkedinOrganizationSwitchTypeResponseJSON     `json:"-"`
}

// accountLinkedinOrganizationSwitchTypeResponseJSON contains the JSON metadata for
// the struct [AccountLinkedinOrganizationSwitchTypeResponse]
type accountLinkedinOrganizationSwitchTypeResponseJSON struct {
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

func (r *AccountLinkedinOrganizationSwitchTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkedinOrganizationSwitchTypeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountLinkedinOrganizationSwitchTypeResponsePlatform string

const (
	AccountLinkedinOrganizationSwitchTypeResponsePlatformTwitter        AccountLinkedinOrganizationSwitchTypeResponsePlatform = "twitter"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformInstagram      AccountLinkedinOrganizationSwitchTypeResponsePlatform = "instagram"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformFacebook       AccountLinkedinOrganizationSwitchTypeResponsePlatform = "facebook"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformLinkedin       AccountLinkedinOrganizationSwitchTypeResponsePlatform = "linkedin"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformTiktok         AccountLinkedinOrganizationSwitchTypeResponsePlatform = "tiktok"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformYoutube        AccountLinkedinOrganizationSwitchTypeResponsePlatform = "youtube"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformPinterest      AccountLinkedinOrganizationSwitchTypeResponsePlatform = "pinterest"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformReddit         AccountLinkedinOrganizationSwitchTypeResponsePlatform = "reddit"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformBluesky        AccountLinkedinOrganizationSwitchTypeResponsePlatform = "bluesky"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformThreads        AccountLinkedinOrganizationSwitchTypeResponsePlatform = "threads"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformTelegram       AccountLinkedinOrganizationSwitchTypeResponsePlatform = "telegram"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformSnapchat       AccountLinkedinOrganizationSwitchTypeResponsePlatform = "snapchat"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformGooglebusiness AccountLinkedinOrganizationSwitchTypeResponsePlatform = "googlebusiness"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformWhatsapp       AccountLinkedinOrganizationSwitchTypeResponsePlatform = "whatsapp"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformMastodon       AccountLinkedinOrganizationSwitchTypeResponsePlatform = "mastodon"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformDiscord        AccountLinkedinOrganizationSwitchTypeResponsePlatform = "discord"
	AccountLinkedinOrganizationSwitchTypeResponsePlatformSMS            AccountLinkedinOrganizationSwitchTypeResponsePlatform = "sms"
)

func (r AccountLinkedinOrganizationSwitchTypeResponsePlatform) IsKnown() bool {
	switch r {
	case AccountLinkedinOrganizationSwitchTypeResponsePlatformTwitter, AccountLinkedinOrganizationSwitchTypeResponsePlatformInstagram, AccountLinkedinOrganizationSwitchTypeResponsePlatformFacebook, AccountLinkedinOrganizationSwitchTypeResponsePlatformLinkedin, AccountLinkedinOrganizationSwitchTypeResponsePlatformTiktok, AccountLinkedinOrganizationSwitchTypeResponsePlatformYoutube, AccountLinkedinOrganizationSwitchTypeResponsePlatformPinterest, AccountLinkedinOrganizationSwitchTypeResponsePlatformReddit, AccountLinkedinOrganizationSwitchTypeResponsePlatformBluesky, AccountLinkedinOrganizationSwitchTypeResponsePlatformThreads, AccountLinkedinOrganizationSwitchTypeResponsePlatformTelegram, AccountLinkedinOrganizationSwitchTypeResponsePlatformSnapchat, AccountLinkedinOrganizationSwitchTypeResponsePlatformGooglebusiness, AccountLinkedinOrganizationSwitchTypeResponsePlatformWhatsapp, AccountLinkedinOrganizationSwitchTypeResponsePlatformMastodon, AccountLinkedinOrganizationSwitchTypeResponsePlatformDiscord, AccountLinkedinOrganizationSwitchTypeResponsePlatformSMS:
		return true
	}
	return false
}

type AccountLinkedinOrganizationSwitchTypeParams struct {
	// Account type to switch to
	AccountType param.Field[AccountLinkedinOrganizationSwitchTypeParamsAccountType] `json:"account_type" api:"required"`
	// LinkedIn organization ID
	OrganizationID param.Field[string] `json:"organization_id" api:"required"`
}

func (r AccountLinkedinOrganizationSwitchTypeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account type to switch to
type AccountLinkedinOrganizationSwitchTypeParamsAccountType string

const (
	AccountLinkedinOrganizationSwitchTypeParamsAccountTypePersonal     AccountLinkedinOrganizationSwitchTypeParamsAccountType = "personal"
	AccountLinkedinOrganizationSwitchTypeParamsAccountTypeOrganization AccountLinkedinOrganizationSwitchTypeParamsAccountType = "organization"
)

func (r AccountLinkedinOrganizationSwitchTypeParamsAccountType) IsKnown() bool {
	switch r {
	case AccountLinkedinOrganizationSwitchTypeParamsAccountTypePersonal, AccountLinkedinOrganizationSwitchTypeParamsAccountTypeOrganization:
		return true
	}
	return false
}
