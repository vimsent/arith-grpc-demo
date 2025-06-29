# ---------- variables editables ----------
# cámbialas una vez por VM
ROLE            ?= processor        # processor | server | client
# Nombre DNS (o IP:puerto) del processor al que se conectará el server
PROCESSOR_ADDR  ?= dist017.inf.santiago.usm.cl:50052
# Nombre DNS (o IP:puerto) del server al que se conectará el client
SERVER_ADDR     ?= dist018.inf.santiago.usm.cl:50051
# Puerto que expondrá cada contenedor en la VM
PORT_PROCESSOR  ?= 50052
PORT_SERVER     ?= 50051

# ---------- no toques desde aquí ----------
PROTO_DIR := proto
IMAGE     := arith-$(ROLE)
NAME      := $(ROLE)

.PHONY: proto build run deploy stop logs clean

proto:
	protoc --go_out=paths=source_relative:. \
	       --go-grpc_out=paths=source_relative:. \
	       $(PROTO_DIR)/*.proto

build: proto
	go mod tidy
	docker build -f Dockerfile.$(ROLE) -t $(IMAGE) .

run:
	@if [ "$(ROLE)" = "processor" ]; then \
	    docker run -d --name $(NAME) -p $(PORT_PROCESSOR):50052 $(IMAGE) ;\
	elif [ "$(ROLE)" = "server" ]; then \
	    docker run -d --name $(NAME) -e PROCESSOR_ADDR=$(PROCESSOR_ADDR) \
	           -p $(PORT_SERVER):50051 $(IMAGE) ;\
	elif [ "$(ROLE)" = "client" ]; then \
	    echo 'Ejemplo de uso:' ;\
	    echo '  make ROLE=client run ARGS="12 * 4"' ;\
	    exit 1 ;\
	fi

# para client permitimos pasar ARGS="7 * 6"
run-client:
	docker run --rm --name client \
	    -e SERVER_ADDR=$(SERVER_ADDR) \
	    $(IMAGE) $(ARGS)

logs:
	docker logs -f $(NAME)

stop:
	-docker rm -f $(NAME)

deploy: stop build run logs

clean: stop
	docker image rm -f $(IMAGE) || true
