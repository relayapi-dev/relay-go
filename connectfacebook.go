// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"github.com/relayapi-dev/relay-go/option"
)

// ConnectFacebookService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectFacebookService] method instead.
type ConnectFacebookService struct {
	Options []option.RequestOption
	Pages   *ConnectFacebookPageService
}

// NewConnectFacebookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectFacebookService(opts ...option.RequestOption) (r *ConnectFacebookService) {
	r = &ConnectFacebookService{}
	r.Options = opts
	r.Pages = NewConnectFacebookPageService(opts...)
	return
}
