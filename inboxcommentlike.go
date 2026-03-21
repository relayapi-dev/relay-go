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

// InboxCommentLikeService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxCommentLikeService] method instead.
type InboxCommentLikeService struct {
	Options []option.RequestOption
}

// NewInboxCommentLikeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboxCommentLikeService(opts ...option.RequestOption) (r *InboxCommentLikeService) {
	r = &InboxCommentLikeService{}
	r.Options = opts
	return
}

// Like a comment
func (r *InboxCommentLikeService) New(ctx context.Context, commentID string, opts ...option.RequestOption) (res *InboxCommentLikeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commentID == "" {
		err = errors.New("missing required comment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s/like", commentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Unlike a comment
func (r *InboxCommentLikeService) Delete(ctx context.Context, commentID string, opts ...option.RequestOption) (res *InboxCommentLikeDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if commentID == "" {
		err = errors.New("missing required comment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/inbox/comments/%s/like", commentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type InboxCommentLikeNewResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                          `json:"comment_id"`
	JSON      inboxCommentLikeNewResponseJSON `json:"-"`
}

// inboxCommentLikeNewResponseJSON contains the JSON metadata for the struct
// [InboxCommentLikeNewResponse]
type inboxCommentLikeNewResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentLikeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentLikeNewResponseJSON) RawJSON() string {
	return r.raw
}

type InboxCommentLikeDeleteResponse struct {
	// Whether the action succeeded
	Success bool `json:"success" api:"required"`
	// Comment ID
	CommentID string                             `json:"comment_id"`
	JSON      inboxCommentLikeDeleteResponseJSON `json:"-"`
}

// inboxCommentLikeDeleteResponseJSON contains the JSON metadata for the struct
// [InboxCommentLikeDeleteResponse]
type inboxCommentLikeDeleteResponseJSON struct {
	Success     apijson.Field
	CommentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboxCommentLikeDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboxCommentLikeDeleteResponseJSON) RawJSON() string {
	return r.raw
}
