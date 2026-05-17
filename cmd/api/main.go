package main

import (
	"log"

	"github.com/Raphsodyz/spacedevs-go/bootstrap"
)

func main() {
	srv, err := bootstrap.NewServer("config-local")
	if err != nil {
		log.Fatalf("failed to bootstrap server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
