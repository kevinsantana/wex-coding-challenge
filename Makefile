include .env
export

SHELL=/bin/bash

all: help

envvars:
	export $(grep -v '^#' .env | xargs)

docker-build: ## docker-build: build the local image
	docker build --force-rm --tag purchase:0.1.0 .

docker-run: ## docker-run: run the local container
	docker run --name purchase --network wex -p 3060:3060 purchase:0.1.0

docker-postgres: ## docker-postgres: run local postgres
	docker run --name postgres --network wex -p 5432:5432 -d -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=postgres -e POSTGRES_DB=purchase postgres:14.2-alpine

compose-build: ## compose-build: build docker images with compose
	docker-compose build

up: ## up: spin-up docker-compose containers.
	docker-compose up -d

down: ## down: Stop docker-compose containers..
	docker-compose down

lint: ## lint: Apply golint.
	golangci-lint run -E gosec -E gofmt -E goimports --skip-dirs tests

help: ## help: Show this help message.
	@echo "usage: make [target] ..."
	@echo ""
	@echo "targets:"
	@grep -Eh '^.+:(\w+)?\ ##\ .+' ${MAKEFILE_LIST} | cut -d ' ' -f '3-' |  column -t -s ':' | egrep --color '^[^ ]*'

build: ## build: Build go executable
	go build .

run: ## run: run purchase api
	go run . api