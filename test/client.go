package test

import (
	"context"
	"sync"

	pb "github.com/philip-bui/grpc-zerolog/protos"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type TestClient struct {
	ExampleReq        *pb.TestMessage
	ExampleInvalidReq *pb.TestMessage
	pb.TestServiceClient
}

func (c *TestClient) SendReq() (*pb.TestMessage, error) {
	return c.TestUnary(context.Background(), c.ExampleReq)
}

func (c *TestClient) SendErr() (*pb.TestMessage, error) {
	return c.TestUnary(context.Background(), c.ExampleInvalidReq)
}

var (
	client     *TestClient
	clientSync sync.Once
)

func GetClient() *TestClient {
	clientSync.Do(func() {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatal().Err(err).Msg("start client")
		}
		client = &TestClient{
			&pb.TestMessage{
				Test: "Hi",
			},
			&pb.TestMessage{
				Test: "",
			},
			pb.NewTestServiceClient(conn),
		}
	})
	return client
}
