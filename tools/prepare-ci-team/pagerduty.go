package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	pagerDutyDefinition = "pagerduty-incidents"
	pagerDutyVersion    = "0.0.1"
	pagerDutyName       = "ci-acceptance-pagerduty"
)

func ensurePagerDutyIntegration(ctx context.Context, deps dependencies, managementKey, pagerDutyKey, pagerDutyEmail string) (string, error) {
	id, status, err := readPagerDutyRegistration(ctx, deps, managementKey)
	if err != nil {
		return "", err
	}
	if id != "" {
		return waitPagerDutyRegistration(ctx, deps, managementKey, id, status)
	}
	if err := verifyPagerDutyDefinition(ctx, deps, managementKey); err != nil {
		return "", err
	}

	metadata := pagerDutyMetadata(pagerDutyKey, pagerDutyEmail)
	var testResult struct {
		Result struct {
			Success *struct{} `json:"success"`
			Failure *struct{} `json:"failure"`
		} `json:"result"`
	}
	if err := postIntegration(ctx, deps, managementKey, "/integrations/metadata/v1/test", map[string]any{"integrationData": metadata}, &testResult); err != nil {
		return "", err
	}
	if testResult.Result.Success == nil || testResult.Result.Failure != nil {
		return "", errors.New("PagerDuty integration validation failed; no integration was created")
	}

	// Another setup may have created the account-level integration during validation.
	id, status, err = readPagerDutyRegistration(ctx, deps, managementKey)
	if err != nil {
		return "", err
	}
	if id != "" {
		return waitPagerDutyRegistration(ctx, deps, managementKey, id, status)
	}

	var created struct {
		IntegrationID string `json:"integrationId"`
	}
	createErr := postIntegration(ctx, deps, managementKey, "/integrations/metadata/v1", map[string]any{"metadata": metadata}, &created)
	// A failed response can follow a successful create. Read before deciding whether to retry.
	id, status, err = readPagerDutyRegistration(ctx, deps, managementKey)
	if err == nil && id != "" {
		return waitPagerDutyRegistration(ctx, deps, managementKey, id, status)
	}
	if createErr != nil {
		return "", createErr
	}
	if err != nil {
		return "", err
	}
	if created.IntegrationID == "" {
		return "", errors.New("PagerDuty create returned no integration ID")
	}
	return waitPagerDutyRegistration(ctx, deps, managementKey, created.IntegrationID, "PENDING")
}

func pagerDutyMetadata(pagerDutyKey, pagerDutyEmail string) map[string]any {
	return map[string]any{
		"integrationKey": pagerDutyDefinition,
		"version":        pagerDutyVersion,
		"integrationParameters": map[string]any{"parameters": []any{
			map[string]any{"key": "IntegrationName", "stringValue": pagerDutyName},
			map[string]any{"key": "ApiKey", "stringValue": pagerDutyKey},
			map[string]any{"key": "Email", "stringValue": pagerDutyEmail},
			map[string]any{"key": "EnableBidirectional", "booleanValue": true},
			map[string]any{"key": "DataProcessingConsent", "booleanValue": true},
		}},
	}
}

func verifyPagerDutyDefinition(ctx context.Context, deps dependencies, managementKey string) error {
	var response struct {
		IntegrationDefinition struct {
			Key       string `json:"key"`
			Revisions []struct {
				ID     string `json:"id"`
				Fields []struct {
					Name     string `json:"templateParamName"`
					Type     string `json:"type"`
					Required bool   `json:"required"`
				} `json:"fields"`
			} `json:"revisions"`
		} `json:"integrationDefinition"`
	}
	if err := integrationGET(ctx, deps, deps.integrationsURL+"/definitions/"+pagerDutyDefinition, managementKey, &response); err != nil {
		return err
	}
	if response.IntegrationDefinition.Key != "" && response.IntegrationDefinition.Key != pagerDutyDefinition {
		return errors.New("PagerDuty definition key changed")
	}
	want := map[string]string{
		"IntegrationName": "TEXT", "ApiKey": "SENSITIVE_DATA", "Email": "TEXT",
		"EnableBidirectional": "BOOLEAN", "DataProcessingConsent": "BOOLEAN",
	}
	for _, revision := range response.IntegrationDefinition.Revisions {
		if revision.ID != pagerDutyVersion {
			continue
		}
		if len(revision.Fields) != len(want) {
			return errors.New("PagerDuty definition fields changed")
		}
		for _, field := range revision.Fields {
			if want[field.Name] != field.Type || !field.Required {
				return errors.New("PagerDuty definition fields changed")
			}
		}
		return nil
	}
	return errors.New("PagerDuty definition version changed")
}

