package pgrepo

import (
	apperrors "github.com/vladimir-kopaliani/tweets/user_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
)

type Configuration struct {
	Logger  interfaces.Logger
	Address string
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
