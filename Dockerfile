FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/bin

FROM alpine:latest

RUN apk upgrade --no-cache && apk add --no-cache libc6-compat
COPY --from=builder /app/bin /app/bin
COPY sqlc /app/sqlc

WORKDIR /app

ENTRYPOINT ["./bin"]