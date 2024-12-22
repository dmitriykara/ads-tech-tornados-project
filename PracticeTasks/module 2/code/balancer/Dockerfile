# Use the official Go image to build and run the app
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and source code into the container
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o ttb-service ./cmd/balancer

# Expose the port the app will listen on (adjust if necessary)
EXPOSE 8080

# Command to run the application
CMD ["./ttb-service"]
