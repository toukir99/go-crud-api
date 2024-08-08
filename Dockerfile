# Use the official Golang 1.22 image
FROM golang:1.22.5-alpine

# set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o go-crud-api ./cmd/server

# EXPOSE the port
EXPOSE 3000

# Run the executable
CMD ["./go-crud-api"]
