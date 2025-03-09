# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git for fetching dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o loadgenerator

# Final stage
FROM alpine:3.16

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/loadgenerator .

# Set the entrypoint
ENTRYPOINT ["/app/loadgenerator"]
