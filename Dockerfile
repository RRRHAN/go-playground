# Use the official Golang image
FROM golang:1.20-alpine as builder

# Set the working directory to /app
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Final image
FROM alpine:latest

# Set the working directory to /root
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose port 9999
EXPOSE 9999

# Command to run the app
CMD ["./app"]