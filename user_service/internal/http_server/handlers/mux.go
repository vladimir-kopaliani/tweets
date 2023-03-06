package handlers

import (
	"errors"
	"fmt"
	"net/http"

	apperrors "github.com/vladimir-kopaliani/tweets/user_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
)

type RestServer struct {
	isProductionMode bool
	service          interfaces.Servicer
	mux              *http.ServeMux
}

type Configuration struct {
	IsProductionMode bool
	Service          interfaces.Servicer
	JWTSecret        string
}

func New(conf Configuration) (http.Handler, error) {
	err := conf.validate()
	if err != nil {
		return nil, fmt.Errorf("validate REST server configuration error: %w", err)
	}

	rest := &RestServer{
		service: conf.Service,
		mux:     http.NewServeMux(),
	}

	rest.mux.HandleFunc("users", rest.handleError(rest.GetUserByID))

	return rest.mux, nil
}

func (conf Configuration) validate() error {
	if conf.Service == nil {
		return apperrors.ErrServiceIsNotSet
	}

	if conf.JWTSecret == "" {
		return fmt.Errorf("validate HTTP mux configuration: %w", apperrors.ErrJWTEmpty)
	}

	return nil
}

func (s RestServer) handleError(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			// show only public errors
			if s.isProductionMode {
				if !apperrors.IsPublicError(err) {
					err = apperrors.ErrInternal
				}
			}

			statusCode := http.StatusInternalServerError

			switch {
			case errors.Is(err, apperrors.ErrWrongCredentials):
				statusCode = http.StatusUnauthorized
			}

			w.WriteHeader(statusCode)
		}
	})
}
