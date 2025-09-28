SHELL := /bin/bash

build: clean
	go build -o bin/api ./cmd/api/

run: build
	set -a && source .env && set +a	
	./bin/api

clean:
	rm -rf bin/*