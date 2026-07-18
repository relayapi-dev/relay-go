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

func TestPostNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.New(context.TODO(), relaygo.PostNewParams{
		ScheduledAt: relaygo.F("now"),
		Targets:     relaygo.F([]string{"string"}),
		Content:     relaygo.F("content"),
		CrossPostActions: relaygo.F([]relaygo.PostNewParamsCrossPostAction{{
			ActionType:      relaygo.F(relaygo.PostNewParamsCrossPostActionsActionTypeRepost),
			TargetAccountID: relaygo.F("target_account_id"),
			Content:         relaygo.F("content"),
			DelayMinutes:    relaygo.F(int64(0)),
		}}),
		IdeaID: relaygo.F("idea_id"),
		Media: relaygo.F([]relaygo.PostNewParamsMedia{{
			URL:       relaygo.F("https://example.com"),
			Thumbnail: relaygo.F("thumbnail"),
			Type:      relaygo.F(relaygo.PostNewParamsMediaTypeImage),
		}}),
		Recycling: relaygo.F(relaygo.PostNewParamsRecycling{
			Gap:               relaygo.F(int64(1)),
			GapFreq:           relaygo.F(relaygo.PostNewParamsRecyclingGapFreqDay),
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

func TestPostGet(t *testing.T) {
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
	_, err := client.Posts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.Update(
		context.TODO(),
		"id",
		relaygo.PostUpdateParams{
			Content: relaygo.F("content"),
			Media: relaygo.F([]relaygo.PostUpdateParamsMedia{{
				URL:       relaygo.F("https://example.com"),
				Thumbnail: relaygo.F("thumbnail"),
				Type:      relaygo.F(relaygo.PostUpdateParamsMediaTypeImage),
			}}),
			Notes: relaygo.F("notes"),
			Recycling: relaygo.F(relaygo.PostUpdateParamsRecycling{
				Gap:               relaygo.F(int64(1)),
				GapFreq:           relaygo.F(relaygo.PostUpdateParamsRecyclingGapFreqDay),
				StartDate:         relaygo.F(time.Now()),
				ContentVariations: relaygo.F([]string{"string"}),
				Enabled:           relaygo.F(true),
				ExpireCount:       relaygo.F(int64(1)),
				ExpireDate:        relaygo.F(time.Now()),
			}),
			ScheduledAt: relaygo.F("now"),
			TargetOptions: relaygo.F(map[string]map[string]interface{}{
				"foo": {
					"foo": "bar",
				},
			}),
			Targets:  relaygo.F([]string{"string"}),
			Timezone: relaygo.F("timezone"),
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

func TestPostListWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.List(context.TODO(), relaygo.PostListParams{
		AccountID:       relaygo.F("account_id"),
		AccountIDs:      relaygo.F("account_ids"),
		Cursor:          relaygo.F("cursor"),
		From:            relaygo.F(time.Now()),
		Include:         relaygo.F("include"),
		IncludeExternal: relaygo.F(relaygo.PostListParamsIncludeExternalTrue),
		Limit:           relaygo.F(int64(1)),
		Status:          relaygo.F(relaygo.PostListParamsStatusDraft),
		To:              relaygo.F(time.Now()),
		WorkspaceID:     relaygo.F("workspace_id"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostDelete(t *testing.T) {
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
	err := client.Posts.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostBulkNew(t *testing.T) {
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
	_, err := client.Posts.BulkNew(context.TODO(), relaygo.PostBulkNewParams{
		Posts: relaygo.F([]relaygo.PostBulkNewParamsPost{{
			ScheduledAt: relaygo.F("now"),
			Targets:     relaygo.F([]string{"string"}),
			Content:     relaygo.F("content"),
			CrossPostActions: relaygo.F([]relaygo.PostBulkNewParamsPostsCrossPostAction{{
				ActionType:      relaygo.F(relaygo.PostBulkNewParamsPostsCrossPostActionsActionTypeRepost),
				TargetAccountID: relaygo.F("target_account_id"),
				Content:         relaygo.F("content"),
				DelayMinutes:    relaygo.F(int64(0)),
			}}),
			IdeaID: relaygo.F("idea_id"),
			Media: relaygo.F([]relaygo.PostBulkNewParamsPostsMedia{{
				URL:       relaygo.F("https://example.com"),
				Thumbnail: relaygo.F("thumbnail"),
				Type:      relaygo.F(relaygo.PostBulkNewParamsPostsMediaTypeImage),
			}}),
			Recycling: relaygo.F(relaygo.PostBulkNewParamsPostsRecycling{
				Gap:               relaygo.F(int64(1)),
				GapFreq:           relaygo.F(relaygo.PostBulkNewParamsPostsRecyclingGapFreqDay),
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
		}}),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostRetry(t *testing.T) {
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
	_, err := client.Posts.Retry(context.TODO(), "id")
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostUnpublishWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.Unpublish(
		context.TODO(),
		"id",
		relaygo.PostUnpublishParams{
			Platforms: relaygo.F([]string{"string"}),
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
