package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Check the integration definitions and connected Slack before any setup writes.
func findIntegrationPrerequisites(ctx context.Context, deps dependencies, managementKey string) (string, error) {
	var catalog struct {
		Integrations []struct {
			Integration struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"integration"`
		} `json:"integrations"`
	}
	if err := integrationGET(ctx, deps, deps.integrationsURL, managementKey, &catalog); err != nil {
		return "", err
	}

	keys := map[string]string{}
	for _, entry := range catalog.Integrations {
		id := strings.ToLower(entry.Integration.ID)
		name := strings.ToLower(entry.Integration.Name)
		if id == "pagerduty-incidents" {
			keys["pagerduty"] = entry.Integration.ID
		}
		if id == "slack-central" || id == "slack" || name == "slack" {
			keys["slack"] = entry.Integration.ID
		}
	}
	if keys["pagerduty"] == "" {
		return "", errors.New("PagerDuty integration definition is unavailable")
	}
	if keys["slack"] == "" {
		return "", errors.New("Slack integration definition is unavailable")
	}
	slackID, err := findActiveRegistration(ctx, deps, managementKey, keys["slack"])
	if err != nil {
		return "", err
	}
	if slackID == "" {
		return "", errors.New("connected Slack integration is missing; complete one-time OAuth for this team")
	}
	return slackID, nil
}

func findActiveRegistration(ctx context.Context, deps dependencies, managementKey, key string) (string, error) {
	var details struct {
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
	if err := integrationGET(ctx, deps, deps.integrationsURL+"/"+key, managementKey, &details); err != nil {
		return "", err
	}
	var active string
	for _, registration := range details.IntegrationDetail.Default.Registered {
		if registration.IntegrationStatus.ConnectionStatus != "ACTIVE" {
			continue
		}
		if active != "" {
			return "", errors.New("multiple active integrations of one type; select one manually")
		}
		active = registration.ID
	}
	return active, nil
}

func integrationGET(ctx context.Context, deps dependencies, url, managementKey string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.New("build integration request")
	}
	req.Header.Set("Authorization", "Bearer "+managementKey)
	resp, err := deps.client.Do(req)
	if err != nil {
		return errors.New("read integrations: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return statusError{operation: "read integrations", status: resp.StatusCode}
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(out); err != nil {
		return errors.New("decode integrations")
	}
	return nil
}
