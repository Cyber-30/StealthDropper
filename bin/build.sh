#!/bin/bash

echo "[*] Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o bin/dropper_windows.exe main.go

echo "[*] Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o bin/dropper_linux main.go

echo "[*] Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o bin/dropper_mac main.go

echo "[+] Builds completed. Check the 'bin/' directory."
