#!/bin/bash

# Build for macOS
env GOOS=darwin GOARCH=amd64 go build -o crawler_mac

# Build for Windows
env GOOS=windows GOARCH=amd64 go build -o crawler_windows.exe

# Build for Linux
env GOOS=linux GOARCH=amd64 go build -o crawler_linux