.PHONY: proto build compose

proto:
	protoc --proto_path=. --go_out=./proto_files --go_opt=paths=source_relative --go-grpc_out=./proto_files --go-grpc_opt=paths=source_relative manifold.proto

build:
	GOOS=linux go build -o ./dist/server ./examples/server.go

compose: build
	docker-compose up --build --scale api=2
