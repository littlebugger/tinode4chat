# Start from the official Go image for building the application
FROM golang:1.22.7 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server/main.go

# Start a new stage from a minimal image
FROM golang:1.22.7

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Make main executable
RUN chmod +x main

# Expose the application port (if needed)
EXPOSE 8080

# Run the executable
CMD ["./main"]