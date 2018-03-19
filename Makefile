SHELL:=/bin/bash
GIT_REV?=$(shell git rev-parse --short HEAD)
BUILD_DATE=`date +%FT%T%z`
VERSION=0.0.1
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Hash=$(GIT_REV) -X main.BuildDate=$(BUILD_DATE)"
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

all: check install

build: test
	vgo fmt ./...
	vgo build $(LDFLAGS) ./cmd/gosvm 
	@mv gosvm ./bin/gosvm

install: build
	@cp ./bin/gosvm $(GOPATH)/bin/gosvm

.PHONY: check
check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v vendor/); do golint $${d}; done
	@go tool vet ${SRC}

.PHONY: test
test:
	vgo test ./...
