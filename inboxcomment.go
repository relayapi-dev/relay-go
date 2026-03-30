// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
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

// InboxCommentService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxCommentService] method instead.
type InboxCommentService struct {
	Options []option.RequestOption
	Hide    *InboxCommentHideService
	Like    *InboxCommentLikeService
}

// NewInboxCommentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxCommentService(opts ...option.RequestOption) (r *InboxCommentService) {
	r = &InboxCommentService{}
	r.Options = opts
	r.Hide = NewInboxCommentHideService(opts...)
	r.Like = NewInboxCommentLikeService(opts...)
	return
}

// Get comments for a specific post
func (r *InboxCommentService) Get(ctx context.Context, postID string, query InboxCommentGetParams, opts ...option.RequestOption) (res *InboxCommentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if postID == "" {
		err = errors.New("missing required post_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s", postID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List comments across platforms
func (r *InboxCommentService) List(ctx context.Context, query InboxCommentListParams, opts ...option.RequestOption) (res *InboxCommentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/inbox/comments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a comment
func (r *InboxCommentService) Delete(ctx context.Context, commentID string, opts ...option.RequestOption) (res *InboxCommentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commentID == "" {
		err = errors.New("missing required comment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s", commentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Send a private reply to a commenter
func (r *InboxCommentService) PrivateReply(ctx context.Context, commentID string, body InboxCommentPrivateReplyParams, opts ...option.RequestOption) (res *InboxCommentPrivateReplyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commentID == "" {
		err = errors.New("missing required comment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s/private-reply", commentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reply to a comment
func (r *InboxCommentService) Reply(ctx context.Context, postID string, body InboxCommentReplyParams, opts ...option.RequestOption) (res *InboxCommentReplyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if postID == "" {
		err = errors.New("missing required post_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s/reply", postID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type InboxCommentGetResponse struct {
	Data       []InboxCommentGetResponseData   `json:"data" api:"required"`
	HasMore    bool                            `json:"has_more"`
	NextCursor string                          `json:"next_cursor" api:"nullable"`
	Platform   InboxCommentGetResponsePlatform `json:"platform"`
	// Post ID if filtered by post
	PostID string                      `json:"post_id"`
	JSON   inboxCommentGetResponseJSON `json:"-"`
}

// inboxCommentGetResponseJSON contains the JSON metadata for the struct
// [InboxCommentGetResponse]
type inboxCommentGetResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Platform    apijson.Field
	PostID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentGetResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentGetResponseData struct {
	// Comment ID
	ID string `json:"id" api:"required"`
	// Comment author name
	AuthorName string `json:"author_name" api:"required"`
	// Comment timestamp
	CreatedAt time.Time                           `json:"created_at" api:"required" format:"date-time"`
	Platform  InboxCommentGetResponseDataPlatform `json:"platform" api:"required"`
	// Comment text
	Text string `json:"text" api:"required"`
	// Social account ID
	AccountID string `json:"account_id"`
	// Author avatar URL
	AuthorAvatar string `json:"author_avatar" api:"nullable"`
	// Whether comment is hidden
	Hidden bool `json:"hidden"`
	// Like count
	Likes float64 `json:"likes"`
	// Platform post/media/video ID
	PostID string `json:"post_id"`
	// URL to the post on the platform
	PostPlatformURL string `json:"post_platform_url" api:"nullable"`
	// Post caption snippet
	PostText string `json:"post_text" api:"nullable"`
	// Post thumbnail URL
	PostThumbnailURL string `json:"post_thumbnail_url" api:"nullable"`
	// Reply count
	RepliesCount float64                         `json:"replies_count"`
	JSON         inboxCommentGetResponseDataJSON `json:"-"`
}

// inboxCommentGetResponseDataJSON contains the JSON metadata for the struct
// [InboxCommentGetResponseData]
type inboxCommentGetResponseDataJSON struct {
	ID               apijson.Field
	AuthorName       apijson.Field
	CreatedAt        apijson.Field
	Platform         apijson.Field
	Text             apijson.Field
	AccountID        apijson.Field
	AuthorAvatar     apijson.Field
	Hidden           apijson.Field
	Likes            apijson.Field
	PostID           apijson.Field
	PostPlatformURL  apijson.Field
	PostText         apijson.Field
	PostThumbnailURL apijson.Field
	RepliesCount     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InboxCommentGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type InboxCommentGetResponseDataPlatform string

const (
	InboxCommentGetResponseDataPlatformTwitter        InboxCommentGetResponseDataPlatform = "twitter"
	InboxCommentGetResponseDataPlatformInstagram      InboxCommentGetResponseDataPlatform = "instagram"
	InboxCommentGetResponseDataPlatformFacebook       InboxCommentGetResponseDataPlatform = "facebook"
	InboxCommentGetResponseDataPlatformLinkedin       InboxCommentGetResponseDataPlatform = "linkedin"
	InboxCommentGetResponseDataPlatformTiktok         InboxCommentGetResponseDataPlatform = "tiktok"
	InboxCommentGetResponseDataPlatformYoutube        InboxCommentGetResponseDataPlatform = "youtube"
	InboxCommentGetResponseDataPlatformPinterest      InboxCommentGetResponseDataPlatform = "pinterest"
	InboxCommentGetResponseDataPlatformReddit         InboxCommentGetResponseDataPlatform = "reddit"
	InboxCommentGetResponseDataPlatformBluesky        InboxCommentGetResponseDataPlatform = "bluesky"
	InboxCommentGetResponseDataPlatformThreads        InboxCommentGetResponseDataPlatform = "threads"
	InboxCommentGetResponseDataPlatformTelegram       InboxCommentGetResponseDataPlatform = "telegram"
	InboxCommentGetResponseDataPlatformSnapchat       InboxCommentGetResponseDataPlatform = "snapchat"
	InboxCommentGetResponseDataPlatformGooglebusiness InboxCommentGetResponseDataPlatform = "googlebusiness"
	InboxCommentGetResponseDataPlatformWhatsapp       InboxCommentGetResponseDataPlatform = "whatsapp"
	InboxCommentGetResponseDataPlatformMastodon       InboxCommentGetResponseDataPlatform = "mastodon"
	InboxCommentGetResponseDataPlatformDiscord        InboxCommentGetResponseDataPlatform = "discord"
	InboxCommentGetResponseDataPlatformSMS            InboxCommentGetResponseDataPlatform = "sms"
)

func (r InboxCommentGetResponseDataPlatform) IsKnown() bool {
	switch r {
	case InboxCommentGetResponseDataPlatformTwitter, InboxCommentGetResponseDataPlatformInstagram, InboxCommentGetResponseDataPlatformFacebook, InboxCommentGetResponseDataPlatformLinkedin, InboxCommentGetResponseDataPlatformTiktok, InboxCommentGetResponseDataPlatformYoutube, InboxCommentGetResponseDataPlatformPinterest, InboxCommentGetResponseDataPlatformReddit, InboxCommentGetResponseDataPlatformBluesky, InboxCommentGetResponseDataPlatformThreads, InboxCommentGetResponseDataPlatformTelegram, InboxCommentGetResponseDataPlatformSnapchat, InboxCommentGetResponseDataPlatformGooglebusiness, InboxCommentGetResponseDataPlatformWhatsapp, InboxCommentGetResponseDataPlatformMastodon, InboxCommentGetResponseDataPlatformDiscord, InboxCommentGetResponseDataPlatformSMS:
		return true
	}
	return false
}

type InboxCommentGetResponsePlatform string

const (
	InboxCommentGetResponsePlatformTwitter        InboxCommentGetResponsePlatform = "twitter"
	InboxCommentGetResponsePlatformInstagram      InboxCommentGetResponsePlatform = "instagram"
	InboxCommentGetResponsePlatformFacebook       InboxCommentGetResponsePlatform = "facebook"
	InboxCommentGetResponsePlatformLinkedin       InboxCommentGetResponsePlatform = "linkedin"
	InboxCommentGetResponsePlatformTiktok         InboxCommentGetResponsePlatform = "tiktok"
	InboxCommentGetResponsePlatformYoutube        InboxCommentGetResponsePlatform = "youtube"
	InboxCommentGetResponsePlatformPinterest      InboxCommentGetResponsePlatform = "pinterest"
	InboxCommentGetResponsePlatformReddit         InboxCommentGetResponsePlatform = "reddit"
	InboxCommentGetResponsePlatformBluesky        InboxCommentGetResponsePlatform = "bluesky"
	InboxCommentGetResponsePlatformThreads        InboxCommentGetResponsePlatform = "threads"
	InboxCommentGetResponsePlatformTelegram       InboxCommentGetResponsePlatform = "telegram"
	InboxCommentGetResponsePlatformSnapchat       InboxCommentGetResponsePlatform = "snapchat"
	InboxCommentGetResponsePlatformGooglebusiness InboxCommentGetResponsePlatform = "googlebusiness"
	InboxCommentGetResponsePlatformWhatsapp       InboxCommentGetResponsePlatform = "whatsapp"
	InboxCommentGetResponsePlatformMastodon       InboxCommentGetResponsePlatform = "mastodon"
	InboxCommentGetResponsePlatformDiscord        InboxCommentGetResponsePlatform = "discord"
	InboxCommentGetResponsePlatformSMS            InboxCommentGetResponsePlatform = "sms"
)

func (r InboxCommentGetResponsePlatform) IsKnown() bool {
	switch r {
	case InboxCommentGetResponsePlatformTwitter, InboxCommentGetResponsePlatformInstagram, InboxCommentGetResponsePlatformFacebook, InboxCommentGetResponsePlatformLinkedin, InboxCommentGetResponsePlatformTiktok, InboxCommentGetResponsePlatformYoutube, InboxCommentGetResponsePlatformPinterest, InboxCommentGetResponsePlatformReddit, InboxCommentGetResponsePlatformBluesky, InboxCommentGetResponsePlatformThreads, InboxCommentGetResponsePlatformTelegram, InboxCommentGetResponsePlatformSnapchat, InboxCommentGetResponsePlatformGooglebusiness, InboxCommentGetResponsePlatformWhatsapp, InboxCommentGetResponsePlatformMastodon, InboxCommentGetResponsePlatformDiscord, InboxCommentGetResponsePlatformSMS:
		return true
	}
	return false
}

type InboxCommentListResponse struct {
	Data       []InboxCommentListResponseData   `json:"data" api:"required"`
	HasMore    bool                             `json:"has_more"`
	NextCursor string                           `json:"next_cursor" api:"nullable"`
	Platform   InboxCommentListResponsePlatform `json:"platform"`
	// Post ID if filtered by post
	PostID string                       `json:"post_id"`
	JSON   inboxCommentListResponseJSON `json:"-"`
}

// inboxCommentListResponseJSON contains the JSON metadata for the struct
// [InboxCommentListResponse]
type inboxCommentListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Platform    apijson.Field
	PostID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentListResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentListResponseData struct {
	// Comment ID
	ID string `json:"id" api:"required"`
	// Comment author name
	AuthorName string `json:"author_name" api:"required"`
	// Comment timestamp
	CreatedAt time.Time                            `json:"created_at" api:"required" format:"date-time"`
	Platform  InboxCommentListResponseDataPlatform `json:"platform" api:"required"`
	// Comment text
	Text string `json:"text" api:"required"`
	// Social account ID
	AccountID string `json:"account_id"`
	// Author avatar URL
	AuthorAvatar string `json:"author_avatar" api:"nullable"`
	// Whether comment is hidden
	Hidden bool `json:"hidden"`
	// Like count
	Likes float64 `json:"likes"`
	// Platform post/media/video ID
	PostID string `json:"post_id"`
	// URL to the post on the platform
	PostPlatformURL string `json:"post_platform_url" api:"nullable"`
	// Post caption snippet
	PostText string `json:"post_text" api:"nullable"`
	// Post thumbnail URL
	PostThumbnailURL string `json:"post_thumbnail_url" api:"nullable"`
	// Reply count
	RepliesCount float64                          `json:"replies_count"`
	JSON         inboxCommentListResponseDataJSON `json:"-"`
}

// inboxCommentListResponseDataJSON contains the JSON metadata for the struct
// [InboxCommentListResponseData]
type inboxCommentListResponseDataJSON struct {
	ID               apijson.Field
	AuthorName       apijson.Field
	CreatedAt        apijson.Field
	Platform         apijson.Field
	Text             apijson.Field
	AccountID        apijson.Field
	AuthorAvatar     apijson.Field
	Hidden           apijson.Field
	Likes            apijson.Field
	PostID           apijson.Field
	PostPlatformURL  apijson.Field
	PostText         apijson.Field
	PostThumbnailURL apijson.Field
	RepliesCount     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InboxCommentListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentListResponseDataJSON) RawJSON() string {
	return r.raw
}

type InboxCommentListResponseDataPlatform string

const (
	InboxCommentListResponseDataPlatformTwitter        InboxCommentListResponseDataPlatform = "twitter"
	InboxCommentListResponseDataPlatformInstagram      InboxCommentListResponseDataPlatform = "instagram"
	InboxCommentListResponseDataPlatformFacebook       InboxCommentListResponseDataPlatform = "facebook"
	InboxCommentListResponseDataPlatformLinkedin       InboxCommentListResponseDataPlatform = "linkedin"
	InboxCommentListResponseDataPlatformTiktok         InboxCommentListResponseDataPlatform = "tiktok"
	InboxCommentListResponseDataPlatformYoutube        InboxCommentListResponseDataPlatform = "youtube"
	InboxCommentListResponseDataPlatformPinterest      InboxCommentListResponseDataPlatform = "pinterest"
	InboxCommentListResponseDataPlatformReddit         InboxCommentListResponseDataPlatform = "reddit"
	InboxCommentListResponseDataPlatformBluesky        InboxCommentListResponseDataPlatform = "bluesky"
	InboxCommentListResponseDataPlatformThreads        InboxCommentListResponseDataPlatform = "threads"
	InboxCommentListResponseDataPlatformTelegram       InboxCommentListResponseDataPlatform = "telegram"
	InboxCommentListResponseDataPlatformSnapchat       InboxCommentListResponseDataPlatform = "snapchat"
	InboxCommentListResponseDataPlatformGooglebusiness InboxCommentListResponseDataPlatform = "googlebusiness"
	InboxCommentListResponseDataPlatformWhatsapp       InboxCommentListResponseDataPlatform = "whatsapp"
	InboxCommentListResponseDataPlatformMastodon       InboxCommentListResponseDataPlatform = "mastodon"
	InboxCommentListResponseDataPlatformDiscord        InboxCommentListResponseDataPlatform = "discord"
	InboxCommentListResponseDataPlatformSMS            InboxCommentListResponseDataPlatform = "sms"
)

func (r InboxCommentListResponseDataPlatform) IsKnown() bool {
	switch r {
	case InboxCommentListResponseDataPlatformTwitter, InboxCommentListResponseDataPlatformInstagram, InboxCommentListResponseDataPlatformFacebook, InboxCommentListResponseDataPlatformLinkedin, InboxCommentListResponseDataPlatformTiktok, InboxCommentListResponseDataPlatformYoutube, InboxCommentListResponseDataPlatformPinterest, InboxCommentListResponseDataPlatformReddit, InboxCommentListResponseDataPlatformBluesky, InboxCommentListResponseDataPlatformThreads, InboxCommentListResponseDataPlatformTelegram, InboxCommentListResponseDataPlatformSnapchat, InboxCommentListResponseDataPlatformGooglebusiness, InboxCommentListResponseDataPlatformWhatsapp, InboxCommentListResponseDataPlatformMastodon, InboxCommentListResponseDataPlatformDiscord, InboxCommentListResponseDataPlatformSMS:
		return true
	}
	return false
}

type InboxCommentListResponsePlatform string

const (
	InboxCommentListResponsePlatformTwitter        InboxCommentListResponsePlatform = "twitter"
	InboxCommentListResponsePlatformInstagram      InboxCommentListResponsePlatform = "instagram"
	InboxCommentListResponsePlatformFacebook       InboxCommentListResponsePlatform = "facebook"
	InboxCommentListResponsePlatformLinkedin       InboxCommentListResponsePlatform = "linkedin"
	InboxCommentListResponsePlatformTiktok         InboxCommentListResponsePlatform = "tiktok"
	InboxCommentListResponsePlatformYoutube        InboxCommentListResponsePlatform = "youtube"
	InboxCommentListResponsePlatformPinterest      InboxCommentListResponsePlatform = "pinterest"
	InboxCommentListResponsePlatformReddit         InboxCommentListResponsePlatform = "reddit"
	InboxCommentListResponsePlatformBluesky        InboxCommentListResponsePlatform = "bluesky"
	InboxCommentListResponsePlatformThreads        InboxCommentListResponsePlatform = "threads"
	InboxCommentListResponsePlatformTelegram       InboxCommentListResponsePlatform = "telegram"
	InboxCommentListResponsePlatformSnapchat       InboxCommentListResponsePlatform = "snapchat"
	InboxCommentListResponsePlatformGooglebusiness InboxCommentListResponsePlatform = "googlebusiness"
	InboxCommentListResponsePlatformWhatsapp       InboxCommentListResponsePlatform = "whatsapp"
	InboxCommentListResponsePlatformMastodon       InboxCommentListResponsePlatform = "mastodon"
	InboxCommentListResponsePlatformDiscord        InboxCommentListResponsePlatform = "discord"
	InboxCommentListResponsePlatformSMS            InboxCommentListResponsePlatform = "sms"
)

func (r InboxCommentListResponsePlatform) IsKnown() bool {
	switch r {
	case InboxCommentListResponsePlatformTwitter, InboxCommentListResponsePlatformInstagram, InboxCommentListResponsePlatformFacebook, InboxCommentListResponsePlatformLinkedin, InboxCommentListResponsePlatformTiktok, InboxCommentListResponsePlatformYoutube, InboxCommentListResponsePlatformPinterest, InboxCommentListResponsePlatformReddit, InboxCommentListResponsePlatformBluesky, InboxCommentListResponsePlatformThreads, InboxCommentListResponsePlatformTelegram, InboxCommentListResponsePlatformSnapchat, InboxCommentListResponsePlatformGooglebusiness, InboxCommentListResponsePlatformWhatsapp, InboxCommentListResponsePlatformMastodon, InboxCommentListResponsePlatformDiscord, InboxCommentListResponsePlatformSMS:
		return true
	}
	return false
}

type InboxCommentDeleteResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                         `json:"comment_id"`
	JSON      inboxCommentDeleteResponseJSON `json:"-"`
}

// inboxCommentDeleteResponseJSON contains the JSON metadata for the struct
// [InboxCommentDeleteResponse]
type inboxCommentDeleteResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentPrivateReplyResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                               `json:"comment_id"`
	JSON      inboxCommentPrivateReplyResponseJSON `json:"-"`
}

// inboxCommentPrivateReplyResponseJSON contains the JSON metadata for the struct
// [InboxCommentPrivateReplyResponse]
type inboxCommentPrivateReplyResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentPrivateReplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentPrivateReplyResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentReplyResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                        `json:"comment_id"`
	JSON      inboxCommentReplyResponseJSON `json:"-"`
}

// inboxCommentReplyResponseJSON contains the JSON metadata for the struct
// [InboxCommentReplyResponse]
type inboxCommentReplyResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentReplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentReplyResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentGetParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items
	Limit param.Field[int64] `query:"limit"`
	// Filter by platform
	Platform param.Field[InboxCommentGetParamsPlatform] `query:"platform"`
}

// URLQuery serializes [InboxCommentGetParams]'s query parameters as `url.Values`.
func (r InboxCommentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type InboxCommentGetParamsPlatform string

const (
	InboxCommentGetParamsPlatformTwitter        InboxCommentGetParamsPlatform = "twitter"
	InboxCommentGetParamsPlatformInstagram      InboxCommentGetParamsPlatform = "instagram"
	InboxCommentGetParamsPlatformFacebook       InboxCommentGetParamsPlatform = "facebook"
	InboxCommentGetParamsPlatformLinkedin       InboxCommentGetParamsPlatform = "linkedin"
	InboxCommentGetParamsPlatformTiktok         InboxCommentGetParamsPlatform = "tiktok"
	InboxCommentGetParamsPlatformYoutube        InboxCommentGetParamsPlatform = "youtube"
	InboxCommentGetParamsPlatformPinterest      InboxCommentGetParamsPlatform = "pinterest"
	InboxCommentGetParamsPlatformReddit         InboxCommentGetParamsPlatform = "reddit"
	InboxCommentGetParamsPlatformBluesky        InboxCommentGetParamsPlatform = "bluesky"
	InboxCommentGetParamsPlatformThreads        InboxCommentGetParamsPlatform = "threads"
	InboxCommentGetParamsPlatformTelegram       InboxCommentGetParamsPlatform = "telegram"
	InboxCommentGetParamsPlatformSnapchat       InboxCommentGetParamsPlatform = "snapchat"
	InboxCommentGetParamsPlatformGooglebusiness InboxCommentGetParamsPlatform = "googlebusiness"
	InboxCommentGetParamsPlatformWhatsapp       InboxCommentGetParamsPlatform = "whatsapp"
	InboxCommentGetParamsPlatformMastodon       InboxCommentGetParamsPlatform = "mastodon"
	InboxCommentGetParamsPlatformDiscord        InboxCommentGetParamsPlatform = "discord"
	InboxCommentGetParamsPlatformSMS            InboxCommentGetParamsPlatform = "sms"
)

func (r InboxCommentGetParamsPlatform) IsKnown() bool {
	switch r {
	case InboxCommentGetParamsPlatformTwitter, InboxCommentGetParamsPlatformInstagram, InboxCommentGetParamsPlatformFacebook, InboxCommentGetParamsPlatformLinkedin, InboxCommentGetParamsPlatformTiktok, InboxCommentGetParamsPlatformYoutube, InboxCommentGetParamsPlatformPinterest, InboxCommentGetParamsPlatformReddit, InboxCommentGetParamsPlatformBluesky, InboxCommentGetParamsPlatformThreads, InboxCommentGetParamsPlatformTelegram, InboxCommentGetParamsPlatformSnapchat, InboxCommentGetParamsPlatformGooglebusiness, InboxCommentGetParamsPlatformWhatsapp, InboxCommentGetParamsPlatformMastodon, InboxCommentGetParamsPlatformDiscord, InboxCommentGetParamsPlatformSMS:
		return true
	}
	return false
}

type InboxCommentListParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items
	Limit param.Field[int64] `query:"limit"`
	// Filter by platform
	Platform param.Field[InboxCommentListParamsPlatform] `query:"platform"`
}

// URLQuery serializes [InboxCommentListParams]'s query parameters as `url.Values`.
func (r InboxCommentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type InboxCommentListParamsPlatform string

const (
	InboxCommentListParamsPlatformTwitter        InboxCommentListParamsPlatform = "twitter"
	InboxCommentListParamsPlatformInstagram      InboxCommentListParamsPlatform = "instagram"
	InboxCommentListParamsPlatformFacebook       InboxCommentListParamsPlatform = "facebook"
	InboxCommentListParamsPlatformLinkedin       InboxCommentListParamsPlatform = "linkedin"
	InboxCommentListParamsPlatformTiktok         InboxCommentListParamsPlatform = "tiktok"
	InboxCommentListParamsPlatformYoutube        InboxCommentListParamsPlatform = "youtube"
	InboxCommentListParamsPlatformPinterest      InboxCommentListParamsPlatform = "pinterest"
	InboxCommentListParamsPlatformReddit         InboxCommentListParamsPlatform = "reddit"
	InboxCommentListParamsPlatformBluesky        InboxCommentListParamsPlatform = "bluesky"
	InboxCommentListParamsPlatformThreads        InboxCommentListParamsPlatform = "threads"
	InboxCommentListParamsPlatformTelegram       InboxCommentListParamsPlatform = "telegram"
	InboxCommentListParamsPlatformSnapchat       InboxCommentListParamsPlatform = "snapchat"
	InboxCommentListParamsPlatformGooglebusiness InboxCommentListParamsPlatform = "googlebusiness"
	InboxCommentListParamsPlatformWhatsapp       InboxCommentListParamsPlatform = "whatsapp"
	InboxCommentListParamsPlatformMastodon       InboxCommentListParamsPlatform = "mastodon"
	InboxCommentListParamsPlatformDiscord        InboxCommentListParamsPlatform = "discord"
	InboxCommentListParamsPlatformSMS            InboxCommentListParamsPlatform = "sms"
)

func (r InboxCommentListParamsPlatform) IsKnown() bool {
	switch r {
	case InboxCommentListParamsPlatformTwitter, InboxCommentListParamsPlatformInstagram, InboxCommentListParamsPlatformFacebook, InboxCommentListParamsPlatformLinkedin, InboxCommentListParamsPlatformTiktok, InboxCommentListParamsPlatformYoutube, InboxCommentListParamsPlatformPinterest, InboxCommentListParamsPlatformReddit, InboxCommentListParamsPlatformBluesky, InboxCommentListParamsPlatformThreads, InboxCommentListParamsPlatformTelegram, InboxCommentListParamsPlatformSnapchat, InboxCommentListParamsPlatformGooglebusiness, InboxCommentListParamsPlatformWhatsapp, InboxCommentListParamsPlatformMastodon, InboxCommentListParamsPlatformDiscord, InboxCommentListParamsPlatformSMS:
		return true
	}
	return false
}

type InboxCommentPrivateReplyParams struct {
	// Account ID to reply from
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Private reply text
	Text param.Field[string] `json:"text" api:"required"`
}

func (r InboxCommentPrivateReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboxCommentReplyParams struct {
	// Account ID to reply from
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Reply text
	Text param.Field[string] `json:"text" api:"required"`
	// Parent comment ID for threaded replies
	CommentID param.Field[string] `json:"comment_id"`
}

func (r InboxCommentReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
