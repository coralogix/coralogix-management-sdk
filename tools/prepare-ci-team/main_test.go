package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

const testApplicationPrefix = "example-ai"

func TestAIApplicationPrefixValidation(t *testing.T) {
	for _, value := range []string{"example-ai", "ai-123"} {
		if !validAIApplicationPrefix(value) {
			t.Fatal("valid application prefix was rejected")
		}
	}
	for _, value := range []string{"", "a", "Uppercase", "a/b", "ai-"} {
		if validAIApplicationPrefix(value) {
			t.Fatal("invalid application prefix was accepted")
		}
	}
}

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func jsonResponse(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(body))}
}

func TestTracePayloads(t *testing.T) {
	for _, test := range []struct {
		name    string
		build   func(time.Time) ([]byte, error)
		service string
		kind    int
		count   int
	}{
		{"AI", func(now time.Time) ([]byte, error) { return aiTracePayload(now, testApplicationPrefix) }, testApplicationPrefix, 3, aiApplicationCount}, {"APM", serverTracePayload, serviceName, 2, 1},
	} {
		payload, err := test.build(time.Now())
		if err != nil {
			t.Fatal(err)
		}
		var doc struct {
			ResourceSpans []struct {
				Resource struct {
					Attributes []struct {
						Key   string `json:"key"`
						Value struct {
							StringValue string `json:"stringValue"`
						} `json:"value"`
					} `json:"attributes"`
				} `json:"resource"`
				ScopeSpans []struct {
					Spans []struct {
						TraceID    string `json:"traceId"`
						SpanID     string `json:"spanId"`
						Kind       int    `json:"kind"`
						Attributes []struct {
							Key   string `json:"key"`
							Value struct {
								StringValue string `json:"stringValue"`
							} `json:"value"`
						} `json:"attributes"`
					} `json:"spans"`
				} `json:"scopeSpans"`
			} `json:"resourceSpans"`
		}
		if err := json.Unmarshal(payload, &doc); err != nil {
			t.Fatal(err)
		}
		if len(doc.ResourceSpans) != test.count {
			t.Fatalf("%s: wrong span count", test.name)
		}
		for index, resourceSpans := range doc.ResourceSpans {
			if len(resourceSpans.ScopeSpans) != 1 || len(resourceSpans.ScopeSpans[0].Spans) != 1 {
				t.Fatalf("%s: wrong scope or span count", test.name)
			}
			attrs := map[string]string{}
			for _, a := range resourceSpans.Resource.Attributes {
				attrs[a.Key] = a.Value.StringValue
			}
			span := resourceSpans.ScopeSpans[0].Spans[0]
			service := test.service
			if test.name == "AI" {
				service = aiApplicationName(testApplicationPrefix, index)
			}
			if attrs["service.name"] != service || test.name == "AI" && attrs["cx.application.name"] != service || attrs["cx.subsystem.name"] != subsystemName || span.Kind != test.kind || len(span.TraceID) != 32 || len(span.SpanID) != 16 {
				t.Fatalf("%s: wrong resource or span", test.name)
			}
			if test.name != "AI" {
				continue
			}
			genAI := map[string]string{}
			for _, a := range span.Attributes {
				genAI[a.Key] = a.Value.StringValue
			}
			for _, key := range []string{"gen_ai.provider.name", "gen_ai.operation.name", "gen_ai.request.model", "gen_ai.input.messages", "gen_ai.output.messages"} {
				if genAI[key] == "" {
					t.Errorf("missing %s", key)
				}
			}
			for _, key := range []string{"gen_ai.input.messages", "gen_ai.output.messages"} {
				var messages []any
				if err := json.Unmarshal([]byte(genAI[key]), &messages); err != nil || len(messages) != 1 {
					t.Errorf("invalid %s", key)
				}
			}
		}
	}
}

