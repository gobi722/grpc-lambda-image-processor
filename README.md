# grpc-lambda-image-processor
A Go microservice using gRPC for client-server communication and AWS Lambda for serverless image processing. Users send requests (e.g., resize, convert) via gRPC; Lambda processes the image, stores it in S3, and returns the processed image’s public URL for easy access and retrieval.
