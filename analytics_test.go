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

func TestAnalyticsGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Analytics.Get(context.TODO(), relaygo.AnalyticsGetParams{
		AccountID: relaygo.F("account_id"),
		FromDate:  relaygo.F("from_date"),
		Limit:     relaygo.F(int64(1)),
		Offset:    relaygo.F(int64(0)),
		Platform:  relaygo.F(relaygo.AnalyticsGetParamsPlatformTwitter),
		PostID:    relaygo.F("post_id"),
		ToDate:    relaygo.F("to_date"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsGetBestTimeWithOptionalParams(t *testing.T) {
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
	_, err := client.Analytics.GetBestTime(context.TODO(), relaygo.AnalyticsGetBestTimeParams{
		AccountID: relaygo.F("account_id"),
		FromDate:  relaygo.F("from_date"),
		Platform:  relaygo.F(relaygo.AnalyticsGetBestTimeParamsPlatformTwitter),
		ToDate:    relaygo.F("to_date"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsGetContentDecayWithOptionalParams(t *testing.T) {
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
	_, err := client.Analytics.GetContentDecay(context.TODO(), relaygo.AnalyticsGetContentDecayParams{
		PostID: relaygo.F("post_id"),
		Days:   relaygo.F(int64(1)),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsGetPostTimelineWithOptionalParams(t *testing.T) {
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
	_, err := client.Analytics.GetPostTimeline(context.TODO(), relaygo.AnalyticsGetPostTimelineParams{
		PostID:   relaygo.F("post_id"),
		FromDate: relaygo.F("from_date"),
		ToDate:   relaygo.F("to_date"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsGetPostingFrequencyWithOptionalParams(t *testing.T) {
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
	_, err := client.Analytics.GetPostingFrequency(context.TODO(), relaygo.AnalyticsGetPostingFrequencyParams{
		AccountID: relaygo.F("account_id"),
		FromDate:  relaygo.F("from_date"),
		Platform:  relaygo.F(relaygo.AnalyticsGetPostingFrequencyParamsPlatformTwitter),
		ToDate:    relaygo.F("to_date"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsListDailyMetricsWithOptionalParams(t *testing.T) {
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
	_, err := client.Analytics.ListDailyMetrics(context.TODO(), relaygo.AnalyticsListDailyMetricsParams{
		AccountID: relaygo.F("account_id"),
		FromDate:  relaygo.F("from_date"),
		Platform:  relaygo.F(relaygo.AnalyticsListDailyMetricsParamsPlatformTwitter),
		ToDate:    relaygo.F("to_date"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
