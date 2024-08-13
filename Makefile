ROOT_PATH = $(shell pwd)
BIN_PATH  = $(ROOT_PATH)/bin

.PHONY: build
build:
	$(info * Building ...)
	env GOOS=linux  CGO_ENABLED=0  go build -o $(BIN_PATH)/service $(ROOT_PATH)/cmd/service

.PHONY: run
run:
	make build
	$(info * Run service ...)
	$(BIN_PATH)/service

.PHONY: up
up:
	make build
	$(info # Start local infrastructure services)
	docker-compose up -d

.PHONY: down
down:
	docker stop outfitbot-app-1

.PHONY: format
format:
	gofumpt -l -w -extra .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: prepare
prepare: format lint build

.PHONY: test
test:
	go test ./... -v -tags unit -race