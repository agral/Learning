.DEFAULT_GOAL := build
.PHONY: build clean fmt vet

build: vet
	go build

clean:
	go clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...