func fakeDependencies(t *testing.T, response func(*http.Request) (*http.Response, error)) dependencies {
	t.Helper()
	return dependencies{
		aiPrefix:  testApplicationPrefix,
		client:    &http.Client{Transport: roundTripFunc(response)},
		ingestURL: "https://test.invalid/v1/traces", aiURL: "https://test.invalid/mgmt/openapi/5/ai/applications/v3",
		evaluationsURL: "https://test.invalid/mgmt/openapi/5/ai/evaluations/v3", integrationsURL: "https://test.invalid/mgmt/openapi/5/integrations/integrations/v1",
		archiveURL:  "https://test.invalid/mgmt/openapi/5/logs/data-setup/v2",
		identityURL: "https://test.invalid/mgmt/openapi/5/aaa/identity/v1/whoami",
		apmURL:      "https://test.invalid/v2/entities/ENTITY_TYPE_SERVICE",
		sloURL:      "https://test.invalid/mgmt/openapi/5/slo/slos/v1",
		seedAPM:     func(context.Context, string, time.Time) (func(), error) { return func() {}, nil },
		pollEvery:   time.Millisecond, timeout: time.Second, currentTime: time.Now,
	}
}

func TestProductionManagementURLsUseSDKBase(t *testing.T) {
	deps := productionDependencies()
	base := "https://api.eu2.coralogix.com/mgmt/openapi/5/"
	for _, url := range []string{deps.aiURL, deps.evaluationsURL, deps.integrationsURL, deps.archiveURL, deps.identityURL} {
		if !strings.HasPrefix(url, base) {
			t.Fatalf("management URL does not use SDK base: %q", url)
		}
	}
	if deps.apmURL != "https://api.eu2.coralogix.com/v2/entities/ENTITY_TYPE_SERVICE" {
		t.Fatalf("wrong APM V2 service URL: %q", deps.apmURL)
	}
	if deps.sloURL != "https://api.eu2.coralogix.com/mgmt/openapi/5/slo/slos/v1" {
		t.Fatalf("wrong SLO URL: %q", deps.sloURL)
	}
}

func TestWriteFixtureOutputKeepsIDsOutOfStdout(t *testing.T) {
	privatePath := filepath.Join(t.TempDir(), "fixtures.json")
	result := fixtureOutput{
		Application: testApplicationPrefix, Subsystem: subsystemName, APMService: serviceName,
		TeamID: "1234", PagerDutyIntegrationID: "fixture-pd", PagerDutyConnectorKey: "fixture-key", SlackIntegrationID: "fixture-slack",
	}
	var stdout bytes.Buffer
	if err := writeFixtureOutput(&stdout, privatePath, result); err != nil {
		t.Fatal(err)
	}
	for _, private := range []string{"1234", "fixture-pd", "fixture-key", "fixture-slack"} {
		if strings.Contains(stdout.String(), private) {
			t.Fatal("stdout contains a private fixture value")
		}
	}
	var public struct {
		Application string `json:"application"`
		Subsystem   string `json:"subsystem"`
		APMService  string `json:"apmService"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &public); err != nil || public.Application != testApplicationPrefix || public.Subsystem != subsystemName || public.APMService != serviceName {
		t.Fatal("stdout does not contain the public fixture names")
	}
	data, err := os.ReadFile(privatePath)
	if err != nil || !strings.Contains(string(data), `"teamId":"1234"`) || !strings.Contains(string(data), "fixture-pd") || !strings.Contains(string(data), "fixture-slack") {
		t.Fatal("private fixture file is incomplete")
	}
	info, err := os.Stat(privatePath)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatal("private fixture file permissions are not 0600")
	}
	if err := writeFixtureOutput(&stdout, privatePath, result); err == nil || strings.Contains(err.Error(), privatePath) {
		t.Fatal("existing private fixture file was overwritten or path was disclosed")
	}
}

func TestLookupFixtureIDsUsesReadsOnly(t *testing.T) {
	var reads int
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet || r.Header.Get("Authorization") != "Bearer management-secret" {
			t.Fatal("lookup used a write or the wrong key")
		}
		reads++
		return fixtureResponse(t, r)
	})
	result, err := lookupFixtureIDs(context.Background(), deps, "management-secret")
	if err != nil || reads != 4 || result.TeamID != "1234" || result.SlackIntegrationID != "fixture-slack" || result.PagerDutyIntegrationID != "fixture-pd" || len(result.PagerDutyConnectorKey) != 36 {
		t.Fatalf("lookup did not return the active registrations: reads=%d err=%v", reads, err)
	}
}

func TestLookupFixtureIDsRequiresActivePagerDuty(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet {
			t.Fatal("lookup sent a write")
		}
		if r.URL.Path == "/mgmt/openapi/5/integrations/integrations/v1/pagerduty-incidents" {
			return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[]}}}`), nil
		}
		return fixtureResponse(t, r)
	})
	_, err := lookupFixtureIDs(context.Background(), deps, "management-secret")
	if err == nil || !strings.Contains(err.Error(), "PagerDuty integration is missing") || strings.Contains(err.Error(), "management-secret") {
		t.Fatalf("wrong missing PagerDuty error: %v", err)
	}
}

