# Stage 1: Build the Go application
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/app/main cmd/app/main.go

# Stage 2: Create a minimal image with Alpine and the compiled Go binary
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/app/main ./bin/app/main

EXPOSE 8080
CMD ["./bin/app/main"]
