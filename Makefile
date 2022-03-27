VERSION := $(shell git describe --tags)

build:
	@go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./tmp/helm-paramstore .
