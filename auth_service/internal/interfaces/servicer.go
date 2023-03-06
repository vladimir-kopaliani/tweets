package interfaces

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"
)

type Servicer interface {
	SignIn(ctx context.Context, input *entities.LoginInput) (*entities.SignInResult, error)
	RefreshTokens(ctx context.Context, input *entities.RefreshTokensInput) (*entities.RefreshTokensResult, error)
}
