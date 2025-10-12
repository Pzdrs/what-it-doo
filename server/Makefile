SHELL := /bin/bash
SWAG_INPUTS := $(shell find cmd/api/main.go internal/apiserver/controller -name "*.go")
SWAG_OUTPUTS := api/docs.go

build: 
	go build -o bin/api ./cmd/api/

sqlc:
	sqlc generate -f infra/db/sqlc.yaml

swag: $(SWAG_OUTPUTS)
	@echo "api docs are up to date"

$(SWAG_OUTPUTS): $(SWAG_INPUTS)
	swag fmt
	swag init -d ./cmd/api,./internal/apiserver -o ./api

clean:
	rm -rf bin/*