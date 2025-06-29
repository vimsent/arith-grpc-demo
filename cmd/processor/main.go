package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vimsent/arith-grpc-demo/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProcessorServiceServer
}

func (*server) Process(_ context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	var res float64
	switch req.Operator {
	case "+":
		res = req.A + req.B
	case "-":
		res = req.A - req.B
	case "*":
		res = req.A * req.B
	case "/":
		if req.B == 0 {
			return &pb.OperationResponse{Error: "division por cero"}, nil
		}
		res = req.A / req.B
	default:
		return &pb.OperationResponse{Error: "operador no soportado"}, nil
	}
	return &pb.OperationResponse{Result: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("error al abrir puerto 50052: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProcessorServiceServer(s, &server{})
	log.Println("PROCESSOR escuchando en :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Servidor gRPC detenido: %v", err)
	}
}
