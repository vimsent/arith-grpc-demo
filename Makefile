# Nombre del directorio con .proto
PROTO_DIR = proto

# Alias para evitar repetir "docker compose …"
COMPOSE   = docker compose

# ---------- gRPC / Protobuf ----------
proto:
	protoc --go_out=. --go-grpc_out=. $(PROTO_DIR)/*.proto

# ---------- Ejecutables locales ----------
run-processor:
	go run ./cmd/processor

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client 7 "*" 6

# ---------- Docker ----------
# 1) Borra la caché global de BuildKit
# 2) Compila todas las imágenes sin usar cache
docker-build:
	docker builder prune -f
	$(COMPOSE) build --no-cache

# 3) Levanta los servicios (siempre recrea)
docker-up: docker-build
	$(COMPOSE) up --force-recreate

# ---------- Limpieza ----------
clean:
	$(COMPOSE) down --remove-orphans
