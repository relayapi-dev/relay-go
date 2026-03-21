// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"github.com/relayapi-dev/relay-go/option"
)

// InboxService contains methods and other services that help with interacting with
// the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxService] method instead.
type InboxService struct {
	Options  []option.RequestOption
	Comments *InboxCommentService
	Messages *InboxMessageService
	Reviews  *InboxReviewService
}

// NewInboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboxService(opts ...option.RequestOption) (r *InboxService) {
	r = &InboxService{}
	r.Options = opts
	r.Comments = NewInboxCommentService(opts...)
	r.Messages = NewInboxMessageService(opts...)
	r.Reviews = NewInboxReviewService(opts...)
	return
}
