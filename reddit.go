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

// RedditService contains methods and other services that help with interacting
// with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRedditService] method instead.
type RedditService struct {
	Options []option.RequestOption
}

// NewRedditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRedditService(opts ...option.RequestOption) (r *RedditService) {
	r = &RedditService{}
	r.Options = opts
	return
}

// Get subreddit feed
func (r *RedditService) GetFeed(ctx context.Context, query RedditGetFeedParams, opts ...option.RequestOption) (res *RedditGetFeedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/reddit/feed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search Reddit posts
func (r *RedditService) Search(ctx context.Context, query RedditSearchParams, opts ...option.RequestOption) (res *RedditSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/reddit/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type RedditGetFeedResponse struct {
	Data       []RedditGetFeedResponseData `json:"data" api:"required"`
	HasMore    bool                        `json:"has_more" api:"required"`
	NextCursor string                      `json:"next_cursor" api:"required,nullable"`
	JSON       redditGetFeedResponseJSON   `json:"-"`
}

// redditGetFeedResponseJSON contains the JSON metadata for the struct
// [RedditGetFeedResponse]
type redditGetFeedResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedditGetFeedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redditGetFeedResponseJSON) RawJSON() string {
	return r.raw
}

type RedditGetFeedResponseData struct {
	// Reddit post ID
	ID string `json:"id" api:"required"`
	// Post author
	Author string `json:"author" api:"required"`
	// Created timestamp (Unix)
	CreatedUtc float64 `json:"created_utc" api:"required"`
	// Whether it's a self post
	IsSelf bool `json:"is_self" api:"required"`
	// Whether NSFW
	Nsfw bool `json:"nsfw" api:"required"`
	// Comment count
	NumComments float64 `json:"num_comments" api:"required"`
	// Post score
	Score float64 `json:"score" api:"required"`
	// Subreddit name
	Subreddit string `json:"subreddit" api:"required"`
	// Post title
	Title string `json:"title" api:"required"`
	// Post URL
	URL string `json:"url" api:"required"`
	// Self text
	Selftext string `json:"selftext" api:"nullable"`
	// Thumbnail URL
	Thumbnail string                        `json:"thumbnail" api:"nullable"`
	JSON      redditGetFeedResponseDataJSON `json:"-"`
}

// redditGetFeedResponseDataJSON contains the JSON metadata for the struct
// [RedditGetFeedResponseData]
type redditGetFeedResponseDataJSON struct {
	ID          apijson.Field
	Author      apijson.Field
	CreatedUtc  apijson.Field
	IsSelf      apijson.Field
	Nsfw        apijson.Field
	NumComments apijson.Field
	Score       apijson.Field
	Subreddit   apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	Selftext    apijson.Field
	Thumbnail   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedditGetFeedResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redditGetFeedResponseDataJSON) RawJSON() string {
	return r.raw
}

type RedditSearchResponse struct {
	Data       []RedditSearchResponseData `json:"data" api:"required"`
	HasMore    bool                       `json:"has_more" api:"required"`
	NextCursor string                     `json:"next_cursor" api:"required,nullable"`
	JSON       redditSearchResponseJSON   `json:"-"`
}

// redditSearchResponseJSON contains the JSON metadata for the struct
// [RedditSearchResponse]
type redditSearchResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedditSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redditSearchResponseJSON) RawJSON() string {
	return r.raw
}

type RedditSearchResponseData struct {
	// Reddit post ID
	ID string `json:"id" api:"required"`
	// Post author
	Author string `json:"author" api:"required"`
	// Created timestamp (Unix)
	CreatedUtc float64 `json:"created_utc" api:"required"`
	// Whether it's a self post
	IsSelf bool `json:"is_self" api:"required"`
	// Whether NSFW
	Nsfw bool `json:"nsfw" api:"required"`
	// Comment count
	NumComments float64 `json:"num_comments" api:"required"`
	// Post score
	Score float64 `json:"score" api:"required"`
	// Subreddit name
	Subreddit string `json:"subreddit" api:"required"`
	// Post title
	Title string `json:"title" api:"required"`
	// Post URL
	URL string `json:"url" api:"required"`
	// Self text
	Selftext string `json:"selftext" api:"nullable"`
	// Thumbnail URL
	Thumbnail string                       `json:"thumbnail" api:"nullable"`
	JSON      redditSearchResponseDataJSON `json:"-"`
}

// redditSearchResponseDataJSON contains the JSON metadata for the struct
// [RedditSearchResponseData]
type redditSearchResponseDataJSON struct {
	ID          apijson.Field
	Author      apijson.Field
	CreatedUtc  apijson.Field
	IsSelf      apijson.Field
	Nsfw        apijson.Field
	NumComments apijson.Field
	Score       apijson.Field
	Subreddit   apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	Selftext    apijson.Field
	Thumbnail   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedditSearchResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redditSearchResponseDataJSON) RawJSON() string {
	return r.raw
}

