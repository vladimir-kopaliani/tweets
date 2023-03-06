package userservice

import (
	apperrors "github.com/vladimir-kopaliani/tweets/auth_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
)

type Configuration struct {
	Address string
	Logger  interfaces.Logger
}

func (conf Configuration) validate() error {
	if conf.Logger == nil {
		return apperrors.ErrLoggerIsNotSet
	}

	if conf.Address == "" {
		return apperrors.ErrAddressIsNotSet
	}

	return nil
}
