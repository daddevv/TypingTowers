#!/usr/bin/env bash
set -e

# Ensure Go tooling uses the vendor directory
export GOFLAGS=-mod=vendor

# Format check
echo "Running gofmt..."
if [ -n "$(gofmt -l .)" ]; then
    gofmt -w .
    if [ $? -ne 0 ]; then
        echo "gofmt found issues. Please fix them before committing."
        exit 1
    else
        echo "gofmt has formatted the code."
    fi
else
    echo "gofmt found no issues."
fi

# Vet check
echo "Running go vet..."
go vet ./...
if [ $? -ne 0 ]; then
    echo "go vet found issues. Please fix them before committing."
    exit 1
fi

# Run tests (headless mode for Ebiten, see next section)
echo "Running go test..."
go test ./...
if [ $? -ne 0 ]; then
    echo "Tests failed. Please fix the issues before committing."
    exit 1
fi