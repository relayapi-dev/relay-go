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

// WhatsappContactService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappContactService] method instead.
type WhatsappContactService struct {
	Options []option.RequestOption
}

// NewWhatsappContactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhatsappContactService(opts ...option.RequestOption) (r *WhatsappContactService) {
	r = &WhatsappContactService{}
	r.Options = opts
	return
}

// Create a contact
func (r *WhatsappContactService) New(ctx context.Context, body WhatsappContactNewParams, opts ...option.RequestOption) (res *WhatsappContactNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get contact details
func (r *WhatsappContactService) Get(ctx context.Context, contactID string, opts ...option.RequestOption) (res *WhatsappContactGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/whatsapp/contacts/%s", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List contacts
func (r *WhatsappContactService) List(ctx context.Context, query WhatsappContactListParams, opts ...option.RequestOption) (res *WhatsappContactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a contact
func (r *WhatsappContactService) Delete(ctx context.Context, contactID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if contactID == "" {
		err = errors.New("missing required contact_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/whatsapp/contacts/%s", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Bulk contact operations (add/remove tags, delete)
func (r *WhatsappContactService) BulkOperations(ctx context.Context, body WhatsappContactBulkOperationsParams, opts ...option.RequestOption) (res *WhatsappContactBulkOperationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/contacts/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Bulk import contacts
func (r *WhatsappContactService) Import(ctx context.Context, body WhatsappContactImportParams, opts ...option.RequestOption) (res *WhatsappContactImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/contacts/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WhatsappContactNewResponse struct {
	// Contact ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether contact has opted in
	OptedIn bool `json:"opted_in" api:"required"`
	// Phone number
	Phone string `json:"phone" api:"required"`
	// Email address
	Email string `json:"email" api:"nullable"`
	// Group IDs
	Groups []string `json:"groups"`
	// Contact name
	Name string `json:"name" api:"nullable"`
	// Tags
	Tags []string                       `json:"tags"`
	JSON whatsappContactNewResponseJSON `json:"-"`
}

// whatsappContactNewResponseJSON contains the JSON metadata for the struct
// [WhatsappContactNewResponse]
type whatsappContactNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	OptedIn     apijson.Field
	Phone       apijson.Field
	Email       apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappContactNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappContactNewResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappContactGetResponse struct {
	// Contact ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether contact has opted in
	OptedIn bool `json:"opted_in" api:"required"`
	// Phone number
	Phone string `json:"phone" api:"required"`
	// Email address
	Email string `json:"email" api:"nullable"`
	// Group IDs
	Groups []string `json:"groups"`
	// Contact name
	Name string `json:"name" api:"nullable"`
	// Tags
	Tags []string                       `json:"tags"`
	JSON whatsappContactGetResponseJSON `json:"-"`
}

// whatsappContactGetResponseJSON contains the JSON metadata for the struct
// [WhatsappContactGetResponse]
type whatsappContactGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	OptedIn     apijson.Field
	Phone       apijson.Field
	Email       apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappContactGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappContactGetResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappContactListResponse struct {
	Data []WhatsappContactListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                          `json:"next_cursor" api:"required,nullable"`
	JSON       whatsappContactListResponseJSON `json:"-"`
}

// whatsappContactListResponseJSON contains the JSON metadata for the struct
// [WhatsappContactListResponse]
type whatsappContactListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappContactListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappContactListResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappContactListResponseData struct {
	// Contact ID
	ID string `json:"id" api:"required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether contact has opted in
	OptedIn bool `json:"opted_in" api:"required"`
	// Phone number
	Phone string `json:"phone" api:"required"`
	// Email address
	Email string `json:"email" api:"nullable"`
	// Group IDs
	Groups []string `json:"groups"`
	// Contact name
	Name string `json:"name" api:"nullable"`
	// Tags
	Tags []string                            `json:"tags"`
	JSON whatsappContactListResponseDataJSON `json:"-"`
}

// whatsappContactListResponseDataJSON contains the JSON metadata for the struct
// [WhatsappContactListResponseData]
type whatsappContactListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	OptedIn     apijson.Field
	Phone       apijson.Field
	Email       apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappContactListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappContactListResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhatsappContactBulkOperationsResponse struct {
	// Number of contacts affected
	Affected float64                                   `json:"affected" api:"required"`
	JSON     whatsappContactBulkOperationsResponseJSON `json:"-"`
}

// whatsappContactBulkOperationsResponseJSON contains the JSON metadata for the
// struct [WhatsappContactBulkOperationsResponse]
type whatsappContactBulkOperationsResponseJSON struct {
	Affected    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappContactBulkOperationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappContactBulkOperationsResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappContactImportResponse struct {
	// Failed count
	Failed float64 `json:"failed" api:"required"`
	// Successfully imported count
	Imported float64 `json:"imported" api:"required"`
	// Skipped (duplicate) count
	Skipped float64                           `json:"skipped" api:"required"`
	JSON    whatsappContactImportResponseJSON `json:"-"`
}

// whatsappContactImportResponseJSON contains the JSON metadata for the struct
// [WhatsappContactImportResponse]
type whatsappContactImportResponseJSON struct {
	Failed      apijson.Field
	Imported    apijson.Field
	Skipped     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappContactImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappContactImportResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappContactNewParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Phone number in E.164 format
	Phone param.Field[string] `json:"phone" api:"required"`
	// Email address
	Email param.Field[string] `json:"email"`
	// Contact name
	Name param.Field[string] `json:"name"`
	// Tags
	Tags param.Field[[]string] `json:"tags"`
}

func (r WhatsappContactNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappContactListParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items
	Limit param.Field[int64] `query:"limit"`
	// Search by name or phone
	Search param.Field[string] `query:"search"`
	// Filter by tag
	Tag param.Field[string] `query:"tag"`
}

// URLQuery serializes [WhatsappContactListParams]'s query parameters as
// `url.Values`.
func (r WhatsappContactListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhatsappContactBulkOperationsParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Action
	Action param.Field[WhatsappContactBulkOperationsParamsAction] `json:"action" api:"required"`
	// Contact IDs
	ContactIDs param.Field[[]string] `json:"contact_ids" api:"required"`
	// Tags (for tag actions)
	Tags param.Field[[]string] `json:"tags"`
}

func (r WhatsappContactBulkOperationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action
type WhatsappContactBulkOperationsParamsAction string

const (
	WhatsappContactBulkOperationsParamsActionAddTags    WhatsappContactBulkOperationsParamsAction = "add_tags"
	WhatsappContactBulkOperationsParamsActionRemoveTags WhatsappContactBulkOperationsParamsAction = "remove_tags"
	WhatsappContactBulkOperationsParamsActionDelete     WhatsappContactBulkOperationsParamsAction = "delete"
)

func (r WhatsappContactBulkOperationsParamsAction) IsKnown() bool {
	switch r {
	case WhatsappContactBulkOperationsParamsActionAddTags, WhatsappContactBulkOperationsParamsActionRemoveTags, WhatsappContactBulkOperationsParamsActionDelete:
		return true
	}
	return false
}

type WhatsappContactImportParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Contacts to import
	Contacts param.Field[[]WhatsappContactImportParamsContact] `json:"contacts" api:"required"`
}

func (r WhatsappContactImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappContactImportParamsContact struct {
	// Phone number
	Phone param.Field[string]   `json:"phone" api:"required"`
	Email param.Field[string]   `json:"email"`
	Name  param.Field[string]   `json:"name"`
	Tags  param.Field[[]string] `json:"tags"`
}

func (r WhatsappContactImportParamsContact) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
