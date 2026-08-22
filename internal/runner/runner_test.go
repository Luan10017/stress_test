package runner

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRunnerRunReturnsExactRequestCount(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    defer server.Close()

    runner, err := NewRunner(server.URL, 25, 5)
    if err != nil {
        t.Fatalf("NewRunner() returned error: %v", err)
    }

    summary, err := runner.Run()
    if err != nil {
        t.Fatalf("Run() returned error: %v", err)
    }

    if summary.Total != 25 {
        t.Fatalf("expected total of 25 requests, got %d", summary.Total)
    }

    if summary.Success != 25 {
        t.Fatalf("expected 25 successful requests, got %d", summary.Success)
    }

    if summary.Errors != 0 {
        t.Fatalf("expected zero network errors, got %d", summary.Errors)
    }
}
