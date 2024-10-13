# backend/Dockerfile
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .
ENV GOOS=linux GOARCH=amd64
RUN go mod tidy
RUN go build -o main .

# Production stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

