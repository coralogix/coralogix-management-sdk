package cxsdk

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRestClientErrorDoesNotExposeResponseBody(t *testing.T) {
	const secret = "synthetic-secret-value"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, `{"token":"`+secret+`"}`)
	}))
	defer server.Close()

	_, err := NewRestClient(server.URL, "synthetic-api-key").Get(context.Background(), "/")
	if err == nil {
		t.Fatal("expected an HTTP error")
	}
	if strings.Contains(err.Error(), secret) {
		t.Fatal("REST error exposed the response body")
	}
	if !strings.Contains(err.Error(), "400") {
		t.Fatalf("REST error omitted the status: %v", err)
	}
}
