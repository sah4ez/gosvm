SHELL="/bin/bash"

build:
	go fmt ./...
	go build ./cmd/gosvm
