#!/bin/bash

gofmt -w .
if [ $? -ne 0 ]; then
  echo "gofmt failed. Please fix the formatting issues."
  exit 1
fi

go vet ./...
if [ $? -ne 0 ]; then
    echo "go vet failed. Please fix the issues."
    exit 1
fi

go test ./...
if [ $? -ne 0 ]; then
    echo "go test failed. Please fix the test issues."
    exit 1
fi

echo "All checks passed successfully."
exit 0