#!/usr/bin/make -f

VERSION := $(shell git describe)

test: fmt
	go test -race -cover -timeout=1s -count=1 ./...

fmt:
	@go version && go fmt ./... && go mod tidy

install: test
	go install -ldflags="-X 'main.Version=$(VERSION)'" github.com/mdwhatcott/ezblog/cmd/...

manual:
	make install && rm hello-world/index.html && ez -source cmd/ez/example.md -dest . && cat hello-world/index.html