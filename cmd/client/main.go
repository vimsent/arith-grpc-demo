package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/vimsent/arith-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func usage() {
	fmt.Println("Uso: client <A> <operador> <B>")
	fmt.Println(`Ejemplo: client 12 "*" 4`)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 4 {
		usage()
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("A no es número: %v", err)
	}
	op := os.Args[2]
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("B no es número: %v", err)
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "server:50051"
	}
	conn, err := grpc.Dial(serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("no se pudo conectar al SERVER (%s): %v", serverAddr, err)
	}
	defer conn.Close()

	client := pb.NewOperationServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Compute(ctx, &pb.OperationRequest{
		A:        a,
		B:        b,
		Operator: op,
	})
	if err != nil {
		log.Fatalf("Compute falló: %v", err)
	}
	if resp.Error != "" {
		log.Fatalf("Error devuelto por el server: %s", resp.Error)
	}
	fmt.Printf("Resultado: %.4f\n", resp.Result)
}
