# Step 1: Use an official Go runtime as a base image
FROM golang:1.22.6-alpine

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files and source code to the container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Step 4: Build the Go application
RUN go build -o main ./cmd/server

# Step 5: Set the command to run the application
CMD ["./main"]
