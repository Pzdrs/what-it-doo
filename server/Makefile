SHELL := /bin/bash

build: clean
	go build -o bin/api ./cmd/api/

run: swag build
	./bin/api

sqlc:
	sqlc generate -f infra/db/sqlc.yaml

swag:
	swag init -d ./cmd/api,./internal/apiserver -o ./api

clean:
	rm -rf bin/*