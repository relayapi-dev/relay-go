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

// InboxReviewService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxReviewService] method instead.
type InboxReviewService struct {
	Options []option.RequestOption
	Reply   *InboxReviewReplyService
}

// NewInboxReviewService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxReviewService(opts ...option.RequestOption) (r *InboxReviewService) {
	r = &InboxReviewService{}
	r.Options = opts
	r.Reply = NewInboxReviewReplyService(opts...)
	return
}

// List reviews across platforms
func (r *InboxReviewService) List(ctx context.Context, query InboxReviewListParams, opts ...option.RequestOption) (res *InboxReviewListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/inbox/reviews"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type InboxReviewListResponse struct {
	Data []InboxReviewListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                      `json:"next_cursor" api:"required,nullable"`
	JSON       inboxReviewListResponseJSON `json:"-"`
}

// inboxReviewListResponseJSON contains the JSON metadata for the struct
// [InboxReviewListResponse]
type inboxReviewListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxReviewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxReviewListResponseJSON) RawJSON() string {
	return r.raw
}

type InboxReviewListResponseData struct {
	// Review ID
	ID string `json:"id" api:"required"`
	// Review author name
	AuthorName string `json:"author_name" api:"required"`
	// Review timestamp
	CreatedAt time.Time                           `json:"created_at" api:"required" format:"date-time"`
	Platform  InboxReviewListResponseDataPlatform `json:"platform" api:"required"`
	// Rating (1-5)
	Rating float64 `json:"rating" api:"required"`
	// Business reply text
	Reply string `json:"reply" api:"nullable"`
	// Review text
	Text string                          `json:"text" api:"nullable"`
	JSON inboxReviewListResponseDataJSON `json:"-"`
}

// inboxReviewListResponseDataJSON contains the JSON metadata for the struct
// [InboxReviewListResponseData]
type inboxReviewListResponseDataJSON struct {
	ID          apijson.Field
	AuthorName  apijson.Field
	CreatedAt   apijson.Field
	Platform    apijson.Field
	Rating      apijson.Field
	Reply       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxReviewListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxReviewListResponseDataJSON) RawJSON() string {
	return r.raw
}

type InboxReviewListResponseDataPlatform string

const (
	InboxReviewListResponseDataPlatformTwitter        InboxReviewListResponseDataPlatform = "twitter"
	InboxReviewListResponseDataPlatformInstagram      InboxReviewListResponseDataPlatform = "instagram"
	InboxReviewListResponseDataPlatformFacebook       InboxReviewListResponseDataPlatform = "facebook"
	InboxReviewListResponseDataPlatformLinkedin       InboxReviewListResponseDataPlatform = "linkedin"
	InboxReviewListResponseDataPlatformTiktok         InboxReviewListResponseDataPlatform = "tiktok"
	InboxReviewListResponseDataPlatformYoutube        InboxReviewListResponseDataPlatform = "youtube"
	InboxReviewListResponseDataPlatformPinterest      InboxReviewListResponseDataPlatform = "pinterest"
	InboxReviewListResponseDataPlatformReddit         InboxReviewListResponseDataPlatform = "reddit"
	InboxReviewListResponseDataPlatformBluesky        InboxReviewListResponseDataPlatform = "bluesky"
	InboxReviewListResponseDataPlatformThreads        InboxReviewListResponseDataPlatform = "threads"
	InboxReviewListResponseDataPlatformTelegram       InboxReviewListResponseDataPlatform = "telegram"
	InboxReviewListResponseDataPlatformSnapchat       InboxReviewListResponseDataPlatform = "snapchat"
	InboxReviewListResponseDataPlatformGooglebusiness InboxReviewListResponseDataPlatform = "googlebusiness"
	InboxReviewListResponseDataPlatformWhatsapp       InboxReviewListResponseDataPlatform = "whatsapp"
	InboxReviewListResponseDataPlatformMastodon       InboxReviewListResponseDataPlatform = "mastodon"
	InboxReviewListResponseDataPlatformDiscord        InboxReviewListResponseDataPlatform = "discord"
	InboxReviewListResponseDataPlatformSMS            InboxReviewListResponseDataPlatform = "sms"
	InboxReviewListResponseDataPlatformBeehiiv        InboxReviewListResponseDataPlatform = "beehiiv"
	InboxReviewListResponseDataPlatformConvertkit     InboxReviewListResponseDataPlatform = "convertkit"
	InboxReviewListResponseDataPlatformMailchimp      InboxReviewListResponseDataPlatform = "mailchimp"
	InboxReviewListResponseDataPlatformListmonk       InboxReviewListResponseDataPlatform = "listmonk"
)

