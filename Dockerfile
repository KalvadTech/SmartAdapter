# Use official Golang image as build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o smartadapter main.go

# Use minimal image for running
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/smartadapter .
# Dynamically expose the port specified by the PORT environment variable (default 8080)
ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}
ENTRYPOINT ["./smartadapter"]
