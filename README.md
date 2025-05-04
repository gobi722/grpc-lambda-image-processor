# grpc-lambda-image-processor

A **Go-based** microservice project that combines **gRPC**, **AWS Lambda**, and **S3** for real-time, scalable image processing.

---

##  Overview

This project demonstrates:
 Go gRPC server and client communication  
 AWS Lambda for serverless image processing (resize, convert, etc.)  
 S3 for storing original and processed images

Clients send a request (e.g., resize an image) to the gRPC server.  
The server triggers a Lambda function to process the image, upload the result to S3, and return the processed image’s URL.

---

##  Project Structure
rpc-lambda-image-processor/
├── proto/
│ └── image_service.proto # gRPC service definition (protobuf)
├── server/
│ ├── main.go # Go gRPC server
│ └── go.mod / go.sum # Go module files
├── client/
│ ├── main.go # Go gRPC client 
│ └── go.mod / go.sum
├── lambda/
│ ├── main.go # AWS Lambda function (Go) for image processing
│ └── go.mod / go.sum
├── README.md # This README file

##  License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

