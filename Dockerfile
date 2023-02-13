# Use an official golang image as the base image
FROM golang:latest

# Set the working directory in the container to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the code to the container
COPY . .

# Build the go application
RUN go build

# Expose port 8080 to the host
EXPOSE 8000

# Specify the command to run when the container starts
CMD ["./assessment"]
