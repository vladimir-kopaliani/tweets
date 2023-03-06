package rpcserver

import (
	"context"
	"fmt"

	templatepb "github.com/vladimir-kopaliani/tweets/user_service/pkg/protobuf/user_service/v1"
)

func (gs *GRPCServer) GetUserByID(ctx context.Context, id *templatepb.GetUserByIDRequest) (*templatepb.GetUserByIDResponse, error) {
	user, err := gs.service.GetUserByID(ctx, id.GetId())
	if err != nil {
		return nil, fmt.Errorf("getting user error: %w", err)
	}

	return &templatepb.GetUserByIDResponse{
		User: &templatepb.User{
			Id: user.ID,
		},
	}, nil
}
