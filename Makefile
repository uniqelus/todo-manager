.PHONY: lint build

GOOS ?= linux
GOARCH ?= amd64
CGO_ENABLED ?= 0

DOCKER_COMPOSE_FILE = docker/docker-compose.yaml
DOCKER_COMPOSE = docker compose -f $(DOCKER_COMPOSE_FILE)

build:
	@mkdir -p bin
	@for dir in cmd/*/; do \
		name=$$(basename $$dir); \
		CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$$name ./$$dir; \
	done

lint:
	golangci-lint run --fix ./...

up: build
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down --remove-orphans -v

start:
	$(DOCKER_COMPOSE) start

stop:
	$(DOCKER_COMPOSE) stop

logs:
	$(DOCKER_COMPOSE) logs -f

ps:
	$(DOCKER_COMPOSE) ps --all --format "table {{.Service}}\t{{.Status}}\t{{.Ports}}"

restart: stop start