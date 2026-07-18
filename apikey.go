// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
	"github.com/tidwall/gjson"
)

// APIKeyService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r *APIKeyService) {
	r = &APIKeyService{}
	r.Options = opts
	return
}

// Create a new API key. The full key is returned only once in the response — store
// it securely.
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys
func (r *APIKeyService) List(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) (res *APIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an API key
func (r *APIKeyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/api-keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type APIKeyNewResponse struct {
	// API key ID
	ID string `json:"id" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Expiration timestamp
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Full API key (shown once, store securely)
	Key string `json:"key" api:"required"`
	// API key name
	Name string `json:"name" api:"required,nullable"`
	// Permission level
	Permission APIKeyNewResponsePermission `json:"permission" api:"required"`
	// Key prefix
	Prefix string `json:"prefix" api:"required"`
	// Workspace access: 'all' or array of workspace IDs
	WorkspaceScope APIKeyNewResponseWorkspaceScopeUnion `json:"workspace_scope" api:"required"`
	JSON           apiKeyNewResponseJSON                `json:"-"`
}

// apiKeyNewResponseJSON contains the JSON metadata for the struct
// [APIKeyNewResponse]
type apiKeyNewResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	ExpiresAt      apijson.Field
	Key            apijson.Field
	Name           apijson.Field
	Permission     apijson.Field
	Prefix         apijson.Field
	WorkspaceScope apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *APIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

// Permission level
type APIKeyNewResponsePermission string

const (
	APIKeyNewResponsePermissionReadWrite APIKeyNewResponsePermission = "read_write"
	APIKeyNewResponsePermissionReadOnly  APIKeyNewResponsePermission = "read_only"
)

func (r APIKeyNewResponsePermission) IsKnown() bool {
	switch r {
	case APIKeyNewResponsePermissionReadWrite, APIKeyNewResponsePermissionReadOnly:
		return true
	}
	return false
}

// Workspace access: 'all' or array of workspace IDs
//
// Union satisfied by [APIKeyNewResponseWorkspaceScopeString] or
// [APIKeyNewResponseWorkspaceScopeArray].
type APIKeyNewResponseWorkspaceScopeUnion interface {
	implementsAPIKeyNewResponseWorkspaceScopeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIKeyNewResponseWorkspaceScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(APIKeyNewResponseWorkspaceScopeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIKeyNewResponseWorkspaceScopeArray{}),
		},
	)
}

type APIKeyNewResponseWorkspaceScopeString string

const (
	APIKeyNewResponseWorkspaceScopeStringAll APIKeyNewResponseWorkspaceScopeString = "all"
)

func (r APIKeyNewResponseWorkspaceScopeString) IsKnown() bool {
	switch r {
	case APIKeyNewResponseWorkspaceScopeStringAll:
		return true
	}
	return false
}

func (r APIKeyNewResponseWorkspaceScopeString) implementsAPIKeyNewResponseWorkspaceScopeUnion() {}

type APIKeyNewResponseWorkspaceScopeArray []string

func (r APIKeyNewResponseWorkspaceScopeArray) implementsAPIKeyNewResponseWorkspaceScopeUnion() {}

type APIKeyListResponse struct {
	Data []APIKeyListResponseData `json:"data" api:"required"`
	// Whether more items exist
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for next page
	NextCursor string                 `json:"next_cursor" api:"required,nullable"`
	JSON       apiKeyListResponseJSON `json:"-"`
}

// apiKeyListResponseJSON contains the JSON metadata for the struct
// [APIKeyListResponse]
type apiKeyListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type APIKeyListResponseData struct {
	// API key ID
	ID string `json:"id" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the key is active
	Enabled bool `json:"enabled" api:"required"`
	// Expiration timestamp
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// API key name
	Name string `json:"name" api:"required,nullable"`
	// Permission level
	Permission APIKeyListResponseDataPermission `json:"permission" api:"required"`
	// Key prefix (e.g. rlay*live*)
	Prefix string `json:"prefix" api:"required,nullable"`
	// First 8 characters of the key (preview)
	Start string `json:"start" api:"required"`
	// Workspace access: 'all' or array of workspace IDs
	WorkspaceScope APIKeyListResponseDataWorkspaceScopeUnion `json:"workspace_scope" api:"required"`
	JSON           apiKeyListResponseDataJSON                `json:"-"`
}

// apiKeyListResponseDataJSON contains the JSON metadata for the struct
// [APIKeyListResponseData]
type apiKeyListResponseDataJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Enabled        apijson.Field
	ExpiresAt      apijson.Field
	Name           apijson.Field
	Permission     apijson.Field
	Prefix         apijson.Field
	Start          apijson.Field
	WorkspaceScope apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *APIKeyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Permission level
type APIKeyListResponseDataPermission string

const (
	APIKeyListResponseDataPermissionReadWrite APIKeyListResponseDataPermission = "read_write"
	APIKeyListResponseDataPermissionReadOnly  APIKeyListResponseDataPermission = "read_only"
)

func (r APIKeyListResponseDataPermission) IsKnown() bool {
	switch r {
	case APIKeyListResponseDataPermissionReadWrite, APIKeyListResponseDataPermissionReadOnly:
		return true
	}
	return false
}

// Workspace access: 'all' or array of workspace IDs
//
// Union satisfied by [APIKeyListResponseDataWorkspaceScopeString] or
// [APIKeyListResponseDataWorkspaceScopeArray].
type APIKeyListResponseDataWorkspaceScopeUnion interface {
	implementsAPIKeyListResponseDataWorkspaceScopeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIKeyListResponseDataWorkspaceScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(APIKeyListResponseDataWorkspaceScopeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIKeyListResponseDataWorkspaceScopeArray{}),
		},
	)
}

