//go:build linux

package exec

import (
	"os/exec"
	"log"
)

func ExecutePayload(decoded []byte) {
	stringPayload := string(decoded)
	cmd := exec.Command("bash", "-c", stringPayload)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
