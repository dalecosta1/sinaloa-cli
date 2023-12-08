# Makefile for Go projects

# Go variables
GO := go
PKG := github.com/dalecosta1/sinaloa-cli
MAIN_FILE := main.go

# Main target
all: build

# Build the executable
build:
	$(GO) build -o sinaloa $(MAIN_FILE)

# Run tests
test:
	$(GO) test ./...

# Clean build artifacts
clean:
	$(GO) clean
	rm -f sinaloa

# Install dependencies
deps:
	$(GO) get -v -t -d ./...

.PHONY: all build test clean deps
