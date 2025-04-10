# Makefile

APP_NAME := auth-service
APP_PORT := 1300
DOCKER_IMAGE := $(APP_NAME):latest
COMPOSE_FILE := docker-compose.yml
MOCKERY_VERSION = v2.41.0

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

mock-install:
	go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)

test:
	go test -cover ./...

test-service:
	go test -cover ./internal/core/service/...

docker-build:
	go test -cover ./...
	docker-compose -f $(COMPOSE_FILE) build

docker-up:
	docker-compose -f $(COMPOSE_FILE) up -d

docker-down:
	docker-compose -f $(COMPOSE_FILE) down

docker-logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

docker-restart:
	docker-compose -f $(COMPOSE_FILE) up -d --build --force-recreate

docker-shell:
	docker exec -it $(APP_NAME) sh

docker-clean:
	docker image rm $(DOCKER_IMAGE)