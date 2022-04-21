MAIN_BINARY := api-server
MAIN_PKG := github.com/edfoh/go-api-starter/cmd/$(MAIN_BINARY)

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
GOLANGCI_LINT_BIN := $(GOBIN)/golangci-lint

.PHONY: build clean ensure_deps lint lint_fix clean

default: build

build:
	@go build $(MAIN_PKG)

ensure_deps:
	@go mod tidy
	@go mod vendor

install-tools:
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

lint: $(GOLANGCI_LINT_BIN)
	"$(GOLANGCI_LINT_BIN)" run --sort-results --verbose

lint_fix: $(GOLANGCI_LINT_BIN)
	"$(GOLANGCI_LINT_BIN)" run --sort-results --verbose --fix

clean:
	@go clean $(MAIN_PKG) && rm -f ./$(MAIN_BINARY)
