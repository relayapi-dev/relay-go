// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/relayapi-dev/relay-go"
	"github.com/relayapi-dev/relay-go/internal/testutil"
	"github.com/relayapi-dev/relay-go/option"
)

func TestAccountGet(t *testing.T) {
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
	_, err := client.Accounts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Update(
		context.TODO(),
		"id",
		relaygo.AccountUpdateParams{
			DisplayName: relaygo.F("display_name"),
			GroupID:     relaygo.F("group_id"),
			Metadata: relaygo.F(map[string]interface{}{
				"foo": "bar",
			}),
		},
	)
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.List(context.TODO(), relaygo.AccountListParams{
		Cursor:    relaygo.F("cursor"),
		From:      relaygo.F(time.Now()),
		GroupID:   relaygo.F("group_id"),
		Limit:     relaygo.F(int64(1)),
		Search:    relaygo.F("search"),
		To:        relaygo.F(time.Now()),
		Ungrouped: relaygo.F(true),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDelete(t *testing.T) {
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
	err := client.Accounts.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
