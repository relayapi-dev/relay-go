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

// InboxMessageService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxMessageService] method instead.
type InboxMessageService struct {
	Options []option.RequestOption
}

// NewInboxMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxMessageService(opts ...option.RequestOption) (r *InboxMessageService) {
	r = &InboxMessageService{}
	r.Options = opts
	return
}

// Get messages in a conversation
func (r *InboxMessageService) Get(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *InboxMessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/messages/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List message conversations
func (r *InboxMessageService) List(ctx context.Context, query InboxMessageListParams, opts ...option.RequestOption) (res *InboxMessageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/inbox/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive a conversation
func (r *InboxMessageService) Archive(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *InboxMessageArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/messages/%s/archive", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Edit a sent message
func (r *InboxMessageService) Edit(ctx context.Context, conversationID string, messageID string, body InboxMessageEditParams, opts ...option.RequestOption) (res *InboxMessageEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/messages/%s/%s", conversationID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Send a message in a conversation
func (r *InboxMessageService) Send(ctx context.Context, conversationID string, body InboxMessageSendParams, opts ...option.RequestOption) (res *InboxMessageSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/messages/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type InboxMessageGetResponse struct {
	Data       []InboxMessageGetResponseData `json:"data" api:"required"`
	HasMore    bool                          `json:"has_more"`
	NextCursor string                        `json:"next_cursor" api:"nullable"`
	JSON       inboxMessageGetResponseJSON   `json:"-"`
}

// inboxMessageGetResponseJSON contains the JSON metadata for the struct
// [InboxMessageGetResponse]
type inboxMessageGetResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageGetResponseJSON) RawJSON() string {
	return r.raw
}

type InboxMessageGetResponseData struct {
	// Message ID
	ID string `json:"id" api:"required"`
	// Message timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Message sender
	Sender InboxMessageGetResponseDataSender `json:"sender" api:"required"`
	// Message text
	Text string `json:"text" api:"required"`
	// Message attachments
	Attachments []InboxMessageGetResponseDataAttachment `json:"attachments"`
	JSON        inboxMessageGetResponseDataJSON         `json:"-"`
}

// inboxMessageGetResponseDataJSON contains the JSON metadata for the struct
// [InboxMessageGetResponseData]
type inboxMessageGetResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Sender      apijson.Field
	Text        apijson.Field
	Attachments apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Message sender
type InboxMessageGetResponseDataSender string

const (
	InboxMessageGetResponseDataSenderUser        InboxMessageGetResponseDataSender = "user"
	InboxMessageGetResponseDataSenderParticipant InboxMessageGetResponseDataSender = "participant"
)

func (r InboxMessageGetResponseDataSender) IsKnown() bool {
	switch r {
	case InboxMessageGetResponseDataSenderUser, InboxMessageGetResponseDataSenderParticipant:
		return true
	}
	return false
}

type InboxMessageGetResponseDataAttachment struct {
	// Attachment MIME type
	Type string `json:"type" api:"required"`
	// Attachment URL
	URL  string                                    `json:"url" api:"required"`
	JSON inboxMessageGetResponseDataAttachmentJSON `json:"-"`
}

// inboxMessageGetResponseDataAttachmentJSON contains the JSON metadata for the
// struct [InboxMessageGetResponseDataAttachment]
type inboxMessageGetResponseDataAttachmentJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageGetResponseDataAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageGetResponseDataAttachmentJSON) RawJSON() string {
	return r.raw
}

type InboxMessageListResponse struct {
	Data []InboxMessageListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                       `json:"next_cursor" api:"required,nullable"`
	JSON       inboxMessageListResponseJSON `json:"-"`
}

// inboxMessageListResponseJSON contains the JSON metadata for the struct
// [InboxMessageListResponse]
type inboxMessageListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageListResponseJSON) RawJSON() string {
	return r.raw
}

type InboxMessageListResponseData struct {
	// Conversation ID
	ID string `json:"id" api:"required"`
	// Account ID
	AccountID string `json:"account_id" api:"required"`
	// Participant display name
	ParticipantName string                               `json:"participant_name" api:"required"`
	Platform        InboxMessageListResponseDataPlatform `json:"platform" api:"required"`
	// Last updated timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Last message text
	LastMessage string `json:"last_message" api:"nullable"`
	// Participant avatar URL
	ParticipantAvatar string `json:"participant_avatar" api:"nullable"`
	// Unread message count
	UnreadCount float64                          `json:"unread_count"`
	JSON        inboxMessageListResponseDataJSON `json:"-"`
}

// inboxMessageListResponseDataJSON contains the JSON metadata for the struct
// [InboxMessageListResponseData]
type inboxMessageListResponseDataJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	ParticipantName   apijson.Field
	Platform          apijson.Field
	UpdatedAt         apijson.Field
	LastMessage       apijson.Field
	ParticipantAvatar apijson.Field
	UnreadCount       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InboxMessageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageListResponseDataJSON) RawJSON() string {
	return r.raw
}

type InboxMessageListResponseDataPlatform string

const (
	InboxMessageListResponseDataPlatformTwitter        InboxMessageListResponseDataPlatform = "twitter"
	InboxMessageListResponseDataPlatformInstagram      InboxMessageListResponseDataPlatform = "instagram"
	InboxMessageListResponseDataPlatformFacebook       InboxMessageListResponseDataPlatform = "facebook"
	InboxMessageListResponseDataPlatformLinkedin       InboxMessageListResponseDataPlatform = "linkedin"
	InboxMessageListResponseDataPlatformTiktok         InboxMessageListResponseDataPlatform = "tiktok"
	InboxMessageListResponseDataPlatformYoutube        InboxMessageListResponseDataPlatform = "youtube"
	InboxMessageListResponseDataPlatformPinterest      InboxMessageListResponseDataPlatform = "pinterest"
	InboxMessageListResponseDataPlatformReddit         InboxMessageListResponseDataPlatform = "reddit"
	InboxMessageListResponseDataPlatformBluesky        InboxMessageListResponseDataPlatform = "bluesky"
	InboxMessageListResponseDataPlatformThreads        InboxMessageListResponseDataPlatform = "threads"
	InboxMessageListResponseDataPlatformTelegram       InboxMessageListResponseDataPlatform = "telegram"
	InboxMessageListResponseDataPlatformSnapchat       InboxMessageListResponseDataPlatform = "snapchat"
	InboxMessageListResponseDataPlatformGooglebusiness InboxMessageListResponseDataPlatform = "googlebusiness"
	InboxMessageListResponseDataPlatformWhatsapp       InboxMessageListResponseDataPlatform = "whatsapp"
	InboxMessageListResponseDataPlatformMastodon       InboxMessageListResponseDataPlatform = "mastodon"
	InboxMessageListResponseDataPlatformDiscord        InboxMessageListResponseDataPlatform = "discord"
	InboxMessageListResponseDataPlatformSMS            InboxMessageListResponseDataPlatform = "sms"
)

func (r InboxMessageListResponseDataPlatform) IsKnown() bool {
	switch r {
	case InboxMessageListResponseDataPlatformTwitter, InboxMessageListResponseDataPlatformInstagram, InboxMessageListResponseDataPlatformFacebook, InboxMessageListResponseDataPlatformLinkedin, InboxMessageListResponseDataPlatformTiktok, InboxMessageListResponseDataPlatformYoutube, InboxMessageListResponseDataPlatformPinterest, InboxMessageListResponseDataPlatformReddit, InboxMessageListResponseDataPlatformBluesky, InboxMessageListResponseDataPlatformThreads, InboxMessageListResponseDataPlatformTelegram, InboxMessageListResponseDataPlatformSnapchat, InboxMessageListResponseDataPlatformGooglebusiness, InboxMessageListResponseDataPlatformWhatsapp, InboxMessageListResponseDataPlatformMastodon, InboxMessageListResponseDataPlatformDiscord, InboxMessageListResponseDataPlatformSMS:
		return true
	}
	return false
}

type InboxMessageArchiveResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Error message if failed
	Error string `json:"error"`
	// Message ID
	MessageID string                          `json:"message_id"`
	JSON      inboxMessageArchiveResponseJSON `json:"-"`
}

