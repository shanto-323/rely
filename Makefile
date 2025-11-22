build: confirm
	echo "Running [go mod tidy]"
	@go mod tidy; 
	echo "Running [go mod vendor]"
	@go mod vendor; 
	echo "Running [docker-compose build --no-cache]"
	docker-compose -f ./scripts/compose.yaml build --no-cache
	echo "Build successful"

run: 
	echo "Running [docker-compose up -d $(service)]"
	docker-compose -f ./scripts/compose.yaml up -d $(service)

stop: confirm 
	echo "Running [docker-compose down -v]"
	docker-compose -f ./scripts/compose.yaml down -v

re_build: confirm
	@if [ -z "$(service)" ]; then \
		echo "Invalid service name"; \
		exit 1; \
	fi
	@echo "Rebuilding service: $(service)"
	@docker-compose -f ./scripts/compose.yaml rm -sf $(service)
	@docker-compose -f ./scripts/compose.yaml build --no-cache $(service)
	@make run service=$(service)


migration_new: 
	@if [ -z "$(name)" ]; then \
		@echo "Error: name parameter is required" \
        @echo "Usage: task migrations:new name=migration_name" \
		exit 1; \
	fi
	@echo "Creating migration file for $(name)"
	@tern new -m ./internal/repository/database/migration ${name}


confirm:
	@read -p "Press [y/Y] to continue: " value; \
	value=$$(echo $$value | tr '[:upper:]' '[:lower:]'); \
	if [ "$$value" = "y" ]; then \
		echo "Running target..."; \
	else \
		exit 1; \
	fi

