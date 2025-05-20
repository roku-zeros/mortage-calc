APP_NAME = calc
SRC_DIR = ./services/calc/cmd
BUILD_DIR = ./bin
CONFIG_DIR =./services/calc/config
LIB_DIR = ./lib
HOST_PORT ?= 8080

IMAGE_NAME = calc-image
CONTAINER_NAME = calc-container


.PHONY: all build run test lint clean docker-build docker-run docker-stop docker-rm

all: build

test:
	@echo "Running tests..."
	go test ./lib/...
	go test ./services/calc/internal/...

lint:
	@echo "Running linter..."
	golangci-lint run


clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)/$(APP_NAME)

deps:
	@echo "Installing dependencies..."
	go mod tidy

docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

docker-run: docker-build
	@echo "Running Docker container..."
	docker run --name $(CONTAINER_NAME) -v $(CURDIR)/$(CONFIG_DIR)/config.yaml:/config.yaml -p $(HOST_PORT):8080 $(IMAGE_NAME)

docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(CONTAINER_NAME)

docker-rm:
	@echo "Removing Docker container..."
	docker rm $(CONTAINER_NAME)
