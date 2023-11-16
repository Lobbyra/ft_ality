#!/bin/bash

if [ "$(uname)" == "Linux" ]; then
    echo "Building for Linux"
    CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -o "$NAME" -tags static -ldflags "-s -w" src/main.go
elif [ "$(uname)" == "Darwin" ]; then
    echo "Building for macOS"
    GOOS=darwin go build -o "$NAME" src/main.go
else
    echo "Unsupported operating system: $(uname)"
    exit 1
fi
