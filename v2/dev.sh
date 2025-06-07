#!/usr/bin/env bash
set -e

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

# Run tests
echo "Running go test..."
go test ./...
if [ $? -ne 0 ]; then
    echo "Tests failed. Please fix the issues before committing."
    exit 1
fi

# Run the game
echo "Running the game..."
go run ./cmd/main.go
if [ $? -ne 0 ]; then
    echo "Game failed to run. Please fix the issues before committing."
    exit 1
fi