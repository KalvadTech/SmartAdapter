# Use official Golang image as build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o smartadapter main.go

# Use minimal image for running
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/smartadapter .
EXPOSE 8080
ENTRYPOINT ["./smartadapter"]
