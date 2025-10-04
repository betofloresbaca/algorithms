# Build and run the game
.PHONY: build run clean test

# Default target
all: build

# Build the game executable
build:
	go build -o algorithms ./cmd/algorithms

# Run the game directly
run:
	go run ./cmd/algorithms

# Clean build artifacts
clean:
	rm -f algorithms

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod tidy
	go mod download

# Development: format, vet, and test
dev: fmt vet test

# Build for release
release: clean fmt vet test build