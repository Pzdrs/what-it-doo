SHELL := /bin/bash

build: clean
	go build -o bin/api ./cmd/api/

run: build
	./bin/api

sqlc:
	sqlc generate -f infra/db/sqlc.yaml

clean:
	rm -rf bin/*