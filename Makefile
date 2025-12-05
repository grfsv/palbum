.PHONY: help build run test clean docker-up docker-down migrate-up migrate-create lint fmt

# デフォルトターゲット
help:
	@echo "利用可能なコマンド:"
	@echo "  make build         - アプリケーションをビルド"
	@echo "  make run           - アプリケーションを起動"
	@echo "  make test          - テストを実行"
	@echo "  make clean         - ビルド成果物を削除"
	@echo "  make docker-up     - Docker環境を起動"
	@echo "  make docker-down   - Docker環境を停止"
	@echo "  make migrate-create  - マイグレーション用SQLの作成"
	@echo "  make migrate-up    - データベースマイグレーション実行"
	@echo "  make migrate-reset - マイグレーションのリセット"
	@echo "  make lint          - コードの静的解析"
	@echo "  make fmt           - コードのフォーマット"

# ビルド
build:
	@echo "Building application..."
	go build -o bin/main cmd/main.go

# 実行
run:
	@echo "Running application..."
	go run cmd/main.go

# テスト
test:
	@echo "Running tests..."
	go test -v ./...

# テストカバレッジ
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# クリーン
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf tmp/
	rm -f coverage.out coverage.html

# Docker環境起動
docker-up:
	@echo "Starting Docker environment..."
	docker-compose up -d

# Docker環境停止
docker-down:
	@echo "Stopping Docker environment..."
	docker-compose down

# Docker環境再構築
docker-rebuild:
	@echo "Rebuilding Docker environment..."
	docker-compose down
	docker-compose up --build -d

# Docker ログ表示
docker-logs:
	docker-compose logs -f

# マイグレーション実行
migrate-up:
	@echo "Running migrations..."
	atlas migrate apply --url "mysql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)"

# マイグレーション作成
migrate-create:
	@echo "Creating migrations..."
	atlas migrate diff --config file://atlas.hcl --env gorm

migrate-reset:
	@echo "Resetting migrations..."
	rm -rf migrations/*.sql
	atlas migrate hash
	make migrate-create

lint:
	@echo "Running linter..."
	golangci-lint run --fix

# コードフォーマット
fmt:
	@echo "Formatting code..."
	golangci-lint fmt

# 環境変数の読み込み
include .env
export
