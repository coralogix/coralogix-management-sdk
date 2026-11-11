package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// The file contains an environment reference, never a credential value.
func collectorConfig(port int) string {
	return fmt.Sprintf(`receivers:
  otlp:
    protocols:
      http:
        endpoint: 127.0.0.1:%d
connectors:
  spanmetrics:
    add_resource_attributes: true
    metrics_flush_interval: 5s
    histogram:
      unit: ms
  spanmetrics/compact:
    add_resource_attributes: true
    namespace: compact
    exclude_dimensions: [span.name]
    metrics_flush_interval: 5s
    histogram:
      unit: ms
exporters:
  otlphttp/coralogix:
    endpoint: https://ingress.eu2.coralogix.com
    headers:
      Authorization: "Bearer ${env:CORALOGIX_SEND_YOUR_DATA_API_KEY}"
    retry_on_failure:
      enabled: true
service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [otlphttp/coralogix, spanmetrics, spanmetrics/compact]
    metrics:
      receivers: [spanmetrics, spanmetrics/compact]
      exporters: [otlphttp/coralogix]
  telemetry:
    logs:
      level: error
`, port)
}

func serverTracePayload(now time.Time) ([]byte, error) {
	traceID, err := randomID(16)
	if err != nil {
		return nil, err
	}
	spanID, err := randomID(8)
	if err != nil {
		return nil, err
	}
	attr := func(key, value string) map[string]any {
		return map[string]any{"key": key, "value": map[string]string{"stringValue": value}}
	}
	return json.Marshal(map[string]any{"resourceSpans": []any{map[string]any{
		"resource": map[string]any{"attributes": []any{
			attr("service.name", serviceName), attr("cx.application.name", "ci-acceptance"),
			attr("cx.subsystem.name", subsystemName), attr("deployment.environment.name", "ci"),
		}},
		"scopeSpans": []any{map[string]any{"scope": map[string]string{"name": "coralogix-ci-team-setup"},
			"spans": []any{map[string]any{
				"traceId": traceID, "spanId": spanID, "name": "ci-acceptance-request", "kind": 2,
				"startTimeUnixNano": strconv.FormatInt(now.Add(-time.Millisecond).UnixNano(), 10),
				"endTimeUnixNano":   strconv.FormatInt(now.UnixNano(), 10),
				"status":            map[string]any{"code": 1},
			}}}},
	}}})
}

func seedAPMService(ctx context.Context, sendDataKey string, now time.Time) error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return errors.New("reserve Collector port")
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()

	config, err := os.CreateTemp("", "prepare-ci-collector-*.yaml")
	if err != nil {
		return errors.New("create Collector configuration")
	}
	defer os.Remove(config.Name())
	if err := config.Chmod(0600); err != nil {
		config.Close()
		return errors.New("protect Collector configuration")
	}
	if _, err := config.WriteString(collectorConfig(port)); err != nil {
		config.Close()
		return errors.New("write Collector configuration")
	}
	if err := config.Close(); err != nil {
		return errors.New("close Collector configuration")
	}

	cmd := exec.CommandContext(ctx, "otelcol-contrib", "--config", config.Name())
	cmd.Env = []string{"CORALOGIX_SEND_YOUR_DATA_API_KEY=" + sendDataKey}
	cmd.Stdout, cmd.Stderr = nil, nil
	if err := cmd.Start(); err != nil {
		return errors.New("start otelcol-contrib")
	}
	defer func() { _ = cmd.Process.Kill(); _ = cmd.Wait() }()

	payload, err := serverTracePayload(now)
	if err != nil {
		return errors.New("build APM trace seed")
	}
	client := &http.Client{Timeout: 2 * time.Second}
	url := fmt.Sprintf("http://127.0.0.1:%d/v1/traces", port)
	deadline := time.NewTimer(15 * time.Second)
	defer deadline.Stop()
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
		if err != nil {
			return errors.New("build local APM trace request")
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return fmt.Errorf("Collector rejected APM trace: HTTP %d", resp.StatusCode)
			}
			break
		}
		select {
		case <-ctx.Done():
			return errors.New("Collector APM seed canceled")
		case <-deadline.C:
			return errors.New("local Collector did not accept the APM trace within 15 seconds; catalog polling has not started")
		case <-time.After(100 * time.Millisecond):
		}
	}
	// The connectors flush on a five-second interval. Give both metric exports time to finish.
	select {
	case <-ctx.Done():
		return errors.New("Collector metrics flush canceled")
	case <-time.After(12 * time.Second):
		return nil
	}
}
