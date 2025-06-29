# ---------------- VARIABLES ----------------
PROTO_DIR := proto
COMPOSE   := docker compose

# ---------------- TARGETS ------------------
.PHONY: proto docker-build docker-up clean \
        run-processor run-server run-client

# Genera *.pb.go en proto/ respetando la ruta fuente
proto:
	protoc --go_out=paths=source_relative:. \
	       --go-grpc_out=paths=source_relative:. \
	       $(PROTO_DIR)/*.proto

# Ejecutables locales (útil para depurar sin Docker)
run-processor:
	go run ./cmd/processor

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client 12 "*" 4

# Construye imágenes SIEMPRE desde cero
docker-build: proto
	go mod tidy
	docker builder prune -f
	$(COMPOSE) build --no-cache

# Levanta el stack; depende de docker-build
docker-up: docker-build
	$(COMPOSE) up --force-recreate

# Limpieza
clean:
	$(COMPOSE) down --remove-orphans
