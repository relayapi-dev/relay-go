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

// TwitterRetweetService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTwitterRetweetService] method instead.
type TwitterRetweetService struct {
	Options []option.RequestOption
}

// NewTwitterRetweetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTwitterRetweetService(opts ...option.RequestOption) (r *TwitterRetweetService) {
	r = &TwitterRetweetService{}
	r.Options = opts
	return
}

// Retweet a tweet
func (r *TwitterRetweetService) New(ctx context.Context, body TwitterRetweetNewParams, opts ...option.RequestOption) (res *TwitterRetweetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/twitter/retweet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Undo a retweet
func (r *TwitterRetweetService) Undo(ctx context.Context, body TwitterRetweetUndoParams, opts ...option.RequestOption) (res *TwitterRetweetUndoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/twitter/retweet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type TwitterRetweetNewResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Action result data from Twitter API
	Data TwitterRetweetNewResponseData `json:"data"`
	// Error details when success is false
	Error TwitterRetweetNewResponseError `json:"error"`
	JSON  twitterRetweetNewResponseJSON  `json:"-"`
}

// twitterRetweetNewResponseJSON contains the JSON metadata for the struct
// [TwitterRetweetNewResponse]
type twitterRetweetNewResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterRetweetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterRetweetNewResponseJSON) RawJSON() string {
	return r.raw
}

// Action result data from Twitter API
type TwitterRetweetNewResponseData struct {
	Bookmarked    bool                              `json:"bookmarked"`
	Following     bool                              `json:"following"`
	PendingFollow bool                              `json:"pending_follow"`
	Retweeted     bool                              `json:"retweeted"`
	JSON          twitterRetweetNewResponseDataJSON `json:"-"`
}

// twitterRetweetNewResponseDataJSON contains the JSON metadata for the struct
// [TwitterRetweetNewResponseData]
type twitterRetweetNewResponseDataJSON struct {
	Bookmarked    apijson.Field
	Following     apijson.Field
	PendingFollow apijson.Field
	Retweeted     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TwitterRetweetNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterRetweetNewResponseDataJSON) RawJSON() string {
	return r.raw
}

// Error details when success is false
type TwitterRetweetNewResponseError struct {
	// Error code (e.g. ACCOUNT_NOT_FOUND, TOKEN_MISSING, TWITTER_API_ERROR)
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Twitter API error code if available
	TwitterErrorCode float64                            `json:"twitter_error_code"`
	JSON             twitterRetweetNewResponseErrorJSON `json:"-"`
}

// twitterRetweetNewResponseErrorJSON contains the JSON metadata for the struct
// [TwitterRetweetNewResponseError]
type twitterRetweetNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	TwitterErrorCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TwitterRetweetNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterRetweetNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TwitterRetweetUndoResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Action result data from Twitter API
	Data TwitterRetweetUndoResponseData `json:"data"`
	// Error details when success is false
	Error TwitterRetweetUndoResponseError `json:"error"`
	JSON  twitterRetweetUndoResponseJSON  `json:"-"`
}

// twitterRetweetUndoResponseJSON contains the JSON metadata for the struct
// [TwitterRetweetUndoResponse]
type twitterRetweetUndoResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterRetweetUndoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterRetweetUndoResponseJSON) RawJSON() string {
	return r.raw
}

// Action result data from Twitter API
type TwitterRetweetUndoResponseData struct {
	Bookmarked    bool                               `json:"bookmarked"`
	Following     bool                               `json:"following"`
	PendingFollow bool                               `json:"pending_follow"`
	Retweeted     bool                               `json:"retweeted"`
	JSON          twitterRetweetUndoResponseDataJSON `json:"-"`
}

// twitterRetweetUndoResponseDataJSON contains the JSON metadata for the struct
// [TwitterRetweetUndoResponseData]
type twitterRetweetUndoResponseDataJSON struct {
	Bookmarked    apijson.Field
	Following     apijson.Field
	PendingFollow apijson.Field
	Retweeted     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TwitterRetweetUndoResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterRetweetUndoResponseDataJSON) RawJSON() string {
	return r.raw
}

// Error details when success is false
type TwitterRetweetUndoResponseError struct {
	// Error code (e.g. ACCOUNT_NOT_FOUND, TOKEN_MISSING, TWITTER_API_ERROR)
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Twitter API error code if available
	TwitterErrorCode float64                             `json:"twitter_error_code"`
	JSON             twitterRetweetUndoResponseErrorJSON `json:"-"`
}

// twitterRetweetUndoResponseErrorJSON contains the JSON metadata for the struct
// [TwitterRetweetUndoResponseError]
type twitterRetweetUndoResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	TwitterErrorCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TwitterRetweetUndoResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterRetweetUndoResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TwitterRetweetNewParams struct {
	// Twitter account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Tweet ID to retweet
	TweetID param.Field[string] `json:"tweet_id" api:"required"`
}

func (r TwitterRetweetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TwitterRetweetUndoParams struct {
	// Twitter account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Tweet ID to retweet
	TweetID param.Field[string] `json:"tweet_id" api:"required"`
}

func (r TwitterRetweetUndoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
