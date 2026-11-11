// prepare-ci-team seeds the small telemetry fixtures used by live CI tests.
package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	applicationName = "ci-acceptance-ai"
	subsystemName   = "ci-acceptance"
	serviceName     = "ci-acceptance-apm"

	ingestURL       = "https://ingress.eu2.coralogix.com/v1/traces"
	aiURL           = "https://api.eu2.coralogix.com/mgmt/openapi/5/ai/applications/v3"
	evaluationsURL  = "https://api.eu2.coralogix.com/mgmt/openapi/5/ai/evaluations/v3"
	integrationsURL = "https://api.eu2.coralogix.com/mgmt/openapi/5/integrations/integrations/v1"
	archiveURL      = "https://api.eu2.coralogix.com/mgmt/openapi/5/logs/data-setup/v2"
	identityURL     = "https://api.eu2.coralogix.com/mgmt/openapi/5/aaa/identity/v1/whoami"
	apmURL          = "https://api.eu2.coralogix.com/v2/entities/ENTITY_TYPE_SERVICE"
)

type dependencies struct {
	client          *http.Client
	ingestURL       string
	aiURL           string
	evaluationsURL  string
	integrationsURL string
	archiveURL      string
	identityURL     string
	apmURL          string
	seedAPM         func(context.Context, string, time.Time) error
	progress        io.Writer
	pollEvery       time.Duration
	timeout         time.Duration
	currentTime     func() time.Time
}

func productionDependencies() dependencies {
	return dependencies{
		client: &http.Client{
			Timeout:       15 * time.Second,
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		},
		ingestURL:       ingestURL,
		aiURL:           aiURL,
		evaluationsURL:  evaluationsURL,
		integrationsURL: integrationsURL,
		archiveURL:      archiveURL,
		identityURL:     identityURL,
		apmURL:          apmURL,
		seedAPM:         seedAPMService,
		progress:        os.Stderr,
		pollEvery:       10 * time.Second,
		timeout:         10 * time.Minute,
		currentTime:     time.Now,
	}
}

func main() {
	if len(os.Args) > 2 || len(os.Args) == 2 && os.Args[1] != "lookup" {
		fmt.Fprintln(os.Stderr, "prepare CI team: usage: prepare-ci-team [lookup]")
		os.Exit(2)
	}
	privatePath := os.Getenv("PREPARE_CI_TEAM_OUTPUT_FILE")
	if len(os.Args) == 2 {
		if privatePath == "" {
			fmt.Fprintln(os.Stderr, "prepare CI team: PREPARE_CI_TEAM_OUTPUT_FILE is required for lookup")
			os.Exit(1)
		}
		managementKey := os.Getenv("CORALOGIX_API_KEY")
		if managementKey == "" {
			fmt.Fprintln(os.Stderr, "prepare CI team: CORALOGIX_API_KEY is required for lookup")
			os.Exit(1)
		}
		result, err := lookupFixtureIDs(context.Background(), productionDependencies(), managementKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "prepare CI team: %v\n", err)
			os.Exit(1)
		}
		if err := writeFixtureOutput(os.Stdout, privatePath, result); err != nil {
			fmt.Fprintf(os.Stderr, "prepare CI team: %v\n", err)
			os.Exit(1)
		}
		return
	}
	if _, err := exec.LookPath("otelcol-contrib"); err != nil {
		fmt.Fprintln(os.Stderr, "prepare CI team: otelcol-contrib is required to generate APM Span Metrics")
		os.Exit(1)
	}

	managementKey := os.Getenv("CORALOGIX_API_KEY")
	sendDataKey := os.Getenv("CORALOGIX_SEND_YOUR_DATA_API_KEY")
	pagerDutyKey := os.Getenv("PAGERDUTY_REST_API_KEY")
	pagerDutyEmail := os.Getenv("PAGERDUTY_USER_EMAIL")
	if managementKey == "" || sendDataKey == "" || pagerDutyKey == "" || pagerDutyEmail == "" {
		fmt.Fprintln(os.Stderr, "prepare CI team: CORALOGIX_API_KEY, CORALOGIX_SEND_YOUR_DATA_API_KEY, PAGERDUTY_REST_API_KEY, and PAGERDUTY_USER_EMAIL are required")
		os.Exit(1)
	}

	result, err := prepare(context.Background(), productionDependencies(), managementKey, sendDataKey, pagerDutyKey, pagerDutyEmail, os.Getenv("ARCHIVE_LOGS_BUCKET"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "prepare CI team: %v\n", err)
		os.Exit(1)
	}

	if err := writeFixtureOutput(os.Stdout, privatePath, result); err != nil {
		fmt.Fprintf(os.Stderr, "prepare CI team: %v\n", err)
		os.Exit(1)
	}
}

