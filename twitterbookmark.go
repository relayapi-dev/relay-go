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

// TwitterBookmarkService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTwitterBookmarkService] method instead.
type TwitterBookmarkService struct {
	Options []option.RequestOption
}

// NewTwitterBookmarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTwitterBookmarkService(opts ...option.RequestOption) (r *TwitterBookmarkService) {
	r = &TwitterBookmarkService{}
	r.Options = opts
	return
}

// Bookmark a tweet
func (r *TwitterBookmarkService) New(ctx context.Context, body TwitterBookmarkNewParams, opts ...option.RequestOption) (res *TwitterBookmarkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/twitter/bookmark"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove a bookmark
func (r *TwitterBookmarkService) Remove(ctx context.Context, body TwitterBookmarkRemoveParams, opts ...option.RequestOption) (res *TwitterBookmarkRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/twitter/bookmark"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type TwitterBookmarkNewResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Action result data from Twitter API
	Data TwitterBookmarkNewResponseData `json:"data"`
	// Error details when success is false
	Error TwitterBookmarkNewResponseError `json:"error"`
	JSON  twitterBookmarkNewResponseJSON  `json:"-"`
}

// twitterBookmarkNewResponseJSON contains the JSON metadata for the struct
// [TwitterBookmarkNewResponse]
type twitterBookmarkNewResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterBookmarkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterBookmarkNewResponseJSON) RawJSON() string {
	return r.raw
}

// Action result data from Twitter API
type TwitterBookmarkNewResponseData struct {
	Bookmarked    bool                               `json:"bookmarked"`
	Following     bool                               `json:"following"`
	PendingFollow bool                               `json:"pending_follow"`
	Retweeted     bool                               `json:"retweeted"`
	JSON          twitterBookmarkNewResponseDataJSON `json:"-"`
}

// twitterBookmarkNewResponseDataJSON contains the JSON metadata for the struct
// [TwitterBookmarkNewResponseData]
type twitterBookmarkNewResponseDataJSON struct {
	Bookmarked    apijson.Field
	Following     apijson.Field
	PendingFollow apijson.Field
	Retweeted     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TwitterBookmarkNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterBookmarkNewResponseDataJSON) RawJSON() string {
	return r.raw
}

// Error details when success is false
type TwitterBookmarkNewResponseError struct {
	// Error code (e.g. ACCOUNT_NOT_FOUND, TOKEN_MISSING, TWITTER_API_ERROR)
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Twitter API error code if available
	TwitterErrorCode float64                             `json:"twitter_error_code"`
	JSON             twitterBookmarkNewResponseErrorJSON `json:"-"`
}

// twitterBookmarkNewResponseErrorJSON contains the JSON metadata for the struct
// [TwitterBookmarkNewResponseError]
type twitterBookmarkNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	TwitterErrorCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TwitterBookmarkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterBookmarkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TwitterBookmarkRemoveResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Action result data from Twitter API
	Data TwitterBookmarkRemoveResponseData `json:"data"`
	// Error details when success is false
	Error TwitterBookmarkRemoveResponseError `json:"error"`
	JSON  twitterBookmarkRemoveResponseJSON  `json:"-"`
}

// twitterBookmarkRemoveResponseJSON contains the JSON metadata for the struct
// [TwitterBookmarkRemoveResponse]
type twitterBookmarkRemoveResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterBookmarkRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterBookmarkRemoveResponseJSON) RawJSON() string {
	return r.raw
}

// Action result data from Twitter API
type TwitterBookmarkRemoveResponseData struct {
	Bookmarked    bool                                  `json:"bookmarked"`
	Following     bool                                  `json:"following"`
	PendingFollow bool                                  `json:"pending_follow"`
	Retweeted     bool                                  `json:"retweeted"`
	JSON          twitterBookmarkRemoveResponseDataJSON `json:"-"`
}

// twitterBookmarkRemoveResponseDataJSON contains the JSON metadata for the struct
// [TwitterBookmarkRemoveResponseData]
type twitterBookmarkRemoveResponseDataJSON struct {
	Bookmarked    apijson.Field
	Following     apijson.Field
	PendingFollow apijson.Field
	Retweeted     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TwitterBookmarkRemoveResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterBookmarkRemoveResponseDataJSON) RawJSON() string {
	return r.raw
}

// Error details when success is false
type TwitterBookmarkRemoveResponseError struct {
	// Error code (e.g. ACCOUNT_NOT_FOUND, TOKEN_MISSING, TWITTER_API_ERROR)
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Twitter API error code if available
	TwitterErrorCode float64                                `json:"twitter_error_code"`
	JSON             twitterBookmarkRemoveResponseErrorJSON `json:"-"`
}

// twitterBookmarkRemoveResponseErrorJSON contains the JSON metadata for the struct
// [TwitterBookmarkRemoveResponseError]
type twitterBookmarkRemoveResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	TwitterErrorCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TwitterBookmarkRemoveResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterBookmarkRemoveResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TwitterBookmarkNewParams struct {
	// Twitter account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Tweet ID to bookmark
	TweetID param.Field[string] `json:"tweet_id" api:"required"`
}

func (r TwitterBookmarkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TwitterBookmarkRemoveParams struct {
	// Twitter account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Tweet ID to bookmark
	TweetID param.Field[string] `json:"tweet_id" api:"required"`
}

func (r TwitterBookmarkRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
