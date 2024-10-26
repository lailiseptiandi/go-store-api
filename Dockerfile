# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

# Tidak perlu copy .env dan set ENV di sini
EXPOSE 8080
CMD ["./main"]