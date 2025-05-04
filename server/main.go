package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "yourmodule/proto"

    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedImageServiceServer
}

func (s *server) ProcessImage(ctx context.Context, req *pb.ImageRequest) (*pb.ImageResponse, error) {
    log.Printf("Received request: action=%s, url=%s", req.Action, req.ImageUrl)

    // Here you'd call AWS Lambda (mocked)
    processedURL := "https://s3.amazonaws.com/yourbucket/processed-image.jpg"

    return &pb.ImageResponse{
        ProcessedImageUrl: processedURL,
        Message:           fmt.Sprintf("Image processed with action: %s", req.Action),
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterImageServiceServer(grpcServer, &server{})

    log.Println("gRPC server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
