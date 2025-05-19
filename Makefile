APP_NAME = calc
SRC_DIR = ./services/calc/cmd
BUILD_DIR = ./bin
CONFIG_DIR = ./services/calc/config
LIB_DIR = ./lib

.PHONY: all build run test clean

all: build

build:
	@echo "Building the application..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

run: build
	@echo "Running the application..."
	$(BUILD_DIR)/$(APP_NAME) --config $(CONFIG_DIR)/config.yaml

test:
	@echo "Running tests..."
	go test ./lib/...
	go test ./services/calc/internal/...

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)/$(APP_NAME)

deps:
	@echo "Installing dependencies..."
	go mod tidy