echo "[*] Building StealthDropper for all OS targets..."

GOOS=windows GOARCH=amd64 go build -o ../bin/stealthdropper_windows.exe ./cmd/stealthdropper
GOOS=linux   GOARCH=amd64 go build -o ../bin/stealthdropper_linux ./cmd/stealthdropper
GOOS=darwin  GOARCH=amd64 go build -o ../bin/stealthdropper_macos ./cmd/stealthdropper

echo "[+] Binaries built successfully in ./bin"
