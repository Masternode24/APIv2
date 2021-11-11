package main

import (
	"log"
	"os/exec"
)

func UpService() string {
	out, err := exec.Command("bash", "-c", "curl -s 'localhost:26657/status' | jq -r '.result.sync_info.catching_up' | tr -d '\n'").Output()
	if err != nil {
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
