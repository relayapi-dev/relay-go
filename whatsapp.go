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

// WhatsappService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappService] method instead.
type WhatsappService struct {
	Options         []option.RequestOption
	Broadcasts      *WhatsappBroadcastService
	Templates       *WhatsappTemplateService
	Contacts        *WhatsappContactService
	Groups          *WhatsappGroupService
	BusinessProfile *WhatsappBusinessProfileService
}

// NewWhatsappService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhatsappService(opts ...option.RequestOption) (r *WhatsappService) {
	r = &WhatsappService{}
	r.Options = opts
	r.Broadcasts = NewWhatsappBroadcastService(opts...)
	r.Templates = NewWhatsappTemplateService(opts...)
	r.Contacts = NewWhatsappContactService(opts...)
	r.Groups = NewWhatsappGroupService(opts...)
	r.BusinessProfile = NewWhatsappBusinessProfileService(opts...)
	return
}

// Send bulk WhatsApp messages via template
func (r *WhatsappService) BulkSend(ctx context.Context, body WhatsappBulkSendParams, opts ...option.RequestOption) (res *WhatsappBulkSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/bulk-send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List registered phone numbers
func (r *WhatsappService) ListPhoneNumbers(ctx context.Context, query WhatsappListPhoneNumbersParams, opts ...option.RequestOption) (res *WhatsappListPhoneNumbersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/phone-numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WhatsappBulkSendResponse struct {
	ID               string                         `json:"id" api:"required"`
	AccountID        string                         `json:"account_id" api:"required"`
	CompletedAt      time.Time                      `json:"completed_at" api:"required,nullable" format:"date-time"`
	CreatedAt        time.Time                      `json:"created_at" api:"required" format:"date-time"`
	Description      string                         `json:"description" api:"required,nullable"`
	FailedCount      int64                          `json:"failed_count" api:"required"`
	MessageText      string                         `json:"message_text" api:"required,nullable"`
	Name             string                         `json:"name" api:"required,nullable"`
	Platform         string                         `json:"platform" api:"required"`
	RecipientCount   int64                          `json:"recipient_count" api:"required"`
	ScheduledAt      time.Time                      `json:"scheduled_at" api:"required,nullable" format:"date-time"`
	SentCount        int64                          `json:"sent_count" api:"required"`
	Status           WhatsappBulkSendResponseStatus `json:"status" api:"required"`
	TemplateLanguage string                         `json:"template_language" api:"required,nullable"`
	TemplateName     string                         `json:"template_name" api:"required,nullable"`
	JSON             whatsappBulkSendResponseJSON   `json:"-"`
}

// whatsappBulkSendResponseJSON contains the JSON metadata for the struct
// [WhatsappBulkSendResponse]
type whatsappBulkSendResponseJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	CompletedAt      apijson.Field
	CreatedAt        apijson.Field
	Description      apijson.Field
	FailedCount      apijson.Field
	MessageText      apijson.Field
	Name             apijson.Field
	Platform         apijson.Field
	RecipientCount   apijson.Field
	ScheduledAt      apijson.Field
	SentCount        apijson.Field
	Status           apijson.Field
	TemplateLanguage apijson.Field
	TemplateName     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WhatsappBulkSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBulkSendResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappBulkSendResponseStatus string

const (
	WhatsappBulkSendResponseStatusDraft           WhatsappBulkSendResponseStatus = "draft"
	WhatsappBulkSendResponseStatusScheduled       WhatsappBulkSendResponseStatus = "scheduled"
	WhatsappBulkSendResponseStatusSending         WhatsappBulkSendResponseStatus = "sending"
	WhatsappBulkSendResponseStatusSent            WhatsappBulkSendResponseStatus = "sent"
	WhatsappBulkSendResponseStatusPartiallyFailed WhatsappBulkSendResponseStatus = "partially_failed"
	WhatsappBulkSendResponseStatusFailed          WhatsappBulkSendResponseStatus = "failed"
	WhatsappBulkSendResponseStatusCancelled       WhatsappBulkSendResponseStatus = "cancelled"
)

func (r WhatsappBulkSendResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappBulkSendResponseStatusDraft, WhatsappBulkSendResponseStatusScheduled, WhatsappBulkSendResponseStatusSending, WhatsappBulkSendResponseStatusSent, WhatsappBulkSendResponseStatusPartiallyFailed, WhatsappBulkSendResponseStatusFailed, WhatsappBulkSendResponseStatusCancelled:
		return true
	}
	return false
}

type WhatsappListPhoneNumbersResponse struct {
	Data []WhatsappListPhoneNumbersResponseData `json:"data" api:"required"`
	JSON whatsappListPhoneNumbersResponseJSON   `json:"-"`
}

// whatsappListPhoneNumbersResponseJSON contains the JSON metadata for the struct
// [WhatsappListPhoneNumbersResponse]
type whatsappListPhoneNumbersResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappListPhoneNumbersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappListPhoneNumbersResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappListPhoneNumbersResponseData struct {
	// Phone number ID
	ID string `json:"id" api:"required"`
	// Phone number
	PhoneNumber string `json:"phone_number" api:"required"`
	// Registration status
	Status WhatsappListPhoneNumbersResponseDataStatus `json:"status" api:"required"`
	// Display name
	DisplayName string                                   `json:"display_name" api:"nullable"`
	JSON        whatsappListPhoneNumbersResponseDataJSON `json:"-"`
}

// whatsappListPhoneNumbersResponseDataJSON contains the JSON metadata for the
// struct [WhatsappListPhoneNumbersResponseData]
type whatsappListPhoneNumbersResponseDataJSON struct {
	ID          apijson.Field
	PhoneNumber apijson.Field
	Status      apijson.Field
	DisplayName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappListPhoneNumbersResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappListPhoneNumbersResponseDataJSON) RawJSON() string {
	return r.raw
}

// Registration status
type WhatsappListPhoneNumbersResponseDataStatus string

const (
	WhatsappListPhoneNumbersResponseDataStatusActive   WhatsappListPhoneNumbersResponseDataStatus = "active"
	WhatsappListPhoneNumbersResponseDataStatusInactive WhatsappListPhoneNumbersResponseDataStatus = "inactive"
	WhatsappListPhoneNumbersResponseDataStatusPending  WhatsappListPhoneNumbersResponseDataStatus = "pending"
)

func (r WhatsappListPhoneNumbersResponseDataStatus) IsKnown() bool {
	switch r {
	case WhatsappListPhoneNumbersResponseDataStatusActive, WhatsappListPhoneNumbersResponseDataStatusInactive, WhatsappListPhoneNumbersResponseDataStatusPending:
		return true
	}
	return false
}

type WhatsappBulkSendParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Recipients
	Recipients param.Field[[]WhatsappBulkSendParamsRecipient] `json:"recipients" api:"required"`
	Template   param.Field[WhatsappBulkSendParamsTemplate]    `json:"template" api:"required"`
}