func readPagerDutyRegistration(ctx context.Context, deps dependencies, managementKey string) (string, string, error) {
	var response struct {
		IntegrationDetail struct {
			Default struct {
				Registered []struct {
					ID                string `json:"id"`
					IntegrationStatus struct {
						ConnectionStatus string `json:"connectionStatus"`
					} `json:"integrationStatus"`
				} `json:"registered"`
			} `json:"default"`
		} `json:"integrationDetail"`
	}
	if err := integrationGET(ctx, deps, deps.integrationsURL+"/"+pagerDutyDefinition, managementKey, &response); err != nil {
		return "", "", err
	}
	registered := response.IntegrationDetail.Default.Registered
	if len(registered) > 1 {
		return "", "", errors.New("multiple account-level PagerDuty integrations found")
	}
	if len(registered) == 0 {
		return "", "", nil
	}
	return registered[0].ID, registered[0].IntegrationStatus.ConnectionStatus, nil
}

func waitPagerDutyRegistration(parent context.Context, deps dependencies, managementKey, expectedID, status string) (string, error) {
	if status == "ACTIVE" {
		return expectedID, nil
	}
	if status == "FAILING" {
		return "", errors.New("account-level PagerDuty integration is failing")
	}
	ctx, cancel := context.WithTimeout(parent, 2*time.Minute)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return "", errors.New("timed out waiting for account-level PagerDuty integration to become active")
		case <-time.After(deps.pollEvery):
		}
		id, current, err := readPagerDutyRegistration(ctx, deps, managementKey)
		if err != nil {
			return "", err
		}
		if id == "" {
			continue
		}
		if id != expectedID {
			return "", errors.New("account-level PagerDuty integration changed during setup")
		}
		if current == "ACTIVE" {
			return id, nil
		}
		if current == "FAILING" {
			return "", errors.New("account-level PagerDuty integration is failing")
		}
	}
}

func postIntegration(ctx context.Context, deps dependencies, managementKey, path string, body, output any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return errors.New("build integration request")
	}
	url := strings.TrimSuffix(deps.integrationsURL, "/integrations/integrations/v1") + path
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return errors.New("build integration request")
	}
	req.Header.Set("Authorization", "Bearer "+managementKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := deps.client.Do(req)
	if err != nil {
		return errors.New("write integration: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		operation := "create PagerDuty integration"
		if strings.HasSuffix(path, "/test") {
			operation = "validate PagerDuty integration"
		}
		field := integrationErrorField(resp.Body)
		if field != "" {
			return fmt.Errorf("%s: HTTP %d (backend field: %s)", operation, resp.StatusCode, field)
		}
		return statusError{operation: operation, status: resp.StatusCode}
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(output); err != nil {
		return errors.New("decode integration response")
	}
	return nil
}

// Only fixed field names can leave the backend error response.
func integrationErrorField(body io.Reader) string {
	var response struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(io.LimitReader(body, 8192)).Decode(&response); err != nil {
		return ""
	}
	for _, field := range []string{
		"apiKeyId", "DataProcessingConsent", "EnableBidirectional",
		"IntegrationName", "ApiKey", "Email", "integrationKey", "version",
	} {
		if strings.Contains(response.Message, field) {
			return field
		}
	}
	return ""
}
