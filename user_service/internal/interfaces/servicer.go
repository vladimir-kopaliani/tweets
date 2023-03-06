package interfaces

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"
)

type Servicer interface {
	RegisterUser(ctx context.Context, user *entities.FullUserInfo) (*entities.User, error)
	GetUserByID(ctx context.Context, id string) (*entities.User, error)
}
