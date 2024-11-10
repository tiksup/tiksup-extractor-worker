package config

import (
	pb "github.com/tiksup/tiksup-extractor-worker/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Conn   *grpc.ClientConn
	Client pb.InteractionsServiceClient
}

func CreateInteractionClient(target string) (*Client, error) {
	conn, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewInteractionsServiceClient(conn)
	return &Client{
		Conn:   conn,
		Client: client,
	}, nil
}

func (c *Client) Close() {
	c.Conn.Close()
}
