// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// QueueSlotService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueSlotService] method instead.
type QueueSlotService struct {
	Options []option.RequestOption
}

// NewQueueSlotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueueSlotService(opts ...option.RequestOption) (r *QueueSlotService) {
	r = &QueueSlotService{}
	r.Options = opts
	return
}

// Create a queue schedule
func (r *QueueSlotService) New(ctx context.Context, body QueueSlotNewParams, opts ...option.RequestOption) (res *QueueSlotNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/queue/slots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update queue schedule
func (r *QueueSlotService) Update(ctx context.Context, body QueueSlotUpdateParams, opts ...option.RequestOption) (res *QueueSlotUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/queue/slots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List queue schedules
func (r *QueueSlotService) List(ctx context.Context, opts ...option.RequestOption) (res *QueueSlotListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/queue/slots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete queue schedule
func (r *QueueSlotService) Delete(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/queue/slots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type QueueSlotNewResponse struct {
	// Queue schedule ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this is the default schedule
	IsDefault bool `json:"is_default" api:"required"`
	// Time slots
	Slots []QueueSlotNewResponseSlot `json:"slots" api:"required"`
	// Updated timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Schedule name
	Name string                   `json:"name" api:"nullable"`
	JSON queueSlotNewResponseJSON `json:"-"`
}

// queueSlotNewResponseJSON contains the JSON metadata for the struct
// [QueueSlotNewResponse]
type queueSlotNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsDefault   apijson.Field
	Slots       apijson.Field
	UpdatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotNewResponseJSON) RawJSON() string {
	return r.raw
}

type QueueSlotNewResponseSlot struct {
	// Day of week (0=Sunday, 6=Saturday)
	DayOfWeek int64 `json:"day_of_week" api:"required"`
	// Time in HH:MM format
	Time string `json:"time" api:"required"`
	// IANA timezone (e.g. America/New_York)
	Timezone string                       `json:"timezone" api:"required"`
	JSON     queueSlotNewResponseSlotJSON `json:"-"`
}

// queueSlotNewResponseSlotJSON contains the JSON metadata for the struct
// [QueueSlotNewResponseSlot]
type queueSlotNewResponseSlotJSON struct {
	DayOfWeek   apijson.Field
	Time        apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotNewResponseSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotNewResponseSlotJSON) RawJSON() string {
	return r.raw
}

type QueueSlotUpdateResponse struct {
	// Queue schedule ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this is the default schedule
	IsDefault bool `json:"is_default" api:"required"`
	// Time slots
	Slots []QueueSlotUpdateResponseSlot `json:"slots" api:"required"`
	// Updated timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Schedule name
	Name string                      `json:"name" api:"nullable"`
	JSON queueSlotUpdateResponseJSON `json:"-"`
}

// queueSlotUpdateResponseJSON contains the JSON metadata for the struct
// [QueueSlotUpdateResponse]
type queueSlotUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsDefault   apijson.Field
	Slots       apijson.Field
	UpdatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type QueueSlotUpdateResponseSlot struct {
	// Day of week (0=Sunday, 6=Saturday)
	DayOfWeek int64 `json:"day_of_week" api:"required"`
	// Time in HH:MM format
	Time string `json:"time" api:"required"`
	// IANA timezone (e.g. America/New_York)
	Timezone string                          `json:"timezone" api:"required"`
	JSON     queueSlotUpdateResponseSlotJSON `json:"-"`
}

// queueSlotUpdateResponseSlotJSON contains the JSON metadata for the struct
// [QueueSlotUpdateResponseSlot]
type queueSlotUpdateResponseSlotJSON struct {
	DayOfWeek   apijson.Field
	Time        apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotUpdateResponseSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotUpdateResponseSlotJSON) RawJSON() string {
	return r.raw
}

type QueueSlotListResponse struct {
	Data []QueueSlotListResponseData `json:"data" api:"required"`
	JSON queueSlotListResponseJSON   `json:"-"`
}

// queueSlotListResponseJSON contains the JSON metadata for the struct
// [QueueSlotListResponse]
type queueSlotListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotListResponseJSON) RawJSON() string {
	return r.raw
}

type QueueSlotListResponseData struct {
	// Queue schedule ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this is the default schedule
	IsDefault bool `json:"is_default" api:"required"`
	// Time slots
	Slots []QueueSlotListResponseDataSlot `json:"slots" api:"required"`
	// Updated timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Schedule name
	Name string                        `json:"name" api:"nullable"`
	JSON queueSlotListResponseDataJSON `json:"-"`
}

// queueSlotListResponseDataJSON contains the JSON metadata for the struct
// [QueueSlotListResponseData]
type queueSlotListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsDefault   apijson.Field
	Slots       apijson.Field
	UpdatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotListResponseDataJSON) RawJSON() string {
	return r.raw
}

type QueueSlotListResponseDataSlot struct {
	// Day of week (0=Sunday, 6=Saturday)
	DayOfWeek int64 `json:"day_of_week" api:"required"`
	// Time in HH:MM format
	Time string `json:"time" api:"required"`
	// IANA timezone (e.g. America/New_York)
	Timezone string                            `json:"timezone" api:"required"`
	JSON     queueSlotListResponseDataSlotJSON `json:"-"`
}

// queueSlotListResponseDataSlotJSON contains the JSON metadata for the struct
// [QueueSlotListResponseDataSlot]
type queueSlotListResponseDataSlotJSON struct {
	DayOfWeek   apijson.Field
	Time        apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueSlotListResponseDataSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueSlotListResponseDataSlotJSON) RawJSON() string {
	return r.raw
}

type QueueSlotNewParams struct {
	// Time slots
	Slots param.Field[[]QueueSlotNewParamsSlot] `json:"slots" api:"required"`
	// Default timezone for slots
	Timezone param.Field[string] `json:"timezone" api:"required"`
	// Schedule name
	Name param.Field[string] `json:"name"`
}

func (r QueueSlotNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueSlotNewParamsSlot struct {
	// Day of week (0=Sunday, 6=Saturday)
	DayOfWeek param.Field[int64] `json:"day_of_week" api:"required"`
	// Time in HH:MM format
	Time param.Field[string] `json:"time" api:"required"`
	// IANA timezone (e.g. America/New_York)
	Timezone param.Field[string] `json:"timezone" api:"required"`
}

func (r QueueSlotNewParamsSlot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueSlotUpdateParams struct {
	// Schedule name
	Name param.Field[string] `json:"name"`
	// Set this schedule as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
	// Updated time slots
	Slots param.Field[[]QueueSlotUpdateParamsSlot] `json:"slots"`
}

func (r QueueSlotUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueueSlotUpdateParamsSlot struct {
	// Day of week (0=Sunday, 6=Saturday)
	DayOfWeek param.Field[int64] `json:"day_of_week" api:"required"`
	// Time in HH:MM format
	Time param.Field[string] `json:"time" api:"required"`
	// IANA timezone (e.g. America/New_York)
	Timezone param.Field[string] `json:"timezone" api:"required"`
}

func (r QueueSlotUpdateParamsSlot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
