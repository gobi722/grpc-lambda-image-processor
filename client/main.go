package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "yourmodule/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file " + err.Error())
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewImageServiceClient(conn)

	req := &pb.ImageRequest{
		ImageUrl: os.Getenv("IMAGE_URL"),
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
