package service

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/tweet_service/internal/interfaces"
)

type Service struct {
	tweetsRepository interfaces.TweetsRepositorier
	jwtSecret        []byte
}

func New(ctx context.Context, cfg Configuration) (interfaces.Servicer, error) {
	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	serv := Service{
		tweetsRepository: cfg.TweetsRepository,
		jwtSecret:        []byte(cfg.JWTSecret),
	}

	return &serv, nil
}
