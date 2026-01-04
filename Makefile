.PHONY: all test lint build generate clean

all: test lint build

test:
	go test -v -race ./...

lint:
	golangci-lint run

build:
	go build ./...

generate:
	./generate.sh

clean:
	rm -rf internal/api/*.go

help:
	@echo "Available targets:"
	@echo "  test     - Run tests with race detection"
	@echo "  lint     - Run golangci-lint"
	@echo "  build    - Build all packages"
	@echo "  generate - Regenerate API client from OpenAPI spec"
	@echo "  clean    - Remove generated files"
