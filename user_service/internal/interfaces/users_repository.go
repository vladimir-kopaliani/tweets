package interfaces

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"
)

// TODO: merge `GetUserByID()` and `GetFullUserInfoByEmail()` together and user dataloader

type UsersRepositorier interface {
	SaveNewUser(ctx context.Context, user *entities.FullUserInfo) error
	GetUserByID(ctx context.Context, id string) (*entities.User, error)
	GetFullUserInfoByEmail(ctx context.Context, email string) (*entities.FullUserInfo, error)

	Close(ctx context.Context) error
}
