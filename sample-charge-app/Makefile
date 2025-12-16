PROJECT_NAME := sample-charge-app
COMPOSE_FILE := docker-compose.yml
MIGRATIONS_DIR := api/internal/database/migration
DB_URL := "mysql://sample:sample@tcp(sample-charge-mysql:3306)/sample_charge?multiStatements=true"
MIGRATE_IMAGE := migrate/migrate:latest

.PHONY: up down logs ps migrate-up migrate-down migrate-force

up:
	docker compose -p $(PROJECT_NAME) -f $(COMPOSE_FILE) up -d

down:
	docker compose -p $(PROJECT_NAME) -f $(COMPOSE_FILE) down

logs:
	docker compose -p $(PROJECT_NAME) -f $(COMPOSE_FILE) logs -f

ps:
	docker compose -p $(PROJECT_NAME) -f $(COMPOSE_FILE) ps

migrate-up:
	docker run --rm \
		--network $(PROJECT_NAME)_default \
		-v $(PWD)/$(MIGRATIONS_DIR):/migrations \
		$(MIGRATE_IMAGE) \
		-path=/migrations -database $(DB_URL) up

migrate-down:
	docker run --rm \
		--network $(PROJECT_NAME)_default \
		-v $(PWD)/$(MIGRATIONS_DIR):/migrations \
		$(MIGRATE_IMAGE) \
		-path=/migrations -database $(DB_URL) down 1

migrate-force:
	@read -p "force version? " ver; \
	docker run --rm \
		--network $(PROJECT_NAME)_default \
		-v $(PWD)/$(MIGRATIONS_DIR):/migrations \
		$(MIGRATE_IMAGE) \
		-path=/migrations -database $(DB_URL) force $$ver

.PHONY: bob-gen
bob-gen:
	cd api && go run github.com/stephenafamo/bob/gen/bobgen-mysql@latest --config bobgen.yaml
SHELL := /bin/sh

APP_NAME := servicechargeservice
PKG := ./...
ENTRY := ./cmd/$(APP_NAME)

.PHONY: gen
gen: proto swagger

.PHONY: proto
proto:
	cd schema && buf generate

.PHONY: swagger
swagger:
	cd /Users/takenariyamamoto/test && oapi-codegen -config schema/swagger/.oapi-codegen.yaml schema/swagger/swagger.yaml

.PHONY: mock
mock:
	cd pkg/pb/account/v1 && go generate
	cd pkg/pb/subscription/v1 && go generate
	cd pkg/pb/servicecharge/v1 && go generate

.PHONY: tidy
tidy:
	cd api && go mod tidy
	cd job && go mod tidy
	cd subscriber && go mod tidy
	cd pkg/pb && go mod tidy

.PHONY: build
build:
	go build -o bin/$(APP_NAME) $(ENTRY)

.PHONY: run
run:
	go run $(ENTRY)

.PHONY: test
test:
	go test -v $(PKG)

.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run:
	docker run --rm -p $${PORT:-8080}:$${PORT:-8080} --env-file .env --name $(APP_NAME) $(APP_NAME):latest

