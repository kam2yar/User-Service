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

# Install Protobuff compiler
RUN apt-get update
RUN apt install -y protobuf-compiler

# Install Protobuff plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

# Expose port for incoming traffic
EXPOSE 80
EXPOSE 8080

ENTRYPOINT ["air"]
