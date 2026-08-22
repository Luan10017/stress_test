package runner

import (
    "context"
    "fmt"
    "net/http"
    "sync"
    "time"

    "github.com/Luan10017/stress_test/internal/report"
)

type Runner struct {
    url        string
    requests   int
    concurrency int
    client     *http.Client
}

func NewRunner(rawURL string, requests, concurrency int) (*Runner, error) {
    if rawURL == "" {
        return nil, fmt.Errorf("URL é obrigatória")
    }
    if requests <= 0 {
        return nil, fmt.Errorf("requests deve ser maior que zero")
    }
    if concurrency <= 0 {
        return nil, fmt.Errorf("concurrency deve ser maior que zero")
    }

    return &Runner{
        url:         rawURL,
        requests:    requests,
        concurrency: concurrency,
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }, nil
}

func (r *Runner) Run() (report.Summary, error) {
    jobs := make(chan struct{}, r.concurrency)
    results := make(chan report.Result, r.requests)

    var wg sync.WaitGroup
    for i := 0; i < r.concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for range jobs {
                result := r.executeRequest()
                results <- result
            }
        }()
    }

    for i := 0; i < r.requests; i++ {
        jobs <- struct{}{}
    }
    close(jobs)
    wg.Wait()
    close(results)

    requestResults := make([]report.Result, 0, r.requests)
    for result := range results {
        requestResults = append(requestResults, result)
    }

    return report.Summarize(requestResults, r.requests, 0), nil
}

func (r *Runner) executeRequest() report.Result {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.url, nil)
    if err != nil {
        return report.Result{Error: err}
    }

    resp, err := r.client.Do(req)
    if err != nil {
        return report.Result{Error: err}
    }
    defer resp.Body.Close()

    return report.Result{StatusCode: resp.StatusCode}
}
