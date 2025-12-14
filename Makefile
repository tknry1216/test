SHELL := /bin/sh

APP_NAME := servicechargeservice
PKG := ./...
ENTRY := ./cmd/$(APP_NAME)

.PHONY: proto
proto:
	cd schema && buf generate

.PHONY: mock
mock:
	cd pkg/pb/account/v1 && go generate

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

