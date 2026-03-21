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

// PostLogService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostLogService] method instead.
type PostLogService struct {
	Options []option.RequestOption
}

// NewPostLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPostLogService(opts ...option.RequestOption) (r *PostLogService) {
	r = &PostLogService{}
	r.Options = opts
	return
}

// Get publishing logs for a post
func (r *PostLogService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PostLogGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/posts/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query publishing logs across all posts with pagination.
func (r *PostLogService) List(ctx context.Context, query PostLogListParams, opts ...option.RequestOption) (res *PostLogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/posts/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type PostLogGetResponse struct {
	Data       []PostLogGetResponseData `json:"data" api:"required"`
	HasMore    bool                     `json:"has_more" api:"required"`
	NextCursor string                   `json:"next_cursor" api:"required,nullable"`
	JSON       postLogGetResponseJSON   `json:"-"`
}

// postLogGetResponseJSON contains the JSON metadata for the struct
// [PostLogGetResponse]
type postLogGetResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type PostLogGetResponseData struct {
	// Log entry ID (post target ID)
	ID string `json:"id" api:"required"`
	// Error message if failed
	Error string `json:"error" api:"required,nullable"`
	// Platform name
	Platform string `json:"platform" api:"required"`
	// Platform post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Published URL
	PlatformURL string `json:"platform_url" api:"required,nullable"`
	// Post ID
	PostID string `json:"post_id" api:"required"`
	// Published timestamp
	PublishedAt time.Time `json:"published_at" api:"required,nullable" format:"date-time"`
	// Social account ID
	SocialAccountID string `json:"social_account_id" api:"required"`
	// Target status
	Status string `json:"status" api:"required"`
	// Last updated
	UpdatedAt time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	JSON      postLogGetResponseDataJSON `json:"-"`
}

// postLogGetResponseDataJSON contains the JSON metadata for the struct
// [PostLogGetResponseData]
type postLogGetResponseDataJSON struct {
	ID              apijson.Field
	Error           apijson.Field
	Platform        apijson.Field
	PlatformPostID  apijson.Field
	PlatformURL     apijson.Field
	PostID          apijson.Field
	PublishedAt     apijson.Field
	SocialAccountID apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PostLogGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postLogGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type PostLogListResponse struct {
	Data       []PostLogListResponseData `json:"data" api:"required"`
	HasMore    bool                      `json:"has_more" api:"required"`
	NextCursor string                    `json:"next_cursor" api:"required,nullable"`
	JSON       postLogListResponseJSON   `json:"-"`
}

// postLogListResponseJSON contains the JSON metadata for the struct
// [PostLogListResponse]
type postLogListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postLogListResponseJSON) RawJSON() string {
	return r.raw
}

type PostLogListResponseData struct {
	// Log entry ID (post target ID)
	ID string `json:"id" api:"required"`
	// Error message if failed
	Error string `json:"error" api:"required,nullable"`
	// Platform name
	Platform string `json:"platform" api:"required"`
	// Platform post ID
	PlatformPostID string `json:"platform_post_id" api:"required,nullable"`
	// Published URL
	PlatformURL string `json:"platform_url" api:"required,nullable"`
	// Post ID
	PostID string `json:"post_id" api:"required"`
	// Published timestamp
	PublishedAt time.Time `json:"published_at" api:"required,nullable" format:"date-time"`
	// Social account ID
	SocialAccountID string `json:"social_account_id" api:"required"`
	// Target status
	Status string `json:"status" api:"required"`
	// Last updated
	UpdatedAt time.Time                   `json:"updated_at" api:"required" format:"date-time"`
	JSON      postLogListResponseDataJSON `json:"-"`
}

// postLogListResponseDataJSON contains the JSON metadata for the struct
// [PostLogListResponseData]
type postLogListResponseDataJSON struct {
	ID              apijson.Field
	Error           apijson.Field
	Platform        apijson.Field
	PlatformPostID  apijson.Field
	PlatformURL     apijson.Field
	PostID          apijson.Field
	PublishedAt     apijson.Field
	SocialAccountID apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PostLogListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postLogListResponseDataJSON) RawJSON() string {
	return r.raw
}

type PostLogListParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PostLogListParams]'s query parameters as `url.Values`.
func (r PostLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
