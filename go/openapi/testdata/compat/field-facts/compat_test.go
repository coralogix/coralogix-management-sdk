package field_facts

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// The new OpenAPI field facts (defaults, presence, unordered lists, keys next
// to a $ref, readOnly) must not change the Go types.
var (
	_ *bool      = Widget{}.Enabled
	_ *string    = Widget{}.Limit
	_ *Kind      = Widget{}.Kind
	_ []string   = Widget{}.Labels
	_ *time.Time = Widget{}.CreateTime
)

// A default is what the server fills in. The constructors must not set it,
// so the client does not send it.
func TestFieldFacts_ConstructorsDoNotSetDefaults(t *testing.T) {
	for name, widget := range map[string]*Widget{
		"NewWidget":             NewWidget("w"),
		"NewWidgetWithDefaults": NewWidgetWithDefaults(),
	} {
		if widget.Enabled != nil || widget.Limit != nil || widget.Kind != nil || widget.Labels != nil {
			t.Fatalf("%s set a default: %+v", name, widget)
		}
	}

	body, err := json.Marshal(NewWidget("w"))
	if err != nil {
		t.Fatalf("json.Marshal returned error: %v", err)
	}
	if string(body) != `{"name":"w"}` {
		t.Fatalf("body = %s, want {\"name\":\"w\"}", body)
	}
}

// A query-parameter default is what the server uses when the parameter is not
// sent. The client must not send it.
func TestFieldFacts_QueryDefaultsAreNotSent(t *testing.T) {
	var rawQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"w"}`))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: server.URL}}
	client := NewAPIClient(cfg)

	if _, _, err := client.DefaultAPI.ListWidgets(context.Background()).Execute(); err != nil {
		t.Fatalf("ListWidgets returned error: %v", err)
	}
	if rawQuery != "" {
		t.Fatalf("query = %q, want no query parameters", rawQuery)
	}
}
