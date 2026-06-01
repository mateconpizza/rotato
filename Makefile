# rotato - simple spinner library for Go.
# See LICENSE file for copyright and license details.

FN		?= .

full: test lint

# Run tests
test:
	@echo '>> Testing go test ./... -race'
	@go test ./... -race
	@echo

# Run tests for a specific function
testfn:
	@echo '>> Testing function $(FN)'
	@go test -run $(FN) ./... -race

testrace:
	@echo '>> Testing'
	@go test ./... -race
	@echo

# Lint code with 'golangci-lint'
lint:
	@echo '>> Linting code'
	@go vet ./...
	golangci-lint run ./...

# run example
demo:
	@go run example/demo.go -demo

# run all example
all:
	@go run example/demo.go -all

ci: lint
	go mod tidy
	git diff --exit-code
	go build ./...
	go test -race ./...

.PHONY: test testfn lint full
