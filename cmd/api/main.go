package main

// @title SpaceDevs API
// @version 1.0
// @description Backend API for SpaceDevs project.
// @BasePath /api/v1
import (
	"log"
	"os"

	"github.com/Raphsodyz/spacedevs-go/bootstrap"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	log.Printf("Starting SpaceDevs API in %s environment", env)

	srv, err := bootstrap.NewServer("config_" + env)
	if err != nil {
		log.Fatalf("failed to bootstrap server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
