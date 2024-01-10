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

# Add a new command with a subcommand and dynamically add flags
# Emaxple: make new-cmd cmd=storj subcmd=add flags="secret:secret:s:Storj secret to connect within:true:|path:path:p:Path where you want store the file on storj bucket:true:|file:file:f:Path of the file where is located:true:"
new-cmd:
	chmod +x ./scripts/create_cmd.sh
	./scripts/create_cmd.sh $(cmd) $(subcmd) "$(flags)" "$(PKG)/cmd/$(cmd)"

.PHONY: all build test clean deps new-cmd
