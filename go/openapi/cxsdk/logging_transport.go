package cxsdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
)

// LoggingTransport logs HTTP requests and responses.
type LoggingTransport struct {
	logger *slog.Logger
}

// NewLoggingTransport creates a new LoggingTransport.
func NewLoggingTransport() *LoggingTransport {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return &LoggingTransport{logger: logger}
}

// RoundTrip implements the http.RoundTripper interface.
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var reqBody []byte
	if req.Body != nil {
		reqBody, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewReader(reqBody))
	}

	t.logger.Debug("SDK HTTP request",
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		slog.Any("body", parseJSONOrRaw(reqBody)),
	)

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		t.logger.Error("SDK HTTP error",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	var respBody []byte
	if resp.Body != nil {
		respBody, _ = io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewReader(respBody))
	}

	t.logger.Debug("SDK HTTP response",
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		slog.Any("body", parseJSONOrRaw(respBody)),
	)

	return resp, nil
}

func parseJSONOrRaw(raw []byte) any {
	if len(raw) == 0 {
		return ""
	}
	var parsed any
	if json.Unmarshal(raw, &parsed) == nil {
		return parsed
	}
	return string(raw)
}
