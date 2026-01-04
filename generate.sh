#!/bin/bash
# generate.sh - Generate Go API client from OpenAPI specification using ogen
#
# Prerequisites:
#   go install github.com/ogen-go/ogen/cmd/ogen@latest
#
# Usage:
#   ./generate.sh
#
# This script:
#   1. Converts OpenAPI 3.1 spec to 3.0.3 for ogen compatibility
#   2. Runs ogen to generate Go code
#   3. Runs go mod tidy to update dependencies
#   4. Verifies the build compiles
#
# To update the OpenAPI spec:
#   cp /path/to/phoenix/schemas/openapi.json openapi/openapi.json

set -e

# Check if ogen is installed
if ! command -v ogen &> /dev/null; then
    echo "Error: ogen is not installed."
    echo "Install with: go install github.com/ogen-go/ogen/cmd/ogen@latest"
    exit 1
fi

# Check if OpenAPI spec exists
if [ ! -f "openapi/openapi.json" ]; then
    echo "Error: openapi/openapi.json not found."
    echo "Copy it from Phoenix repo:"
    echo "  cp /path/to/phoenix/schemas/openapi.json openapi/openapi.json"
    exit 1
fi

echo "Converting OpenAPI 3.1 to 3.0.3 for ogen compatibility..."
go run ./cmd/openapi-convert openapi/openapi.json openapi/openapi-v3.0.json

echo "Generating API code with ogen..."
ogen --package api --target internal/api --clean openapi/openapi-v3.0.json

echo "Running go mod tidy..."
go mod tidy

echo "Verifying build..."
go build ./...

echo ""
echo "Done! API client regenerated successfully."
echo ""
echo "Next steps:"
echo "  1. Review changes in internal/api/"
echo "  2. Update SDK wrapper code if needed for new/changed endpoints"
echo "  3. Run tests: go test ./..."
echo "  4. Run linter: golangci-lint run"
