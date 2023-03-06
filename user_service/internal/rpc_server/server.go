package rpcserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
	templatepb "github.com/vladimir-kopaliani/tweets/user_service/pkg/protobuf/user_service/v1"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	logger  interfaces.Logger
	address string
	server  *grpc.Server
	service interfaces.Servicer

	templatepb.UnimplementedUserServiceServer
}

func New(ctx context.Context, conf Configuration) (*GRPCServer, error) {
	server := &GRPCServer{
		service: conf.Service,
		logger:  conf.Logger,
	}

	err := conf.validate()
	if err != nil {
		return nil, fmt.Errorf("validation configuration error: %w", err)
	}

	server.address = conf.Address
	server.server = grpc.NewServer()

	templatepb.RegisterUserServiceServer(server.server, server)

	return server, nil
}

// Launch starts gRPC server
func (gs GRPCServer) Launch(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":"+gs.address)
	if err != nil {
		return fmt.Errorf("failed listening port: %w", err)
	}

	gs.logger.Info(ctx, fmt.Sprintf("gRPC is listening on :%s", gs.address))

	err = gs.server.Serve(lis)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("launching gRPC server error: %w", err)
	}

	return nil
}

// Shutdown graceful closing gRPC Server
func (gs GRPCServer) Shutdown(ctx context.Context) error {
	gs.logger.Info(ctx, "gRPC server is shutting down...")
	gs.server.GracefulStop()
	gs.logger.Info(ctx, "gRPC server is off.")

	return nil
}
