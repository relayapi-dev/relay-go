// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// WhatsappBusinessProfileService contains methods and other services that help
// with interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappBusinessProfileService] method instead.
type WhatsappBusinessProfileService struct {
	Options []option.RequestOption
}

// NewWhatsappBusinessProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappBusinessProfileService(opts ...option.RequestOption) (r *WhatsappBusinessProfileService) {
	r = &WhatsappBusinessProfileService{}
	r.Options = opts
	return
}

// Get WhatsApp Business profile
func (r *WhatsappBusinessProfileService) Get(ctx context.Context, query WhatsappBusinessProfileGetParams, opts ...option.RequestOption) (res *WhatsappBusinessProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/business-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update WhatsApp Business profile
func (r *WhatsappBusinessProfileService) Update(ctx context.Context, body WhatsappBusinessProfileUpdateParams, opts ...option.RequestOption) (res *WhatsappBusinessProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/business-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type WhatsappBusinessProfileGetResponse struct {
	// About text
	About string `json:"about" api:"nullable"`
	// Business address
	Address string `json:"address" api:"nullable"`
	// Description
	Description string `json:"description" api:"nullable"`
	// Business email
	Email string `json:"email" api:"nullable"`
	// Profile picture URL
	ProfilePictureURL string `json:"profile_picture_url" api:"nullable"`
	// Website URLs
	Websites []string                               `json:"websites"`
	JSON     whatsappBusinessProfileGetResponseJSON `json:"-"`
}

// whatsappBusinessProfileGetResponseJSON contains the JSON metadata for the struct
// [WhatsappBusinessProfileGetResponse]
type whatsappBusinessProfileGetResponseJSON struct {
	About             apijson.Field
	Address           apijson.Field
	Description       apijson.Field
	Email             apijson.Field
	ProfilePictureURL apijson.Field
	Websites          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WhatsappBusinessProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBusinessProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappBusinessProfileUpdateResponse struct {
	// About text
	About string `json:"about" api:"nullable"`
	// Business address
	Address string `json:"address" api:"nullable"`
	// Description
	Description string `json:"description" api:"nullable"`
	// Business email
	Email string `json:"email" api:"nullable"`
	// Profile picture URL
	ProfilePictureURL string `json:"profile_picture_url" api:"nullable"`
	// Website URLs
	Websites []string                                  `json:"websites"`
	JSON     whatsappBusinessProfileUpdateResponseJSON `json:"-"`
}

// whatsappBusinessProfileUpdateResponseJSON contains the JSON metadata for the
// struct [WhatsappBusinessProfileUpdateResponse]
type whatsappBusinessProfileUpdateResponseJSON struct {
	About             apijson.Field
	Address           apijson.Field
	Description       apijson.Field
	Email             apijson.Field
	ProfilePictureURL apijson.Field
	Websites          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WhatsappBusinessProfileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappBusinessProfileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappBusinessProfileGetParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappBusinessProfileGetParams]'s query parameters as
// `url.Values`.
func (r WhatsappBusinessProfileGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhatsappBusinessProfileUpdateParams struct {
	// WhatsApp account ID
	AccountID   param.Field[string]   `json:"account_id" api:"required"`
	About       param.Field[string]   `json:"about"`
	Address     param.Field[string]   `json:"address"`
	Description param.Field[string]   `json:"description"`
	Email       param.Field[string]   `json:"email"`
	Websites    param.Field[[]string] `json:"websites"`
}

func (r WhatsappBusinessProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
