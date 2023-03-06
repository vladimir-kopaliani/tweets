package pgrepo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PGRepo struct {
	Logger interfaces.Logger
	db     *bun.DB
}

func NewPGRepo(ctx context.Context, conf Configuration) (interfaces.UsersRepositorier, error) {
	var err error

	repo := &PGRepo{
		Logger: conf.Logger,
	}

	err = conf.validate()
	if err != nil {
		return nil, fmt.Errorf("validation configuration error: %w", err)
	}

	repo.db = bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.Address))),
		pgdialect.New(),
	)

	if err = repo.db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	if _, err := repo.db.NewCreateTable().
		Model((*entities.FullUserInfo)(nil)).
		Exec(ctx); err != nil {
		// ignore error: ERROR: relation "..." already exists (SQLSTATE=42P07)
		if strings.HasPrefix(err.Error(), "ERROR: relation \"") &&
			strings.HasSuffix(err.Error(), "\" already exists (SQLSTATE=42P07)") {
			err = nil
		}

		if err != nil {
			return nil, fmt.Errorf("creating relation error: %w", err)
		}
	}

	return repo, nil
}

func (repo PGRepo) Close(ctx context.Context) error {
	return repo.db.Close()
}
