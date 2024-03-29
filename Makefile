# Go variables
GO := go
PKG := github.com/dalecosta1/sinaloa-cli
MAIN_FILE := main.go

# Main target
all: build

# Build the executable
build:
	$(GO) build -o sinaloa $(MAIN_FILE)
	@file="sinaloa"; \
	folder="build"; \
	if [ ! -d "$$folder" ]; then \
		mkdir -p "$$folder"; \
	fi; \
	mv "$$file" "$$folder/"

# Clean build artifacts
build-clean:
	$(GO) clean
	rm -rf build

# Install dependencies
deps:
	$(GO) get -v -t -d ./...

# Cleans up the module's go.mod and go.sum
mod-tidy:
	$(GO) mod tidy

# Run tests
test:
	@cd tests && $(GO) test ./...

# Add env vars from .env file
add-env:
	@cd scripts && chmod +x ./set_env_var.sh && . set_env_var.sh

# Add a new command with a subcommand and dynamically add flags
# Example: make new-cmd cmd=storj subcmd=add flags="secret:secret:s:Storj secret to connect within:true:|path:path:p:Path where you want store the file on storj bucket:true:|file:file:f:Path of the file where is located:true:"
new-cmd:
	chmod +x ./scripts/create_cmd.sh
	./scripts/create_cmd.sh $(cmd) $(subcmd) "$(PKG)/cmd/$(cmd)" "$(flags)"

# Add a new command with a subcommand and dynamically add flags
# Example make new_sub cmd=storj subcmd=put flags="msg:msg:m:Message to receive:true:|path:path:p:Path where you want store the file on storj bucket:true:"
new-sub:
	chmod +x ./scripts/create_sub.sh
	./scripts/create_sub.sh $(cmd) $(subcmd) "$(PKG)/cmd/$(cmd)" "$(flags)"

# Set the version of the project, used only to generate thefinal build of the cli
set-version:
	chmod +x ./scripts/set_version.sh
	./scripts/set_version.sh

.PHONY: all build test clean deps new-cmd new-sub set-version
