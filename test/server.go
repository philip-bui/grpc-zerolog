package test

import (
	"context"
	"net"
	"sync"

	pb "github.com/philip-bui/grpc-zerolog/protos"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	address = "localhost:7070"
)

var (
	server     *grpc.Server
	serverSync sync.Once
)

type TestServer struct{}

func (s *TestServer) TestUnary(ctx context.Context, t *pb.TestMessage) (*pb.TestMessage, error) {
	if t.Test == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty message")
	}
	return t, nil
}

//func init() {
//	log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
//}

func StartServer(interceptor grpc.UnaryServerInterceptor) {
	serverSync.Do(func() {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatal().Err(err).Msg("start server")
		} else {
			log.Info().Msg("start server")
		}
		server = grpc.NewServer(
			grpc.UnaryInterceptor(interceptor),
		)
		pb.RegisterTestServiceServer(server, &TestServer{})
		if err := server.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("start server")
		}
	})
}
