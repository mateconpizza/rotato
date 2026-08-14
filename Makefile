# rotato - simple spinner library for Go.
# See LICENSE file for copyright and license details.

FN		?= .

full: test lint

# Run tests
test:
	@go test -race -count=1 ./...
	@echo

# Run tests for a specific function
testfn:
	@echo '>> Testing function $(FN)'
	@go test -race -run $(FN) ./...

# Run tests with gotestsum
testsum:
	@echo '>> Testing with gotestsum'
	@gotestsum --format=testname --hide-summary=skipped

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
