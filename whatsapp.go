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

// List purchased phone numbers
func (r *WhatsappService) ListPhoneNumbers(ctx context.Context, query WhatsappListPhoneNumbersParams, opts ...option.RequestOption) (res *WhatsappListPhoneNumbersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/phone-numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WhatsappBulkSendResponse struct {
	Results []WhatsappBulkSendResponseResult `json:"results" api:"required"`
	Summary WhatsappBulkSendResponseSummary  `json:"summary" api:"required"`
	JSON    whatsappBulkSendResponseJSON     `json:"-"`
}

// whatsappBulkSendResponseJSON contains the JSON metadata for the struct
// [WhatsappBulkSendResponse]
type whatsappBulkSendResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappBulkSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBulkSendResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappBulkSendResponseResult struct {
	// Recipient phone number
	Phone string `json:"phone" api:"required"`
	// Send status
	Status WhatsappBulkSendResponseResultsStatus `json:"status" api:"required"`
	// Error message if failed
	Error string                             `json:"error" api:"nullable"`
	JSON  whatsappBulkSendResponseResultJSON `json:"-"`
}

// whatsappBulkSendResponseResultJSON contains the JSON metadata for the struct
// [WhatsappBulkSendResponseResult]
type whatsappBulkSendResponseResultJSON struct {
	Phone       apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappBulkSendResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBulkSendResponseResultJSON) RawJSON() string {
	return r.raw
}

// Send status
type WhatsappBulkSendResponseResultsStatus string

const (
	WhatsappBulkSendResponseResultsStatusSent   WhatsappBulkSendResponseResultsStatus = "sent"
	WhatsappBulkSendResponseResultsStatusFailed WhatsappBulkSendResponseResultsStatus = "failed"
)

func (r WhatsappBulkSendResponseResultsStatus) IsKnown() bool {
	switch r {
	case WhatsappBulkSendResponseResultsStatusSent, WhatsappBulkSendResponseResultsStatusFailed:
		return true
	}
	return false
}

type WhatsappBulkSendResponseSummary struct {
	// Failed count
	Failed float64 `json:"failed" api:"required"`
	// Successfully sent count
	Sent float64                             `json:"sent" api:"required"`
	JSON whatsappBulkSendResponseSummaryJSON `json:"-"`
}

// whatsappBulkSendResponseSummaryJSON contains the JSON metadata for the struct
// [WhatsappBulkSendResponseSummary]
type whatsappBulkSendResponseSummaryJSON struct {
	Failed      apijson.Field
	Sent        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappBulkSendResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBulkSendResponseSummaryJSON) RawJSON() string {
	return r.raw
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
	// Phone number resource ID
	ID string `json:"id" api:"required"`
	// ISO country code
	Country string `json:"country" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Monthly cost in cents
	MonthlyCostCents float64 `json:"monthly_cost_cents" api:"required"`
	// E.164 phone number
	PhoneNumber string `json:"phone_number" api:"required"`
	// Carrier provider
	Provider string `json:"provider" api:"required"`
	// Provisioning status
	Status WhatsappListPhoneNumbersResponseDataStatus `json:"status" api:"required"`
	// Linked RelayAPI social account ID
	SocialAccountID string `json:"social_account_id" api:"nullable"`
	// Meta WhatsApp phone number ID
	WaPhoneNumberID string                                   `json:"wa_phone_number_id" api:"nullable"`
	JSON            whatsappListPhoneNumbersResponseDataJSON `json:"-"`
}

// whatsappListPhoneNumbersResponseDataJSON contains the JSON metadata for the
// struct [WhatsappListPhoneNumbersResponseData]
type whatsappListPhoneNumbersResponseDataJSON struct {
	ID               apijson.Field
	Country          apijson.Field
	CreatedAt        apijson.Field
	MonthlyCostCents apijson.Field
	PhoneNumber      apijson.Field
	Provider         apijson.Field
	Status           apijson.Field
	SocialAccountID  apijson.Field
	WaPhoneNumberID  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WhatsappListPhoneNumbersResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappListPhoneNumbersResponseDataJSON) RawJSON() string {
	return r.raw
}

// Provisioning status
type WhatsappListPhoneNumbersResponseDataStatus string

const (
	WhatsappListPhoneNumbersResponseDataStatusPurchasing          WhatsappListPhoneNumbersResponseDataStatus = "purchasing"
	WhatsappListPhoneNumbersResponseDataStatusPendingVerification WhatsappListPhoneNumbersResponseDataStatus = "pending_verification"
	WhatsappListPhoneNumbersResponseDataStatusVerified            WhatsappListPhoneNumbersResponseDataStatus = "verified"
	WhatsappListPhoneNumbersResponseDataStatusActive              WhatsappListPhoneNumbersResponseDataStatus = "active"
	WhatsappListPhoneNumbersResponseDataStatusReleasing           WhatsappListPhoneNumbersResponseDataStatus = "releasing"
	WhatsappListPhoneNumbersResponseDataStatusReleased            WhatsappListPhoneNumbersResponseDataStatus = "released"
)

func (r WhatsappListPhoneNumbersResponseDataStatus) IsKnown() bool {
	switch r {
	case WhatsappListPhoneNumbersResponseDataStatusPurchasing, WhatsappListPhoneNumbersResponseDataStatusPendingVerification, WhatsappListPhoneNumbersResponseDataStatusVerified, WhatsappListPhoneNumbersResponseDataStatusActive, WhatsappListPhoneNumbersResponseDataStatusReleasing, WhatsappListPhoneNumbersResponseDataStatusReleased:
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
	// Filter by provisioning status
	Status param.Field[WhatsappListPhoneNumbersParamsStatus] `query:"status"`
}

// URLQuery serializes [WhatsappListPhoneNumbersParams]'s query parameters as
// `url.Values`.
func (r WhatsappListPhoneNumbersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by provisioning status
type WhatsappListPhoneNumbersParamsStatus string

const (
	WhatsappListPhoneNumbersParamsStatusPurchasing          WhatsappListPhoneNumbersParamsStatus = "purchasing"
	WhatsappListPhoneNumbersParamsStatusPendingVerification WhatsappListPhoneNumbersParamsStatus = "pending_verification"
	WhatsappListPhoneNumbersParamsStatusVerified            WhatsappListPhoneNumbersParamsStatus = "verified"
	WhatsappListPhoneNumbersParamsStatusActive              WhatsappListPhoneNumbersParamsStatus = "active"
	WhatsappListPhoneNumbersParamsStatusReleasing           WhatsappListPhoneNumbersParamsStatus = "releasing"
	WhatsappListPhoneNumbersParamsStatusReleased            WhatsappListPhoneNumbersParamsStatus = "released"
)

func (r WhatsappListPhoneNumbersParamsStatus) IsKnown() bool {
	switch r {
	case WhatsappListPhoneNumbersParamsStatusPurchasing, WhatsappListPhoneNumbersParamsStatusPendingVerification, WhatsappListPhoneNumbersParamsStatusVerified, WhatsappListPhoneNumbersParamsStatusActive, WhatsappListPhoneNumbersParamsStatusReleasing, WhatsappListPhoneNumbersParamsStatusReleased:
		return true
	}
	return false
}
