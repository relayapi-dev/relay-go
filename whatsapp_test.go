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

func TestWhatsappBulkSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Whatsapp.BulkSend(context.TODO(), relaygo.WhatsappBulkSendParams{
		AccountID: relaygo.F("account_id"),
		Recipients: relaygo.F([]relaygo.WhatsappBulkSendParamsRecipient{{
			Phone: relaygo.F("phone"),
			Variables: relaygo.F(map[string]string{
				"foo": "string",
			}),
		}}),
		Template: relaygo.F(relaygo.WhatsappBulkSendParamsTemplate{
			Language: relaygo.F("language"),
			Name:     relaygo.F("name"),
			Components: relaygo.F([]relaygo.WhatsappBulkSendParamsTemplateComponent{{
				Type: relaygo.F(relaygo.WhatsappBulkSendParamsTemplateComponentsTypeHeader),
				Parameters: relaygo.F([]map[string]interface{}{{
					"foo": "bar",
				}}),
			}}),
		}),
	})
	if err != nil {
		var apierr *relaygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWhatsappListPhoneNumbers(t *testing.T) {
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
	_, err := client.Whatsapp.ListPhoneNumbers(context.TODO(), relaygo.WhatsappListPhoneNumbersParams{
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
