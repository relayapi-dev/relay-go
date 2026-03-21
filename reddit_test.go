// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/relayapi-dev/relay-go"
	"github.com/relayapi-dev/relay-go/internal/testutil"
	"github.com/relayapi-dev/relay-go/option"
)

func TestRedditGetFeedWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := relaygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Reddit.GetFeed(context.TODO(), relaygo.RedditGetFeedParams{
		AccountID: relaygo.F("account_id"),
		Subreddit: relaygo.F("subreddit"),
		Cursor:    relaygo.F("cursor"),
		Limit:     relaygo.F(int64(1)),
		Sort:      relaygo.F(relaygo.RedditGetFeedParamsSortHot),
		Time:      relaygo.F(relaygo.RedditGetFeedParamsTimeHour),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRedditSearchWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := relaygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Reddit.Search(context.TODO(), relaygo.RedditSearchParams{
		AccountID: relaygo.F("account_id"),
		Query:     relaygo.F("query"),
		Cursor:    relaygo.F("cursor"),
		Limit:     relaygo.F(int64(1)),
		Sort:      relaygo.F(relaygo.RedditSearchParamsSortRelevance),
		Subreddit: relaygo.F("subreddit"),
		Time:      relaygo.F(relaygo.RedditSearchParamsTimeHour),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
