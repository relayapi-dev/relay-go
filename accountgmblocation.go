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

// AccountGmbLocationService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGmbLocationService] method instead.
type AccountGmbLocationService struct {
	Options []option.RequestOption
}

// NewAccountGmbLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGmbLocationService(opts ...option.RequestOption) (r *AccountGmbLocationService) {
	r = &AccountGmbLocationService{}
	r.Options = opts
	return
}

// Fetch Google My Business locations
func (r *AccountGmbLocationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGmbLocationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/gmb-locations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set default GMB location
func (r *AccountGmbLocationService) SetDefault(ctx context.Context, id string, body AccountGmbLocationSetDefaultParams, opts ...option.RequestOption) (res *AccountGmbLocationSetDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/gmb-locations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AccountGmbLocationGetResponse struct {
	Data []AccountGmbLocationGetResponseData `json:"data" api:"required"`
	JSON accountGmbLocationGetResponseJSON   `json:"-"`
}

// accountGmbLocationGetResponseJSON contains the JSON metadata for the struct
// [AccountGmbLocationGetResponse]
type accountGmbLocationGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGmbLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGmbLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGmbLocationGetResponseData struct {
	ID      string                                `json:"id" api:"required"`
	Address string                                `json:"address" api:"required,nullable"`
	Name    string                                `json:"name" api:"required"`
	JSON    accountGmbLocationGetResponseDataJSON `json:"-"`
}

// accountGmbLocationGetResponseDataJSON contains the JSON metadata for the struct
// [AccountGmbLocationGetResponseData]
type accountGmbLocationGetResponseDataJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGmbLocationGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGmbLocationGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountGmbLocationSetDefaultResponse struct {
	// Account ID
	ID                string                                       `json:"id" api:"required"`
	AvatarURL         string                                       `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                    `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                       `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                       `json:"metadata" api:"required,nullable"`
	Platform          AccountGmbLocationSetDefaultResponsePlatform `json:"platform" api:"required"`
	PlatformAccountID string                                       `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                    `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                       `json:"username" api:"required,nullable"`
	JSON              accountGmbLocationSetDefaultResponseJSON     `json:"-"`
}

// accountGmbLocationSetDefaultResponseJSON contains the JSON metadata for the
// struct [AccountGmbLocationSetDefaultResponse]
type accountGmbLocationSetDefaultResponseJSON struct {
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

func (r *AccountGmbLocationSetDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGmbLocationSetDefaultResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGmbLocationSetDefaultResponsePlatform string

const (
	AccountGmbLocationSetDefaultResponsePlatformTwitter        AccountGmbLocationSetDefaultResponsePlatform = "twitter"
	AccountGmbLocationSetDefaultResponsePlatformInstagram      AccountGmbLocationSetDefaultResponsePlatform = "instagram"
	AccountGmbLocationSetDefaultResponsePlatformFacebook       AccountGmbLocationSetDefaultResponsePlatform = "facebook"
	AccountGmbLocationSetDefaultResponsePlatformLinkedin       AccountGmbLocationSetDefaultResponsePlatform = "linkedin"
	AccountGmbLocationSetDefaultResponsePlatformTiktok         AccountGmbLocationSetDefaultResponsePlatform = "tiktok"
	AccountGmbLocationSetDefaultResponsePlatformYoutube        AccountGmbLocationSetDefaultResponsePlatform = "youtube"
	AccountGmbLocationSetDefaultResponsePlatformPinterest      AccountGmbLocationSetDefaultResponsePlatform = "pinterest"
	AccountGmbLocationSetDefaultResponsePlatformReddit         AccountGmbLocationSetDefaultResponsePlatform = "reddit"
	AccountGmbLocationSetDefaultResponsePlatformBluesky        AccountGmbLocationSetDefaultResponsePlatform = "bluesky"
	AccountGmbLocationSetDefaultResponsePlatformThreads        AccountGmbLocationSetDefaultResponsePlatform = "threads"
	AccountGmbLocationSetDefaultResponsePlatformTelegram       AccountGmbLocationSetDefaultResponsePlatform = "telegram"
	AccountGmbLocationSetDefaultResponsePlatformSnapchat       AccountGmbLocationSetDefaultResponsePlatform = "snapchat"
	AccountGmbLocationSetDefaultResponsePlatformGooglebusiness AccountGmbLocationSetDefaultResponsePlatform = "googlebusiness"
	AccountGmbLocationSetDefaultResponsePlatformWhatsapp       AccountGmbLocationSetDefaultResponsePlatform = "whatsapp"
	AccountGmbLocationSetDefaultResponsePlatformMastodon       AccountGmbLocationSetDefaultResponsePlatform = "mastodon"
	AccountGmbLocationSetDefaultResponsePlatformDiscord        AccountGmbLocationSetDefaultResponsePlatform = "discord"
	AccountGmbLocationSetDefaultResponsePlatformSMS            AccountGmbLocationSetDefaultResponsePlatform = "sms"
)

func (r AccountGmbLocationSetDefaultResponsePlatform) IsKnown() bool {
	switch r {
	case AccountGmbLocationSetDefaultResponsePlatformTwitter, AccountGmbLocationSetDefaultResponsePlatformInstagram, AccountGmbLocationSetDefaultResponsePlatformFacebook, AccountGmbLocationSetDefaultResponsePlatformLinkedin, AccountGmbLocationSetDefaultResponsePlatformTiktok, AccountGmbLocationSetDefaultResponsePlatformYoutube, AccountGmbLocationSetDefaultResponsePlatformPinterest, AccountGmbLocationSetDefaultResponsePlatformReddit, AccountGmbLocationSetDefaultResponsePlatformBluesky, AccountGmbLocationSetDefaultResponsePlatformThreads, AccountGmbLocationSetDefaultResponsePlatformTelegram, AccountGmbLocationSetDefaultResponsePlatformSnapchat, AccountGmbLocationSetDefaultResponsePlatformGooglebusiness, AccountGmbLocationSetDefaultResponsePlatformWhatsapp, AccountGmbLocationSetDefaultResponsePlatformMastodon, AccountGmbLocationSetDefaultResponsePlatformDiscord, AccountGmbLocationSetDefaultResponsePlatformSMS:
		return true
	}
	return false
}

type AccountGmbLocationSetDefaultParams struct {
	// Google My Business location ID to set as default
	LocationID param.Field[string] `json:"location_id" api:"required"`
}

func (r AccountGmbLocationSetDefaultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
