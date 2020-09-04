package zerolog

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// UnaryInterceptor is a gRPC Server Option that uses NewUnaryServerInterceptor() to log gRPC Requests.
func UnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(NewUnaryServerInterceptor())
}

func UnaryInterceptorWithLogger(log *zerolog.Logger) grpc.ServerOption {
	return grpc.UnaryInterceptor(NewUnaryServerInterceptorWithLogger(log))
}

// NewUnaryServerInterceptor that logs gRPC Requests using Zerolog.
//	{
//		ServiceField: "ExampleService",
//		MethodField: "ExampleMethod",
//		DurationField: 1.00
//
//		IpField: "127.0.0.1",
//
//		MetadataField: {},
//
//		UserAgentField: "ExampleClientUserAgent",
//		ReqField: {}, // JSON representation of Request Protobuf
//
//		Err: "An unexpected error occurred",
//		CodeField: "Unknown",
//		MsgField: "Error message returned from the server",
//		DetailsField: [Errors],
//
//		RespField: {}, // JSON representation of Response Protobuf
//
//		ZerologMessageField: "UnaryMessageDefault",
//	}
func NewUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return NewUnaryServerInterceptorWithLogger(&log.Logger)
}

func NewUnaryServerInterceptorWithLogger(log *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		now := time.Now()
		resp, err := handler(ctx, req)
		if log.Error().Enabled() {
			if err != nil {
				logger := log.Error()
				LogIncomingCall(ctx, logger, info.FullMethod, now, req)
				LogStatusError(logger, err)
				logger.Msg(UnaryMessageDefault)
			} else if log.Info().Enabled() {
				logger := log.Info()
				LogIncomingCall(ctx, logger, info.FullMethod, now, req)
				LogResponse(logger, resp)
				logger.Msg(UnaryMessageDefault)
			}
		}
		return resp, err
	}
}
