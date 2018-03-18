SHELL:=/bin/bash
GIT_REV?=$(shell git rev-parse --short HEAD)
BUILD_DATE=`date +%FT%T%z`
VERSION=0.0.1
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Hash=$(GIT_REV) -X main.BuildDate=$(BUILD_DATE)"

.PHONY: build
build:
	vgo fmt ./...
	vgo build $(LDFLAGS) ./cmd/gosvm 
	@mv gosvm ./bin/gosvm

install: build
	@cp ./bin/gosvm $(GOPATH)/bin/gosvm

.PHONY: test
test:
	vgo test ./...
