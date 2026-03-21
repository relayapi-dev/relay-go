// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"github.com/relayapi-dev/relay-go/option"
)

// ConnectLinkedinService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectLinkedinService] method instead.
type ConnectLinkedinService struct {
	Options       []option.RequestOption
	Organizations *ConnectLinkedinOrganizationService
}

// NewConnectLinkedinService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectLinkedinService(opts ...option.RequestOption) (r *ConnectLinkedinService) {
	r = &ConnectLinkedinService{}
	r.Options = opts
	r.Organizations = NewConnectLinkedinOrganizationService(opts...)
	return
}
