package httpserver

import (
	"context"
	"net/http"

	apperrors "github.com/vladimir-kopaliani/tweets/auth_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
)

type Handler struct {
	Path    string
	Handler http.Handler
}

type Configuration struct {
	Logger   interfaces.Logger
	Address  string
	Handlers []Handler
	Service  interfaces.Servicer
}

func (conf Configuration) validate() error {
	if conf.Logger == nil {
		return apperrors.ErrLoggerIsNotSet
	}

	if conf.Address == "" {
		return apperrors.ErrAddressIsNotSet
	}

	if conf.Service == nil {
		return apperrors.ErrServiceIsNotSet
	}

	return nil
}

type Servicer interface {
	GetUserByID(ctx context.Context, id string) (*entities.User, error)
}
