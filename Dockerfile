# Build stage — compile the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o perennial-wisdom .

# Run stage — minimal image, single binary + templates
FROM alpine:3.21

RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/perennial-wisdom .
COPY --from=builder /app/templates ./templates

# Persistent volume for SQLite data
RUN mkdir -p /data

ENV DB_PATH=/data/wisdom.db
ENV PORT=8080
EXPOSE 8080

CMD ["./perennial-wisdom"]
