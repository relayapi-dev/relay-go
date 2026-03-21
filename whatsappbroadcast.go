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

// WhatsappBroadcastService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappBroadcastService] method instead.
type WhatsappBroadcastService struct {
	Options []option.RequestOption
}

// NewWhatsappBroadcastService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappBroadcastService(opts ...option.RequestOption) (r *WhatsappBroadcastService) {
	r = &WhatsappBroadcastService{}
	r.Options = opts
	return
}

// Create a broadcast
func (r *WhatsappBroadcastService) New(ctx context.Context, body WhatsappBroadcastNewParams, opts ...option.RequestOption) (res *WhatsappBroadcastNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/broadcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get broadcast details
func (r *WhatsappBroadcastService) Get(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *WhatsappBroadcastGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcast_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/whatsapp/broadcasts/%s", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List broadcasts
func (r *WhatsappBroadcastService) List(ctx context.Context, query WhatsappBroadcastListParams, opts ...option.RequestOption) (res *WhatsappBroadcastListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/broadcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a broadcast
func (r *WhatsappBroadcastService) Delete(ctx context.Context, broadcastID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if broadcastID == "" {
		err = errors.New("missing required broadcast_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/whatsapp/broadcasts/%s", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Schedule a broadcast
func (r *WhatsappBroadcastService) Schedule(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *WhatsappBroadcastScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcast_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/whatsapp/broadcasts/%s/schedule", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Send a broadcast immediately
func (r *WhatsappBroadcastService) Send(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *WhatsappBroadcastSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcast_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/whatsapp/broadcasts/%s/send", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WhatsappBroadcastNewResponse struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Broadcast name
	Name string `json:"name" api:"required"`
	// Total recipients
	RecipientCount float64 `json:"recipient_count" api:"required"`
	// Broadcast status
	Status WhatsappBroadcastNewResponseStatus `json:"status" api:"required"`
	// Template name
	Template string `json:"template" api:"required"`
	// Failed sends
	Failed float64 `json:"failed"`
	// Scheduled time
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// Successfully sent
	Sent float64                          `json:"sent"`
	JSON whatsappBroadcastNewResponseJSON `json:"-"`
}

// whatsappBroadcastNewResponseJSON contains the JSON metadata for the struct
// [WhatsappBroadcastNewResponse]
type whatsappBroadcastNewResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	RecipientCount apijson.Field
	Status         apijson.Field
	Template       apijson.Field
	Failed         apijson.Field
	ScheduledAt    apijson.Field
	Sent           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WhatsappBroadcastNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBroadcastNewResponseJSON) RawJSON() string {
	return r.raw
}

// Broadcast status
type WhatsappBroadcastNewResponseStatus string

const (
	WhatsappBroadcastNewResponseStatusDraft     WhatsappBroadcastNewResponseStatus = "draft"
	WhatsappBroadcastNewResponseStatusScheduled WhatsappBroadcastNewResponseStatus = "scheduled"
	WhatsappBroadcastNewResponseStatusSending   WhatsappBroadcastNewResponseStatus = "sending"
	WhatsappBroadcastNewResponseStatusSent      WhatsappBroadcastNewResponseStatus = "sent"
	WhatsappBroadcastNewResponseStatusFailed    WhatsappBroadcastNewResponseStatus = "failed"
)

func (r WhatsappBroadcastNewResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappBroadcastNewResponseStatusDraft, WhatsappBroadcastNewResponseStatusScheduled, WhatsappBroadcastNewResponseStatusSending, WhatsappBroadcastNewResponseStatusSent, WhatsappBroadcastNewResponseStatusFailed:
		return true
	}
	return false
}

type WhatsappBroadcastGetResponse struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Broadcast name
	Name string `json:"name" api:"required"`
	// Total recipients
	RecipientCount float64 `json:"recipient_count" api:"required"`
	// Broadcast status
	Status WhatsappBroadcastGetResponseStatus `json:"status" api:"required"`
	// Template name
	Template string `json:"template" api:"required"`
	// Failed sends
	Failed float64 `json:"failed"`
	// Scheduled time
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// Successfully sent
	Sent float64                          `json:"sent"`
	JSON whatsappBroadcastGetResponseJSON `json:"-"`
}

// whatsappBroadcastGetResponseJSON contains the JSON metadata for the struct
// [WhatsappBroadcastGetResponse]
type whatsappBroadcastGetResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	RecipientCount apijson.Field
	Status         apijson.Field
	Template       apijson.Field
	Failed         apijson.Field
	ScheduledAt    apijson.Field
	Sent           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WhatsappBroadcastGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBroadcastGetResponseJSON) RawJSON() string {
	return r.raw
}

// Broadcast status
type WhatsappBroadcastGetResponseStatus string

const (
	WhatsappBroadcastGetResponseStatusDraft     WhatsappBroadcastGetResponseStatus = "draft"
	WhatsappBroadcastGetResponseStatusScheduled WhatsappBroadcastGetResponseStatus = "scheduled"
	WhatsappBroadcastGetResponseStatusSending   WhatsappBroadcastGetResponseStatus = "sending"
	WhatsappBroadcastGetResponseStatusSent      WhatsappBroadcastGetResponseStatus = "sent"
	WhatsappBroadcastGetResponseStatusFailed    WhatsappBroadcastGetResponseStatus = "failed"
)

func (r WhatsappBroadcastGetResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappBroadcastGetResponseStatusDraft, WhatsappBroadcastGetResponseStatusScheduled, WhatsappBroadcastGetResponseStatusSending, WhatsappBroadcastGetResponseStatusSent, WhatsappBroadcastGetResponseStatusFailed:
		return true
	}
	return false
}

type WhatsappBroadcastListResponse struct {
	Data []WhatsappBroadcastListResponseData `json:"data" api:"required"`
	JSON whatsappBroadcastListResponseJSON   `json:"-"`
}

// whatsappBroadcastListResponseJSON contains the JSON metadata for the struct
// [WhatsappBroadcastListResponse]
type whatsappBroadcastListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappBroadcastListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBroadcastListResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappBroadcastListResponseData struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Broadcast name
	Name string `json:"name" api:"required"`
	// Total recipients
	RecipientCount float64 `json:"recipient_count" api:"required"`
	// Broadcast status
	Status WhatsappBroadcastListResponseDataStatus `json:"status" api:"required"`
	// Template name
	Template string `json:"template" api:"required"`
	// Failed sends
	Failed float64 `json:"failed"`
	// Scheduled time
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// Successfully sent
	Sent float64                               `json:"sent"`
	JSON whatsappBroadcastListResponseDataJSON `json:"-"`
}

// whatsappBroadcastListResponseDataJSON contains the JSON metadata for the struct
// [WhatsappBroadcastListResponseData]
type whatsappBroadcastListResponseDataJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	RecipientCount apijson.Field
	Status         apijson.Field
	Template       apijson.Field
	Failed         apijson.Field
	ScheduledAt    apijson.Field
	Sent           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WhatsappBroadcastListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBroadcastListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Broadcast status
type WhatsappBroadcastListResponseDataStatus string

const (
	WhatsappBroadcastListResponseDataStatusDraft     WhatsappBroadcastListResponseDataStatus = "draft"
	WhatsappBroadcastListResponseDataStatusScheduled WhatsappBroadcastListResponseDataStatus = "scheduled"
	WhatsappBroadcastListResponseDataStatusSending   WhatsappBroadcastListResponseDataStatus = "sending"
	WhatsappBroadcastListResponseDataStatusSent      WhatsappBroadcastListResponseDataStatus = "sent"
	WhatsappBroadcastListResponseDataStatusFailed    WhatsappBroadcastListResponseDataStatus = "failed"
)

func (r WhatsappBroadcastListResponseDataStatus) IsKnown() bool {
	switch r {
	case WhatsappBroadcastListResponseDataStatusDraft, WhatsappBroadcastListResponseDataStatusScheduled, WhatsappBroadcastListResponseDataStatusSending, WhatsappBroadcastListResponseDataStatusSent, WhatsappBroadcastListResponseDataStatusFailed:
		return true
	}
	return false
}

type WhatsappBroadcastScheduleResponse struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Broadcast name
	Name string `json:"name" api:"required"`
	// Total recipients
	RecipientCount float64 `json:"recipient_count" api:"required"`
	// Broadcast status
	Status WhatsappBroadcastScheduleResponseStatus `json:"status" api:"required"`
	// Template name
	Template string `json:"template" api:"required"`
	// Failed sends
	Failed float64 `json:"failed"`
	// Scheduled time
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// Successfully sent
	Sent float64                               `json:"sent"`
	JSON whatsappBroadcastScheduleResponseJSON `json:"-"`
}

