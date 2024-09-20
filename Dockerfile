# Use official Go image as a build stage
FROM golang:1.23.1 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go App
RUN go build -o main .

RUN ls -l

# Use minimal base image
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port of the App
EXPOSE 8080

# Command to run the App
CMD [ "./main" ]