// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"github.com/relayapi-dev/relay-go/option"
)

// ConnectGooglebusinessService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectGooglebusinessService] method instead.
type ConnectGooglebusinessService struct {
	Options   []option.RequestOption
	Locations *ConnectGooglebusinessLocationService
}

// NewConnectGooglebusinessService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectGooglebusinessService(opts ...option.RequestOption) (r *ConnectGooglebusinessService) {
	r = &ConnectGooglebusinessService{}
	r.Options = opts
	r.Locations = NewConnectGooglebusinessLocationService(opts...)
	return
}
