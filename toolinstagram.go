// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// ToolInstagramService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolInstagramService] method instead.
type ToolInstagramService struct {
	Options []option.RequestOption
}

// NewToolInstagramService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolInstagramService(opts ...option.RequestOption) (r *ToolInstagramService) {
	r = &ToolInstagramService{}
	r.Options = opts
	return
}

// Check Instagram hashtag safety status
func (r *ToolInstagramService) CheckHashtagSafety(ctx context.Context, body ToolInstagramCheckHashtagSafetyParams, opts ...option.RequestOption) (res *ToolInstagramCheckHashtagSafetyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/instagram/hashtag-checker"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ToolInstagramCheckHashtagSafetyResponse struct {
	Results []ToolInstagramCheckHashtagSafetyResponseResult `json:"results" api:"required"`
	JSON    toolInstagramCheckHashtagSafetyResponseJSON     `json:"-"`
}

// toolInstagramCheckHashtagSafetyResponseJSON contains the JSON metadata for the
// struct [ToolInstagramCheckHashtagSafetyResponse]
type toolInstagramCheckHashtagSafetyResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolInstagramCheckHashtagSafetyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolInstagramCheckHashtagSafetyResponseJSON) RawJSON() string {
	return r.raw
}

type ToolInstagramCheckHashtagSafetyResponseResult struct {
	// Hashtag checked
	Hashtag string `json:"hashtag" api:"required"`
	// Hashtag safety status
	Status ToolInstagramCheckHashtagSafetyResponseResultsStatus `json:"status" api:"required"`
	JSON   toolInstagramCheckHashtagSafetyResponseResultJSON    `json:"-"`
}

// toolInstagramCheckHashtagSafetyResponseResultJSON contains the JSON metadata for
// the struct [ToolInstagramCheckHashtagSafetyResponseResult]
type toolInstagramCheckHashtagSafetyResponseResultJSON struct {
	Hashtag     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolInstagramCheckHashtagSafetyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolInstagramCheckHashtagSafetyResponseResultJSON) RawJSON() string {
	return r.raw
}

// Hashtag safety status
type ToolInstagramCheckHashtagSafetyResponseResultsStatus string

const (
	ToolInstagramCheckHashtagSafetyResponseResultsStatusSafe       ToolInstagramCheckHashtagSafetyResponseResultsStatus = "safe"
	ToolInstagramCheckHashtagSafetyResponseResultsStatusRestricted ToolInstagramCheckHashtagSafetyResponseResultsStatus = "restricted"
	ToolInstagramCheckHashtagSafetyResponseResultsStatusBanned     ToolInstagramCheckHashtagSafetyResponseResultsStatus = "banned"
)

func (r ToolInstagramCheckHashtagSafetyResponseResultsStatus) IsKnown() bool {
	switch r {
	case ToolInstagramCheckHashtagSafetyResponseResultsStatusSafe, ToolInstagramCheckHashtagSafetyResponseResultsStatusRestricted, ToolInstagramCheckHashtagSafetyResponseResultsStatusBanned:
		return true
	}
	return false
}

type ToolInstagramCheckHashtagSafetyParams struct {
	// Hashtags to check (without # prefix)
	Hashtags param.Field[[]string] `json:"hashtags" api:"required"`
}

func (r ToolInstagramCheckHashtagSafetyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