func TestLookupFixtureIDsRejectsMissingTeamID(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet || r.URL.Path != "/mgmt/openapi/5/aaa/identity/v1/whoami" {
			t.Fatal("lookup continued without a team ID")
		}
		return jsonResponse(200, `{"teamName":"fixture-team"}`), nil
	})
	_, err := lookupFixtureIDs(context.Background(), deps, "management-secret")
	if err == nil || !strings.Contains(err.Error(), "decode team identity") || strings.Contains(err.Error(), "management-secret") {
		t.Fatalf("wrong missing team ID error: %v", err)
	}
}

func TestLookupFixtureIDsRejectsIdentityAuthorization(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet || r.URL.Path != "/mgmt/openapi/5/aaa/identity/v1/whoami" {
			t.Fatal("lookup continued after identity authorization failure")
		}
		return jsonResponse(403, `{"message":"private response"}`), nil
	})
	_, err := lookupFixtureIDs(context.Background(), deps, "management-secret")
	if err == nil || !strings.Contains(err.Error(), "HTTP 403") || strings.Contains(err.Error(), "private response") {
		t.Fatalf("wrong identity authorization error: %v", err)
	}
}

func fixtureResponse(t *testing.T, r *http.Request) (*http.Response, error) {
	t.Helper()
	switch r.URL.Path {
	case "/mgmt/openapi/5/aaa/identity/v1/whoami":
		return jsonResponse(200, `{"teamId":1234,"teamName":"fixture-team"}`), nil
	case "/mgmt/openapi/5/integrations/integrations/v1":
		return jsonResponse(200, `{"integrations":[{"integration":{"id":"pagerduty-incidents"}},{"integration":{"id":"slack-central"}}]}`), nil
	case "/mgmt/openapi/5/integrations/integrations/v1/pagerduty-incidents":
		return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[{"id":"fixture-pd","integrationStatus":{"connectionStatus":"ACTIVE"}}]}}}`), nil
	case "/mgmt/openapi/5/integrations/integrations/v1/slack-central":
		return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[{"id":"fixture-slack","integrationStatus":{"connectionStatus":"ACTIVE"}}]}}}`), nil
	case "/mgmt/openapi/5/logs/data-setup/v2":
		return jsonResponse(200, `{"target":{"archiveSpec":{"isActive":true},"s3":{"bucket":"fixture-bucket"}}}`), nil
	case "/v1/traces":
		return jsonResponse(200, `{}`), nil
	case "/mgmt/openapi/5/ai/applications/v3":
		applications := make([]map[string]string, 0, aiApplicationCount)
		for index := range aiApplicationCount {
			applications = append(applications, map[string]string{"application": aiApplicationName(testApplicationPrefix, index), "subsystem": subsystemName})
		}
		body, err := json.Marshal(map[string]any{"aiApplications": applications})
		if err != nil {
			t.Fatal(err)
		}
		return jsonResponse(200, string(body)), nil
	case "/mgmt/openapi/5/ai/evaluations/v3":
		return jsonResponse(200, `{"aiEvaluations":[]}`), nil
	case "/v2/entities/ENTITY_TYPE_SERVICE":
		return jsonResponse(200, `{"entities":[{"name":"ci-acceptance-apm"}]}`), nil
	case "/mgmt/openapi/5/slo/slos/v1":
		if r.Method != http.MethodPost || r.Header.Get("Authorization") != "Bearer management-secret" {
			t.Errorf("wrong SLO catalog probe request")
		}
		var body struct {
			Name   string `json:"name"`
			APMSli struct {
				Services []string `json:"services"`
			} `json:"apmSli"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name != sloProbeName || len(body.APMSli.Services) != 1 || body.APMSli.Services[0] != serviceName {
			t.Errorf("wrong SLO catalog probe body")
		}
		return jsonResponse(200, `{"slo":{"id":"fixture-slo"}}`), nil
	case "/mgmt/openapi/5/slo/slos/v1/fixture-slo":
		if r.Method != http.MethodDelete || r.Header.Get("Authorization") != "Bearer management-secret" {
			t.Errorf("wrong SLO catalog probe delete")
		}
		return jsonResponse(200, `{}`), nil
	default:
		t.Errorf("unexpected path %q", r.URL.Path)
		return jsonResponse(404, `{}`), nil
	}
}

func TestPreparePollsAndReturnsFixtures(t *testing.T) {
	var aiReads, apmReads, seeds int
	var collectorRunning bool
	var progress bytes.Buffer
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/ai/applications/v3" {
			aiReads++
			if aiReads == 1 {
				return jsonResponse(200, `{"aiApplications":[]}`), nil
			}
		}
		if r.URL.Path == "/v2/entities/ENTITY_TYPE_SERVICE" {
			if !collectorRunning {
				t.Error("Collector stopped before catalog readiness")
			}
			apmReads++
			if r.Header.Get("Authorization") != "Bearer management-secret" {
				t.Error("wrong APM catalog key")
			}
			if apmReads == 1 {
				return jsonResponse(200, `{"entities":[]}`), nil
			}
		}
		if r.URL.Path == "/v1/traces" && r.Header.Get("Authorization") != "Bearer send-secret" {
			t.Error("wrong telemetry key")
		}
		if strings.HasPrefix(r.URL.Path, "/mgmt/openapi/5/integrations/") && r.Header.Get("Authorization") != "Bearer management-secret" {
			t.Error("wrong management key")
		}
		return fixtureResponse(t, r)
	})
	deps.progress = &progress
	deps.seedAPM = func(_ context.Context, key string, _ time.Time) (func(), error) {
		seeds++
		collectorRunning = true
		if key != "send-secret" {
			t.Error("wrong APM key")
		}
		return func() { collectorRunning = false }, nil
	}
	result, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err != nil {
		t.Fatal(err)
	}
	if collectorRunning {
		t.Fatal("Collector was not stopped after catalog readiness")
	}
	if seeds != 1 || aiReads != 2 || apmReads != 2 || result.PagerDutyIntegrationID != "fixture-pd" || result.SlackIntegrationID != "fixture-slack" || len(result.PagerDutyConnectorKey) != 36 {
		t.Fatalf("unexpected fixture result or call counts")
	}
	for _, want := range []string{
		"catalog check 1: AI applications=missing, free evaluation targets=not checked, APM service (V2)=missing, SLO catalog=not checked",
		"catalog check 2: AI applications=ready, free evaluation targets=ready, APM service (V2)=ready, SLO catalog=ready",
	} {
		if !strings.Contains(progress.String(), want) {
			t.Errorf("missing progress state %q", want)
		}
	}
	for _, secret := range []string{"management-secret", "send-secret", "pd-secret", "actor@example.invalid"} {
		if strings.Contains(progress.String(), secret) {
			t.Fatal("progress contains a credential")
		}
	}
}

func TestAIApplicationsRequireAllFixtures(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != "/mgmt/openapi/5/ai/applications/v3" {
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
		return jsonResponse(200, `{"aiApplications":[{"application":"example-ai","subsystem":"ci-acceptance"}]}`), nil
	})
	ready, err := hasAIApplications(context.Background(), deps, "management-secret")
	if err != nil || ready {
		t.Fatalf("one application passed the twenty-application readiness check: ready=%t err=%v", ready, err)
	}
}

func TestAIApplicationsRequireFreeCapacity(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != "/mgmt/openapi/5/ai/evaluations/v3" {
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
		if r.URL.Query().Get("application") == aiApplicationName(testApplicationPrefix, 19) {
			return jsonResponse(200, `{"aiEvaluations":[]}`), nil
		}
		return jsonResponse(200, `{"aiEvaluations":[{"target":"RESPONSE"},{"target":"PROMPT"}]}`), nil
	})
	ready, err := hasUnusedEvaluationTarget(context.Background(), deps, "management-secret")
	if err != nil || ready {
		t.Fatalf("one free application passed the capacity check: ready=%t err=%v", ready, err)
	}
}

const pagerDutyDefinitionResponse = `{"integrationDefinition":{"key":"pagerduty-incidents","revisions":[{"id":"0.0.1","fields":[{"templateParamName":"IntegrationName","type":"TEXT","required":true},{"templateParamName":"ApiKey","type":"SENSITIVE_DATA","required":true},{"templateParamName":"Email","type":"TEXT","required":true},{"templateParamName":"EnableBidirectional","type":"BOOLEAN","required":true},{"templateParamName":"DataProcessingConsent","type":"BOOLEAN","required":true}]}]}}`

func TestCreatePagerDutyAndReuseIt(t *testing.T) {
	var validated, created bool
	var testCalls, createCalls int
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		switch r.URL.Path {
		case "/mgmt/openapi/5/integrations/integrations/v1/pagerduty-incidents":
			if created {
				return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[{"id":"fixture-created","integrationStatus":{"connectionStatus":"ACTIVE"}}]}}}`), nil
			}
			return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[]}}}`), nil
		case "/mgmt/openapi/5/integrations/integrations/v1/definitions/pagerduty-incidents":
			return jsonResponse(200, pagerDutyDefinitionResponse), nil
		case "/mgmt/openapi/5/integrations/metadata/v1/test", "/mgmt/openapi/5/integrations/metadata/v1":
			body, _ := io.ReadAll(r.Body)
			if strings.Contains(string(body), "management-secret") || r.Header.Get("Authorization") != "Bearer management-secret" {
				t.Fatal("wrong integration authorization")
			}
			if strings.Contains(string(body), `"apiKey":`) {
				t.Fatal("PagerDuty API key used the rejected apiKey value arm")
			}
			var request struct {
				Metadata struct {
					IntegrationParameters struct {
						Parameters []struct {
							Key          string `json:"key"`
							StringValue  string `json:"stringValue"`
							BooleanValue *bool  `json:"booleanValue"`
							ApiKey       struct {
								Value string `json:"value"`
							} `json:"apiKey"`
						} `json:"parameters"`
					} `json:"integrationParameters"`
				} `json:"metadata"`
				IntegrationData struct {
					IntegrationParameters struct {
						Parameters []struct {
							Key          string `json:"key"`
							StringValue  string `json:"stringValue"`
							BooleanValue *bool  `json:"booleanValue"`
							ApiKey       struct {
								Value string `json:"value"`
							} `json:"apiKey"`
						} `json:"parameters"`
					} `json:"integrationParameters"`
				} `json:"integrationData"`
			}
			if err := json.Unmarshal(body, &request); err != nil {
				t.Fatal(err)
			}
			parameters := request.IntegrationData.IntegrationParameters.Parameters
			if r.URL.Path == "/mgmt/openapi/5/integrations/metadata/v1" {
				parameters = request.Metadata.IntegrationParameters.Parameters
			}
			seen := map[string]bool{}
			for _, p := range parameters {
				switch p.Key {
				case "ApiKey":
					seen[p.Key] = p.StringValue == "pd-secret" && p.ApiKey.Value == ""
				case "Email":
					seen[p.Key] = p.StringValue == "actor@example.invalid"
				case "DataProcessingConsent", "EnableBidirectional":
					seen[p.Key] = p.BooleanValue != nil && *p.BooleanValue
				case "IntegrationName":
					seen[p.Key] = p.StringValue == pagerDutyName
				}
			}
			if len(seen) != 5 {
				t.Fatal("missing PagerDuty parameters")
			}
			for _, valid := range seen {
				if !valid {
					t.Fatal("wrong PagerDuty parameter value or type")
				}
			}
			if strings.HasSuffix(r.URL.Path, "/test") {
				testCalls++
				validated = true
				return jsonResponse(200, `{"result":{"success":{}}}`), nil
			}
			if !validated {
				t.Fatal("created before validation")
			}
			createCalls++
			created = true
			return jsonResponse(200, `{"integrationId":"fixture-created"}`), nil
		}
		return fixtureResponse(t, r)
	})
	first, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err != nil || first.PagerDutyIntegrationID != "fixture-created" {
		t.Fatalf("create failed: %v", err)
	}
	second, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err != nil || second.PagerDutyIntegrationID != "fixture-created" || testCalls != 1 || createCalls != 1 {
		t.Fatalf("repeat run was not idempotent: %v", err)
	}
}

func TestPagerDutyValidationFailureDoesNotCreate(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		switch r.URL.Path {
		case "/mgmt/openapi/5/integrations/integrations/v1/pagerduty-incidents":
			return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[]}}}`), nil
		case "/mgmt/openapi/5/integrations/integrations/v1/definitions/pagerduty-incidents":
			return jsonResponse(200, pagerDutyDefinitionResponse), nil
		case "/mgmt/openapi/5/integrations/metadata/v1/test":
			return jsonResponse(200, `{"result":{"failure":{}}}`), nil
		case "/mgmt/openapi/5/integrations/metadata/v1", "/v1/traces":
			t.Fatal("wrote after validation failure")
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "validation failed") || strings.Contains(err.Error(), "pd-secret") {
		t.Fatalf("wrong validation failure: %v", err)
	}
}

