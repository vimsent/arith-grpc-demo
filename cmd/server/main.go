package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/vimsent/arith-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedOperationServiceServer
	processor pb.ProcessorServiceClient
}

func newProcessorClient(addr string) pb.ProcessorServiceClient {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("no se pudo conectar al PROCESSOR (%s): %v", addr, err)
	}
	return pb.NewProcessorServiceClient(conn)
}

func (s *server) Compute(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	return s.processor.Process(ctx, req)
}

func main() {
	processorAddr := os.Getenv("PROCESSOR_ADDR")
	if processorAddr == "" {
		processorAddr = "processor:50052"
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error al abrir puerto 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOperationServiceServer(s, &server{
		processor: newProcessorClient(processorAddr),
	})
	log.Println("SERVER escuchando en :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Servidor gRPC detenido: %v", err)
	}
}
