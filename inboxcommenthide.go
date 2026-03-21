// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// InboxCommentHideService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxCommentHideService] method instead.
type InboxCommentHideService struct {
	Options []option.RequestOption
}

// NewInboxCommentHideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboxCommentHideService(opts ...option.RequestOption) (r *InboxCommentHideService) {
	r = &InboxCommentHideService{}
	r.Options = opts
	return
}

// Hide a comment
func (r *InboxCommentHideService) New(ctx context.Context, commentID string, opts ...option.RequestOption) (res *InboxCommentHideNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commentID == "" {
		err = errors.New("missing required comment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s/hide", commentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Unhide a comment
func (r *InboxCommentHideService) Delete(ctx context.Context, commentID string, opts ...option.RequestOption) (res *InboxCommentHideDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commentID == "" {
		err = errors.New("missing required comment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s/hide", commentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type InboxCommentHideNewResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                          `json:"comment_id"`
	JSON      inboxCommentHideNewResponseJSON `json:"-"`
}

// inboxCommentHideNewResponseJSON contains the JSON metadata for the struct
// [InboxCommentHideNewResponse]
type inboxCommentHideNewResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentHideNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentHideNewResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentHideDeleteResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                             `json:"comment_id"`
	JSON      inboxCommentHideDeleteResponseJSON `json:"-"`
}

// inboxCommentHideDeleteResponseJSON contains the JSON metadata for the struct
// [InboxCommentHideDeleteResponse]
type inboxCommentHideDeleteResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentHideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentHideDeleteResponseJSON) RawJSON() string {
	return r.raw
}
