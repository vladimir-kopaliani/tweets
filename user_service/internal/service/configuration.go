package service

import (
	"errors"
	"fmt"

	apperrors "github.com/vladimir-kopaliani/tweets/user_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
)

const (
	serviceConfigurationError = "validate service configuration: %w"
)

type Configuration struct {
	Logger       interfaces.Logger
	PGRepository interfaces.UsersRepositorier
	JWTSecret    string
}

func (conf Configuration) validate() error {
	if conf.Logger == nil {
		return apperrors.ErrLoggerIsNotSet
	}

	if conf.PGRepository == nil {
		return fmt.Errorf(serviceConfigurationError, apperrors.ErrRepositoryIsNotSet)
	}

	if conf.JWTSecret == "" {
		return fmt.Errorf(serviceConfigurationError, errors.New("jwt is empty"))
	}

	return nil
}
