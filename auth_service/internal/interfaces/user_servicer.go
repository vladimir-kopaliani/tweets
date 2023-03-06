package interfaces

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"
)

type UserServicer interface {
	GetUserByID(ctx context.Context, userID string) (entities.FullUserInfo, error)
	CheckRegisteredUser(ctx context.Context, email, password string) (userID string, err error)

	Close(ctx context.Context) error
}
