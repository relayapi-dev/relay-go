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
	Success bool                         `json:"success" api:"required"`
	JSON    twitterFollowNewResponseJSON `json:"-"`
}

// twitterFollowNewResponseJSON contains the JSON metadata for the struct
// [TwitterFollowNewResponse]
type twitterFollowNewResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterFollowNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowNewResponseJSON) RawJSON() string {
	return r.raw
}

type TwitterFollowUnfollowResponse struct {
	// Whether the action succeeded
	Success bool                              `json:"success" api:"required"`
	JSON    twitterFollowUnfollowResponseJSON `json:"-"`
}

// twitterFollowUnfollowResponseJSON contains the JSON metadata for the struct
// [TwitterFollowUnfollowResponse]
type twitterFollowUnfollowResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TwitterFollowUnfollowResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r twitterFollowUnfollowResponseJSON) RawJSON() string {
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