type RedditGetFeedParams struct {
	// Reddit account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
	// Subreddit name
	Subreddit param.Field[string] `query:"subreddit" api:"required"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Sort order
	Sort param.Field[RedditGetFeedParamsSort] `query:"sort"`
	// Time filter (for top sort)
	Time param.Field[RedditGetFeedParamsTime] `query:"time"`
}

// URLQuery serializes [RedditGetFeedParams]'s query parameters as `url.Values`.
func (r RedditGetFeedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type RedditGetFeedParamsSort string

const (
	RedditGetFeedParamsSortHot    RedditGetFeedParamsSort = "hot"
	RedditGetFeedParamsSortNew    RedditGetFeedParamsSort = "new"
	RedditGetFeedParamsSortTop    RedditGetFeedParamsSort = "top"
	RedditGetFeedParamsSortRising RedditGetFeedParamsSort = "rising"
)

func (r RedditGetFeedParamsSort) IsKnown() bool {
	switch r {
	case RedditGetFeedParamsSortHot, RedditGetFeedParamsSortNew, RedditGetFeedParamsSortTop, RedditGetFeedParamsSortRising:
		return true
	}
	return false
}

// Time filter (for top sort)
type RedditGetFeedParamsTime string

const (
	RedditGetFeedParamsTimeHour  RedditGetFeedParamsTime = "hour"
	RedditGetFeedParamsTimeDay   RedditGetFeedParamsTime = "day"
	RedditGetFeedParamsTimeWeek  RedditGetFeedParamsTime = "week"
	RedditGetFeedParamsTimeMonth RedditGetFeedParamsTime = "month"
	RedditGetFeedParamsTimeYear  RedditGetFeedParamsTime = "year"
	RedditGetFeedParamsTimeAll   RedditGetFeedParamsTime = "all"
)

func (r RedditGetFeedParamsTime) IsKnown() bool {
	switch r {
	case RedditGetFeedParamsTimeHour, RedditGetFeedParamsTimeDay, RedditGetFeedParamsTimeWeek, RedditGetFeedParamsTimeMonth, RedditGetFeedParamsTimeYear, RedditGetFeedParamsTimeAll:
		return true
	}
	return false
}

type RedditSearchParams struct {
	// Reddit account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
	// Search query
	Query param.Field[string] `query:"query" api:"required"`
	// Pagination cursor
	Cursor param.Field[string] `query:"cursor"`
	// Number of items per page
	Limit param.Field[int64] `query:"limit"`
	// Sort order
	Sort param.Field[RedditSearchParamsSort] `query:"sort"`
	// Limit to subreddit
	Subreddit param.Field[string] `query:"subreddit"`
	// Time filter
	Time param.Field[RedditSearchParamsTime] `query:"time"`
}

// URLQuery serializes [RedditSearchParams]'s query parameters as `url.Values`.
func (r RedditSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type RedditSearchParamsSort string

const (
	RedditSearchParamsSortRelevance RedditSearchParamsSort = "relevance"
	RedditSearchParamsSortHot       RedditSearchParamsSort = "hot"
	RedditSearchParamsSortTop       RedditSearchParamsSort = "top"
	RedditSearchParamsSortNew       RedditSearchParamsSort = "new"
	RedditSearchParamsSortComments  RedditSearchParamsSort = "comments"
)

func (r RedditSearchParamsSort) IsKnown() bool {
	switch r {
	case RedditSearchParamsSortRelevance, RedditSearchParamsSortHot, RedditSearchParamsSortTop, RedditSearchParamsSortNew, RedditSearchParamsSortComments:
		return true
	}
	return false
}

// Time filter
type RedditSearchParamsTime string

const (
	RedditSearchParamsTimeHour  RedditSearchParamsTime = "hour"
	RedditSearchParamsTimeDay   RedditSearchParamsTime = "day"
	RedditSearchParamsTimeWeek  RedditSearchParamsTime = "week"
	RedditSearchParamsTimeMonth RedditSearchParamsTime = "month"
	RedditSearchParamsTimeYear  RedditSearchParamsTime = "year"
	RedditSearchParamsTimeAll   RedditSearchParamsTime = "all"
)

func (r RedditSearchParamsTime) IsKnown() bool {
	switch r {
	case RedditSearchParamsTimeHour, RedditSearchParamsTimeDay, RedditSearchParamsTimeWeek, RedditSearchParamsTimeMonth, RedditSearchParamsTimeYear, RedditSearchParamsTimeAll:
		return true
	}
	return false
}