type APIKeyListResponseDataWorkspaceScopeString string

const (
	APIKeyListResponseDataWorkspaceScopeStringAll APIKeyListResponseDataWorkspaceScopeString = "all"
)

func (r APIKeyListResponseDataWorkspaceScopeString) IsKnown() bool {
	switch r {
	case APIKeyListResponseDataWorkspaceScopeStringAll:
		return true
	}
	return false
}

func (r APIKeyListResponseDataWorkspaceScopeString) implementsAPIKeyListResponseDataWorkspaceScopeUnion() {
}

type APIKeyListResponseDataWorkspaceScopeArray []string

func (r APIKeyListResponseDataWorkspaceScopeArray) implementsAPIKeyListResponseDataWorkspaceScopeUnion() {
}

type APIKeyNewParams struct {
	// Name for the API key
	Name param.Field[string] `json:"name" api:"required"`
	// Number of days until the key expires
	ExpiresInDays param.Field[int64] `json:"expires_in_days"`
	// Permission level: read_write (default) or read_only
	Permission param.Field[APIKeyNewParamsPermission] `json:"permission"`
	// Workspace access: 'all' for unrestricted, or array of workspace IDs
	WorkspaceScope param.Field[APIKeyNewParamsWorkspaceScopeUnion] `json:"workspace_scope"`
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Permission level: read_write (default) or read_only
type APIKeyNewParamsPermission string

const (
	APIKeyNewParamsPermissionReadWrite APIKeyNewParamsPermission = "read_write"
	APIKeyNewParamsPermissionReadOnly  APIKeyNewParamsPermission = "read_only"
)

func (r APIKeyNewParamsPermission) IsKnown() bool {
	switch r {
	case APIKeyNewParamsPermissionReadWrite, APIKeyNewParamsPermissionReadOnly:
		return true
	}
	return false
}

// Workspace access: 'all' for unrestricted, or array of workspace IDs
//
// Satisfied by [APIKeyNewParamsWorkspaceScopeString],
// [APIKeyNewParamsWorkspaceScopeArray].
type APIKeyNewParamsWorkspaceScopeUnion interface {
	implementsAPIKeyNewParamsWorkspaceScopeUnion()
}

type APIKeyNewParamsWorkspaceScopeString string

const (
	APIKeyNewParamsWorkspaceScopeStringAll APIKeyNewParamsWorkspaceScopeString = "all"
)

func (r APIKeyNewParamsWorkspaceScopeString) IsKnown() bool {
	switch r {
	case APIKeyNewParamsWorkspaceScopeStringAll:
		return true
	}
	return false
}

func (r APIKeyNewParamsWorkspaceScopeString) implementsAPIKeyNewParamsWorkspaceScopeUnion() {}

type APIKeyNewParamsWorkspaceScopeArray []string

func (r APIKeyNewParamsWorkspaceScopeArray) implementsAPIKeyNewParamsWorkspaceScopeUnion() {}

type APIKeyListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Filter: start date (ISO 8601)
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Filter: end date (ISO 8601)
	To param.Field[time.Time] `query:"to" format:"date-time"`
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
