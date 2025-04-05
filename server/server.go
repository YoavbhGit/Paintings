package main

import (
	"context"
	"log"
	"net"

	pb "inventory-service/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedInventoryServiceServer
}

func (s *server) AddPainting(ctx context.Context, in *pb.Painting) (*pb.Painting, error) {
    // Implement your logic here
    return in, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterInventoryServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
