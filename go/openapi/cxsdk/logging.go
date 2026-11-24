// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cxsdk

import (
	"log/slog"
	"net/http"
)

// LoggingTransport logs HTTP requests and responses.
type LoggingTransport struct {
	logger *slog.Logger
}

// NewLoggingTransport creates a new LoggingTransport.
func NewLoggingTransport(logger *slog.Logger) *LoggingTransport {
	return &LoggingTransport{logger: logger}
}

// RoundTrip implements the http.RoundTripper interface.
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.logger.Debug("SDK HTTP request", slog.String("method", req.Method))

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		t.logger.Error("SDK HTTP error", slog.String("method", req.Method))
		return nil, err
	}

	t.logger.Debug("SDK HTTP response",
		slog.String("method", req.Method),
		slog.Int("status_code", resp.StatusCode),
	)
	return resp, nil
}
