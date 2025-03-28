# Makefile for Docker Compose Operations

# Default docker-compose file
COMPOSE_FILE ?= docker-compose.yml

# Phony targets to ensure commands always run
.PHONY: up down psql help

# Start all services defined in docker-compose.yml
up:
	docker compose -f $(COMPOSE_FILE) up --remove-orphans -d

# Stop and remove containers, networks, volumes
down:
	docker compose -f $(COMPOSE_FILE) down --remove-orphans

# Connect to postgres container using psql
psql:
	docker exec -it postgres-container psql -U postgres

# Target to show help information
help:
	@echo "Docker Compose Makefile Commands:"
	@echo "  make up       - Start all services in detached mode"
	@echo "  make down     - Stop and remove containers and networks"
	