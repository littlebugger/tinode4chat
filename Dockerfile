# Step 1: Use an official Go runtime as the base image
FROM golang:1.22.7-alpine

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the source code into the container
COPY . .

# Step 5: Build the Go application
RUN go build -o main ./cmd/server/main.go

RUN ls -l /app/main

# Step 5.1: Make binary executable
RUN chmod +x main

# Step 6: Expose the application port (e.g., 8080)
EXPOSE 8080

# Step 7: Run the application
CMD ["./main"]