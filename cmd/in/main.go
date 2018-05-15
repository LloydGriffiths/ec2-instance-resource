package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/LloydGriffiths/ec2-instance-resource/in"
)

func main() {
	var req in.Request
	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		log.Fatalf("error reading request from stdin: %s", err)
	}

	r, err := in.Run(os.Args[1], &req)
	if err != nil {
		log.Fatalf("error running command: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(r); err != nil {
		log.Fatalf("error writing response to stdout: %s", err)
	}
}
