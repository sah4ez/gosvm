SHELL:=/bin/bash

.PHONY: build
build:
	vgo fmt ./...
	vgo build ./cmd/gosvm 
	@mv gosvm ./bin/gosvm

install: build
	@cp ./bin/gosvm $(GOPATH)/bin/gosvm
