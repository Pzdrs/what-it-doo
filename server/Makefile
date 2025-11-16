SHELL := /bin/bash
SWAG_INPUTS := $(shell find cmd/api/main.go internal/apiserver/controller -name "*.go")
SWAG_OUTPUTS := api/docs.go

TAG_NAME := $(shell git describe --abbrev=0 --tags)
SHA := $(shell git rev-parse HEAD)
VERSION_GIT := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
VERSION := $(if $(VERSION),$(VERSION),$(VERSION_GIT))

BIN_NAME := server
DOCKER_BUILD_PLATFORMS ?= linux/amd64,linux/arm64

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

build: 
	go build -o bin/api ./cmd/api/

run: 
	go run ./cmd/api/

worker: 
	go run ./cmd/worker/

sqlc:
	sqlc generate -f infra/db/sqlc.yaml

migrate:
	goose up

swag: $(SWAG_OUTPUTS)
	@echo "api docs are up to date"

$(SWAG_OUTPUTS): $(SWAG_INPUTS)
	swag fmt
	swag init -d ./cmd/api,./internal/apiserver -o ./api

clean:
	rm -rf bin/*

multi-arch-image-%:
	docker buildx build \
		--platform=$(DOCKER_BUILD_PLATFORMS) \
		-t registry.gitlab.pycrs.cz/tul/what-it-doo/server:$* \
		--build-arg VERSION=$(VERSION) \
		--provenance=false \
		..
