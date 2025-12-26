# Use the official Go image to create a binary
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app 

# Copy go.mod and go.sum first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download 

# Copy the rest of the source code
COPY *.go ./   
 
# Compile the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Bind to a TCP port (optional)
EXPOSE 8080

# Specify command to run when image is used
CMD ["/docker-gs-ping"]
