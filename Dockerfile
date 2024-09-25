# Builder stage
FROM golang:1.23-alpine AS builder

ARG APP_NAME
ARG VERSION

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest of the source code
COPY . .

# Build the Go app with version argument
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}" -o ./bin/${APP_NAME} ./cmd/${APP_NAME}/main.go

# Final stage
#FROM debian:buster-slim
FROM alpine:latest

ARG APP_NAME
ENV APP_NAME=${APP_NAME}

# Install CA certificates
#RUN apt-get update && apt-get install -y apt-transport-https ca-certificates gnupg curl
RUN apk update && apk add --no-cache ca-certificates curl

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin .

# Copy the entrypoint script
COPY scripts/docker-entrypoint.sh /app/docker-entrypoint.sh
RUN chmod +x /app/docker-entrypoint.sh

# Command to run the executable
ENTRYPOINT ["./docker-entrypoint.sh"]
