NAME := aws-sso-role-switcher

.DEFAULT_GOAL := build

.PHONY: init
init:
	go mod init aws-sso-role-switcher

.PHONY: deps
deps:
	export GOPROXY=direct
	go mod download
	go mod tidy

.PHONY: build
build: deps
	go build -o bin/$(NAME)

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test: build
	go test ./...

.PHONY: check
check: lint test

.PHONY: run
run:
	go run main.go
