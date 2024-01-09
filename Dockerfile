# Use an official Golang runtime as a parent image
FROM golang:1.21.5

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Install air
RUN go install github.com/cosmtrek/air@latest

# Expose port 8080 for incoming traffic
EXPOSE 8080

ENTRYPOINT ["air"]