func TestMissingSlackFailsBeforePagerDutyCreation(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/integrations/integrations/v1/slack-central" {
			return jsonResponse(200, `{"integrationDetail":{"default":{"registered":[]}}}`), nil
		}
		if r.Method == http.MethodPost || r.URL.Path == "/v1/traces" {
			t.Fatal("setup wrote before Slack check")
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "Slack") {
		t.Fatalf("wrong missing Slack error: %v", err)
	}
}

func TestAuthorizationFailsFast(t *testing.T) {
	deps := fakeDependencies(t, func(*http.Request) (*http.Response, error) { return jsonResponse(403, `{}`), nil })
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "HTTP 403") {
		t.Fatalf("wrong authorization error: %v", err)
	}
}

func TestTimeoutNamesMissingCatalogAndTarget(t *testing.T) {
	var collectorStopped bool
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/ai/evaluations/v3" {
			return jsonResponse(200, `{"aiEvaluations":[{"target":"RESPONSE"},{"target":"PROMPT"}]}`), nil
		}
		if r.URL.Path == "/v2/entities/ENTITY_TYPE_SERVICE" {
			return jsonResponse(200, `{"entities":[]}`), nil
		}
		return fixtureResponse(t, r)
	})
	deps.seedAPM = func(context.Context, string, time.Time) (func(), error) {
		return func() { collectorStopped = true }, nil
	}
	deps.timeout = 15 * time.Millisecond
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if !collectorStopped {
		t.Fatal("Collector was not stopped after catalog timeout")
	}
	if err == nil || !strings.Contains(err.Error(), "free targets=false") || !strings.Contains(err.Error(), serviceName) {
		t.Fatalf("wrong timeout error: %v", err)
	}
	for _, secret := range []string{"management-secret", "send-secret"} {
		if strings.Contains(err.Error(), secret) {
			t.Fatal("error disclosed credential")
		}
	}
}

