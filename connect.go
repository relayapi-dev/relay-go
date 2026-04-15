// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"fmt"
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

// ConnectService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectService] method instead.
type ConnectService struct {
	Options        []option.RequestOption
	Telegram       *ConnectTelegramService
	Whatsapp       *ConnectWhatsappService
	Facebook       *ConnectFacebookService
	Linkedin       *ConnectLinkedinService
	Pinterest      *ConnectPinterestService
	Googlebusiness *ConnectGooglebusinessService
	Snapchat       *ConnectSnapchatService
}

// NewConnectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConnectService(opts ...option.RequestOption) (r *ConnectService) {
	r = &ConnectService{}
	r.Options = opts
	r.Telegram = NewConnectTelegramService(opts...)
	r.Whatsapp = NewConnectWhatsappService(opts...)
	r.Facebook = NewConnectFacebookService(opts...)
	r.Linkedin = NewConnectLinkedinService(opts...)
	r.Pinterest = NewConnectPinterestService(opts...)
	r.Googlebusiness = NewConnectGooglebusinessService(opts...)
	r.Snapchat = NewConnectSnapchatService(opts...)
	return
}

// Exchange OAuth code for tokens and save the account.
func (r *ConnectService) CompleteOAuthCallback(ctx context.Context, platform ConnectCompleteOAuthCallbackParamsPlatform, body ConnectCompleteOAuthCallbackParams, opts ...option.RequestOption) (res *ConnectCompleteOAuthCallbackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/connect/%v", platform)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Connect Bluesky via app password
func (r *ConnectService) NewBlueskyConnection(ctx context.Context, body ConnectNewBlueskyConnectionParams, opts ...option.RequestOption) (res *ConnectNewBlueskyConnectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/bluesky"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// One-time use, expires after 10 minutes. For headless OAuth flows.
func (r *ConnectService) FetchPendingData(ctx context.Context, query ConnectFetchPendingDataParams, opts ...option.RequestOption) (res *ConnectFetchPendingDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/pending-data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns an auth_url to redirect the user for OAuth authorization.
func (r *ConnectService) StartOAuthFlow(ctx context.Context, platform ConnectStartOAuthFlowParamsPlatform, query ConnectStartOAuthFlowParams, opts ...option.RequestOption) (res *ConnectStartOAuthFlowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/connect/%v", platform)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConnectCompleteOAuthCallbackResponse struct {
	Account ConnectCompleteOAuthCallbackResponseAccount `json:"account" api:"required"`
	JSON    connectCompleteOAuthCallbackResponseJSON    `json:"-"`
}

// connectCompleteOAuthCallbackResponseJSON contains the JSON metadata for the
// struct [ConnectCompleteOAuthCallbackResponse]
type connectCompleteOAuthCallbackResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectCompleteOAuthCallbackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectCompleteOAuthCallbackResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectCompleteOAuthCallbackResponseAccount struct {
	// Account ID
	ID                string                                              `json:"id" api:"required"`
	AvatarURL         string                                              `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                           `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                              `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                              `json:"metadata" api:"required,nullable"`
	Platform          ConnectCompleteOAuthCallbackResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                              `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                           `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                              `json:"username" api:"required,nullable"`
	// Account workspace
	Workspace ConnectCompleteOAuthCallbackResponseAccountWorkspace `json:"workspace" api:"required,nullable"`
	JSON      connectCompleteOAuthCallbackResponseAccountJSON      `json:"-"`
}

// connectCompleteOAuthCallbackResponseAccountJSON contains the JSON metadata for
// the struct [ConnectCompleteOAuthCallbackResponseAccount]
type connectCompleteOAuthCallbackResponseAccountJSON struct {
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

func (r *ConnectCompleteOAuthCallbackResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectCompleteOAuthCallbackResponseAccountJSON) RawJSON() string {
	return r.raw
}

type ConnectCompleteOAuthCallbackResponseAccountPlatform string

const (
	ConnectCompleteOAuthCallbackResponseAccountPlatformTwitter        ConnectCompleteOAuthCallbackResponseAccountPlatform = "twitter"
	ConnectCompleteOAuthCallbackResponseAccountPlatformInstagram      ConnectCompleteOAuthCallbackResponseAccountPlatform = "instagram"
	ConnectCompleteOAuthCallbackResponseAccountPlatformFacebook       ConnectCompleteOAuthCallbackResponseAccountPlatform = "facebook"
	ConnectCompleteOAuthCallbackResponseAccountPlatformLinkedin       ConnectCompleteOAuthCallbackResponseAccountPlatform = "linkedin"
	ConnectCompleteOAuthCallbackResponseAccountPlatformTiktok         ConnectCompleteOAuthCallbackResponseAccountPlatform = "tiktok"
	ConnectCompleteOAuthCallbackResponseAccountPlatformYoutube        ConnectCompleteOAuthCallbackResponseAccountPlatform = "youtube"
	ConnectCompleteOAuthCallbackResponseAccountPlatformPinterest      ConnectCompleteOAuthCallbackResponseAccountPlatform = "pinterest"
	ConnectCompleteOAuthCallbackResponseAccountPlatformReddit         ConnectCompleteOAuthCallbackResponseAccountPlatform = "reddit"
	ConnectCompleteOAuthCallbackResponseAccountPlatformBluesky        ConnectCompleteOAuthCallbackResponseAccountPlatform = "bluesky"
	ConnectCompleteOAuthCallbackResponseAccountPlatformThreads        ConnectCompleteOAuthCallbackResponseAccountPlatform = "threads"
	ConnectCompleteOAuthCallbackResponseAccountPlatformTelegram       ConnectCompleteOAuthCallbackResponseAccountPlatform = "telegram"
	ConnectCompleteOAuthCallbackResponseAccountPlatformSnapchat       ConnectCompleteOAuthCallbackResponseAccountPlatform = "snapchat"
	ConnectCompleteOAuthCallbackResponseAccountPlatformGooglebusiness ConnectCompleteOAuthCallbackResponseAccountPlatform = "googlebusiness"
	ConnectCompleteOAuthCallbackResponseAccountPlatformWhatsapp       ConnectCompleteOAuthCallbackResponseAccountPlatform = "whatsapp"
	ConnectCompleteOAuthCallbackResponseAccountPlatformMastodon       ConnectCompleteOAuthCallbackResponseAccountPlatform = "mastodon"
	ConnectCompleteOAuthCallbackResponseAccountPlatformDiscord        ConnectCompleteOAuthCallbackResponseAccountPlatform = "discord"
	ConnectCompleteOAuthCallbackResponseAccountPlatformSMS            ConnectCompleteOAuthCallbackResponseAccountPlatform = "sms"
	ConnectCompleteOAuthCallbackResponseAccountPlatformBeehiiv        ConnectCompleteOAuthCallbackResponseAccountPlatform = "beehiiv"
	ConnectCompleteOAuthCallbackResponseAccountPlatformConvertkit     ConnectCompleteOAuthCallbackResponseAccountPlatform = "convertkit"
	ConnectCompleteOAuthCallbackResponseAccountPlatformMailchimp      ConnectCompleteOAuthCallbackResponseAccountPlatform = "mailchimp"
	ConnectCompleteOAuthCallbackResponseAccountPlatformListmonk       ConnectCompleteOAuthCallbackResponseAccountPlatform = "listmonk"
)

func (r ConnectCompleteOAuthCallbackResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectCompleteOAuthCallbackResponseAccountPlatformTwitter, ConnectCompleteOAuthCallbackResponseAccountPlatformInstagram, ConnectCompleteOAuthCallbackResponseAccountPlatformFacebook, ConnectCompleteOAuthCallbackResponseAccountPlatformLinkedin, ConnectCompleteOAuthCallbackResponseAccountPlatformTiktok, ConnectCompleteOAuthCallbackResponseAccountPlatformYoutube, ConnectCompleteOAuthCallbackResponseAccountPlatformPinterest, ConnectCompleteOAuthCallbackResponseAccountPlatformReddit, ConnectCompleteOAuthCallbackResponseAccountPlatformBluesky, ConnectCompleteOAuthCallbackResponseAccountPlatformThreads, ConnectCompleteOAuthCallbackResponseAccountPlatformTelegram, ConnectCompleteOAuthCallbackResponseAccountPlatformSnapchat, ConnectCompleteOAuthCallbackResponseAccountPlatformGooglebusiness, ConnectCompleteOAuthCallbackResponseAccountPlatformWhatsapp, ConnectCompleteOAuthCallbackResponseAccountPlatformMastodon, ConnectCompleteOAuthCallbackResponseAccountPlatformDiscord, ConnectCompleteOAuthCallbackResponseAccountPlatformSMS, ConnectCompleteOAuthCallbackResponseAccountPlatformBeehiiv, ConnectCompleteOAuthCallbackResponseAccountPlatformConvertkit, ConnectCompleteOAuthCallbackResponseAccountPlatformMailchimp, ConnectCompleteOAuthCallbackResponseAccountPlatformListmonk:
		return true
	}
	return false
}

// Account workspace
type ConnectCompleteOAuthCallbackResponseAccountWorkspace struct {
	ID   string                                                   `json:"id" api:"required"`
	Name string                                                   `json:"name" api:"required"`
	JSON connectCompleteOAuthCallbackResponseAccountWorkspaceJSON `json:"-"`
}

// connectCompleteOAuthCallbackResponseAccountWorkspaceJSON contains the JSON
// metadata for the struct [ConnectCompleteOAuthCallbackResponseAccountWorkspace]
type connectCompleteOAuthCallbackResponseAccountWorkspaceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectCompleteOAuthCallbackResponseAccountWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectCompleteOAuthCallbackResponseAccountWorkspaceJSON) RawJSON() string {
	return r.raw
}

type ConnectNewBlueskyConnectionResponse struct {
	Account ConnectNewBlueskyConnectionResponseAccount `json:"account" api:"required"`
	JSON    connectNewBlueskyConnectionResponseJSON    `json:"-"`
}

// connectNewBlueskyConnectionResponseJSON contains the JSON metadata for the
// struct [ConnectNewBlueskyConnectionResponse]
type connectNewBlueskyConnectionResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectNewBlueskyConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectNewBlueskyConnectionResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectNewBlueskyConnectionResponseAccount struct {
	// Account ID
	ID                string                                             `json:"id" api:"required"`
	AvatarURL         string                                             `json:"avatar_url" api:"required,nullable"`
	ConnectedAt       time.Time                                          `json:"connected_at" api:"required" format:"date-time"`
	DisplayName       string                                             `json:"display_name" api:"required,nullable"`
	Metadata          map[string]interface{}                             `json:"metadata" api:"required,nullable"`
	Platform          ConnectNewBlueskyConnectionResponseAccountPlatform `json:"platform" api:"required"`
	PlatformAccountID string                                             `json:"platform_account_id" api:"required"`
	UpdatedAt         time.Time                                          `json:"updated_at" api:"required" format:"date-time"`
	Username          string                                             `json:"username" api:"required,nullable"`
	// Account workspace
	Workspace ConnectNewBlueskyConnectionResponseAccountWorkspace `json:"workspace" api:"required,nullable"`
	JSON      connectNewBlueskyConnectionResponseAccountJSON      `json:"-"`
}

// connectNewBlueskyConnectionResponseAccountJSON contains the JSON metadata for
// the struct [ConnectNewBlueskyConnectionResponseAccount]
type connectNewBlueskyConnectionResponseAccountJSON struct {
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

func (r *ConnectNewBlueskyConnectionResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectNewBlueskyConnectionResponseAccountJSON) RawJSON() string {
	return r.raw
}

type ConnectNewBlueskyConnectionResponseAccountPlatform string

const (
	ConnectNewBlueskyConnectionResponseAccountPlatformTwitter        ConnectNewBlueskyConnectionResponseAccountPlatform = "twitter"
	ConnectNewBlueskyConnectionResponseAccountPlatformInstagram      ConnectNewBlueskyConnectionResponseAccountPlatform = "instagram"
	ConnectNewBlueskyConnectionResponseAccountPlatformFacebook       ConnectNewBlueskyConnectionResponseAccountPlatform = "facebook"
	ConnectNewBlueskyConnectionResponseAccountPlatformLinkedin       ConnectNewBlueskyConnectionResponseAccountPlatform = "linkedin"
	ConnectNewBlueskyConnectionResponseAccountPlatformTiktok         ConnectNewBlueskyConnectionResponseAccountPlatform = "tiktok"
	ConnectNewBlueskyConnectionResponseAccountPlatformYoutube        ConnectNewBlueskyConnectionResponseAccountPlatform = "youtube"
	ConnectNewBlueskyConnectionResponseAccountPlatformPinterest      ConnectNewBlueskyConnectionResponseAccountPlatform = "pinterest"
	ConnectNewBlueskyConnectionResponseAccountPlatformReddit         ConnectNewBlueskyConnectionResponseAccountPlatform = "reddit"
	ConnectNewBlueskyConnectionResponseAccountPlatformBluesky        ConnectNewBlueskyConnectionResponseAccountPlatform = "bluesky"
	ConnectNewBlueskyConnectionResponseAccountPlatformThreads        ConnectNewBlueskyConnectionResponseAccountPlatform = "threads"
	ConnectNewBlueskyConnectionResponseAccountPlatformTelegram       ConnectNewBlueskyConnectionResponseAccountPlatform = "telegram"
	ConnectNewBlueskyConnectionResponseAccountPlatformSnapchat       ConnectNewBlueskyConnectionResponseAccountPlatform = "snapchat"
	ConnectNewBlueskyConnectionResponseAccountPlatformGooglebusiness ConnectNewBlueskyConnectionResponseAccountPlatform = "googlebusiness"
	ConnectNewBlueskyConnectionResponseAccountPlatformWhatsapp       ConnectNewBlueskyConnectionResponseAccountPlatform = "whatsapp"
	ConnectNewBlueskyConnectionResponseAccountPlatformMastodon       ConnectNewBlueskyConnectionResponseAccountPlatform = "mastodon"
	ConnectNewBlueskyConnectionResponseAccountPlatformDiscord        ConnectNewBlueskyConnectionResponseAccountPlatform = "discord"
	ConnectNewBlueskyConnectionResponseAccountPlatformSMS            ConnectNewBlueskyConnectionResponseAccountPlatform = "sms"
	ConnectNewBlueskyConnectionResponseAccountPlatformBeehiiv        ConnectNewBlueskyConnectionResponseAccountPlatform = "beehiiv"
	ConnectNewBlueskyConnectionResponseAccountPlatformConvertkit     ConnectNewBlueskyConnectionResponseAccountPlatform = "convertkit"
	ConnectNewBlueskyConnectionResponseAccountPlatformMailchimp      ConnectNewBlueskyConnectionResponseAccountPlatform = "mailchimp"
	ConnectNewBlueskyConnectionResponseAccountPlatformListmonk       ConnectNewBlueskyConnectionResponseAccountPlatform = "listmonk"
)

func (r ConnectNewBlueskyConnectionResponseAccountPlatform) IsKnown() bool {
	switch r {
	case ConnectNewBlueskyConnectionResponseAccountPlatformTwitter, ConnectNewBlueskyConnectionResponseAccountPlatformInstagram, ConnectNewBlueskyConnectionResponseAccountPlatformFacebook, ConnectNewBlueskyConnectionResponseAccountPlatformLinkedin, ConnectNewBlueskyConnectionResponseAccountPlatformTiktok, ConnectNewBlueskyConnectionResponseAccountPlatformYoutube, ConnectNewBlueskyConnectionResponseAccountPlatformPinterest, ConnectNewBlueskyConnectionResponseAccountPlatformReddit, ConnectNewBlueskyConnectionResponseAccountPlatformBluesky, ConnectNewBlueskyConnectionResponseAccountPlatformThreads, ConnectNewBlueskyConnectionResponseAccountPlatformTelegram, ConnectNewBlueskyConnectionResponseAccountPlatformSnapchat, ConnectNewBlueskyConnectionResponseAccountPlatformGooglebusiness, ConnectNewBlueskyConnectionResponseAccountPlatformWhatsapp, ConnectNewBlueskyConnectionResponseAccountPlatformMastodon, ConnectNewBlueskyConnectionResponseAccountPlatformDiscord, ConnectNewBlueskyConnectionResponseAccountPlatformSMS, ConnectNewBlueskyConnectionResponseAccountPlatformBeehiiv, ConnectNewBlueskyConnectionResponseAccountPlatformConvertkit, ConnectNewBlueskyConnectionResponseAccountPlatformMailchimp, ConnectNewBlueskyConnectionResponseAccountPlatformListmonk:
		return true
	}
	return false
}

// Account workspace
type ConnectNewBlueskyConnectionResponseAccountWorkspace struct {
	ID   string                                                  `json:"id" api:"required"`
	Name string                                                  `json:"name" api:"required"`
	JSON connectNewBlueskyConnectionResponseAccountWorkspaceJSON `json:"-"`
}

// connectNewBlueskyConnectionResponseAccountWorkspaceJSON contains the JSON
// metadata for the struct [ConnectNewBlueskyConnectionResponseAccountWorkspace]
type connectNewBlueskyConnectionResponseAccountWorkspaceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectNewBlueskyConnectionResponseAccountWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectNewBlueskyConnectionResponseAccountWorkspaceJSON) RawJSON() string {
	return r.raw
}

type ConnectFetchPendingDataResponse struct {
	Platform ConnectFetchPendingDataResponsePlatform `json:"platform" api:"required"`
	// Token to use for secondary selection
	TempToken string `json:"temp_token" api:"required"`
	// Basic user profile from the platform
	UserProfile ConnectFetchPendingDataResponseUserProfile `json:"user_profile" api:"required"`
	// Pinterest boards available
	Boards []map[string]interface{} `json:"boards"`
	// Google Business locations available
	Locations []map[string]interface{} `json:"locations"`
	// LinkedIn organizations available
	Organizations []map[string]interface{} `json:"organizations"`
	// Facebook pages available
	Pages []map[string]interface{} `json:"pages"`
	// Snapchat profiles available
	Profiles []map[string]interface{}            `json:"profiles"`
	JSON     connectFetchPendingDataResponseJSON `json:"-"`
}

// connectFetchPendingDataResponseJSON contains the JSON metadata for the struct
// [ConnectFetchPendingDataResponse]
type connectFetchPendingDataResponseJSON struct {
	Platform      apijson.Field
	TempToken     apijson.Field
	UserProfile   apijson.Field
	Boards        apijson.Field
	Locations     apijson.Field
	Organizations apijson.Field
	Pages         apijson.Field
	Profiles      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectFetchPendingDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFetchPendingDataResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectFetchPendingDataResponsePlatform string

const (
	ConnectFetchPendingDataResponsePlatformTwitter        ConnectFetchPendingDataResponsePlatform = "twitter"
	ConnectFetchPendingDataResponsePlatformInstagram      ConnectFetchPendingDataResponsePlatform = "instagram"
	ConnectFetchPendingDataResponsePlatformFacebook       ConnectFetchPendingDataResponsePlatform = "facebook"
	ConnectFetchPendingDataResponsePlatformLinkedin       ConnectFetchPendingDataResponsePlatform = "linkedin"
	ConnectFetchPendingDataResponsePlatformTiktok         ConnectFetchPendingDataResponsePlatform = "tiktok"
	ConnectFetchPendingDataResponsePlatformYoutube        ConnectFetchPendingDataResponsePlatform = "youtube"
	ConnectFetchPendingDataResponsePlatformPinterest      ConnectFetchPendingDataResponsePlatform = "pinterest"
	ConnectFetchPendingDataResponsePlatformReddit         ConnectFetchPendingDataResponsePlatform = "reddit"
	ConnectFetchPendingDataResponsePlatformBluesky        ConnectFetchPendingDataResponsePlatform = "bluesky"
	ConnectFetchPendingDataResponsePlatformThreads        ConnectFetchPendingDataResponsePlatform = "threads"
	ConnectFetchPendingDataResponsePlatformTelegram       ConnectFetchPendingDataResponsePlatform = "telegram"
	ConnectFetchPendingDataResponsePlatformSnapchat       ConnectFetchPendingDataResponsePlatform = "snapchat"
	ConnectFetchPendingDataResponsePlatformGooglebusiness ConnectFetchPendingDataResponsePlatform = "googlebusiness"
	ConnectFetchPendingDataResponsePlatformWhatsapp       ConnectFetchPendingDataResponsePlatform = "whatsapp"
	ConnectFetchPendingDataResponsePlatformMastodon       ConnectFetchPendingDataResponsePlatform = "mastodon"
	ConnectFetchPendingDataResponsePlatformDiscord        ConnectFetchPendingDataResponsePlatform = "discord"
	ConnectFetchPendingDataResponsePlatformSMS            ConnectFetchPendingDataResponsePlatform = "sms"
	ConnectFetchPendingDataResponsePlatformBeehiiv        ConnectFetchPendingDataResponsePlatform = "beehiiv"
	ConnectFetchPendingDataResponsePlatformConvertkit     ConnectFetchPendingDataResponsePlatform = "convertkit"
	ConnectFetchPendingDataResponsePlatformMailchimp      ConnectFetchPendingDataResponsePlatform = "mailchimp"
	ConnectFetchPendingDataResponsePlatformListmonk       ConnectFetchPendingDataResponsePlatform = "listmonk"
)

func (r ConnectFetchPendingDataResponsePlatform) IsKnown() bool {
	switch r {
	case ConnectFetchPendingDataResponsePlatformTwitter, ConnectFetchPendingDataResponsePlatformInstagram, ConnectFetchPendingDataResponsePlatformFacebook, ConnectFetchPendingDataResponsePlatformLinkedin, ConnectFetchPendingDataResponsePlatformTiktok, ConnectFetchPendingDataResponsePlatformYoutube, ConnectFetchPendingDataResponsePlatformPinterest, ConnectFetchPendingDataResponsePlatformReddit, ConnectFetchPendingDataResponsePlatformBluesky, ConnectFetchPendingDataResponsePlatformThreads, ConnectFetchPendingDataResponsePlatformTelegram, ConnectFetchPendingDataResponsePlatformSnapchat, ConnectFetchPendingDataResponsePlatformGooglebusiness, ConnectFetchPendingDataResponsePlatformWhatsapp, ConnectFetchPendingDataResponsePlatformMastodon, ConnectFetchPendingDataResponsePlatformDiscord, ConnectFetchPendingDataResponsePlatformSMS, ConnectFetchPendingDataResponsePlatformBeehiiv, ConnectFetchPendingDataResponsePlatformConvertkit, ConnectFetchPendingDataResponsePlatformMailchimp, ConnectFetchPendingDataResponsePlatformListmonk:
		return true
	}
	return false
}

// Basic user profile from the platform
type ConnectFetchPendingDataResponseUserProfile struct {
	ID        string                                         `json:"id" api:"required"`
	AvatarURL string                                         `json:"avatar_url" api:"required,nullable"`
	Name      string                                         `json:"name" api:"required,nullable"`
	Username  string                                         `json:"username" api:"required,nullable"`
	JSON      connectFetchPendingDataResponseUserProfileJSON `json:"-"`
}

// connectFetchPendingDataResponseUserProfileJSON contains the JSON metadata for
// the struct [ConnectFetchPendingDataResponseUserProfile]
type connectFetchPendingDataResponseUserProfileJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	Name        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectFetchPendingDataResponseUserProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectFetchPendingDataResponseUserProfileJSON) RawJSON() string {
	return r.raw
}

type ConnectStartOAuthFlowResponse struct {
	// URL to redirect the user for OAuth authorization
	AuthURL string                            `json:"auth_url" api:"required" format:"uri"`
	JSON    connectStartOAuthFlowResponseJSON `json:"-"`
}

// connectStartOAuthFlowResponseJSON contains the JSON metadata for the struct
// [ConnectStartOAuthFlowResponse]
type connectStartOAuthFlowResponseJSON struct {
	AuthURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectStartOAuthFlowResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectStartOAuthFlowResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectCompleteOAuthCallbackParams struct {
	// OAuth authorization code
	Code param.Field[string] `json:"code" api:"required"`
	// Redirect URL used during the OAuth flow (must match)
	RedirectURL param.Field[string] `json:"redirect_url" format:"uri"`
	// OAuth state token for direct KV lookup
	State param.Field[string] `json:"state"`
}

func (r ConnectCompleteOAuthCallbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OAuth platform to complete
type ConnectCompleteOAuthCallbackParamsPlatform string

const (
	ConnectCompleteOAuthCallbackParamsPlatformTwitter        ConnectCompleteOAuthCallbackParamsPlatform = "twitter"
	ConnectCompleteOAuthCallbackParamsPlatformInstagram      ConnectCompleteOAuthCallbackParamsPlatform = "instagram"
	ConnectCompleteOAuthCallbackParamsPlatformFacebook       ConnectCompleteOAuthCallbackParamsPlatform = "facebook"
	ConnectCompleteOAuthCallbackParamsPlatformLinkedin       ConnectCompleteOAuthCallbackParamsPlatform = "linkedin"
	ConnectCompleteOAuthCallbackParamsPlatformTiktok         ConnectCompleteOAuthCallbackParamsPlatform = "tiktok"
	ConnectCompleteOAuthCallbackParamsPlatformYoutube        ConnectCompleteOAuthCallbackParamsPlatform = "youtube"
	ConnectCompleteOAuthCallbackParamsPlatformPinterest      ConnectCompleteOAuthCallbackParamsPlatform = "pinterest"
	ConnectCompleteOAuthCallbackParamsPlatformReddit         ConnectCompleteOAuthCallbackParamsPlatform = "reddit"
	ConnectCompleteOAuthCallbackParamsPlatformThreads        ConnectCompleteOAuthCallbackParamsPlatform = "threads"
	ConnectCompleteOAuthCallbackParamsPlatformSnapchat       ConnectCompleteOAuthCallbackParamsPlatform = "snapchat"
	ConnectCompleteOAuthCallbackParamsPlatformGooglebusiness ConnectCompleteOAuthCallbackParamsPlatform = "googlebusiness"
	ConnectCompleteOAuthCallbackParamsPlatformMastodon       ConnectCompleteOAuthCallbackParamsPlatform = "mastodon"
)

func (r ConnectCompleteOAuthCallbackParamsPlatform) IsKnown() bool {
	switch r {
	case ConnectCompleteOAuthCallbackParamsPlatformTwitter, ConnectCompleteOAuthCallbackParamsPlatformInstagram, ConnectCompleteOAuthCallbackParamsPlatformFacebook, ConnectCompleteOAuthCallbackParamsPlatformLinkedin, ConnectCompleteOAuthCallbackParamsPlatformTiktok, ConnectCompleteOAuthCallbackParamsPlatformYoutube, ConnectCompleteOAuthCallbackParamsPlatformPinterest, ConnectCompleteOAuthCallbackParamsPlatformReddit, ConnectCompleteOAuthCallbackParamsPlatformThreads, ConnectCompleteOAuthCallbackParamsPlatformSnapchat, ConnectCompleteOAuthCallbackParamsPlatformGooglebusiness, ConnectCompleteOAuthCallbackParamsPlatformMastodon:
		return true
	}
	return false
}

type ConnectNewBlueskyConnectionParams struct {
	// Bluesky app password
	AppPassword param.Field[string] `json:"app_password" api:"required"`
	// Bluesky handle (e.g. user.bsky.social)
	Handle param.Field[string] `json:"handle" api:"required"`
}

func (r ConnectNewBlueskyConnectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectFetchPendingDataParams struct {
	// Temporary token from headless OAuth flow
	Token param.Field[string] `query:"token" api:"required"`
}

// URLQuery serializes [ConnectFetchPendingDataParams]'s query parameters as
// `url.Values`.
func (r ConnectFetchPendingDataParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectStartOAuthFlowParams struct {
	// Set to "true" for headless mode (returns data instead of redirecting)
	Headless param.Field[string] `query:"headless"`
	// Auth method variant (e.g. "direct" for Instagram Login instead of Facebook
	// Login)
	Method param.Field[string] `query:"method"`
	// URL to redirect after OAuth completes
	RedirectURL param.Field[string] `query:"redirect_url" format:"uri"`
}

// URLQuery serializes [ConnectStartOAuthFlowParams]'s query parameters as
// `url.Values`.
func (r ConnectStartOAuthFlowParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// OAuth platform to connect
type ConnectStartOAuthFlowParamsPlatform string

const (
	ConnectStartOAuthFlowParamsPlatformTwitter        ConnectStartOAuthFlowParamsPlatform = "twitter"
	ConnectStartOAuthFlowParamsPlatformInstagram      ConnectStartOAuthFlowParamsPlatform = "instagram"
	ConnectStartOAuthFlowParamsPlatformFacebook       ConnectStartOAuthFlowParamsPlatform = "facebook"
	ConnectStartOAuthFlowParamsPlatformLinkedin       ConnectStartOAuthFlowParamsPlatform = "linkedin"
	ConnectStartOAuthFlowParamsPlatformTiktok         ConnectStartOAuthFlowParamsPlatform = "tiktok"
	ConnectStartOAuthFlowParamsPlatformYoutube        ConnectStartOAuthFlowParamsPlatform = "youtube"
	ConnectStartOAuthFlowParamsPlatformPinterest      ConnectStartOAuthFlowParamsPlatform = "pinterest"
	ConnectStartOAuthFlowParamsPlatformReddit         ConnectStartOAuthFlowParamsPlatform = "reddit"
	ConnectStartOAuthFlowParamsPlatformThreads        ConnectStartOAuthFlowParamsPlatform = "threads"
	ConnectStartOAuthFlowParamsPlatformSnapchat       ConnectStartOAuthFlowParamsPlatform = "snapchat"
	ConnectStartOAuthFlowParamsPlatformGooglebusiness ConnectStartOAuthFlowParamsPlatform = "googlebusiness"
	ConnectStartOAuthFlowParamsPlatformMastodon       ConnectStartOAuthFlowParamsPlatform = "mastodon"
)

func (r ConnectStartOAuthFlowParamsPlatform) IsKnown() bool {
	switch r {
	case ConnectStartOAuthFlowParamsPlatformTwitter, ConnectStartOAuthFlowParamsPlatformInstagram, ConnectStartOAuthFlowParamsPlatformFacebook, ConnectStartOAuthFlowParamsPlatformLinkedin, ConnectStartOAuthFlowParamsPlatformTiktok, ConnectStartOAuthFlowParamsPlatformYoutube, ConnectStartOAuthFlowParamsPlatformPinterest, ConnectStartOAuthFlowParamsPlatformReddit, ConnectStartOAuthFlowParamsPlatformThreads, ConnectStartOAuthFlowParamsPlatformSnapchat, ConnectStartOAuthFlowParamsPlatformGooglebusiness, ConnectStartOAuthFlowParamsPlatformMastodon:
		return true
	}
	return false
}
