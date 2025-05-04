package main

import (
    "context"
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
    ImageURL string `json:"image_url"`
    Action   string `json:"action"`
}

type Response struct {
    ProcessedImageURL string `json:"processed_image_url"`
}

func handler(ctx context.Context, event Event) (Response, error) {
    fmt.Printf("Processing image: %s with action: %s\n", event.ImageURL, event.Action)

    // Here you'd process the image and upload to S3 (mocked)
    processedURL := "https://s3.amazonaws.com/yourbucket/processed-image.jpg"

    return Response{ProcessedImageURL: processedURL}, nil
}

func main() {
    lambda.Start(handler)
}
