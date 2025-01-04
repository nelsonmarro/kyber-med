project_name = kybermed
image_name = kybermed:latest
dev_image_name = kybermed-dev 

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run ./cmd/kybermed/main.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

up: ## Run the project in a local container
	make up-silent
	make shell

build: ## Generate docker image
	docker build -t $(image_name) -f deploy/Dockerfile .

build-no-cache: ## Generate docker image with no cache
	docker build --no-cache -t $(image_name) -f deploy/Dockerfile .

up-silent: ## Run local container in background
	make delete-container-if-exist
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app

up-silent-prefork: ## Run local container in background with prefork
	make delete-container-if-exist
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app -prod

delete-container-if-exist: ## Delete container if it exists
	docker stop $(project_name) || true && docker rm $(project_name) || true

shell: ## Run interactive shell in the container
	docker exec -it $(project_name) /bin/sh

stop: ## Stop the container
	docker stop $(project_name)

start: ## Start the container
	docker start $(project_name)

build-dev:
	docker build --target dev -t $(dev_image_name) -f deploy/Dockerfile .

up-dev: ## Run dev container con volumen montado y modo interactivo
	docker run -it --rm -p 3000:3000 -v "$$(pwd)":/app $(dev_image_name)

