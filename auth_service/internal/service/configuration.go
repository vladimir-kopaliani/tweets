package service

import (
	"fmt"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"

	apperrors "github.com/vladimir-kopaliani/tweets/auth_service/internal/app_errors"
)

const (
	serviceConfigurationError = "validate service configuration: %w"
)

type Configuration struct {
	Logger      interfaces.Logger
	UserService interfaces.UserServicer
	JWTSecret   string
}

func (conf Configuration) validate() error {
	if conf.Logger == nil {
		return apperrors.ErrLoggerIsNotSet
	}

	if conf.UserService == nil {
		return apperrors.ErrRepositoryIsNotSet
	}

	if conf.JWTSecret == "" {
		return fmt.Errorf(serviceConfigurationError, apperrors.ErrJWTEmpty)
	}

	return nil
}
