// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// UsageService contains methods and other services that help with interacting with
// the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
	r = &UsageService{}
	r.Options = opts
	return
}

// Returns current subscription details and usage statistics for the organization.
func (r *UsageService) Get(ctx context.Context, opts ...option.RequestOption) (res *UsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UsageGetResponse struct {
	APICalls UsageGetResponseAPICalls `json:"api_calls" api:"required"`
	Plan     UsageGetResponsePlan     `json:"plan" api:"required"`
	Usage    UsageGetResponseUsage    `json:"usage" api:"required"`
	JSON     usageGetResponseJSON     `json:"-"`
}

// usageGetResponseJSON contains the JSON metadata for the struct
// [UsageGetResponse]
type usageGetResponseJSON struct {
	APICalls    apijson.Field
	Plan        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponseAPICalls struct {
	// API calls in the current minute
	CurrentMinute float64 `json:"current_minute" api:"required"`
	// Max API calls per minute
	LimitPerMinute float64                      `json:"limit_per_minute" api:"required"`
	JSON           usageGetResponseAPICallsJSON `json:"-"`
}

// usageGetResponseAPICallsJSON contains the JSON metadata for the struct
// [UsageGetResponseAPICalls]
type usageGetResponseAPICallsJSON struct {
	CurrentMinute  apijson.Field
	LimitPerMinute apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UsageGetResponseAPICalls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseAPICallsJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponsePlan struct {
	// API calls allowed per minute
	APICallsPerMin float64 `json:"api_calls_per_min" api:"required"`
	// Plan name
	Name string `json:"name" api:"required"`
	// Max posts per billing cycle
	PostsLimit float64                  `json:"posts_limit" api:"required"`
	JSON       usageGetResponsePlanJSON `json:"-"`
}

// usageGetResponsePlanJSON contains the JSON metadata for the struct
// [UsageGetResponsePlan]
type usageGetResponsePlanJSON struct {
	APICallsPerMin apijson.Field
	Name           apijson.Field
	PostsLimit     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UsageGetResponsePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponsePlanJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponseUsage struct {
	// Current billing cycle end
	CycleEnd time.Time `json:"cycle_end" api:"required" format:"date-time"`
	// When the cycle resets
	CycleResetsAt time.Time `json:"cycle_resets_at" api:"required" format:"date-time"`
	// Current billing cycle start
	CycleStart time.Time `json:"cycle_start" api:"required" format:"date-time"`
	// Max posts per billing cycle
	PostsLimit float64 `json:"posts_limit" api:"required"`
	// Posts used this cycle
	PostsUsed float64                   `json:"posts_used" api:"required"`
	JSON      usageGetResponseUsageJSON `json:"-"`
}

// usageGetResponseUsageJSON contains the JSON metadata for the struct
// [UsageGetResponseUsage]
type usageGetResponseUsageJSON struct {
	CycleEnd      apijson.Field
	CycleResetsAt apijson.Field
	CycleStart    apijson.Field
	PostsLimit    apijson.Field
	PostsUsed     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UsageGetResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseUsageJSON) RawJSON() string {
	return r.raw
}
