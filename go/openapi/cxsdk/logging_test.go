package cxsdk

import (
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
)

func TestLoggingTransportDoesNotLogSensitiveData(t *testing.T) {
	const secret = "synthetic-secret-value"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil || !bytes.Contains(body, []byte(secret)) {
			t.Errorf("request body was not preserved: %v", err)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = io.WriteString(w, `{"value":"`+secret+`"}`)
	}))
	defer server.Close()

	var output bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}))
	client := &http.Client{Transport: NewLoggingTransport(logger)}
	request, err := http.NewRequest(http.MethodPost, server.URL+"?token="+secret, strings.NewReader(`{"apiKey":"`+secret+`"}`))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Authorization", "Bearer "+secret)
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if body, err := io.ReadAll(response.Body); err != nil || !bytes.Contains(body, []byte(secret)) {
		t.Fatalf("response body was not preserved: %v", err)
	}
	if strings.Contains(output.String(), secret) {
		t.Fatal("HTTP logger exposed sensitive data")
	}
	if !strings.Contains(output.String(), `"status_code":201`) {
		t.Fatal("HTTP logger did not record the response status")
	}
}

func TestGeneratedClientDebugDoesNotLogSensitiveData(t *testing.T) {
	const secret = "synthetic-secret-value"
	var output bytes.Buffer
	previous := log.Writer()
	log.SetOutput(&output)
	defer log.SetOutput(previous)

	cfg := apikeys.NewConfiguration()
	cfg.Debug = true
	cfg.Servers = apikeys.ServerConfigurations{{URL: "https://example.test"}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+secret)
	cfg.HTTPClient = &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(request.Body)
		if err != nil || !bytes.Contains(body, []byte(secret)) {
			t.Errorf("generated request body was not preserved: %v", err)
		}
		return &http.Response{
			StatusCode: http.StatusCreated,
			Status:     "201 Created",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"keyId":"test-id","name":"test-name","value":"` + secret + `"}`)),
			Request:    request,
		}, nil
	})}
	client := apikeys.NewAPIClient(cfg)
	name := secret
	request := apikeys.CreateApiKeyRequest{Name: name}
	response, _, err := client.APIKeysServiceAPI.ApiKeysServiceCreateApiKey(context.Background()).CreateApiKeyRequest(request).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if response.GetValue() != secret {
		t.Fatal("generated client did not decode the response")
	}
	if strings.Contains(output.String(), secret) {
		t.Fatal("generated debug logger exposed sensitive data")
	}
	if !strings.Contains(output.String(), "status=201") {
		t.Fatal("generated debug logger did not record the response status")
	}
}
