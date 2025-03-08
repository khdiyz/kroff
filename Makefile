GOOSE_BIN=goose
ENV_FILE=.env
MIGRATIONS_DIR=./schema

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

DB_DRIVER=postgres
DB_DSN=postgres://$(POSTGRE_USER):$(POSTGRE_PASSWORD)@$(POSTGRE_HOST):$(POSTGRE_PORT)/$(POSTGRE_DB)?sslmode=disable

.PHONY: up down create-migration status help

create-migration:
	@read -p "Enter migration name: " name && \
	mkdir -p $(MIGRATIONS_DIR) && \
	$(GOOSE_BIN) -s create "$$name" sql --dir $(MIGRATIONS_DIR)

up:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" up

down:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" down

redo:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" redo

status:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" status


run:
	@go run cmd/main.go

tidy:
	@go mod tidy

swag:
	@swag init -g cmd/main.go