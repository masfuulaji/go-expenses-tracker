# Use the official Golang image as the base
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install Air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Run the Go app
CMD ["air","cmd/main.go"]
