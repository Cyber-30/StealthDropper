//go:build windows

package main

import "os/exec"
import "log"

func executepayload(decoded []byte) {
	string_payload := string(decoded)
	cmd := exec.Command("cmd", "/C", string_payload)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
