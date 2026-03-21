// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// AccountRedditFlairService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRedditFlairService] method instead.
type AccountRedditFlairService struct {
	Options []option.RequestOption
}

// NewAccountRedditFlairService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRedditFlairService(opts ...option.RequestOption) (r *AccountRedditFlairService) {
	r = &AccountRedditFlairService{}
	r.Options = opts
	return
}

// Fetch Reddit flairs for a subreddit
func (r *AccountRedditFlairService) Get(ctx context.Context, id string, query AccountRedditFlairGetParams, opts ...option.RequestOption) (res *AccountRedditFlairGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/reddit-flairs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AccountRedditFlairGetResponse struct {
	Data []AccountRedditFlairGetResponseData `json:"data" api:"required"`
	JSON accountRedditFlairGetResponseJSON   `json:"-"`
}

// accountRedditFlairGetResponseJSON contains the JSON metadata for the struct
// [AccountRedditFlairGetResponse]
type accountRedditFlairGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRedditFlairGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRedditFlairGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRedditFlairGetResponseData struct {
	ID   string                                `json:"id" api:"required"`
	Text string                                `json:"text" api:"required"`
	JSON accountRedditFlairGetResponseDataJSON `json:"-"`
}

// accountRedditFlairGetResponseDataJSON contains the JSON metadata for the struct
// [AccountRedditFlairGetResponseData]
type accountRedditFlairGetResponseDataJSON struct {
	ID          apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRedditFlairGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRedditFlairGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountRedditFlairGetParams struct {
	// Subreddit name
	Subreddit param.Field[string] `query:"subreddit" api:"required"`
}

// URLQuery serializes [AccountRedditFlairGetParams]'s query parameters as
// `url.Values`.
func (r AccountRedditFlairGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
