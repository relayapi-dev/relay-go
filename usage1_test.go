// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo_test

import (
	"context"
	"os"
	"testing"

	"github.com/relayapi-dev/relay-go"
	"github.com/relayapi-dev/relay-go/internal/testutil"
	"github.com/relayapi-dev/relay-go/option"
)

func TestUsage(t *testing.T) {
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
	posts, err := client.Posts.List(context.TODO(), relaygo.PostListParams{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", posts.Data)
}
