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

// Returns current plan details and API call usage statistics for the organization.
func (r *UsageService) Get(ctx context.Context, opts ...option.RequestOption) (res *UsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UsageGetResponse struct {
	Plan         UsageGetResponsePlan         `json:"plan" api:"required"`
	RateLimit    UsageGetResponseRateLimit    `json:"rate_limit" api:"required"`
	Subscription UsageGetResponseSubscription `json:"subscription" api:"required"`
	Usage        UsageGetResponseUsage        `json:"usage" api:"required"`
	JSON         usageGetResponseJSON         `json:"-"`
}

// usageGetResponseJSON contains the JSON metadata for the struct
// [UsageGetResponse]
type usageGetResponseJSON struct {
	Plan         apijson.Field
	RateLimit    apijson.Field
	Subscription apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UsageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponsePlan struct {
	// API calls included per billing cycle
	APICallsLimit float64 `json:"api_calls_limit" api:"required"`
	// API calls allowed per minute
	APICallsPerMin float64                      `json:"api_calls_per_min" api:"required"`
	Features       UsageGetResponsePlanFeatures `json:"features" api:"required"`
	// Current plan
	Name UsageGetResponsePlanName `json:"name" api:"required"`
	JSON usageGetResponsePlanJSON `json:"-"`
}

// usageGetResponsePlanJSON contains the JSON metadata for the struct
// [UsageGetResponsePlan]
type usageGetResponsePlanJSON struct {
	APICallsLimit  apijson.Field
	APICallsPerMin apijson.Field
	Features       apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UsageGetResponsePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponsePlanJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponsePlanFeatures struct {
	// Access to /v1/analytics
	Analytics bool `json:"analytics" api:"required"`
	// Access to /v1/inbox
	Inbox bool                             `json:"inbox" api:"required"`
	JSON  usageGetResponsePlanFeaturesJSON `json:"-"`
}

// usageGetResponsePlanFeaturesJSON contains the JSON metadata for the struct
// [UsageGetResponsePlanFeatures]
type usageGetResponsePlanFeaturesJSON struct {
	Analytics   apijson.Field
	Inbox       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponsePlanFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponsePlanFeaturesJSON) RawJSON() string {
	return r.raw
}

// Current plan
type UsageGetResponsePlanName string

const (
	UsageGetResponsePlanNameFree UsageGetResponsePlanName = "free"
	UsageGetResponsePlanNamePro  UsageGetResponsePlanName = "pro"
)

func (r UsageGetResponsePlanName) IsKnown() bool {
	switch r {
	case UsageGetResponsePlanNameFree, UsageGetResponsePlanNamePro:
		return true
	}
	return false
}

type UsageGetResponseRateLimit struct {
	// Max API calls per rate-limit window
	LimitPerMinute float64                       `json:"limit_per_minute" api:"required"`
	JSON           usageGetResponseRateLimitJSON `json:"-"`
}

// usageGetResponseRateLimitJSON contains the JSON metadata for the struct
// [UsageGetResponseRateLimit]
type usageGetResponseRateLimitJSON struct {
	LimitPerMinute apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UsageGetResponseRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseRateLimitJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponseSubscription struct {
	// Base monthly price in cents
	MonthlyPriceCents float64 `json:"monthly_price_cents" api:"required"`
	// Overage price per 1K API calls in cents
	PricePerThousandCallsCents float64 `json:"price_per_thousand_calls_cents" api:"required"`
	// Subscription status
	Status string                           `json:"status" api:"required"`
	JSON   usageGetResponseSubscriptionJSON `json:"-"`
}

// usageGetResponseSubscriptionJSON contains the JSON metadata for the struct
// [UsageGetResponseSubscription]
type usageGetResponseSubscriptionJSON struct {
	MonthlyPriceCents          apijson.Field
	PricePerThousandCallsCents apijson.Field
	Status                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UsageGetResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponseUsage struct {
	// API calls remaining this cycle. Null for pro plan (unlimited, overage billed).
	APICallsRemaining float64 `json:"api_calls_remaining" api:"required,nullable"`
	// API calls used this cycle
	APICallsUsed float64 `json:"api_calls_used" api:"required"`
	// Current billing cycle end
	CycleEnd time.Time `json:"cycle_end" api:"required" format:"date-time"`
	// Current billing cycle start
	CycleStart time.Time `json:"cycle_start" api:"required" format:"date-time"`
	// API calls exceeding included amount
	OverageCalls float64 `json:"overage_calls" api:"required"`
	// Overage cost in cents
	OverageCostCents float64                   `json:"overage_cost_cents" api:"required"`
	JSON             usageGetResponseUsageJSON `json:"-"`
}

// usageGetResponseUsageJSON contains the JSON metadata for the struct
// [UsageGetResponseUsage]
type usageGetResponseUsageJSON struct {
	APICallsRemaining apijson.Field
	APICallsUsed      apijson.Field
	CycleEnd          apijson.Field
	CycleStart        apijson.Field
	OverageCalls      apijson.Field
	OverageCostCents  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UsageGetResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseUsageJSON) RawJSON() string {
	return r.raw
}