func TestCatalogAuthorizationFailsFast(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/ai/evaluations/v3" {
			return jsonResponse(401, `{}`), nil
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "HTTP 401") {
		t.Fatalf("wrong catalog authorization error: %v", err)
	}
}

func TestCatalogRateLimitRetries(t *testing.T) {
	var reads int
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/ai/applications/v3" {
			reads++
			if reads == 1 {
				return jsonResponse(429, `{}`), nil
			}
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err != nil || reads != 2 {
		t.Fatalf("rate limit was not retried: reads=%d err=%v", reads, err)
	}
}

func TestCollectorConfigUsesEnvironmentReference(t *testing.T) {
	config := collectorConfig(4318)
	for _, want := range []string{"spanmetrics/compact", "namespace: compact", "receivers: [spanmetrics, spanmetrics/compact]", "histogram:\n      unit: ms", "${env:CORALOGIX_SEND_YOUR_DATA_API_KEY}"} {
		if !strings.Contains(config, want) {
			t.Errorf("missing %q", want)
		}
	}
	if strings.Contains(config, "send-secret") {
		t.Fatal("config contains credential value")
	}
}

func TestAPMV2CatalogResponse(t *testing.T) {
	for _, test := range []struct {
		body   string
		status int
		found  bool
		error  string
	}{
		{`{"entities":[{"name":"other"},{"name":"ci-acceptance-apm"}]}`, 200, true, ""},
		{`{"entities":[]}`, 200, false, ""},
		{`{}`, 200, false, "omitted entities"},
		{`{"entities":`, 200, false, "decode APM V2 services"},
		{`{"message":"secret response"}`, 403, false, "HTTP 403"},
	} {
		deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/v2/entities/ENTITY_TYPE_SERVICE" || r.Header.Get("Authorization") != "Bearer management-secret" {
				t.Fatal("wrong APM V2 request")
			}
			return jsonResponse(test.status, test.body), nil
		})
		found, err := hasAPMService(context.Background(), deps, "management-secret")
		if found != test.found || (test.error == "" && err != nil) || (test.error != "" && (err == nil || !strings.Contains(err.Error(), test.error))) {
			t.Fatalf("APM V2 found=%t, err=%v", found, err)
		}
		if err != nil && (strings.Contains(err.Error(), "management-secret") || strings.Contains(err.Error(), "secret response")) {
			t.Fatal("APM V2 error disclosed response or key")
		}
	}
}

