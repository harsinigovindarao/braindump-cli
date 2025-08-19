package nlp

import (
	"context"
	//"log"
	"time"

	pb "github.com/harsinigovindarao/braindump-cli/internal/nlp/proto"

	"google.golang.org/grpc"
)

type NLPClient struct {
	client pb.NLPServiceClient
}

func NewNLPClient(address string) (*NLPClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewNLPServiceClient(conn)
	return &NLPClient{client: client}, nil
}

func (c *NLPClient) ClassifyText(text string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TextRequest{Text: text}
	res, err := c.client.ClassifyText(ctx, req)
	if err != nil {
		return "", "", err
	}

	return res.Category, res.Tone, nil
}