func lookupFixtureIDs(ctx context.Context, deps dependencies, managementKey string) (fixtureOutput, error) {
	var output fixtureOutput
	if managementKey == "" {
		return output, errors.New("CORALOGIX_API_KEY is required for lookup")
	}
	teamID, err := lookupTeamID(ctx, deps, managementKey)
	if err != nil {
		return output, err
	}
	slackID, err := findIntegrationPrerequisites(ctx, deps, managementKey)
	if err != nil {
		return output, err
	}
	pagerDutyID, status, err := readPagerDutyRegistration(ctx, deps, managementKey)
	if err != nil {
		return output, err
	}
	if pagerDutyID == "" || status != "ACTIVE" {
		return output, errors.New("active account-level PagerDuty integration is missing")
	}
	connectorKey, err := randomUUID()
	if err != nil {
		return output, errors.New("generate connector test value")
	}
	return fixtureOutput{
		Application: applicationName, Subsystem: subsystemName, APMService: serviceName,
		TeamID: teamID, PagerDutyIntegrationID: pagerDutyID,
		PagerDutyConnectorKey: connectorKey, SlackIntegrationID: slackID,
	}, nil
}

type fixtureOutput struct {
	Application            string `json:"application"`
	Subsystem              string `json:"subsystem"`
	APMService             string `json:"apmService"`
	TeamID                 string `json:"teamId,omitempty"`
	PagerDutyIntegrationID string `json:"pagerDutyIntegrationId"`
	PagerDutyConnectorKey  string `json:"pagerDutyConnectorKey"`
	SlackIntegrationID     string `json:"slackIntegrationId"`
}

func writeFixtureOutput(stdout io.Writer, privatePath string, result fixtureOutput) error {
	if privatePath != "" {
		file, err := os.OpenFile(privatePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			return errors.New("create private fixture output")
		}
		if err := json.NewEncoder(file).Encode(result); err != nil {
			file.Close()
			os.Remove(privatePath)
			return errors.New("write private fixture output")
		}
		if err := file.Close(); err != nil {
			os.Remove(privatePath)
			return errors.New("close private fixture output")
		}
	}
	if err := json.NewEncoder(stdout).Encode(struct {
		Application string `json:"application"`
		Subsystem   string `json:"subsystem"`
		APMService  string `json:"apmService"`
	}{result.Application, result.Subsystem, result.APMService}); err != nil {
		return errors.New("write fixture names")
	}
	return nil
}

