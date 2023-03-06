package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"

	"github.com/uptrace/bun"
)

func (repo *PGRepo) SaveNewUser(ctx context.Context, user *entities.FullUserInfo) error {
	_, err := repo.db.NewInsert().
		Model(user).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("insert new user into db error: %w", err)
	}

	return nil
}

func (repo *PGRepo) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	user := new(entities.User)

	err := repo.db.NewSelect().
		Model(user).
		Where("? = ?", bun.Ident("id"), id).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get user by id from db error: %w", err)
	}

	return user, nil
}

func (repo *PGRepo) GetFullUserInfoByEmail(ctx context.Context, email string) (*entities.FullUserInfo, error) {
	user := new(entities.FullUserInfo)

	err := repo.db.NewSelect().
		Model(user).
		Where("? = ?", bun.Ident("email"), email).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get user by email from db error: %w", err)
	}

	return user, nil
}
