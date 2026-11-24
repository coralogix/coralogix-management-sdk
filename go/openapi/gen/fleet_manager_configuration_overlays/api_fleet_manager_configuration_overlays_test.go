package fleet_manager_configuration_overlays

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListConfigurationOverlaysBasicView(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("view"); got != string(CONFIGURATIONOVERLAYVIEW_CONFIGURATION_OVERLAY_VIEW_BASIC) {
			t.Errorf("view query = %q, want %q", got, CONFIGURATIONOVERLAYVIEW_CONFIGURATION_OVERLAY_VIEW_BASIC)
		}
		if got := r.URL.Query().Get("include_archived"); got != "true" {
			t.Errorf("include_archived query = %q, want true", got)
		}
		if got := r.URL.Query().Get("page_token"); got != "page-2" {
			t.Errorf("page_token query = %q, want page-2", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"overlays":[{"id":"01975ed4-9c1a-7b3f-8d4e-426614174000","name":"basic","tags":[],"archived":false,"priorityOrder":1,"active":true}]}`)
	}))
	defer server.Close()

	configuration := NewConfiguration()
	configuration.Servers = ServerConfigurations{{URL: server.URL}}
	configuration.HTTPClient = server.Client()
	api := NewAPIClient(configuration).FleetManagerConfigurationOverlaysAPI
	response, httpResponse, err := api.ConfigurationOverlayServiceListConfigurationOverlays(t.Context()).
		IncludeArchived(true).
		PageToken("page-2").
		View(CONFIGURATIONOVERLAYVIEW_CONFIGURATION_OVERLAY_VIEW_BASIC).
		Execute()
	if err != nil {
		t.Fatalf("list request failed: %v", err)
	}
	if httpResponse == nil || httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("HTTP response = %#v, want status %d", httpResponse, http.StatusOK)
	}
	if response == nil || len(response.GetOverlays()) != 1 {
		t.Fatalf("response = %#v, want one overlay", response)
	}
	overlay := response.GetOverlays()[0]
	if overlay.GetName() != "basic" || !overlay.GetActive() {
		t.Fatalf("overlay metadata = %#v, want name basic and active", overlay)
	}
	if overlay.Custom != nil || overlay.Preset != nil || overlay.Version != nil || overlay.VersionId != nil || overlay.OverlayHash != nil || len(overlay.Schedules) != 0 {
		t.Fatalf("BASIC overlay contains fields omitted by the API: %#v", overlay)
	}
}
