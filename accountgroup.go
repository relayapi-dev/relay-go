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

// AccountGroupService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGroupService] method instead.
type AccountGroupService struct {
	Options []option.RequestOption
}

// NewAccountGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountGroupService(opts ...option.RequestOption) (r *AccountGroupService) {
	r = &AccountGroupService{}
	r.Options = opts
	return
}

// Create an account group
func (r *AccountGroupService) New(ctx context.Context, body AccountGroupNewParams, opts ...option.RequestOption) (res *AccountGroupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an account group
func (r *AccountGroupService) Update(ctx context.Context, id string, body AccountGroupUpdateParams, opts ...option.RequestOption) (res *AccountGroupUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account-groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List account groups
func (r *AccountGroupService) List(ctx context.Context, query AccountGroupListParams, opts ...option.RequestOption) (res *AccountGroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an account group
func (r *AccountGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/account-groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AccountGroupNewResponse struct {
	// Group ID
	ID string `json:"id" api:"required"`
	// Number of accounts in this group
	AccountCount float64 `json:"account_count" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Group description
	Description string `json:"description" api:"required,nullable"`
	// Group name
	Name string `json:"name" api:"required"`
	// Last updated timestamp
	UpdatedAt time.Time                   `json:"updated_at" api:"required" format:"date-time"`
	JSON      accountGroupNewResponseJSON `json:"-"`
}

// accountGroupNewResponseJSON contains the JSON metadata for the struct
// [AccountGroupNewResponse]
type accountGroupNewResponseJSON struct {
	ID           apijson.Field
	AccountCount apijson.Field
	CreatedAt    apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGroupUpdateResponse struct {
	// Group ID
	ID string `json:"id" api:"required"`
	// Number of accounts in this group
	AccountCount float64 `json:"account_count" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Group description
	Description string `json:"description" api:"required,nullable"`
	// Group name
	Name string `json:"name" api:"required"`
	// Last updated timestamp
	UpdatedAt time.Time                      `json:"updated_at" api:"required" format:"date-time"`
	JSON      accountGroupUpdateResponseJSON `json:"-"`
}

// accountGroupUpdateResponseJSON contains the JSON metadata for the struct
// [AccountGroupUpdateResponse]
type accountGroupUpdateResponseJSON struct {
	ID           apijson.Field
	AccountCount apijson.Field
	CreatedAt    apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGroupListResponse struct {
	Data       []AccountGroupListResponseData `json:"data" api:"required"`
	HasMore    bool                           `json:"has_more" api:"required"`
	NextCursor string                         `json:"next_cursor" api:"required,nullable"`
	JSON       accountGroupListResponseJSON   `json:"-"`
}

// accountGroupListResponseJSON contains the JSON metadata for the struct
// [AccountGroupListResponse]
type accountGroupListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGroupListResponseData struct {
	// Group ID
	ID string `json:"id" api:"required"`
	// Number of accounts in this group
	AccountCount float64 `json:"account_count" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Group description
	Description string `json:"description" api:"required,nullable"`
	// Group name
	Name string `json:"name" api:"required"`
	// Last updated timestamp
	UpdatedAt time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	JSON      accountGroupListResponseDataJSON `json:"-"`
}

// accountGroupListResponseDataJSON contains the JSON metadata for the struct
// [AccountGroupListResponseData]
type accountGroupListResponseDataJSON struct {
	ID           apijson.Field
	AccountCount apijson.Field
	CreatedAt    apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGroupListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGroupListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountGroupNewParams struct {
	// Group name
	Name param.Field[string] `json:"name" api:"required"`
	// Group description
	Description param.Field[string] `json:"description"`
}

func (r AccountGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGroupUpdateParams struct {
	// Group description
	Description param.Field[string] `json:"description"`
	// Group name
	Name param.Field[string] `json:"name"`
}

func (r AccountGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGroupListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Page size
	Limit param.Field[float64] `query:"limit"`
	// Search groups by name
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountGroupListParams]'s query parameters as `url.Values`.
func (r AccountGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
