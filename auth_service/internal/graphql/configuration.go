package graphql

import (
	"fmt"

	apperrors "github.com/vladimir-kopaliani/tweets/auth_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
)

type Configuration struct {
	Service          interfaces.Servicer
	JWTSecret        string
	Logger           interfaces.Logger
	IsProductionMode bool
}

func (conf Configuration) validate() error {
	if conf.Logger == nil {
		return apperrors.ErrLoggerIsNotSet
	}

	if conf.Service == nil {
		return apperrors.ErrServiceIsNotSet
	}

	if conf.JWTSecret == "" {
		return fmt.Errorf("validate GraphQL server configuration: %w", apperrors.ErrJWTEmpty)
	}

	return nil
}
