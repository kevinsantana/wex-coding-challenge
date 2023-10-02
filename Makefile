include .env
export

SHELL=/bin/bash

.PHONY: build
build:
	go build .

.PHONY: run-api
run-api:
	go run . api