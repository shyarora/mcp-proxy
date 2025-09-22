BUILD_DIR=./build
BUILD=$(shell git rev-parse --short HEAD)@$(shell date +%s)
CURRENT_OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
CURRENT_ARCH := $(shell uname -m | tr '[:upper:]' '[:lower:]')
LD_FLAGS=-ldflags "-X main.BuildVersion=$(BUILD)"
GO_BUILD=CGO_ENABLED=0 go build $(LD_FLAGS)

.PHONY: build
build:
	@echo "This is a library-only project. Use 'go get github.com/shyarora/mcp-proxy' to import."
	go build ./...

.PHONY: buildLinuxX86
buildLinuxX86:
	@echo "This is a library-only project. Use 'go get github.com/shyarora/mcp-proxy' to import."
	GOOS=linux GOARCH=amd64 go build ./...

.PHONY: buildImage
buildImage:
	docker buildx build --platform=linux/amd64,linux/arm64 -t ghcr.io/tbxark/map-proxy:latest . --push --provenance=false

.PHONY: format
format:
	golangci-lint fmt --no-config --enable gofmt,goimports
	golangci-lint run --no-config --fix
	go fmt ./...
	go mod tidy