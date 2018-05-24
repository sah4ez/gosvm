NAME=gosvm
SHELL:=/bin/bash
GIT_REV?=$(shell git rev-parse --short HEAD)
BUILD_DATE=`date +%FT%T%z`
VERSION=$(shell cat VERSION)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Hash=$(GIT_REV) -X main.BuildDate=$(BUILD_DATE)"
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

GO=CC=gcc vgo
CC=gcc

all: check install

.PHONY: clean
clean:
	rm -rf dist/

build: test
	@$(GO) fmt ./...
	$(GO) build $(LDFLAGS) -o ./bin/$(NAME) ./cmd/$(NAME) 

.PHONY: test
test:
	$(GO) test ./...

install: build
	@cp ./bin/$(NAME) $(GOPATH)/bin/$(NAME)

.PHONY: check
check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v vendor/); do golint $${d}; done
	@$(GO) tool vet ${SRC}

.PHONY: build-all
build-all:
	mkdir -p _build
	GOOS=darwin  GOARCH=amd64 $(GO) build $(LDFLAGS) -o _build/$(NAME)-$(VERSION)-darwin-amd64 ./cmd/$(NAME)
	GOOS=linux   GOARCH=amd64 $(GO) build $(LDFLAGS) -o _build/$(NAME)-$(VERSION)-linux-amd64 ./cmd/$(NAME)
	GOOS=linux   GOARCH=arm   $(GO) build $(LDFLAGS) -o _build/$(NAME)-$(VERSION)-linux-arm ./cmd/$(NAME)
	GOOS=linux   GOARCH=arm64 $(GO) build $(LDFLAGS) -o _build/$(NAME)-$(VERSION)-linux-arm64 ./cmd/$(NAME)
	cd _build; sha256sum * > sha256sums.txt
#GOOS=windows GOARCH=amd64 $(GO) build $(LD_FLAGS) -o _build/$(NAME)-$(VERSION)-windows-amd64 ./cmd/$(NAME)

release: clean 
	go get github.com/goreleaser/goreleaser/...
	goreleaser
