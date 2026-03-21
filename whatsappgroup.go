// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
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

// Create a contact group
func (r *WhatsappGroupService) New(ctx context.Context, body WhatsappGroupNewParams, opts ...option.RequestOption) (res *WhatsappGroupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List contact groups
func (r *WhatsappGroupService) List(ctx context.Context, query WhatsappGroupListParams, opts ...option.RequestOption) (res *WhatsappGroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a contact group
func (r *WhatsappGroupService) Delete(ctx context.Context, groupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/whatsapp/groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WhatsappGroupNewResponse struct {
	// Group ID
	ID string `json:"id" api:"required"`
	// Number of contacts
	ContactCount float64 `json:"contact_count" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Group name
	Name        string                       `json:"name" api:"required"`
	Description string                       `json:"description" api:"nullable"`
	JSON        whatsappGroupNewResponseJSON `json:"-"`
}

// whatsappGroupNewResponseJSON contains the JSON metadata for the struct
// [WhatsappGroupNewResponse]
type whatsappGroupNewResponseJSON struct {
	ID           apijson.Field
	ContactCount apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WhatsappGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappGroupListResponse struct {
	Data []WhatsappGroupListResponseData `json:"data" api:"required"`
	JSON whatsappGroupListResponseJSON   `json:"-"`
}

// whatsappGroupListResponseJSON contains the JSON metadata for the struct
// [WhatsappGroupListResponse]
type whatsappGroupListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappGroupListResponseData struct {
	// Group ID
	ID string `json:"id" api:"required"`
	// Number of contacts
	ContactCount float64 `json:"contact_count" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Group name
	Name        string                            `json:"name" api:"required"`
	Description string                            `json:"description" api:"nullable"`
	JSON        whatsappGroupListResponseDataJSON `json:"-"`
}

// whatsappGroupListResponseDataJSON contains the JSON metadata for the struct
// [WhatsappGroupListResponseData]
type whatsappGroupListResponseDataJSON struct {
	ID           apijson.Field
	ContactCount apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WhatsappGroupListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappGroupListResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhatsappGroupNewParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Group name
	Name param.Field[string] `json:"name" api:"required"`
	// Initial contact IDs
	ContactIDs param.Field[[]string] `json:"contact_ids"`
	// Group description
	Description param.Field[string] `json:"description"`
}

func (r WhatsappGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappGroupListParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappGroupListParams]'s query parameters as
// `url.Values`.
func (r WhatsappGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
