include .env

DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_NAME ?= postgres
DATABASE_USER ?= postgres
DATABASE_PASSWORD ?= postgres

MIGRATIONS_DIR=migrations

COLOR_INFO=\033[1;34m

new-migration:
	@if [ -z "$(name)" ]; then \
		echo "Please specify a name for the migration."; \
		exit 1; \
	fi
	@echo "${COLOR_INFO}Creating new migration: ${name}..."
	migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq $(name)

migrate-up:
	@echo "${COLOR_INFO}Running migrations..."
	migrate -path ${MIGRATIONS_DIR} -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" up

migrate-down:
	@echo "${COLOR_INFO}Rolling back migrations..."
	migrate -path ${MIGRATIONS_DIR} -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" down

server:
	@echo "${COLOR_INFO}Starting server..."
	go run cmd/server/main.go

test:
	@echo "${COLOR_INFO}Running tests..."
	go test -v ./...

dc-up:
	@echo "${COLOR_INFO}Starting docker-compose..."
	docker-compose up

dc-down:
	@echo "${COLOR_INFO}Stopping docker-compose..."
	docker-compose down -v


