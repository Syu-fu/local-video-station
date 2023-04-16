COMPOSE_FILE := docker-compose.dev.yml

.PHONY: dev

dev:
	@docker-compose -f $(COMPOSE_FILE) up
