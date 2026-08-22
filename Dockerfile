FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o stress-test .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/stress-test /app/stress-test

ENTRYPOINT ["/app/stress-test"]
