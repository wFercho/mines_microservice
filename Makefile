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

run-docker-dev:
	@docker compose --env-file .env.local --profile development -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml up -d

run-docker-prod:
	@docker compose --env-file .env.production --profile production -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml up -d

stop-docker:
	@docker compose -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml down

.PHONY: build run test clean