func prepare(parent context.Context, deps dependencies, managementKey, sendDataKey, pagerDutyKey, pagerDutyEmail, archiveBucket string) (fixtureOutput, error) {
	var output fixtureOutput
	if managementKey == "" || sendDataKey == "" || pagerDutyKey == "" || pagerDutyEmail == "" {
		return output, errors.New("all four credential environment variables are required")
	}
	slackID, err := findIntegrationPrerequisites(parent, deps, managementKey)
	if err != nil {
		return output, err
	}
	if err := ensureArchiveTarget(parent, deps, managementKey, archiveBucket); err != nil {
		return output, err
	}
	pdID, err := ensurePagerDutyIntegration(parent, deps, managementKey, pagerDutyKey, pagerDutyEmail)
	if err != nil {
		return output, err
	}

	if err := seedAITrace(parent, deps, sendDataKey); err != nil {
		return output, err
	}
	if err := deps.seedAPM(parent, sendDataKey, deps.currentTime()); err != nil {
		return output, err
	}
	ctx, cancel := context.WithTimeout(parent, deps.timeout)
	defer cancel()
	poll := time.NewTicker(deps.pollEvery)
	defer poll.Stop()

	var aiReady, targetReady, apmReady bool
	var lastError error
	attempt := 0
	for {
		attempt++
		aiReady, targetReady, apmReady = false, false, false
		lastError = nil
		aiState, targetState, apmState := "missing", "not checked", "missing"
		aiReady, err = hasAIApplication(ctx, deps, managementKey)
		if err != nil {
			if !isRetryableCatalogError(err) {
				return output, err
			}
			lastError = err
			aiState = "retrying"
		} else if aiReady {
			aiState = "ready"
			targetReady, err = hasUnusedEvaluationTarget(ctx, deps, managementKey)
			if err != nil {
				if !isRetryableCatalogError(err) {
					return output, err
				}
				lastError = err
				targetState = "retrying"
			} else if targetReady {
				targetState = "ready"
			} else {
				targetState = "missing"
			}
		}
		apmReady, err = hasAPMService(ctx, deps, managementKey)
		if err != nil {
			if !isRetryableCatalogError(err) {
				return output, err
			}
			lastError = err
			apmState = "retrying"
		} else if apmReady {
			apmState = "ready"
		}
		if deps.progress != nil {
			fmt.Fprintf(deps.progress, "prepare CI team: catalog check %d: AI application=%s, unused evaluation target=%s, APM service (V2)=%s\n", attempt, aiState, targetState, apmState)
		}
		if aiReady && targetReady && apmReady {
			key, err := randomUUID()
			if err != nil {
				return output, errors.New("generate connector test value")
			}
			return fixtureOutput{
				Application: applicationName, Subsystem: subsystemName, APMService: serviceName,
				PagerDutyIntegrationID: pdID, PagerDutyConnectorKey: key, SlackIntegrationID: slackID,
			}, nil
		}

		select {
		case <-ctx.Done():
			if lastError != nil {
				return output, fmt.Errorf("timed out waiting for AI application %q (ready=%t, unused target=%t) and APM service %q (V2 ready=%t); last catalog error: %v", applicationName, aiReady, targetReady, serviceName, apmReady, lastError)
			}
			return output, fmt.Errorf("timed out waiting for AI application %q (ready=%t, unused target=%t) and APM service %q (V2 ready=%t); check trace archive and APM service discovery", applicationName, aiReady, targetReady, serviceName, apmReady)
		case <-poll.C:
		}
	}
}

