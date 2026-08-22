package report

import (
    "net/http"
    "sort"
    "time"
)

type Result struct {
    StatusCode int
    Error      error
}

type Summary struct {
    Total        int
    Success      int
    Errors       int
    Duration     time.Duration
    StatusCounts map[int]int
}

func Summarize(results []Result, total int, duration time.Duration) Summary {
    statusCounts := make(map[int]int)
    success := 0
    errors := 0

    for _, result := range results {
        if result.Error != nil {
            errors++
            continue
        }

        if result.StatusCode == http.StatusOK {
            success++
        }

        statusCounts[result.StatusCode]++
    }

    codes := make([]int, 0, len(statusCounts))
    for code := range statusCounts {
        codes = append(codes, code)
    }
    sort.Ints(codes)

    orderedCounts := make(map[int]int, len(codes))
    for _, code := range codes {
        orderedCounts[code] = statusCounts[code]
    }

    return Summary{
        Total:        total,
        Success:      success,
        Errors:       errors,
        Duration:     duration,
        StatusCounts: orderedCounts,
    }
}
