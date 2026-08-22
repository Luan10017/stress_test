package main

import (
    "errors"
    "flag"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "time"

    "github.com/Luan10017/stress_test/internal/report"
    "github.com/Luan10017/stress_test/internal/runner"
)

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    urlFlag := flag.String("url", "", "URL do serviço a ser testado")
    requestsFlag := flag.Int("requests", 0, "Número total de requisições")
    concurrencyFlag := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
    flag.Parse()

    if err := validateInputs(*urlFlag, *requestsFlag, *concurrencyFlag); err != nil {
        return err
    }

    start := time.Now()
    r, err := runner.NewRunner(*urlFlag, *requestsFlag, *concurrencyFlag)
    if err != nil {
        return err
    }

    summary, err := r.Run()
    if err != nil {
        return err
    }

    duration := time.Since(start)
    summary.Duration = duration

    printReport(summary)

    return nil
}

func validateInputs(rawURL string, requests, concurrency int) error {
    if rawURL == "" {
        return errors.New("--url é obrigatório")
    }

    validURLPattern := regexp.MustCompile(`^https?://(localhost|127\.0\.0\.1|[A-Za-z0-9.-]+(?:\.[A-Za-z]{2,})+)(:[0-9]+)?(/.*)?$`)
    if !validURLPattern.MatchString(rawURL) {
        return errors.New("--url deve ser uma URL válida, com http:// ou https://")
    }

    if requests <= 0 {
        return errors.New("--requests deve ser maior que zero")
    }

    if concurrency <= 0 {
        return errors.New("--concurrency deve ser maior que zero")
    }

    if _, err := strconv.Atoi(strconv.Itoa(requests)); err != nil {
        return fmt.Errorf("valor inválido para --requests: %w", err)
    }

    return nil
}

func printReport(summary report.Summary) {
    fmt.Println("\n=== Relatório de Execução ===")
    fmt.Printf("Tempo total gasto: %s\n", summary.Duration)
    fmt.Printf("Quantidade total de requests: %d\n", summary.Total)
    fmt.Printf("Quantidade de requests com status 200: %d\n", summary.Success)
    fmt.Println("Distribuição de outros códigos de status HTTP:")

    if len(summary.StatusCounts) == 0 {
        fmt.Println("Nenhum status adicional registrado.")
        return
    }

    for code, count := range summary.StatusCounts {
        fmt.Printf("  HTTP %d: %d\n", code, count)
    }

    if summary.Errors > 0 {
        fmt.Printf("Erros de rede/timeout: %d\n", summary.Errors)
    }
}
