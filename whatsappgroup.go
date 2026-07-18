// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"github.com/relayapi-dev/relay-go/option"
)

// WhatsappGroupService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappGroupService] method instead.
type WhatsappGroupService struct {
	Options []option.RequestOption
}

// NewWhatsappGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhatsappGroupService(opts ...option.RequestOption) (r *WhatsappGroupService) {
	r = &WhatsappGroupService{}
	r.Options = opts
	return
}
