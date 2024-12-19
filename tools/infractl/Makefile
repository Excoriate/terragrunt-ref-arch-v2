# Makefile for InfraCTL CLI

# Go parameters
GO ?= go
GOBUILD = $(GO) build
GOTEST = $(GO) test
GOMOD = $(GO) mod
BINARY_NAME = infractl
TARGET_DIR = ./target
BINARY_PATH = $(TARGET_DIR)/$(BINARY_NAME)

# Directories
SRC_DIR = .
INTERNAL_DIR = ./internal
PKG_DIR = ./pkg
TEST_DIR = ./tests

# Build flags
VERSION ?= $(shell git describe --tags --always --dirty)
LDFLAGS = -s -w \
	-X github.com/your-org/infractl/pkg/version.Version=$(VERSION) \
	-X github.com/your-org/infractl/pkg/version.BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

GOFLAGS = -trimpath
GO_PACKAGES = ./...

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[0;33m
RED = \033[0;31m
NC = \033[0m

# Default target
.PHONY: all
all: clean deps lint test build

# Install dependencies
.PHONY: deps
deps:
	@echo "${YELLOW}Installing dependencies...${NC}"
	$(GOMOD) tidy
	$(GOMOD) verify
	$(GO) mod download

# Run tests with coverage
.PHONY: test
test:
	@echo "${YELLOW}Running tests...${NC}"
	@mkdir -p coverage
	$(GOTEST) -v -race -covermode=atomic -coverprofile=coverage/coverage.out $(GO_PACKAGES)
	$(GO) tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "${GREEN}✓ Tests completed${NC}"

# Build the binary
.PHONY: build
build:
	@echo "${YELLOW}Building binary...${NC}"
	@mkdir -p $(TARGET_DIR)
	CGO_ENABLED=0 $(GOBUILD) $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BINARY_PATH) $(SRC_DIR)/main.go
	@echo "${GREEN}✓ Binary built at $(BINARY_PATH)${NC}"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "${YELLOW}Cleaning build artifacts...${NC}"
	rm -rf $(TARGET_DIR)
	rm -rf coverage
	$(GO) clean -cache -testcache
	@echo "${GREEN}✓ Clean completed${NC}"

# Run linters
.PHONY: lint
lint:
	@echo "${YELLOW}Running linters...${NC}"
	golangci-lint run $(GO_PACKAGES)
	@echo "${GREEN}✓ Linting completed${NC}"

# Format code
.PHONY: fmt
fmt:
	@echo "${YELLOW}Formatting code...${NC}"
	$(GO) fmt $(GO_PACKAGES)
	@echo "${GREEN}✓ Code formatted${NC}"

# Install development tools
.PHONY: tools
tools:
	@echo "${YELLOW}Installing development tools...${NC}"
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install golang.org/x/tools/cmd/goimports@latest
	@echo "${GREEN}✓ Development tools installed${NC}"

# Run the binary
.PHONY: run
run: build
	@echo "${YELLOW}Running InfraCTL CLI...${NC}"
	$(BINARY_PATH)

# Cross-compilation for different platforms
.PHONY: build-all
build-all:
	@echo "${YELLOW}Cross-compiling for multiple platforms...${NC}"
	@mkdir -p $(TARGET_DIR)
	GOOS=linux   GOARCH=amd64 $(GOBUILD) $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(TARGET_DIR)/$(BINARY_NAME)-linux-amd64
	GOOS=darwin  GOARCH=amd64 $(GOBUILD) $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(TARGET_DIR)/$(BINARY_NAME)-darwin-amd64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(TARGET_DIR)/$(BINARY_NAME)-windows-amd64.exe
	@echo "${GREEN}✓ Cross-compilation completed${NC}"

# Generate documentation
.PHONY: docs
docs:
	@echo "${YELLOW}Generating documentation...${NC}"
	$(GO) run ./cmd/gendocs/main.go
	@echo "${GREEN}✓ Documentation generated${NC}"

# Validate configuration
.PHONY: validate
validate:
	@echo "${YELLOW}Validating infrastructure configurations...${NC}"
	$(BINARY_PATH) validate
	@echo "${GREEN}✓ Configuration validation completed${NC}"

# Help target
.PHONY: help
help:
	@echo "InfraCTL CLI Development Makefile"
	@echo "================================="
	@echo "Available targets:"
	@echo "  all         - Clean, install deps, lint, test, and build"
	@echo "  deps        - Install dependencies"
	@echo "  test        - Run tests with coverage"
	@echo "  build       - Build the binary"
	@echo "  clean       - Remove build artifacts"
	@echo "  lint        - Run linters"
	@echo "  fmt         - Format code"
	@echo "  tools       - Install development tools"
	@echo "  run         - Build and run the binary"
	@echo "  build-all   - Cross-compile for multiple platforms"
	@echo "  docs        - Generate documentation"
	@echo "  validate    - Validate infrastructure configurations"
