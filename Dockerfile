FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/flibot-urt

FROM alpine:latest

RUN apk upgrade --no-cache && apk add --no-cache libc6-compat
COPY --from=builder /app/flibot-urt /app/flibot-urt
COPY sqlc /app/sqlc

WORKDIR /app

ENTRYPOINT ["./flibot-urt"]