func (r InboxReviewListResponseDataPlatform) IsKnown() bool {
	switch r {
	case InboxReviewListResponseDataPlatformTwitter, InboxReviewListResponseDataPlatformInstagram, InboxReviewListResponseDataPlatformFacebook, InboxReviewListResponseDataPlatformLinkedin, InboxReviewListResponseDataPlatformTiktok, InboxReviewListResponseDataPlatformYoutube, InboxReviewListResponseDataPlatformPinterest, InboxReviewListResponseDataPlatformReddit, InboxReviewListResponseDataPlatformBluesky, InboxReviewListResponseDataPlatformThreads, InboxReviewListResponseDataPlatformTelegram, InboxReviewListResponseDataPlatformSnapchat, InboxReviewListResponseDataPlatformGooglebusiness, InboxReviewListResponseDataPlatformWhatsapp, InboxReviewListResponseDataPlatformMastodon, InboxReviewListResponseDataPlatformDiscord, InboxReviewListResponseDataPlatformSMS, InboxReviewListResponseDataPlatformBeehiiv, InboxReviewListResponseDataPlatformConvertkit, InboxReviewListResponseDataPlatformMailchimp, InboxReviewListResponseDataPlatformListmonk:
		return true
	}
	return false
}

type InboxReviewListParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items
	Limit     param.Field[int64] `query:"limit"`
	MaxRating param.Field[int64] `query:"max_rating"`
	MinRating param.Field[int64] `query:"min_rating"`
	// Filter by platform
	Platform param.Field[InboxReviewListParamsPlatform] `query:"platform"`
}

// URLQuery serializes [InboxReviewListParams]'s query parameters as `url.Values`.
func (r InboxReviewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type InboxReviewListParamsPlatform string

const (
	InboxReviewListParamsPlatformTwitter        InboxReviewListParamsPlatform = "twitter"
	InboxReviewListParamsPlatformInstagram      InboxReviewListParamsPlatform = "instagram"
	InboxReviewListParamsPlatformFacebook       InboxReviewListParamsPlatform = "facebook"
	InboxReviewListParamsPlatformLinkedin       InboxReviewListParamsPlatform = "linkedin"
	InboxReviewListParamsPlatformTiktok         InboxReviewListParamsPlatform = "tiktok"
	InboxReviewListParamsPlatformYoutube        InboxReviewListParamsPlatform = "youtube"
	InboxReviewListParamsPlatformPinterest      InboxReviewListParamsPlatform = "pinterest"
	InboxReviewListParamsPlatformReddit         InboxReviewListParamsPlatform = "reddit"
	InboxReviewListParamsPlatformBluesky        InboxReviewListParamsPlatform = "bluesky"
	InboxReviewListParamsPlatformThreads        InboxReviewListParamsPlatform = "threads"
	InboxReviewListParamsPlatformTelegram       InboxReviewListParamsPlatform = "telegram"
	InboxReviewListParamsPlatformSnapchat       InboxReviewListParamsPlatform = "snapchat"
	InboxReviewListParamsPlatformGooglebusiness InboxReviewListParamsPlatform = "googlebusiness"
	InboxReviewListParamsPlatformWhatsapp       InboxReviewListParamsPlatform = "whatsapp"
	InboxReviewListParamsPlatformMastodon       InboxReviewListParamsPlatform = "mastodon"
	InboxReviewListParamsPlatformDiscord        InboxReviewListParamsPlatform = "discord"
	InboxReviewListParamsPlatformSMS            InboxReviewListParamsPlatform = "sms"
	InboxReviewListParamsPlatformBeehiiv        InboxReviewListParamsPlatform = "beehiiv"
	InboxReviewListParamsPlatformConvertkit     InboxReviewListParamsPlatform = "convertkit"
	InboxReviewListParamsPlatformMailchimp      InboxReviewListParamsPlatform = "mailchimp"
	InboxReviewListParamsPlatformListmonk       InboxReviewListParamsPlatform = "listmonk"
)

func (r InboxReviewListParamsPlatform) IsKnown() bool {
	switch r {
	case InboxReviewListParamsPlatformTwitter, InboxReviewListParamsPlatformInstagram, InboxReviewListParamsPlatformFacebook, InboxReviewListParamsPlatformLinkedin, InboxReviewListParamsPlatformTiktok, InboxReviewListParamsPlatformYoutube, InboxReviewListParamsPlatformPinterest, InboxReviewListParamsPlatformReddit, InboxReviewListParamsPlatformBluesky, InboxReviewListParamsPlatformThreads, InboxReviewListParamsPlatformTelegram, InboxReviewListParamsPlatformSnapchat, InboxReviewListParamsPlatformGooglebusiness, InboxReviewListParamsPlatformWhatsapp, InboxReviewListParamsPlatformMastodon, InboxReviewListParamsPlatformDiscord, InboxReviewListParamsPlatformSMS, InboxReviewListParamsPlatformBeehiiv, InboxReviewListParamsPlatformConvertkit, InboxReviewListParamsPlatformMailchimp, InboxReviewListParamsPlatformListmonk:
		return true
	}
	return false
}