func (r WhatsappBulkSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappBulkSendParamsRecipient struct {
	// Phone number in E.164 format
	Phone param.Field[string] `json:"phone" api:"required"`
	// Template variable substitutions
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r WhatsappBulkSendParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappBulkSendParamsTemplate struct {
	// Template language code
	Language param.Field[string] `json:"language" api:"required"`
	// Template name
	Name param.Field[string] `json:"name" api:"required"`
	// Template components
	Components param.Field[[]WhatsappBulkSendParamsTemplateComponent] `json:"components"`
}

func (r WhatsappBulkSendParamsTemplate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappBulkSendParamsTemplateComponent struct {
	// Component type
	Type param.Field[WhatsappBulkSendParamsTemplateComponentsType] `json:"type" api:"required"`
	// Component parameters
	Parameters param.Field[[]map[string]interface{}] `json:"parameters"`
}

func (r WhatsappBulkSendParamsTemplateComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Component type
type WhatsappBulkSendParamsTemplateComponentsType string

const (
	WhatsappBulkSendParamsTemplateComponentsTypeHeader WhatsappBulkSendParamsTemplateComponentsType = "header"
	WhatsappBulkSendParamsTemplateComponentsTypeBody   WhatsappBulkSendParamsTemplateComponentsType = "body"
	WhatsappBulkSendParamsTemplateComponentsTypeButton WhatsappBulkSendParamsTemplateComponentsType = "button"
)

func (r WhatsappBulkSendParamsTemplateComponentsType) IsKnown() bool {
	switch r {
	case WhatsappBulkSendParamsTemplateComponentsTypeHeader, WhatsappBulkSendParamsTemplateComponentsTypeBody, WhatsappBulkSendParamsTemplateComponentsTypeButton:
		return true
	}
	return false
}

type WhatsappListPhoneNumbersParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappListPhoneNumbersParams]'s query parameters as
// `url.Values`.
func (r WhatsappListPhoneNumbersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
