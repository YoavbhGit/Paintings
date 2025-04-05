package main

import (
	"context"
	"log"
	"time"

	pb "inventory-service/proto"

	"google.golang.org/grpc"
)

func main() {
    // Set up a connection to the server with context support.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewInventoryServiceClient(conn)

    // Make a call to the ListPaintings service
    paintings, err := c.ListPaintings(ctx, &pb.Empty{})
    if err != nil {
        log.Fatalf("Could not list paintings: %v", err)
    }
    log.Printf("Paintings: %v", paintings)
}
