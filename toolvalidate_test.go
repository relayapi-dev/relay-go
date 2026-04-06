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

func TestToolValidateCheckPostLength(t *testing.T) {
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
	_, err := client.Tools.Validate.CheckPostLength(context.TODO(), relaygo.ToolValidateCheckPostLengthParams{
		Content: relaygo.F("content"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolValidateGetSubreddit(t *testing.T) {
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
	_, err := client.Tools.Validate.GetSubreddit(context.TODO(), relaygo.ToolValidateGetSubredditParams{
		Name: relaygo.F("name"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolValidateValidateMedia(t *testing.T) {
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
	_, err := client.Tools.Validate.ValidateMedia(context.TODO(), relaygo.ToolValidateValidateMediaParams{
		URL: relaygo.F("https://example.com"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolValidateValidatePostWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Validate.ValidatePost(context.TODO(), relaygo.ToolValidateValidatePostParams{
		ScheduledAt: relaygo.F("now"),
		Targets:     relaygo.F([]string{"string"}),
		Content:     relaygo.F("content"),
		CrossPostActions: relaygo.F([]relaygo.ToolValidateValidatePostParamsCrossPostAction{{
			ActionType:      relaygo.F(relaygo.ToolValidateValidatePostParamsCrossPostActionsActionTypeRepost),
			TargetAccountID: relaygo.F("target_account_id"),
			Content:         relaygo.F("content"),
			DelayMinutes:    relaygo.F(int64(0)),
		}}),
		Media: relaygo.F([]relaygo.ToolValidateValidatePostParamsMedia{{
			URL:  relaygo.F("https://example.com"),
			Type: relaygo.F(relaygo.ToolValidateValidatePostParamsMediaTypeImage),
		}}),
		Recycling: relaygo.F(relaygo.ToolValidateValidatePostParamsRecycling{
			Gap:               relaygo.F(int64(1)),
			GapFreq:           relaygo.F(relaygo.ToolValidateValidatePostParamsRecyclingGapFreqDay),
			StartDate:         relaygo.F(time.Now()),
			ContentVariations: relaygo.F([]string{"string"}),
			Enabled:           relaygo.F(true),
			ExpireCount:       relaygo.F(int64(1)),
			ExpireDate:        relaygo.F(time.Now()),
		}),
		ShortenURLs:   relaygo.F(true),
		SkipSignature: relaygo.F(true),
		TargetOptions: relaygo.F(map[string]map[string]interface{}{
			"foo": {
				"foo": "bar",
			},
		}),
		TemplateID: relaygo.F("template_id"),
		TemplateVariables: relaygo.F(map[string]string{
			"foo": "string",
		}),
		Timezone:    relaygo.F("timezone"),
		WorkspaceID: relaygo.F("workspace_id"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
