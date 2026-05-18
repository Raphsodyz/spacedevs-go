package main

// @title SpaceDevs API
// @version 1.0
// @description Backend API for SpaceDevs project.
// @BasePath /api/v1

import (
	"log"

	"github.com/Raphsodyz/spacedevs-go/bootstrap"
)

func main() {
	srv, err := bootstrap.NewServer("config_development")
	if err != nil {
		log.Fatalf("failed to bootstrap server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
