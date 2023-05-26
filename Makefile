SECONDS_SINCE_EPOCH=$(shell date +%s)

new_migration:
	@if [ -z "$(name)" ]; then \
		echo "Please specify a name for the migration."; \
		exit 1; \
	fi
	@echo "Creating new empty up migration file..."
	touch ./migrations/$(SECONDS_SINCE_EPOCH)-$(name).up.sql

	@echo "Creating new empty down migration file..."
	touch ./migrations/$(SECONDS_SINCE_EPOCH)-$(name).down.sql

