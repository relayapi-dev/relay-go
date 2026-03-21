// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// InboxReviewReplyService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxReviewReplyService] method instead.
type InboxReviewReplyService struct {
	Options []option.RequestOption
}

// NewInboxReviewReplyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboxReviewReplyService(opts ...option.RequestOption) (r *InboxReviewReplyService) {
	r = &InboxReviewReplyService{}
	r.Options = opts
	return
}

// Reply to a review
func (r *InboxReviewReplyService) New(ctx context.Context, reviewID string, body InboxReviewReplyNewParams, opts ...option.RequestOption) (res *InboxReviewReplyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reviewID == "" {
		err = errors.New("missing required review_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/reviews/%s/reply", reviewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a review reply
func (r *InboxReviewReplyService) Delete(ctx context.Context, reviewID string, opts ...option.RequestOption) (res *InboxReviewReplyDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reviewID == "" {
		err = errors.New("missing required review_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/reviews/%s/reply", reviewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type InboxReviewReplyNewResponse struct {
	// Whether the action succeeded
	Success bool                            `json:"success" api:"required"`
	JSON    inboxReviewReplyNewResponseJSON `json:"-"`
}

// inboxReviewReplyNewResponseJSON contains the JSON metadata for the struct
// [InboxReviewReplyNewResponse]
type inboxReviewReplyNewResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxReviewReplyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxReviewReplyNewResponseJSON) RawJSON() string {
	return r.raw
}

type InboxReviewReplyDeleteResponse struct {
	// Whether the action succeeded
	Success bool                               `json:"success" api:"required"`
	JSON    inboxReviewReplyDeleteResponseJSON `json:"-"`
}

// inboxReviewReplyDeleteResponseJSON contains the JSON metadata for the struct
// [InboxReviewReplyDeleteResponse]
type inboxReviewReplyDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxReviewReplyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxReviewReplyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type InboxReviewReplyNewParams struct {
	// Account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Reply text
	Text param.Field[string] `json:"text" api:"required"`
}

func (r InboxReviewReplyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
