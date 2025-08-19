package nlp

import (
	"context"
	"log"
	"net"

	pb "github.com/harsinigovindarao/braindump-cli/internal/nlp/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNLPServiceServer
}

// This is the main logic to handle the gRPC request
func (s *server) ClassifyText(ctx context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
	log.Printf("Received text: %s", req.Text)

	// Placeholder logic: Replace this with actual call to Python
	category := "test-category"
	tone := "test-tone"

	// Return dummy response
	return &pb.TextResponse{
		Category: category,
		Tone:     tone,
	}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNLPServiceServer(s, &server{})

	log.Println("🚀 Go gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
