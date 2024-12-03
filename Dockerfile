
# Start from the official Golang base image
FROM golang:1.23.1-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

# Copy any config files if needed
# COPY --from=builder /app/config ./config

# Expose the port the API runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
