DOCKER_COMPOSE = docker-compose -f deploy/docker-compose.yml
APP_SERVICE = app
GENERATOR_IMAGE = ghcr.io/a-h/templ:latest
PROJECT_DIR = $(PWD)

.PHONY: help up down build generate dev test clean logs shell

help: ## This help dialog
	@grep -E '^[a-zA-Z_-]+:.*?##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run-local-server: ## Run the app locally
	go run ./cmd/kybermed_server/main.go

run-ui-desktop: ## Run the app locally
	go run ./cmd/kybermed_UI/main.go

run-ui-mobile: ## Run the app locally
	go run -tags mobile ./cmd/kybermed_UI/main.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

generate: ## Generate code with Templ
	docker run --rm \
		-v $(PWD):/app \
		-w /app \
		--user $(shell id -u):$(shell id -g) \
		$(GENERATOR_IMAGE) generate

up: generate ## Run the project in a local container
	$(DOCKER_COMPOSE) up --build

down: ## Stop all running containers
	$(DOCKER_COMPOSE) down

build: ## Build the production Docker image
	docker build --no-cache -t kybermed_server:latest -f deploy/Dockerfile .

dev: generate ## Run development environment with Air
	$(DOCKER_COMPOSE) up $(APP_SERVICE) --build

test:
	go test -v ./...

logs: ## View logs from the app container
	$(DOCKER_COMPOSE) logs -f $(APP_SERVICE)

shell: ## Open an interactive shell in the app container
	docker exec -it $(APP_SERVICE) /bin/sh

clean: ## Clean up Docker volumes and containers
	$(DOCKER_COMPOSE) down -v
