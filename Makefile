VERSION ?= v0.0.1

REGISTRY ?= tmaxcloudck

SERVER_NAME = mail-sender-server
SERVER_IMG = $(REGISTRY)/$(SERVER_NAME):$(VERSION)

CLIENT_NAME = mail-sender-client
CLIENT_IMG = $(REGISTRY)/$(CLIENT_NAME):$(VERSION)


all: build

build: build-server build-client

build-server:
	CGO_ENABLED=0 go build -o ./bin/server ./cmd/server

build-client:
	CGO_ENABLED=0 go build -o ./bin/client ./cmd/client


.PHONY: image image-server image-client
image: image-server image-client

image-server:
	docker build -f build/server/Dockerfile -t $(SERVER_IMG) .

image-client:
	docker build -f build/client/Dockerfile -t $(CLIENT_IMG) .


.PHONY: push push-server push-client
push: push-server push-client

push-server:
	docker push $(SERVER_IMG)

push-client:
	docker push $(CLIENT_IMG)


.PHONY: test test-verify save-sha-mod compare-sha-mod verify test-unit test-lint
test: test-verify test-unit test-lint

test-verify: save-sha-mod verify compare-sha-mod

save-sha-mod:
	$(eval MODSHA=$(shell sha512sum go.mod))
	$(eval SUMSHA=$(shell sha512sum go.sum))

verify:
	go mod verify

compare-sha-mod:
	$(eval MODSHA_AFTER=$(shell sha512sum go.mod))
	$(eval SUMSHA_AFTER=$(shell sha512sum go.sum))
	@if [ "${MODSHA_AFTER}" = "${MODSHA}" ]; then echo "go.mod is not changed"; else echo "go.mod file is changed"; exit 1; fi
	@if [ "${SUMSHA_AFTER}" = "${SUMSHA}" ]; then echo "go.sum is not changed"; else echo "go.sum file is changed"; exit 1; fi

test-unit:
	go test -v ./pkg/...

test-lint:
	golangci-lint run ./... -v -E gofmt --timeout 1h0m0s