func seedAITrace(ctx context.Context, deps dependencies, sendDataKey string) error {
	payload, err := aiTracePayload(deps.currentTime())
	if err != nil {
		return fmt.Errorf("build CI trace seed: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, deps.ingestURL, bytes.NewReader(payload))
	if err != nil {
		return errors.New("build trace request")
	}
	req.Header.Set("Authorization", "Bearer "+sendDataKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := deps.client.Do(req)
	if err != nil {
		return errors.New("send CI trace seed: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("send CI trace seed: HTTP %d", resp.StatusCode)
	}
	var result struct {
		PartialSuccess struct {
			RejectedSpans json.Number `json:"rejectedSpans"`
		} `json:"partialSuccess"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4096)).Decode(&result); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode trace seed response: %w", err)
	}
	if rejected, _ := result.PartialSuccess.RejectedSpans.Int64(); rejected > 0 {
		return fmt.Errorf("trace seed rejected %d spans", rejected)
	}
	return nil
}

func aiTracePayload(now time.Time) ([]byte, error) {
	traceID, err := randomID(16)
	if err != nil {
		return nil, err
	}
	aiID, err := randomID(8)
	if err != nil {
		return nil, err
	}
	start := strconv.FormatInt(now.Add(-time.Millisecond).UnixNano(), 10)
	end := strconv.FormatInt(now.UnixNano(), 10)
	attribute := func(key, value string) map[string]any {
		return map[string]any{"key": key, "value": map[string]string{"stringValue": value}}
	}
	spans := []map[string]any{
		{
			"traceId": traceID, "spanId": aiID,
			"name": "chat ci-fixture", "kind": 3,
			"startTimeUnixNano": start, "endTimeUnixNano": end,
			"attributes": []map[string]any{
				attribute("gen_ai.operation.name", "chat"),
				attribute("gen_ai.provider.name", "openai"),
				attribute("gen_ai.request.model", "ci-fixture"),
				attribute("gen_ai.input.messages", `[{"role":"user","parts":[{"type":"text","content":"CI fixture"}]}]`),
				attribute("gen_ai.output.messages", `[{"role":"assistant","parts":[{"type":"text","content":"Fixture ready"}]}]`),
			},
		},
	}
	return json.Marshal(map[string]any{"resourceSpans": []any{map[string]any{
		"resource": map[string]any{"attributes": []map[string]any{
			attribute("service.name", applicationName),
			attribute("cx.application.name", applicationName),
			attribute("cx.subsystem.name", subsystemName),
			attribute("deployment.environment.name", "ci"),
		}},
		"scopeSpans": []any{map[string]any{"scope": map[string]string{"name": "coralogix-ci-team-setup"}, "spans": spans}},
	}}})
}

func randomID(size int) (string, error) {
	value := make([]byte, size)
	if _, err := rand.Read(value); err != nil {
		return "", err
	}
	return hex.EncodeToString(value), nil
}

func randomUUID() (string, error) {
	value := make([]byte, 16)
	if _, err := rand.Read(value); err != nil {
		return "", err
	}
	value[6] = (value[6] & 0x0f) | 0x40
	value[8] = (value[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", value[:4], value[4:6], value[6:8], value[8:10], value[10:]), nil
}

func hasAIApplication(ctx context.Context, deps dependencies, managementKey string) (bool, error) {
	for page := 0; page < 10; page++ {
		url := deps.aiURL + "?page_size=1000&page_offset=" + strconv.Itoa(page)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return false, errors.New("build AI application request")
		}
		req.Header.Set("Authorization", "Bearer "+managementKey)
		resp, err := deps.client.Do(req)
		if err != nil {
			return false, errors.New("list AI applications: request failed")
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			resp.Body.Close()
			return false, statusError{operation: "list AI applications", status: resp.StatusCode}
		}
		var result struct {
			AIApplications []struct {
				Application string `json:"application"`
				Subsystem   string `json:"subsystem"`
			} `json:"aiApplications"`
		}
		err = json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&result)
		resp.Body.Close()
		if err != nil {
			return false, fmt.Errorf("decode AI applications: %w", err)
		}
		for _, app := range result.AIApplications {
			if app.Application == applicationName && app.Subsystem == subsystemName {
				return true, nil
			}
		}
		if len(result.AIApplications) < 1000 {
			return false, nil
		}
	}
	return false, errors.New("AI application list exceeded 10 pages")
}

func hasUnusedEvaluationTarget(ctx context.Context, deps dependencies, managementKey string) (bool, error) {
	used := map[string]bool{}
	for page := 0; page < 10; page++ {
		url := deps.evaluationsURL + "?application=" + applicationName + "&subsystem=" + subsystemName + "&page_size=200&page_offset=" + strconv.Itoa(page)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return false, errors.New("build AI evaluation request")
		}
		req.Header.Set("Authorization", "Bearer "+managementKey)
		resp, err := deps.client.Do(req)
		if err != nil {
			return false, errors.New("list AI evaluations: request failed")
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			resp.Body.Close()
			return false, statusError{operation: "list AI evaluations", status: resp.StatusCode}
		}
		var result struct {
			AIEvaluations []struct {
				Target string `json:"target"`
			} `json:"aiEvaluations"`
		}
		err = json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&result)
		resp.Body.Close()
		if err != nil {
			return false, errors.New("decode AI evaluations")
		}
		for _, evaluation := range result.AIEvaluations {
			if evaluation.Target == "" {
				return false, nil
			}
			used[strings.ToLower(evaluation.Target)] = true
		}
		if len(result.AIEvaluations) < 200 {
			return !used["response"] || !used["prompt"], nil
		}
	}
	return false, errors.New("AI evaluation list exceeded 10 pages")
}

type statusError struct {
	operation string
	status    int
}

func (e statusError) Error() string {
	return fmt.Sprintf("%s: HTTP %d", e.operation, e.status)
}

func isRetryableCatalogError(err error) bool {
	var status statusError
	if errors.As(err, &status) {
		return status.status == http.StatusTooManyRequests || status.status >= 500
	}
	return true
}

func hasAPMService(ctx context.Context, deps dependencies, managementKey string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, deps.apmURL, nil)
	if err != nil {
		return false, errors.New("build APM V2 service request")
	}
	req.Header.Set("Authorization", "Bearer "+managementKey)
	resp, err := deps.client.Do(req)
	if err != nil {
		return false, errors.New("list APM V2 services: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, statusError{operation: "list APM V2 services", status: resp.StatusCode}
	}
	var result struct {
		Services []struct {
			Name string `json:"name"`
		} `json:"entities"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 8<<20)).Decode(&result); err != nil {
		return false, errors.New("decode APM V2 services")
	}
	if result.Services == nil {
		return false, errors.New("APM V2 service list omitted entities")
	}
	for _, service := range result.Services {
		if service.Name == serviceName {
			return true, nil
		}
	}
	return false, nil
}
