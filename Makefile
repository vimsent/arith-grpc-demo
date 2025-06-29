PROTO_DIR = proto

proto:
	protoc --go_out=. --go-grpc_out=. $(PROTO_DIR)/*.proto

run-processor:
	go run ./cmd/processor

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client 7 "*" 6

docker-up:
	docker compose up --build

clean:
	docker compose down --remove-orphans
