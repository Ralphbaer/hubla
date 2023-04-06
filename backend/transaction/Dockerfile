# Use the official Go image as the base image
FROM golang:1.17

# Set up the working directory for your project
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Install the mockgen tool
RUN go get -u github.com/golang/mock/mockgen

# Make sure the $PATH includes the installed mockgen binary
ENV PATH="/go/bin:${PATH}"

# Copy the rest of your project files into the container
COPY . .

# Run the go generate command
RUN go generate ./...

# Build the transaction-service binary
RUN go build -o .bin/transaction-service ./transaction/main.go

# Expose the port 3000 for the transaction-service
EXPOSE 3000

# Set the entrypoint for the container
CMD [".bin/transaction-service", "0.0.0.0"]
