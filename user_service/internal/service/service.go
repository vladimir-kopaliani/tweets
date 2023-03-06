package service

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
)

type Service struct {
	pgRepository interfaces.UsersRepositorier
	jwtSecret    []byte
}

func New(ctx context.Context, cfg Configuration) (interfaces.Servicer, error) {
	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	serv := Service{
		pgRepository: cfg.PGRepository,
		jwtSecret:    []byte(cfg.JWTSecret),
	}

	return &serv, nil
}
