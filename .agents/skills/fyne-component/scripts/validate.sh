#!/usr/bin/env bash
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

echo "Checking formatting..."
unformatted="$(gofmt -l components showcase theme)"

if [[ -n "$unformatted" ]]; then
    echo "$unformatted"
    echo "FAIL: unformatted Go files"
    exit 1
fi

echo "Running unit tests..."
go test ./...

echo "Running static analysis..."
go vet ./...

echo "Running race tests..."
go test -race ./...

echo "All quality checks passed."
