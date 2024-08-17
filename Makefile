ROOT_PATH = $(shell pwd)
BIN_PATH  = $(ROOT_PATH)/bin

GO_MODULES_FINDER = find $(ROOT_PATH) -type f -name go.mod \
	-not -path "$(ROOT_PATH)/cmd/*" \
	-not -path "$(ROOT_PATH)/.*"

GO_PACKAGES = $(shell $(GO_MODULES_FINDER) -exec sh -c "cd \$$(dirname {}) && go list ./..." \; | sort | uniq | grep -v "mock" | head -c-1 | tr "\n" " ")


.PHONY: build
build:
	$(info * Building ...)
	env GOOS=linux  CGO_ENABLED=0 IS_DEV=1 go build -o $(BIN_PATH)/service $(ROOT_PATH)/cmd/service

.PHONY: run
run:
	make build
	$(info * Run service ...)
	go run $(ROOT_PATH)/cmd/service -dev

.PHONY: up
up:
	make build
	$(info # Start local infrastructure services)
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: format
format:
	gofumpt -l -w -extra .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: prepare
prepare: format lint build test

.PHONY: test
test:
	env TG_LOGGER_TOKEN=test TG_LOGGER_CHAT_ID=3333 go test ./... -v -tags unit -race $(GO_PACKAGES)