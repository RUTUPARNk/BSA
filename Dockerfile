# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY bsa-core-go/go.mod ./
# COPY bsa-core-go/go.sum ./ # No dependencies yet

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY bsa-core-go/ .

# Build the Go app
RUN go build -o bsa-core main.go

# Run Stage
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bsa-core .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./bsa-core"]
