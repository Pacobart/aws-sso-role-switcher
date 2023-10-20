NAME := aws-sso-role-switcher

.DEFAULT_GOAL := build

.PHONY: deps
deps:
	go mod download

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
