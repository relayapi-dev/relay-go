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

// TwitterFollowService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTwitterFollowService] method instead.
type TwitterFollowService struct {
	Options []option.RequestOption
}

// NewTwitterFollowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTwitterFollowService(opts ...option.RequestOption) (r *TwitterFollowService) {
	r = &TwitterFollowService{}
	r.Options = opts
	return
}

// Follow a user
func (r *TwitterFollowService) New(ctx context.Context, body TwitterFollowNewParams, opts ...option.RequestOption) (res *TwitterFollowNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/twitter/follow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unfollow a user
func (r *TwitterFollowService) Unfollow(ctx context.Context, body TwitterFollowUnfollowParams, opts ...option.RequestOption) (res *TwitterFollowUnfollowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/twitter/follow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type TwitterFollowNewResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Action result data from Twitter API
	Data TwitterFollowNewResponseData `json:"data"`
	// Error details when success is false
	Error TwitterFollowNewResponseError `json:"error"`
	JSON  twitterFollowNewResponseJSON  `json:"-"`
}

// twitterFollowNewResponseJSON contains the JSON metadata for the struct
// [TwitterFollowNewResponse]
type twitterFollowNewResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterFollowNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowNewResponseJSON) RawJSON() string {
	return r.raw
}

// Action result data from Twitter API
type TwitterFollowNewResponseData struct {
	Bookmarked    bool                             `json:"bookmarked"`
	Following     bool                             `json:"following"`
	PendingFollow bool                             `json:"pending_follow"`
	Retweeted     bool                             `json:"retweeted"`
	JSON          twitterFollowNewResponseDataJSON `json:"-"`
}

// twitterFollowNewResponseDataJSON contains the JSON metadata for the struct
// [TwitterFollowNewResponseData]
type twitterFollowNewResponseDataJSON struct {
	Bookmarked    apijson.Field
	Following     apijson.Field
	PendingFollow apijson.Field
	Retweeted     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TwitterFollowNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowNewResponseDataJSON) RawJSON() string {
	return r.raw
}

// Error details when success is false
type TwitterFollowNewResponseError struct {
	// Error code (e.g. ACCOUNT_NOT_FOUND, TOKEN_MISSING, TWITTER_API_ERROR)
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Twitter API error code if available
	TwitterErrorCode float64                           `json:"twitter_error_code"`
	JSON             twitterFollowNewResponseErrorJSON `json:"-"`
}

// twitterFollowNewResponseErrorJSON contains the JSON metadata for the struct
// [TwitterFollowNewResponseError]
type twitterFollowNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	TwitterErrorCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TwitterFollowNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TwitterFollowUnfollowResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Action result data from Twitter API
	Data TwitterFollowUnfollowResponseData `json:"data"`
	// Error details when success is false
	Error TwitterFollowUnfollowResponseError `json:"error"`
	JSON  twitterFollowUnfollowResponseJSON  `json:"-"`
}

// twitterFollowUnfollowResponseJSON contains the JSON metadata for the struct
// [TwitterFollowUnfollowResponse]
type twitterFollowUnfollowResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterFollowUnfollowResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowUnfollowResponseJSON) RawJSON() string {
	return r.raw
}

// Action result data from Twitter API
type TwitterFollowUnfollowResponseData struct {
	Bookmarked    bool                                  `json:"bookmarked"`
	Following     bool                                  `json:"following"`
	PendingFollow bool                                  `json:"pending_follow"`
	Retweeted     bool                                  `json:"retweeted"`
	JSON          twitterFollowUnfollowResponseDataJSON `json:"-"`
}

// twitterFollowUnfollowResponseDataJSON contains the JSON metadata for the struct
// [TwitterFollowUnfollowResponseData]
type twitterFollowUnfollowResponseDataJSON struct {
	Bookmarked    apijson.Field
	Following     apijson.Field
	PendingFollow apijson.Field
	Retweeted     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TwitterFollowUnfollowResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowUnfollowResponseDataJSON) RawJSON() string {
	return r.raw
}

// Error details when success is false
type TwitterFollowUnfollowResponseError struct {
	// Error code (e.g. ACCOUNT_NOT_FOUND, TOKEN_MISSING, TWITTER_API_ERROR)
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Twitter API error code if available
	TwitterErrorCode float64                                `json:"twitter_error_code"`
	JSON             twitterFollowUnfollowResponseErrorJSON `json:"-"`
}

// twitterFollowUnfollowResponseErrorJSON contains the JSON metadata for the struct
// [TwitterFollowUnfollowResponseError]
type twitterFollowUnfollowResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	TwitterErrorCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TwitterFollowUnfollowResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowUnfollowResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TwitterFollowNewParams struct {
	// Twitter account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// User ID to follow
	TargetUserID param.Field[string] `json:"target_user_id" api:"required"`
}

func (r TwitterFollowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TwitterFollowUnfollowParams struct {
	// Twitter account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// User ID to follow
	TargetUserID param.Field[string] `json:"target_user_id" api:"required"`
}

func (r TwitterFollowUnfollowParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
