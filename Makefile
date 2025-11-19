BIN_NAME ?= qa-service
CONFIG_PATH ?= config/config.toml

.PHONY: build run docker-up docker-down test lint

build:
	go build -o bin/$(BIN_NAME) ./cmd

run:
	CONFIG_PATH=$(CONFIG_PATH) go run ./cmd

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

test:
	GOCACHE=$$(pwd)/.cache go test ./...

lint:
	golangci-lint run ./...
