VERSION ?= 0.0.1

REGISTRY ?= 172.22.11.2:30500

SERVER_NAME = mail-sender-server
SERVER_IMG = $(REGISTRY)/$(SERVER_NAME):$(VERSION)

CLIENT_NAME = mail-sender-client
CLIENT_IMG = $(REGISTRY)/$(CLIENT_NAME):$(VERSION)


all: build

build: build-server build-client

build-server:
	go build -o ./bin/client ./cmd/client

build-client:
	go build -o ./bin/server ./cmd/server


image: image-server image-client

image-server:
	docker build -f build/server/Dockerfile -t $(SERVER_IMG) .

image-client:
	docker build -f build/client/Dockerfile -t $(CLIENT_IMG) .


push: push-server push-client

push-server:
	docker push $(SERVER_IMG)

push-client:
	docker push $(CLIENT_IMG)
