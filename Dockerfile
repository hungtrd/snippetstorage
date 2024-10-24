# Base image for both development and production
FROM golang:1.21-alpine AS base

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# ============ Development Stage ============
FROM base AS dev

# Install air for live reload in development
RUN apk update && apk add --no-cache make bash && \
    go install github.com/a-h/templ/cmd/templ@latest && \
    go install github.com/cosmtrek/air@v1.49.0

# Expose the port the app will run on
EXPOSE 8080

# Command to run in development mode with live reload
CMD ["air"]

# ============ Production Stage ============
FROM base AS prod

# Build the Go app
RUN go build -o /app/main cmd/api/main.go

# Expose the port for the production environment
EXPOSE 8080

# Command to run in production mode
CMD ["/app/main"]
