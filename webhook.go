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

// WebhookService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Create a new webhook endpoint. The signing secret is returned only once in the
// response.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a webhook endpoint
func (r *WebhookService) Update(ctx context.Context, id string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List webhook endpoints
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a webhook endpoint
func (r *WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns delivery logs from the last 7 days.
func (r *WebhookService) ListLogs(ctx context.Context, query WebhookListLogsParams, opts ...option.RequestOption) (res *WebhookListLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Send a test POST request to the webhook URL to verify it is reachable.
func (r *WebhookService) SendTest(ctx context.Context, body WebhookSendTestParams, opts ...option.RequestOption) (res *WebhookSendTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks/test"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WebhookNewResponse struct {
	// Webhook ID
	ID string `json:"id" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events
	Events []string `json:"events" api:"required"`
	// Webhook signing secret (shown only once)
	Secret string `json:"secret" api:"required"`
	// Endpoint URL
	URL  string                 `json:"url" api:"required" format:"uri"`
	JSON webhookNewResponseJSON `json:"-"`
}

// webhookNewResponseJSON contains the JSON metadata for the struct
// [WebhookNewResponse]
type webhookNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Secret      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookUpdateResponse struct {
	// Webhook ID
	ID string `json:"id" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events
	Events []string `json:"events" api:"required"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Endpoint URL
	URL  string                    `json:"url" api:"required" format:"uri"`
	JSON webhookUpdateResponseJSON `json:"-"`
}

// webhookUpdateResponseJSON contains the JSON metadata for the struct
// [WebhookUpdateResponse]
type webhookUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookListResponse struct {
	Data []WebhookListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                  `json:"next_cursor" api:"required,nullable"`
	JSON       webhookListResponseJSON `json:"-"`
}

// webhookListResponseJSON contains the JSON metadata for the struct
// [WebhookListResponse]
type webhookListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookListResponseData struct {
	// Webhook ID
	ID string `json:"id" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events
	Events []string `json:"events" api:"required"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Endpoint URL
	URL  string                      `json:"url" api:"required" format:"uri"`
	JSON webhookListResponseDataJSON `json:"-"`
}

// webhookListResponseDataJSON contains the JSON metadata for the struct
// [WebhookListResponseData]
type webhookListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookListLogsResponse struct {
	Data       []WebhookListLogsResponseData `json:"data" api:"required"`
	HasMore    bool                          `json:"has_more" api:"required"`
	NextCursor string                        `json:"next_cursor" api:"required,nullable"`
	JSON       webhookListLogsResponseJSON   `json:"-"`
}

// webhookListLogsResponseJSON contains the JSON metadata for the struct
// [WebhookListLogsResponse]
type webhookListLogsResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListLogsResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookListLogsResponseData struct {
	ID             string                          `json:"id" api:"required"`
	CreatedAt      time.Time                       `json:"created_at" api:"required" format:"date-time"`
	Error          string                          `json:"error" api:"required,nullable"`
	Event          string                          `json:"event" api:"required"`
	ResponseTimeMs float64                         `json:"response_time_ms" api:"required,nullable"`
	StatusCode     float64                         `json:"status_code" api:"required,nullable"`
	Success        bool                            `json:"success" api:"required"`
	WebhookID      string                          `json:"webhook_id" api:"required"`
	JSON           webhookListLogsResponseDataJSON `json:"-"`
}

// webhookListLogsResponseDataJSON contains the JSON metadata for the struct
// [WebhookListLogsResponseData]
type webhookListLogsResponseDataJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Error          apijson.Field
	Event          apijson.Field
	ResponseTimeMs apijson.Field
	StatusCode     apijson.Field
	Success        apijson.Field
	WebhookID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WebhookListLogsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListLogsResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookSendTestResponse struct {
	// Response time in milliseconds
	ResponseTimeMs int64 `json:"response_time_ms" api:"required,nullable"`
	// HTTP status code from the test delivery
	StatusCode int64 `json:"status_code" api:"required,nullable"`
	// Whether the test delivery succeeded
	Success bool                        `json:"success" api:"required"`
	JSON    webhookSendTestResponseJSON `json:"-"`
}

// webhookSendTestResponseJSON contains the JSON metadata for the struct
// [WebhookSendTestResponse]
type webhookSendTestResponseJSON struct {
	ResponseTimeMs apijson.Field
	StatusCode     apijson.Field
	Success        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WebhookSendTestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookSendTestResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookNewParams struct {
	// Events to subscribe to
	Events param.Field[[]WebhookNewParamsEvent] `json:"events" api:"required"`
	// Webhook endpoint URL
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Workspace ID to scope this webhook to
	WorkspaceID param.Field[string] `json:"workspace_id"`
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookNewParamsEvent string

const (
	WebhookNewParamsEventPostPublished       WebhookNewParamsEvent = "post.published"
	WebhookNewParamsEventPostPartial         WebhookNewParamsEvent = "post.partial"
	WebhookNewParamsEventPostFailed          WebhookNewParamsEvent = "post.failed"
	WebhookNewParamsEventPostScheduled       WebhookNewParamsEvent = "post.scheduled"
	WebhookNewParamsEventAccountConnected    WebhookNewParamsEvent = "account.connected"
	WebhookNewParamsEventAccountDisconnected WebhookNewParamsEvent = "account.disconnected"
	WebhookNewParamsEventCommentReceived     WebhookNewParamsEvent = "comment.received"
	WebhookNewParamsEventMessageReceived     WebhookNewParamsEvent = "message.received"
)

func (r WebhookNewParamsEvent) IsKnown() bool {
	switch r {
	case WebhookNewParamsEventPostPublished, WebhookNewParamsEventPostPartial, WebhookNewParamsEventPostFailed, WebhookNewParamsEventPostScheduled, WebhookNewParamsEventAccountConnected, WebhookNewParamsEventAccountDisconnected, WebhookNewParamsEventCommentReceived, WebhookNewParamsEventMessageReceived:
		return true
	}
	return false
}

type WebhookUpdateParams struct {
	// Enable or disable the webhook
	Enabled param.Field[bool] `json:"enabled"`
	// Updated events
	Events param.Field[[]WebhookUpdateParamsEvent] `json:"events"`
	// Updated endpoint URL
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookUpdateParamsEvent string

const (
	WebhookUpdateParamsEventPostPublished       WebhookUpdateParamsEvent = "post.published"
	WebhookUpdateParamsEventPostPartial         WebhookUpdateParamsEvent = "post.partial"
	WebhookUpdateParamsEventPostFailed          WebhookUpdateParamsEvent = "post.failed"
	WebhookUpdateParamsEventPostScheduled       WebhookUpdateParamsEvent = "post.scheduled"
	WebhookUpdateParamsEventAccountConnected    WebhookUpdateParamsEvent = "account.connected"
	WebhookUpdateParamsEventAccountDisconnected WebhookUpdateParamsEvent = "account.disconnected"
	WebhookUpdateParamsEventCommentReceived     WebhookUpdateParamsEvent = "comment.received"
	WebhookUpdateParamsEventMessageReceived     WebhookUpdateParamsEvent = "message.received"
)

func (r WebhookUpdateParamsEvent) IsKnown() bool {
	switch r {
	case WebhookUpdateParamsEventPostPublished, WebhookUpdateParamsEventPostPartial, WebhookUpdateParamsEventPostFailed, WebhookUpdateParamsEventPostScheduled, WebhookUpdateParamsEventAccountConnected, WebhookUpdateParamsEventAccountDisconnected, WebhookUpdateParamsEventCommentReceived, WebhookUpdateParamsEventMessageReceived:
		return true
	}
	return false
}

type WebhookListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Filter: start date (ISO 8601)
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Filter: end date (ISO 8601)
	To param.Field[time.Time] `query:"to" format:"date-time"`
	// Filter by workspace ID
	WorkspaceID param.Field[string] `query:"workspace_id"`
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookListLogsParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Filter: start date (ISO 8601)
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Filter: end date (ISO 8601)
	To param.Field[time.Time] `query:"to" format:"date-time"`
}

// URLQuery serializes [WebhookListLogsParams]'s query parameters as `url.Values`.
func (r WebhookListLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookSendTestParams struct {
	// ID of the webhook to test
	WebhookID param.Field[string] `json:"webhook_id" api:"required"`
}

func (r WebhookSendTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
