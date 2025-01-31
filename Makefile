APP_NAME := goBackend
CMD_DIR := ./cmd/$(APP_NAME)
BUILD_DIR := ./build
GO_FILES := $(shell find . -type f -name '*.go')

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go
	@echo "Build completed! Binary is located at $(BUILD_DIR)/$(APP_NAME)"

# Run the application
.PHONY: run
run:
	@echo "Running the application..."
	@go run $(CMD_DIR)/main.go || true
	@echo "Server exited cleanly." || exit 0

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Lint code (requires golangci-lint to be installed)
.PHONY: lint
lint:
	@echo "Linting Go code..."
	@golangci-lint run

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned!"

# Run everything (build, format, lint, and test)
.PHONY: all-checks
all-checks: fmt lint test build

# Help menu
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build      Build the application"
	@echo "  run        Run the application"
	@echo "  test       Run tests"
	@echo "  fmt        Format Go code"
	@echo "  lint       Lint Go code (requires golangci-lint)"
	@echo "  clean      Clean build files"
	@echo "  all-checks Run formatting, linting, testing, and build"
	@echo "  help       Show this help message"

