.PHONY: all build clean test test-cover lint run docker-build docker-run help

# 变量定义
APP_NAME := monitor
VERSION := 1.0.0
MAIN_PKG := ./cmd/server
BUILD_DIR := ./bin
DOCKER_IMAGE := $(APP_NAME):$(VERSION)

# Go 相关变量
GO := go
GOFLAGS := -v
LDFLAGS := -X main.Version=$(VERSION) -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# 默认目标
all: lint test build

# 构建
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)

# 清理
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.txt coverage.html

# 测试
test:
	@echo "Running tests..."
	$(GO) test ./... -v -race

# 测试覆盖率
test-cover:
	@echo "Running tests with coverage..."
	$(GO) test ./... -v -race -coverprofile=coverage.txt -covermode=atomic
	$(GO) tool cover -html=coverage.txt -o coverage.html

# 代码检查
lint:
	@echo "Running linter..."
	golangci-lint run ./...

# 运行
run:
	@echo "Running $(APP_NAME)..."
	$(GO) run $(MAIN_PKG)

# Docker 构建
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) -f build/docker/Dockerfile .

# Docker 运行
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# 生成 wire
wire:
	@echo "Generating wire..."
	cd internal/app && wire

# 生成 mock
mock:
	@echo "Generating mocks..."
	mockgen -source=internal/domain/repository/monitor.go -destination=test/mocks/repository/monitor_mock.go
	mockgen -source=internal/domain/repository/notification.go -destination=test/mocks/repository/notification_mock.go

# 数据库迁移
migrate-up:
	@echo "Running database migrations..."
	migrate -path scripts/migrations -database "mysql://root:root@tcp(localhost:3306)/monitor" up

migrate-down:
	@echo "Rolling back database migrations..."
	migrate -path scripts/migrations -database "mysql://root:root@tcp(localhost:3306)/monitor" down

# 帮助
help:
	@echo "Make targets:"
	@echo "  all          - Run lint, test, and build"
	@echo "  build        - Build the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  test-cover   - Run tests with coverage"
	@echo "  lint         - Run linter"
	@echo "  run          - Run the application"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  wire         - Generate wire"
	@echo "  mock         - Generate mocks"
	@echo "  migrate-up   - Run database migrations"
	@echo "  migrate-down - Rollback database migrations"
	@echo "  help         - Show this help" 