// inboxMessageArchiveResponseJSON contains the JSON metadata for the struct
// [InboxMessageArchiveResponse]
type inboxMessageArchiveResponseJSON struct {
	Success     apijson.Field
	Error       apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type InboxMessageEditResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Error message if failed
	Error string `json:"error"`
	// Message ID
	MessageID string                       `json:"message_id"`
	JSON      inboxMessageEditResponseJSON `json:"-"`
}

// inboxMessageEditResponseJSON contains the JSON metadata for the struct
// [InboxMessageEditResponse]
type inboxMessageEditResponseJSON struct {
	Success     apijson.Field
	Error       apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageEditResponseJSON) RawJSON() string {
	return r.raw
}

type InboxMessageSendResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Error message if failed
	Error string `json:"error"`
	// Message ID
	MessageID string                       `json:"message_id"`
	JSON      inboxMessageSendResponseJSON `json:"-"`
}

// inboxMessageSendResponseJSON contains the JSON metadata for the struct
// [InboxMessageSendResponse]
type inboxMessageSendResponseJSON struct {
	Success     apijson.Field
	Error       apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxMessageSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxMessageSendResponseJSON) RawJSON() string {
	return r.raw
}

type InboxMessageListParams struct {
	// Filter by account ID
	AccountID param.Field[string] `query:"account_id"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items
	Limit param.Field[int64] `query:"limit"`
	// Filter by platform
	Platform param.Field[InboxMessageListParamsPlatform] `query:"platform"`
	// Filter by workspace ID
	WorkspaceID param.Field[string] `query:"workspace_id"`
}

// URLQuery serializes [InboxMessageListParams]'s query parameters as `url.Values`.
func (r InboxMessageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type InboxMessageListParamsPlatform string

const (
	InboxMessageListParamsPlatformTwitter        InboxMessageListParamsPlatform = "twitter"
	InboxMessageListParamsPlatformInstagram      InboxMessageListParamsPlatform = "instagram"
	InboxMessageListParamsPlatformFacebook       InboxMessageListParamsPlatform = "facebook"
	InboxMessageListParamsPlatformLinkedin       InboxMessageListParamsPlatform = "linkedin"
	InboxMessageListParamsPlatformTiktok         InboxMessageListParamsPlatform = "tiktok"
	InboxMessageListParamsPlatformYoutube        InboxMessageListParamsPlatform = "youtube"
	InboxMessageListParamsPlatformPinterest      InboxMessageListParamsPlatform = "pinterest"
	InboxMessageListParamsPlatformReddit         InboxMessageListParamsPlatform = "reddit"
	InboxMessageListParamsPlatformBluesky        InboxMessageListParamsPlatform = "bluesky"
	InboxMessageListParamsPlatformThreads        InboxMessageListParamsPlatform = "threads"
	InboxMessageListParamsPlatformTelegram       InboxMessageListParamsPlatform = "telegram"
	InboxMessageListParamsPlatformSnapchat       InboxMessageListParamsPlatform = "snapchat"
	InboxMessageListParamsPlatformGooglebusiness InboxMessageListParamsPlatform = "googlebusiness"
	InboxMessageListParamsPlatformWhatsapp       InboxMessageListParamsPlatform = "whatsapp"
	InboxMessageListParamsPlatformMastodon       InboxMessageListParamsPlatform = "mastodon"
	InboxMessageListParamsPlatformDiscord        InboxMessageListParamsPlatform = "discord"
	InboxMessageListParamsPlatformSMS            InboxMessageListParamsPlatform = "sms"
)

func (r InboxMessageListParamsPlatform) IsKnown() bool {
	switch r {
	case InboxMessageListParamsPlatformTwitter, InboxMessageListParamsPlatformInstagram, InboxMessageListParamsPlatformFacebook, InboxMessageListParamsPlatformLinkedin, InboxMessageListParamsPlatformTiktok, InboxMessageListParamsPlatformYoutube, InboxMessageListParamsPlatformPinterest, InboxMessageListParamsPlatformReddit, InboxMessageListParamsPlatformBluesky, InboxMessageListParamsPlatformThreads, InboxMessageListParamsPlatformTelegram, InboxMessageListParamsPlatformSnapchat, InboxMessageListParamsPlatformGooglebusiness, InboxMessageListParamsPlatformWhatsapp, InboxMessageListParamsPlatformMastodon, InboxMessageListParamsPlatformDiscord, InboxMessageListParamsPlatformSMS:
		return true
	}
	return false
}

type InboxMessageEditParams struct {
	// Updated message text
	Text param.Field[string] `json:"text" api:"required"`
}

func (r InboxMessageEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboxMessageSendParams struct {
	// Account ID to send from
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Attachments
	Attachments param.Field[[]InboxMessageSendParamsAttachment] `json:"attachments"`
	// Message tag for sending outside the 24h window (Facebook only)
	MessageTag param.Field[InboxMessageSendParamsMessageTag] `json:"message_tag"`
	// Quick reply buttons (Facebook/Instagram, max 13)
	QuickReplies param.Field[[]InboxMessageSendParamsQuickReply] `json:"quick_replies"`
	// Message ID to reply to
	ReplyTo param.Field[string] `json:"reply_to"`
	// Structured template message (Facebook/Instagram)
	Template param.Field[InboxMessageSendParamsTemplate] `json:"template"`
	// Message text
	Text param.Field[string] `json:"text"`
}

func (r InboxMessageSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboxMessageSendParamsAttachment struct {
	// Attachment MIME type
	Type param.Field[string] `json:"type" api:"required"`
	// Attachment URL
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
}

func (r InboxMessageSendParamsAttachment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Message tag for sending outside the 24h window (Facebook only)
type InboxMessageSendParamsMessageTag string

const (
	InboxMessageSendParamsMessageTagHumanAgent       InboxMessageSendParamsMessageTag = "HUMAN_AGENT"
	InboxMessageSendParamsMessageTagCustomerFeedback InboxMessageSendParamsMessageTag = "CUSTOMER_FEEDBACK"
)

func (r InboxMessageSendParamsMessageTag) IsKnown() bool {
	switch r {
	case InboxMessageSendParamsMessageTagHumanAgent, InboxMessageSendParamsMessageTagCustomerFeedback:
		return true
	}
	return false
}

type InboxMessageSendParamsQuickReply struct {
	// Quick reply type
	ContentType param.Field[InboxMessageSendParamsQuickRepliesContentType] `json:"content_type"`
	// Icon URL for the button
	ImageURL param.Field[string] `json:"image_url" format:"uri"`
	// Postback payload
	Payload param.Field[string] `json:"payload"`
	// Button label (required for text type)
	Title param.Field[string] `json:"title"`
}

func (r InboxMessageSendParamsQuickReply) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Quick reply type
type InboxMessageSendParamsQuickRepliesContentType string

const (
	InboxMessageSendParamsQuickRepliesContentTypeText            InboxMessageSendParamsQuickRepliesContentType = "text"
	InboxMessageSendParamsQuickRepliesContentTypeUserPhoneNumber InboxMessageSendParamsQuickRepliesContentType = "user_phone_number"
	InboxMessageSendParamsQuickRepliesContentTypeUserEmail       InboxMessageSendParamsQuickRepliesContentType = "user_email"
)

func (r InboxMessageSendParamsQuickRepliesContentType) IsKnown() bool {
	switch r {
	case InboxMessageSendParamsQuickRepliesContentTypeText, InboxMessageSendParamsQuickRepliesContentTypeUserPhoneNumber, InboxMessageSendParamsQuickRepliesContentTypeUserEmail:
		return true
	}
	return false
}

// Structured template message (Facebook/Instagram)
type InboxMessageSendParamsTemplate struct {
	// Template elements (max 10 for carousel)
	Elements param.Field[[]InboxMessageSendParamsTemplateElement] `json:"elements" api:"required"`
	// Template type
	Type param.Field[InboxMessageSendParamsTemplateType] `json:"type" api:"required"`
}

func (r InboxMessageSendParamsTemplate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboxMessageSendParamsTemplateElement struct {
	// Element title
	Title param.Field[string] `json:"title" api:"required"`
	// Element buttons (max 3)
	Buttons param.Field[[]InboxMessageSendParamsTemplateElementsButton] `json:"buttons"`
	// Element image URL
	ImageURL param.Field[string] `json:"image_url" format:"uri"`
	// Element subtitle
	Subtitle param.Field[string] `json:"subtitle"`
}

func (r InboxMessageSendParamsTemplateElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboxMessageSendParamsTemplateElementsButton struct {
	// Button label
	Title param.Field[string] `json:"title" api:"required"`
	// Button type
	Type param.Field[InboxMessageSendParamsTemplateElementsButtonsType] `json:"type" api:"required"`
	// Payload for postback buttons
	Payload param.Field[string] `json:"payload"`
	// URL for web_url buttons
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r InboxMessageSendParamsTemplateElementsButton) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Button type
type InboxMessageSendParamsTemplateElementsButtonsType string

const (
	InboxMessageSendParamsTemplateElementsButtonsTypeWebURL   InboxMessageSendParamsTemplateElementsButtonsType = "web_url"
	InboxMessageSendParamsTemplateElementsButtonsTypePostback InboxMessageSendParamsTemplateElementsButtonsType = "postback"
)

func (r InboxMessageSendParamsTemplateElementsButtonsType) IsKnown() bool {
	switch r {
	case InboxMessageSendParamsTemplateElementsButtonsTypeWebURL, InboxMessageSendParamsTemplateElementsButtonsTypePostback:
		return true
	}
	return false
}

// Template type
type InboxMessageSendParamsTemplateType string

const (
	InboxMessageSendParamsTemplateTypeGeneric InboxMessageSendParamsTemplateType = "generic"
	InboxMessageSendParamsTemplateTypeButton  InboxMessageSendParamsTemplateType = "button"
)

func (r InboxMessageSendParamsTemplateType) IsKnown() bool {
	switch r {
	case InboxMessageSendParamsTemplateTypeGeneric, InboxMessageSendParamsTemplateTypeButton:
		return true
	}
	return false
}
