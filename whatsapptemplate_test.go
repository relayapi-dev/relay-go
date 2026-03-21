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

func TestWhatsappTemplateNew(t *testing.T) {
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
	_, err := client.Whatsapp.Templates.New(context.TODO(), relaygo.WhatsappTemplateNewParams{
		AccountID: relaygo.F("account_id"),
		Category:  relaygo.F(relaygo.WhatsappTemplateNewParamsCategoryMarketing),
		Components: relaygo.F([]relaygo.WhatsappTemplateNewParamsComponent{{
			Type: relaygo.F(relaygo.WhatsappTemplateNewParamsComponentsTypeHeader),
			Buttons: relaygo.F([]relaygo.WhatsappTemplateNewParamsComponentsButton{{
				Text:        relaygo.F("text"),
				Type:        relaygo.F("type"),
				PhoneNumber: relaygo.F("phone_number"),
				URL:         relaygo.F("url"),
			}}),
			Format: relaygo.F("format"),
			Text:   relaygo.F("text"),
		}}),
		Language: relaygo.F("language"),
		Name:     relaygo.F("name"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWhatsappTemplateGet(t *testing.T) {
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
	_, err := client.Whatsapp.Templates.Get(
		context.TODO(),
		"template_name",
		relaygo.WhatsappTemplateGetParams{
			AccountID: relaygo.F("account_id"),
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

func TestWhatsappTemplateList(t *testing.T) {
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
	_, err := client.Whatsapp.Templates.List(context.TODO(), relaygo.WhatsappTemplateListParams{
		AccountID: relaygo.F("account_id"),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWhatsappTemplateDelete(t *testing.T) {
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
	err := client.Whatsapp.Templates.Delete(
		context.TODO(),
		"template_name",
		relaygo.WhatsappTemplateDeleteParams{
			AccountID: relaygo.F("account_id"),
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
