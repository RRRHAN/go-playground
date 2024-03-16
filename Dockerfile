# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY ./back-end/go.mod ./back-end/go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code from the host to the container
COPY ./back-end ./

# Build the Go app
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
