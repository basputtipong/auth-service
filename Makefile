# Makefile

APP_NAME := auth-service
APP_PORT := 1300
DOCKER_IMAGE := $(APP_NAME):latest
COMPOSE_FILE := docker-compose.yml

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

test:
	go test ./...

docker-build:
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
	docker exec -it auth_service sh

docker-clean:
	docker image rm $(DOCKER_IMAGE)