package web
import (
	"context"
	"log"
	"time"

	pb "github.com/harsinigovindarao/braindump-cli/internal/nlp/proto"
	"google.golang.org/grpc"
	"github.com/harsinigovindarao/braindump-cli/utils"
	"github.com/harsinigovindarao/braindump-cli/internal/models"
	"github.com/harsinigovindarao/braindump-cli/internal/storage"
)

// Channels
var ThoughtInputChan = make(chan models.Thought)
var ProcessedThoughtChan = make(chan models.Thought)

func StartWorker() {
	go func() {
		// Connect to gRPC server
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		defer conn.Close()
		client := pb.NewNLPServiceClient(conn)

		for t := range ThoughtInputChan {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

    resp, err := client.ClassifyText(ctx, &pb.TextRequest{Text: t.Text})
    if err != nil {
        log.Printf("gRPC error: %v", err)
        t.Category = "unknown"
        t.Tone = "neutral"
    } else {
        t.Category = resp.Category
        t.Tone = resp.Tone
    }

    // Always call cancel() here, not defer
    cancel()

    t.Priority = utils.ScorePriority(t, storage.LoadThoughts())
    ProcessedThoughtChan <- t
}

	}()

	// Worker for saving processed thoughts
	go func() {
		for pt := range ProcessedThoughtChan {
			storage.SaveThought(pt)
		}
	}()
}
