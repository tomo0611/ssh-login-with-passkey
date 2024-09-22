package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// 1. Generate a login token
	// https://passkey.tomo0611.dev/api/v1/generateLoginToken
	// 2. Show QR Code in the terminal (https://passkey.tomo0611.dev/login/{loginToken})
	// 3. Start Long Polling to check if the login is successful
	// https://passkey.tomo0611.dev/api/v1/loginLongPolling
	// 3.1. If the login is successful, the server return username
	// 4. Start a shell

	// Create a new command to run /bin/bash
	cmd := exec.Command("/bin/bash")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Wait for the command to finish
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
