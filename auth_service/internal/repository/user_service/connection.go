package userservice

import (
	"context"
	"fmt"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
	userrpc "github.com/vladimir-kopaliani/tweets/user_service/pkg/protobuf/user_service/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceRPC struct {
	Logger interfaces.Logger
	client userrpc.UserServiceClient
	conn   *grpc.ClientConn
}

func New(ctx context.Context, conf Configuration) (interfaces.UserServicer, error) {
	err := conf.validate()
	if err != nil {
		return nil, fmt.Errorf("validation configuration error: %w", err)
	}

	service := &UserServiceRPC{}

	service.conn, err = grpc.DialContext(
		ctx,
		conf.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("dial gRPC error: %w", err)
	}

	service.client = userrpc.NewUserServiceClient(service.conn)

	return service, nil
}

func (s UserServiceRPC) Close(ctx context.Context) error {
	return s.conn.Close()
}
