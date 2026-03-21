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

// QueueService contains methods and other services that help with interacting with
// the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueService] method instead.
type QueueService struct {
	Options []option.RequestOption
	Slots   *QueueSlotService
}

// NewQueueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueueService(opts ...option.RequestOption) (r *QueueService) {
	r = &QueueService{}
	r.Options = opts
	r.Slots = NewQueueSlotService(opts...)
	return
}

// Get next available queue slot
func (r *QueueService) GetNextSlot(ctx context.Context, opts ...option.RequestOption) (res *QueueGetNextSlotResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/queue/next-slot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Preview upcoming queue slots
func (r *QueueService) Preview(ctx context.Context, query QueuePreviewParams, opts ...option.RequestOption) (res *QueuePreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/queue/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type QueueGetNextSlotResponse struct {
	// Next available slot (ISO 8601)
	NextSlotAt time.Time `json:"next_slot_at" api:"required" format:"date-time"`
	// Queue schedule ID
	QueueID string                       `json:"queue_id" api:"required"`
	JSON    queueGetNextSlotResponseJSON `json:"-"`
}

// queueGetNextSlotResponseJSON contains the JSON metadata for the struct
// [QueueGetNextSlotResponse]
type queueGetNextSlotResponseJSON struct {
	NextSlotAt  apijson.Field
	QueueID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueGetNextSlotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueGetNextSlotResponseJSON) RawJSON() string {
	return r.raw
}

type QueuePreviewResponse struct {
	// Upcoming slot timestamps (ISO 8601)
	Slots []time.Time              `json:"slots" api:"required" format:"date-time"`
	JSON  queuePreviewResponseJSON `json:"-"`
}

// queuePreviewResponseJSON contains the JSON metadata for the struct
// [QueuePreviewResponse]
type queuePreviewResponseJSON struct {
	Slots       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueuePreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queuePreviewResponseJSON) RawJSON() string {
	return r.raw
}

type QueuePreviewParams struct {
	// Number of upcoming slots to preview
	Count param.Field[int64] `query:"count"`
}

// URLQuery serializes [QueuePreviewParams]'s query parameters as `url.Values`.
func (r QueuePreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
