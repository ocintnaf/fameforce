DATABASE_HOST=$(shell grep -oP '(?<=DATABASE_HOST=).*' .env)
DATABASE_PORT=$(shell grep -oP '(?<=DATABASE_PORT=).*' .env)
DATABASE_USER=$(shell grep -oP '(?<=DATABASE_USER=).*' .env)
DATABASE_PASSWORD=$(shell grep -oP '(?<=DATABASE_PASSWORD=).*' .env)
DATABASE_NAME=$(shell grep -oP '(?<=DATABASE_NAME=).*' .env)

new_migration:
	@if [ -z "$(name)" ]; then \
		echo "Please specify a name for the migration."; \
		exit 1; \
	fi
	@echo "Creating new empty up migration file..."
	migrate create -ext sql -dir migrations/ -seq $(name)

