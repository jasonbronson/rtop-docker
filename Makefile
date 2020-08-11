PACKAGE   := rtop
SERVICE_PKG_DIR := ./cmd/api
MIGRATION_DIR := ./migrations
MIGRATION_PKG := migrate
MIGRATIONS_PKG_DIR := ./cmd/migrations

setup:
	go mod download

build: ## Build go binary
	go build -o bin/${PACKAGE}-service ${SERVICE_PKG_DIR} && go build -tags netgo -a -v -o bin/${MIGRATION_PKG} ${MIGRATIONS_PKG_DIR}

local: 
	docker-compose up

	

