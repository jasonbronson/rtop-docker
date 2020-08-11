PACKAGE   := rtop
SERVICE_PKG_DIR := ./cmd/api

setup:
	go mod download

build: ## Build go binary
	go build -o bin/${PACKAGE}-service ${SERVICE_PKG_DIR}

local: 
	docker-compose up

	

