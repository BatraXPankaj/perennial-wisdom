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

# Persistent storage directory
# - Azure App Service: /home is persistent, set DB_PATH=/home/data/wisdom.db
# - Fly.io / local: /data via mounted volume
RUN mkdir -p /data /home/data

ENV DB_PATH=/home/data/wisdom.db
ENV PORT=8080
EXPOSE 8080

CMD ["./perennial-wisdom"]
