include .env
export

SHELL=/bin/bash

all: help

up:
	docker-compose up -d

down:
	docker-compose down

lint: ## lint: Apply golint.
	golangci-lint run -E gosec -E gofmt -E goimports --skip-dirs tests

help: ## help: Show this help message.
	@echo "usage: make [target] ..."
	@echo ""
	@echo "targets:"
	@grep -Eh '^.+:(\w+)?\ ##\ .+' ${MAKEFILE_LIST} | cut -d ' ' -f '3-' |  column -t -s ':' | egrep --color '^[^ ]*'

build:
	go build .

run:
	go run . api