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

.PHONY: build run test clean

