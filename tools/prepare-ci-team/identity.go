package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// lookupTeamID resolves the team authenticated by the management key.
func lookupTeamID(ctx context.Context, deps dependencies, managementKey string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, deps.identityURL, nil)
	if err != nil {
		return "", errors.New("build team identity request")
	}
	req.Header.Set("Authorization", "Bearer "+managementKey)
	resp, err := deps.client.Do(req)
	if err != nil {
		return "", errors.New("read team identity: request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", statusError{operation: "read team identity", status: resp.StatusCode}
	}
	var identity struct {
		TeamID *uint32 `json:"teamId"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4096)).Decode(&identity); err != nil || identity.TeamID == nil {
		return "", errors.New("decode team identity")
	}
	return strconv.FormatUint(uint64(*identity.TeamID), 10), nil
}
