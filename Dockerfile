# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Install build tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o go-run-reports ./cms/server

# Stage 2: Create a minimal image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/go-run-reports .

# Expose the application port (customize if needed)
EXPOSE 8080

# Run the binary
CMD ["./go-run-reports"]