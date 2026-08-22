package report

import (
    "net/http"
    "testing"
    "time"
)

func TestSummarizeCountsStatusCodes(t *testing.T) {
    results := []Result{
        {StatusCode: http.StatusOK},
        {StatusCode: http.StatusNotFound},
        {StatusCode: http.StatusOK},
        {StatusCode: http.StatusInternalServerError},
    }

    summary := Summarize(results, 4, time.Second)

    if summary.Total != 4 {
        t.Fatalf("expected total 4, got %d", summary.Total)
    }

    if summary.Success != 2 {
        t.Fatalf("expected 2 successful requests, got %d", summary.Success)
    }

    if summary.StatusCounts[http.StatusOK] != 2 {
        t.Fatalf("expected 2 OK requests, got %d", summary.StatusCounts[http.StatusOK])
    }

    if summary.StatusCounts[http.StatusNotFound] != 1 {
        t.Fatalf("expected 1 not found request, got %d", summary.StatusCounts[http.StatusNotFound])
    }

    if summary.StatusCounts[http.StatusInternalServerError] != 1 {
        t.Fatalf("expected 1 internal server error, got %d", summary.StatusCounts[http.StatusInternalServerError])
    }
}
