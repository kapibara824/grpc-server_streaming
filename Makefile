

.PHONY: all


all:
	make build-server
	make build-client

build-server:
	go build -o ./bin/server ./server/main.go

build-client:
	go build -o ./bin/client ./client/main.go

protoc:
	protoc --go_out=plugins=grpc:./pb ./proto/*.proto