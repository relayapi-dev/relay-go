// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"github.com/relayapi-dev/relay-go/option"
)

// ConnectPinterestService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectPinterestService] method instead.
type ConnectPinterestService struct {
	Options []option.RequestOption
	Boards  *ConnectPinterestBoardService
}

// NewConnectPinterestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectPinterestService(opts ...option.RequestOption) (r *ConnectPinterestService) {
	r = &ConnectPinterestService{}
	r.Options = opts
	r.Boards = NewConnectPinterestBoardService(opts...)
	return
}
