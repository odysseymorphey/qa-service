package main

import (
	"log"
	"os"
	"qa-service/internal/config"
	"qa-service/internal/di"
	"qa-service/internal/server"
)

func main() {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = "config/config.toml"
	}

	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	container := di.NewContainer(cfg)

	app := server.New(container, cfg)

	app.Run()
}
