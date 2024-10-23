# Start with a small base image of Go on Alpine for smaller container size
FROM golang:1.18-alpine

# Set environment variables to make the Go app build statically
ENV CGO_ENABLED=0 GOOS=linux

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project files into the container
COPY . .

# Load dependencies effectively
RUN go mod tidy

# Build the Go binary
RUN go build -o /test-in-go ./main.go

EXPOSE 8060

# Define the command to run the application
CMD ["/test-in-go"]
