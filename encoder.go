package main

import "fmt"
import	"os"


func XOR(data []byte, key byte) []byte {
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ key
	}
	return result
}

func main() {
	// 🔁 Read actual command content from payload.txt
	payload, err := os.ReadFile("payload.txt")
	if err != nil {
		fmt.Println("Failed to read payload.txt:", err)
		return
	}

	key := byte(0xAA)
	encoded := XOR(payload, key)

	// 💾 Overwrite payload.txt with the encoded version
	err = os.WriteFile("payload.txt", encoded, 0644)
	if err != nil {
		fmt.Println("Failed to write payload:", err)
		return
	}

	fmt.Println("Payload encoded and saved to payload.txt")
}
