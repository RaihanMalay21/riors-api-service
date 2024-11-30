# Stage 1: Build the application
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# Stage 2: Run the application
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
