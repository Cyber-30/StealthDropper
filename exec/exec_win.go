//go:build windows

package exec

import (
	"os/exec"
	"log"
)

func ExecutePayload(decoded []byte) {
	stringPayload := string(decoded)
	cmd := exec.Command("cmd", "/C", stringPayload)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
