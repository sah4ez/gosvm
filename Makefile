SHELL:=/bin/bash

.PHONY: build
build:
	go fmt ./...
	go build ./cmd/gosvm 
	@mv gosvm ./bin/gosvm

dev-build: build
	@cp ./bin/gosvm $(GOPATH)/bin/gosvm
