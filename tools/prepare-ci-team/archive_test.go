package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

const archivePath = "/mgmt/openapi/5/logs/data-setup/v2"

func TestArchiveTargetCreatedOnce(t *testing.T) {
	var created bool
	var writes int
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != archivePath || r.Header.Get("Authorization") != "Bearer management-secret" {
			t.Fatal("wrong archive request")
		}
		switch r.Method {
		case http.MethodGet:
			if !created {
				return jsonResponse(404, `{}`), nil
			}
			return jsonResponse(200, `{"target":{"archiveSpec":{"isActive":true},"s3":{"bucket":"fixture-bucket"}}}`), nil
		case http.MethodPost:
			writes++
			var body struct {
				IsActive bool `json:"isActive"`
				S3       struct {
					Bucket string `json:"bucket"`
					Region string `json:"region"`
				} `json:"s3"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil || !body.IsActive || body.S3.Bucket != "fixture-bucket" || body.S3.Region != "eu-north-1" {
				t.Fatal("wrong archive request body")
			}
			created = true
			return jsonResponse(200, `{}`), nil
		default:
			t.Fatal("wrong archive method")
			return nil, nil
		}
	})
	for range 2 {
		if err := ensureArchiveTarget(context.Background(), deps, "management-secret", "fixture-bucket"); err != nil {
			t.Fatal(err)
		}
	}
	if writes != 1 {
		t.Fatalf("archive written %d times", writes)
	}
}

func TestArchiveTargetReusesActiveS3(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet || r.URL.Path != archivePath {
			t.Fatal("changed active archive")
		}
		return jsonResponse(200, `{"target":{"archiveSpec":{"isActive":true},"s3":{"bucket":"existing-bucket"}}}`), nil
	})
	if err := ensureArchiveTarget(context.Background(), deps, "management-secret", "fixture-bucket"); err != nil {
		t.Fatal(err)
	}
}

func TestArchiveTargetRequiresBucketWhenMissing(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet || r.URL.Path != archivePath {
			t.Fatal("wrote without archive bucket")
		}
		return jsonResponse(404, `{}`), nil
	})
	err := ensureArchiveTarget(context.Background(), deps, "management-secret", "")
	if err == nil || !strings.Contains(err.Error(), "ARCHIVE_LOGS_BUCKET") {
		t.Fatalf("wrong missing bucket error: %v", err)
	}
}

func TestArchiveTargetRefusesReplacement(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet || r.URL.Path != archivePath {
			t.Fatal("replaced existing archive")
		}
		return jsonResponse(200, `{"target":{"archiveSpec":{"isActive":false},"s3":{"bucket":"existing-bucket"}}}`), nil
	})
	err := ensureArchiveTarget(context.Background(), deps, "management-secret", "fixture-bucket")
	if err == nil || !strings.Contains(err.Error(), "refusing to replace") {
		t.Fatalf("wrong archive conflict error: %v", err)
	}
}

func TestArchiveTargetAuthorizationFailsWithoutWrite(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet {
			t.Fatal("wrote after archive authorization failure")
		}
		return jsonResponse(403, `{}`), nil
	})
	err := ensureArchiveTarget(context.Background(), deps, "management-secret", "fixture-bucket")
	if err == nil || !strings.Contains(err.Error(), "HTTP 403") || strings.Contains(err.Error(), "management-secret") {
		t.Fatalf("wrong archive authorization error: %v", err)
	}
}

func TestArchiveTargetChecksAfterUncertainWrite(t *testing.T) {
	var written bool
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.Method == http.MethodGet {
			if written {
				return jsonResponse(200, `{"target":{"archiveSpec":{"isActive":true},"s3":{"bucket":"fixture-bucket"}}}`), nil
			}
			return jsonResponse(404, `{}`), nil
		}
		written = true
		return jsonResponse(503, `{}`), nil
	})
	if err := ensureArchiveTarget(context.Background(), deps, "management-secret", "fixture-bucket"); err != nil {
		t.Fatalf("archive state was active after uncertain write: %v", err)
	}
}

func TestMissingArchiveStopsBeforePagerDutyAndTelemetry(t *testing.T) {
	deps := fakeDependencies(t, func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == archivePath {
			return jsonResponse(404, `{}`), nil
		}
		if r.Method == http.MethodPost || r.URL.Path == "/v1/traces" {
			t.Fatal("setup wrote before archive check")
		}
		return fixtureResponse(t, r)
	})
	_, err := prepare(context.Background(), deps, "management-secret", "send-secret", "pd-secret", "actor@example.invalid", "")
	if err == nil || !strings.Contains(err.Error(), "ARCHIVE_LOGS_BUCKET") {
		t.Fatalf("wrong missing archive error: %v", err)
	}
}
