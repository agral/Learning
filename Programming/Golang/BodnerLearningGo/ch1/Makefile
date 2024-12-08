.DEFAULT_GOAL := build
.PHONY: build fmt vet

build: vet
	go build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...
