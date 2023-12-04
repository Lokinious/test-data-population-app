# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code into the container
COPY . .

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
