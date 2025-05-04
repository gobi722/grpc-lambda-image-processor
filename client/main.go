package main

import (
    "context"
    "log"
    "time"

    pb "yourmodule/proto"

    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewImageServiceClient(conn)

    req := &pb.ImageRequest{
        ImageUrl: "https://example.com/image.jpg",
        Action:   "resize",
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    res, err := client.ProcessImage(ctx, req)
    if err != nil {
        log.Fatalf("Error calling ProcessImage: %v", err)
    }

    log.Printf("Response: %s (URL: %s)", res.Message, res.ProcessedImageUrl)
}
