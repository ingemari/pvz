ifneq (,$(wildcard .env))
	include .env
	export
endif

MIGRATIONS_DIR ?= ./migrations
DATABASE_URL := postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(SSL_MODE)
PG_CONTAINER_NAME := postgres-dev

.PHONY: go
go:
	@go run cmd/main.go


.PHONY: pg
pg:
	@if [ $$(docker ps -aq -f name=$(PG_CONTAINER_NAME)) ]; then \
		echo "Postgres container already exists"; \
		docker start -a $(PG_CONTAINER_NAME); \
	else \
		echo "Starting new postgres container..."; \
		docker run --name $(PG_CONTAINER_NAME) \
			-e POSTGRES_PASSWORD=$(DATABASE_PASSWORD) \
			-e POSTGRES_USER=$(DATABASE_USER) \
			-e POSTGRES_DB=$(DATABASE_NAME) \
			-p $(DATABASE_PORT):$(DATABASE_PORT) \
			-d postgres; \
	fi

.PHONY: psql
psql:
	docker exec -it $(PG_CONTAINER_NAME) psql -U $(DATABASE_USER) -d $(DATABASE_NAME)

.PHONY: migrate_up
migrate_up:
	migrate -path $(MIGRATIONS_DIR) -database '$(DATABASE_URL)' up

.PHONY: migrate_down
migrate_down:
	migrate -path $(MIGRATIONS_DIR) -database '$(DATABASE_URL)' down

.PHONY: pg_rm
pg_rm:
	-docker stop $(PG_CONTAINER_NAME)
	-docker rm $(PG_CONTAINER_NAME)