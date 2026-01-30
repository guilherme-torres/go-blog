FROM golang:1.25-alpine AS builder

# Instala o GCC e outras ferramentas de build necessárias
RUN apk add --no-cache build-base

ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app ./cmd/server

# Final lightweight stage
FROM alpine:3.21 AS final

COPY --from=builder /app .
COPY --from=builder /go/bin/goose .

COPY --from=builder /build/migrations /migrations

EXPOSE 8000

CMD ["./app"]