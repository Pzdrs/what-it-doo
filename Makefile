dev:
	@trap 'make dev-down' EXIT; docker compose -f ./docker/compose.dev.yaml up --remove-orphans

dev-down:
	docker compose -f ./docker/compose.dev.yaml down --remove-orphans