BIN_DIR=build/bin
BIN_NAME=mines_microservice
CMD_PATH=cmd/mines_microservice/main.go

build:
	@echo "Building the project..."
	@mkdir -p $(BIN_DIR)
	@go fmt ./...
	@go vet ./...
	@go build -o $(BIN_DIR)/$(BIN_NAME) $(CMD_PATH)

run: build
	@echo "Running the project..."
	@$(BIN_DIR)/$(BIN_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

DOCKER_COMPOSE_DIR = build/docker

build-docker:
	@echo "Building Docker images without cache..."
	@docker compose -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml build --no-cache

build-docker-dev:
	@echo "Building Docker images without cache..."
	@docker compose --env-file .env.local --profile development -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml build --no-cache

run-docker-dev: stop-docker-dev build-docker-dev
	@echo "Starting Docker containers for development..."
	@docker compose --env-file .env.local --profile development -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml up -d
	
stop-docker-dev-volumes:
	@echo "Stopping Docker containers and removing volumes..."
	@docker compose --env-file .env.local --profile development -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml down -v || true

stop-docker-dev:
	@echo "Stopping Docker containers"
	@docker compose --env-file .env.local --profile development -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml down || true

run-docker-prod: build-docker
	@echo "Starting Docker containers for production..."
	@docker compose --env-file .env.production --profile production -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml up -d

stop-docker:
	@echo "Stopping all Docker containers..."
	@docker compose -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml down

reset-db: stop-docker-dev run-docker-dev
	@echo "Database has been reset!"

.PHONY: build run test clean build-docker run-docker-dev stop-docker-dev run-docker-prod stop-docker reset-db
