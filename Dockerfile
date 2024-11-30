# Stage 1: Build
FROM golang:1.22-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy application files
WORKDIR /app
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest

# Copy binary from the build stage
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the port used by your app
EXPOSE 8080

# Run the binary
CMD ["./main"]