func TestSLOCatalogMissRetries(t *testing.T) {
	var creates int
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/slo/slos/v1" && r.Method == http.MethodPost {
			creates++
			if creates == 1 {
				return jsonResponse(400, `{"message":"One or more APM SLO services not found in catalog"}`), nil
			}
		}
		return fixtureResponse(t, r)
	})
	if _, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", ""); err != nil || creates != 2 {
		t.Fatalf("catalog miss was not retried: creates=%d err=%v", creates, err)
	}
}

func TestSLOCatalogOtherErrorStops(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/slo/slos/v1" && r.Method == http.MethodPost {
			return jsonResponse(400, `{"message":"secret response"}`), nil
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "create SLO catalog probe: HTTP 400") || strings.Contains(err.Error(), "secret response") {
		t.Fatalf("wrong SLO catalog error: %v", err)
	}
}

func TestSLOCatalogDeleteFailureStops(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/mgmt/openapi/5/slo/slos/v1/fixture-slo" {
			return jsonResponse(500, `{"message":"secret response"}`), nil
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "delete SLO catalog probe: HTTP 500") || strings.Contains(err.Error(), "secret response") || strings.Contains(err.Error(), "fixture-slo") {
		t.Fatalf("wrong SLO delete error: %v", err)
	}
}

func TestIntegrationErrorNamesStageWithoutResponseValues(t *testing.T) {
	for _, test := range []struct {
		path  string
		stage string
		field string
		body  string
	}{
		{"/integrations/metadata/v1/test", "validate PagerDuty integration", "apiKeyId", `{"message":"apiKeyId must not be empty; key=fake-sensitive-value"}`},
		{"/integrations/metadata/v1", "create PagerDuty integration", "ApiKey", `{"message":"ApiKey is invalid; key=fake-sensitive-value"}`},
	} {
		deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
			return jsonResponse(400, test.body), nil
		})
		var output any
		err := postIntegration(context.Background(), deps, "management-secret", test.path, map[string]any{}, &output)
		if err == nil || !strings.Contains(err.Error(), test.stage) || !strings.Contains(err.Error(), test.field) || strings.Contains(err.Error(), "fake-sensitive-value") {
			t.Fatalf("unsafe or unclear integration error: %v", err)
		}
	}
}
