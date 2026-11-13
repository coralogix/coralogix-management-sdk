package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// The logs archive target also stores traces used by AI Center.
type archiveTarget struct {
	ArchiveSpec *struct {
		IsActive bool `json:"isActive"`
	} `json:"archiveSpec"`
	S3 *struct {
		Bucket string `json:"bucket"`
	} `json:"s3"`
}

func ensureArchiveTarget(ctx context.Context, deps dependencies, managementKey, bucket string) error {
	current, err := readArchiveTarget(ctx, deps, managementKey)
	if err != nil {
		return err
	}
	if current != nil && current.ArchiveSpec != nil && current.ArchiveSpec.IsActive {
		if current.S3 == nil || current.S3.Bucket == "" {
			return errors.New("active logs and traces archive is not an S3 target")
		}
		return nil
	}

	bucket = strings.TrimSpace(bucket)
	if bucket == "" {
		return errors.New("logs and traces archive is missing or inactive; set ARCHIVE_LOGS_BUCKET to an existing EU2 archive bucket")
	}
	if current != nil && (current.S3 == nil || current.S3.Bucket != bucket) {
		return errors.New("inactive logs and traces archive uses another target; refusing to replace it")
	}

	writeErr := setArchiveTarget(ctx, deps, managementKey, bucket)
	// A failed response can follow a successful write. Read before reporting failure.
	updated, readErr := readArchiveTarget(ctx, deps, managementKey)
	if readErr == nil && updated != nil && updated.ArchiveSpec != nil && updated.ArchiveSpec.IsActive && updated.S3 != nil && updated.S3.Bucket == bucket {
		return nil
	}
	if writeErr != nil {
		return writeErr
	}
	if readErr != nil {
		return readErr
	}
	return errors.New("logs and traces archive did not become active")
}

func readArchiveTarget(ctx context.Context, deps dependencies, managementKey string) (*archiveTarget, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, deps.archiveURL, nil)
	if err != nil {
		return nil, errors.New("build archive target request")
	}
	req.Header.Set("Authorization", "Bearer "+managementKey)
	resp, err := deps.client.Do(req)
	if err != nil {
		return nil, errors.New("read archive target: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, statusError{operation: "read archive target", status: resp.StatusCode}
	}
	var result struct {
		Target *archiveTarget `json:"target"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&result); err != nil {
		return nil, errors.New("decode archive target")
	}
	return result.Target, nil
}

func setArchiveTarget(ctx context.Context, deps dependencies, managementKey, bucket string) error {
	payload, err := json.Marshal(map[string]any{
		"isActive": true,
		"s3":       map[string]string{"bucket": bucket, "region": "eu-north-1"},
	})
	if err != nil {
		return errors.New("build archive target request")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, deps.archiveURL, bytes.NewReader(payload))
	if err != nil {
		return errors.New("build archive target request")
	}
	req.Header.Set("Authorization", "Bearer "+managementKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := deps.client.Do(req)
	if err != nil {
		return errors.New("set archive target: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return statusError{operation: "set archive target", status: resp.StatusCode}
	}
	return nil
}
