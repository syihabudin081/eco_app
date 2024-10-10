# Use the official Go image as the base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application and output it to /app/main
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Install necessary packages (if needed) for running the app
RUN apt-get update && apt-get install -y ca-certificates

# Command to run the executable
CMD ["./main"]

# Expose the port on which your app runs
EXPOSE 3000
