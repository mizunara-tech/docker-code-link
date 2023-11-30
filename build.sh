#!/bin/bash

# Linux amd64
echo "Building for Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o build/docker-code-link-linux-amd64 ./cmd/docker-code-link

# Windows amd64
echo "Building for Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o build/docker-code-link-windows-amd64.exe ./cmd/docker-code-link

# MacOS Intel
echo "Building for MacOS Intel..."
GOOS=darwin GOARCH=amd64 go build -o build/docker-code-link-mac-intel ./cmd/docker-code-link

# MacOS M1
echo "Building for MacOS M1..."
GOOS=darwin GOARCH=arm64 go build -o build/docker-code-link-mac-m1 ./cmd/docker-code-link

echo "Build complete."
