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
	// Key prefix
	Prefix string                `json:"prefix" api:"required"`
	JSON   apiKeyNewResponseJSON `json:"-"`
}

// apiKeyNewResponseJSON contains the JSON metadata for the struct
// [APIKeyNewResponse]
type apiKeyNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ExpiresAt   apijson.Field
	Key         apijson.Field
	Name        apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

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
	// Key prefix (e.g. rlay*live*)
	Prefix string `json:"prefix" api:"required,nullable"`
	// First 8 characters of the key (preview)
	Start string                     `json:"start" api:"required"`
	JSON  apiKeyListResponseDataJSON `json:"-"`
}

// apiKeyListResponseDataJSON contains the JSON metadata for the struct
// [APIKeyListResponseData]
type apiKeyListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ExpiresAt   apijson.Field
	Name        apijson.Field
	Prefix      apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIKeyNewParams struct {
	// Name for the API key
	Name param.Field[string] `json:"name" api:"required"`
	// Number of days until the key expires
	ExpiresInDays param.Field[int64] `json:"expires_in_days"`
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
