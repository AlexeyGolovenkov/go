# Go application Makefile
APP_NAME ?= account
BINARY_NAME ?= bin/$(APP_NAME)
MAIN_PATH ?= ./cmd
VERSION ?= $(shell git decribe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS ?= -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Go file to format and lint
GOFILES ?= $(shell find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*")

# Default target
.DEFAULT_GOAL := build

# Build application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p bin
	go build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	./$(BINARY_NAME)

# Running without building (if binary exists)
run-only:
	@echo "Running $(APP_NAME)..."
	./$(BINARY_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Update dependencies
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Format code
ftm:
	@echo "Formatting code..."
	gofmt -s -w $(GOFILES)
	gofumpt -l -w -s $(GOFILES)
	goimports -l -w $(GOFILES)

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run --timeout=10m -v ./...

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v coverprofile=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. ./...

# Clear build artefacts
clean:
	@echo "Cleaning build artefacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install application
install: build
	@echo "Installing $(APP_NAME)..."
	go install $(LDFLAGS) $(MAIN_PATH)

# Show help
help:
	@echo "Available targets:"
	@echo "  build    		- Build the application"
	@echo "  run 			- Build and run the application"
	@echo "  run-only 		- Run the application without building"
	@echo "  deps 			- Install dependencies"
	@echo "  deps-update 	- Update dependencies"
	@echo "  fmt 			- Format code"
	@echo "  lint 			- Run linter"
	@echo "  test 			- Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  bench 			- Run benchmarks"
	@echo "  clean 			- Clean build artefacts"
	@echo "  install 		- Install the application"
	@echo "  docs 			- Generate documentation"
	@echo "  help 			- Show this help"

