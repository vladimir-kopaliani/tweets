package service

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
)

type Service struct {
	jwtSecret   []byte
	userService interfaces.UserServicer
}

func New(ctx context.Context, cfg Configuration) (interfaces.Servicer, error) {
	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	serv := Service{
		jwtSecret:   []byte(cfg.JWTSecret),
		userService: cfg.UserService,
	}

	return &serv, nil
}