// whatsappBroadcastScheduleResponseJSON contains the JSON metadata for the struct
// [WhatsappBroadcastScheduleResponse]
type whatsappBroadcastScheduleResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	RecipientCount apijson.Field
	Status         apijson.Field
	Template       apijson.Field
	Failed         apijson.Field
	ScheduledAt    apijson.Field
	Sent           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WhatsappBroadcastScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBroadcastScheduleResponseJSON) RawJSON() string {
	return r.raw
}

// Broadcast status
type WhatsappBroadcastScheduleResponseStatus string

const (
	WhatsappBroadcastScheduleResponseStatusDraft     WhatsappBroadcastScheduleResponseStatus = "draft"
	WhatsappBroadcastScheduleResponseStatusScheduled WhatsappBroadcastScheduleResponseStatus = "scheduled"
	WhatsappBroadcastScheduleResponseStatusSending   WhatsappBroadcastScheduleResponseStatus = "sending"
	WhatsappBroadcastScheduleResponseStatusSent      WhatsappBroadcastScheduleResponseStatus = "sent"
	WhatsappBroadcastScheduleResponseStatusFailed    WhatsappBroadcastScheduleResponseStatus = "failed"
)

func (r WhatsappBroadcastScheduleResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappBroadcastScheduleResponseStatusDraft, WhatsappBroadcastScheduleResponseStatusScheduled, WhatsappBroadcastScheduleResponseStatusSending, WhatsappBroadcastScheduleResponseStatusSent, WhatsappBroadcastScheduleResponseStatusFailed:
		return true
	}
	return false
}

