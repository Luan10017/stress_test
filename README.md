# Stress Test CLI

Aplicação em Go para realizar testes de carga em serviços web usando concorrência configurável.

## Requisitos

- Go 1.22+
- Docker (opcional, para executar em container)

## Como executar

```bash
go run . --url=http://google.com --requests=100 --concurrency=10
```

## Como compilar

```bash
go build -o stress-test .
```

## Como executar o binário compilado

```bash
./stress-test --url=http://google.com --requests=100 --concurrency=10
```

## Executando com Docker

### Build da imagem

```bash
docker build -t stress-test-cli .
```

### Execução

```bash
docker run --rm stress-test-cli --url=http://google.com --requests=1000 --concurrency=10
```

## Observações

- O número total de requisições será distribuído entre as goroutines de acordo com a concorrência informada.
- O relatório final apresenta duração total, quantidade total de requisições, status 200 e distribuição dos demais códigos HTTP.
- Todas as chamadas usam timeout para evitar bloqueios prolongados.
