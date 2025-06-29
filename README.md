# arith-grpc-demo

Pequeño ejemplo en Go + gRPC + Protocol Buffers con tres servicios contenedorizados.

## Estructura

* **CLIENT** → envía la operación al SERVER.
* **SERVER** → reenvía al PROCESSOR y devuelve el resultado.
* **PROCESSOR** → calcula la operación.

## Requisitos

* Go ≥ 1.22  
* `protoc` ≥ 3.21 con los plugins `protoc-gen-go` y `protoc-gen-go-grpc`  
* Docker + Docker Compose v2

## Pasos rápidos

```bash
git clone https://github.com/your-github-user/arith-grpc-demo.git
cd arith-grpc-demo

# (opcional) Regenerar código protobuf
make proto

# Construir imágenes y lanzar todo
make docker-up

# Lanzar otro cliente de prueba
docker compose run --rm client 100 / 5
