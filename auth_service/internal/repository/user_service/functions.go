package userservice

import (
	"context"
	"fmt"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"
	userrpc "github.com/vladimir-kopaliani/tweets/user_service/pkg/protobuf/user_service/v1"
)

func (s UserServiceRPC) GetUserByID(ctx context.Context, userID string) (entities.FullUserInfo, error) {
	in := &userrpc.GetUserByIDRequest{Id: userID}

	response, err := s.client.GetUserByID(ctx, in)
	if err != nil {
		return entities.FullUserInfo{}, fmt.Errorf("calling `GetUserByID` error: %w", err)
	}

	return entities.FullUserInfo{
		User: entities.User{
			ID:        response.User.Id,
			FirstName: response.User.FirstName,
			LastName:  response.User.LastName,
			// Email:
			// CreatedAt:
		},
		// Password:
	}, nil
}

func (s UserServiceRPC) CheckRegisteredUser(ctx context.Context, email, password string) (userID string, err error) {
	in := &userrpc.CheckRegisteredUserRequest{}

	response, err := s.client.CheckRegisteredUser(ctx, in)
	if err != nil {
		return "", fmt.Errorf("calling `CheckRegisteredUser` error: %w", err)
	}

	return response.UserId, nil
}
