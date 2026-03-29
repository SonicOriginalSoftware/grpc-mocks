package slog

import (
	"context"
	"log/slog"
	"sync"
)

// CaptureHandler is a test helper that captures log records.
type CaptureHandler struct {
	records *[]slog.Record
	mu      sync.Mutex
}

// NewCaptureHandler creates a new CaptureHandler that appends records to the provided slice.
func NewCaptureHandler(records *[]slog.Record) *CaptureHandler {
	return &CaptureHandler{records: records}
}

func (h *CaptureHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *CaptureHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	*h.records = append(*h.records, r)
	return nil
}

func (h *CaptureHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *CaptureHandler) WithGroup(_ string) slog.Handler {
	return h
}