type WhatsappBroadcastSendResponse struct {
	// Broadcast ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Broadcast name
	Name string `json:"name" api:"required"`
	// Total recipients
	RecipientCount float64 `json:"recipient_count" api:"required"`
	// Broadcast status
	Status WhatsappBroadcastSendResponseStatus `json:"status" api:"required"`
	// Template name
	Template string `json:"template" api:"required"`
	// Failed sends
	Failed float64 `json:"failed"`
	// Scheduled time
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// Successfully sent
	Sent float64                           `json:"sent"`
	JSON whatsappBroadcastSendResponseJSON `json:"-"`
}

// whatsappBroadcastSendResponseJSON contains the JSON metadata for the struct
// [WhatsappBroadcastSendResponse]
type whatsappBroadcastSendResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	RecipientCount apijson.Field
	Status         apijson.Field
	Template       apijson.Field
	Failed         apijson.Field
	ScheduledAt    apijson.Field
	Sent           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WhatsappBroadcastSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBroadcastSendResponseJSON) RawJSON() string {
	return r.raw
}

// Broadcast status
type WhatsappBroadcastSendResponseStatus string

const (
	WhatsappBroadcastSendResponseStatusDraft     WhatsappBroadcastSendResponseStatus = "draft"
	WhatsappBroadcastSendResponseStatusScheduled WhatsappBroadcastSendResponseStatus = "scheduled"
	WhatsappBroadcastSendResponseStatusSending   WhatsappBroadcastSendResponseStatus = "sending"
	WhatsappBroadcastSendResponseStatusSent      WhatsappBroadcastSendResponseStatus = "sent"
	WhatsappBroadcastSendResponseStatusFailed    WhatsappBroadcastSendResponseStatus = "failed"
)

func (r WhatsappBroadcastSendResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappBroadcastSendResponseStatusDraft, WhatsappBroadcastSendResponseStatusScheduled, WhatsappBroadcastSendResponseStatusSending, WhatsappBroadcastSendResponseStatusSent, WhatsappBroadcastSendResponseStatusFailed:
		return true
	}
	return false
}

type WhatsappBroadcastNewParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Broadcast name
	Name param.Field[string] `json:"name" api:"required"`
	// Recipient list
	Recipients param.Field[[]WhatsappBroadcastNewParamsRecipient] `json:"recipients" api:"required"`
	Template   param.Field[WhatsappBroadcastNewParamsTemplate]    `json:"template" api:"required"`
	// ISO 8601 timestamp to schedule send
	ScheduledAt param.Field[string] `json:"scheduled_at"`
}

func (r WhatsappBroadcastNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappBroadcastNewParamsRecipient struct {
	// Phone number in E.164 format
	Phone param.Field[string] `json:"phone" api:"required"`
	// Template variable substitutions
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r WhatsappBroadcastNewParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappBroadcastNewParamsTemplate struct {
	// Template language code
	Language param.Field[string] `json:"language" api:"required"`
	// Template name
	Name       param.Field[string]                                        `json:"name" api:"required"`
	Components param.Field[[]WhatsappBroadcastNewParamsTemplateComponent] `json:"components"`
}

func (r WhatsappBroadcastNewParamsTemplate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappBroadcastNewParamsTemplateComponent struct {
	// Component type
	Type param.Field[WhatsappBroadcastNewParamsTemplateComponentsType] `json:"type" api:"required"`
	// Component parameters
	Parameters param.Field[[]map[string]interface{}] `json:"parameters"`
}

func (r WhatsappBroadcastNewParamsTemplateComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Component type
type WhatsappBroadcastNewParamsTemplateComponentsType string

const (
	WhatsappBroadcastNewParamsTemplateComponentsTypeHeader WhatsappBroadcastNewParamsTemplateComponentsType = "header"
	WhatsappBroadcastNewParamsTemplateComponentsTypeBody   WhatsappBroadcastNewParamsTemplateComponentsType = "body"
	WhatsappBroadcastNewParamsTemplateComponentsTypeButton WhatsappBroadcastNewParamsTemplateComponentsType = "button"
)

func (r WhatsappBroadcastNewParamsTemplateComponentsType) IsKnown() bool {
	switch r {
	case WhatsappBroadcastNewParamsTemplateComponentsTypeHeader, WhatsappBroadcastNewParamsTemplateComponentsTypeBody, WhatsappBroadcastNewParamsTemplateComponentsTypeButton:
		return true
	}
	return false
}

type WhatsappBroadcastListParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappBroadcastListParams]'s query parameters as
// `url.Values`.
func (r WhatsappBroadcastListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
