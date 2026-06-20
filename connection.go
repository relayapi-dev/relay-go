// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
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

// ConnectionService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

// Returns connection event history for the organization.
func (r *ConnectionService) ListLogs(ctx context.Context, query ConnectionListLogsParams, opts ...option.RequestOption) (res *ConnectionListLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connections/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConnectionListLogsResponse struct {
	Data       []ConnectionListLogsResponseData `json:"data" api:"required"`
	HasMore    bool                             `json:"has_more" api:"required"`
	NextCursor string                           `json:"next_cursor" api:"required,nullable"`
	// Total matching log entries (ignores pagination)
	Total float64                        `json:"total" api:"required"`
	JSON  connectionListLogsResponseJSON `json:"-"`
}

// connectionListLogsResponseJSON contains the JSON metadata for the struct
// [ConnectionListLogsResponse]
type connectionListLogsResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListLogsResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionListLogsResponseData struct {
	// Log entry ID
	ID string `json:"id" api:"required"`
	// Social account ID
	AccountID string `json:"account_id" api:"required,nullable"`
	// Timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Event type
	Event ConnectionListLogsResponseDataEvent `json:"event" api:"required"`
	// Event details
	Message string `json:"message" api:"required,nullable"`
	// Platform name
	Platform string                             `json:"platform" api:"required"`
	JSON     connectionListLogsResponseDataJSON `json:"-"`
}

// connectionListLogsResponseDataJSON contains the JSON metadata for the struct
// [ConnectionListLogsResponseData]
type connectionListLogsResponseDataJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	CreatedAt   apijson.Field
	Event       apijson.Field
	Message     apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListLogsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListLogsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Event type
type ConnectionListLogsResponseDataEvent string

const (
	ConnectionListLogsResponseDataEventConnected      ConnectionListLogsResponseDataEvent = "connected"
	ConnectionListLogsResponseDataEventDisconnected   ConnectionListLogsResponseDataEvent = "disconnected"
	ConnectionListLogsResponseDataEventTokenRefreshed ConnectionListLogsResponseDataEvent = "token_refreshed"
	ConnectionListLogsResponseDataEventError          ConnectionListLogsResponseDataEvent = "error"
)

func (r ConnectionListLogsResponseDataEvent) IsKnown() bool {
	switch r {
	case ConnectionListLogsResponseDataEventConnected, ConnectionListLogsResponseDataEventDisconnected, ConnectionListLogsResponseDataEventTokenRefreshed, ConnectionListLogsResponseDataEventError:
		return true
	}
	return false
}

type ConnectionListLogsParams struct {
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Filter: start date (ISO 8601)
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip for offset-based pagination. Enables random access to
	// any page; takes precedence over `cursor` when provided.
	Offset param.Field[int64] `query:"offset"`
	// Filter: end date (ISO 8601)
	To param.Field[time.Time] `query:"to" format:"date-time"`
}

// URLQuery serializes [ConnectionListLogsParams]'s query parameters as
// `url.Values`.
func (r ConnectionListLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
