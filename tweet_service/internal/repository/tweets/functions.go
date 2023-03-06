package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/vladimir-kopaliani/tweets/tweet_service/internal/entities"

	"github.com/uptrace/bun"
)

func (repo *PGRepo) SaveNewTweet(ctx context.Context, tweet *entities.Tweet) error {
	_, err := repo.db.NewInsert().
		Model(tweet).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("insert new tweet into db error: %w", err)
	}

	return nil
}

func (repo *PGRepo) GetTweetByID(ctx context.Context, id string) (*entities.Tweet, error) {
	tweet := new(entities.Tweet)

	err := repo.db.NewSelect().
		Model(tweet).
		Where("? = ?", bun.Ident("id"), id).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get tweet by id from db error: %w", err)
	}

	return tweet, nil
}

func (repo *PGRepo) GetTweetsByUserID(ctx context.Context, userID string) ([]*entities.Tweet, error) {
	var tweets ([]*entities.Tweet)

	err := repo.db.NewSelect().
		Model(tweets).
		Where("? = ?", bun.Ident("authorId"), userID).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get tweets by user id from db error: %w", err)
	}

	return tweets, nil
}

func (repo *PGRepo) DeleteTweetByID(ctx context.Context, id string) error {
	_, err := repo.db.NewDelete().
		Where("? = ?", bun.Ident("id"), id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("deleting tweet by id from db error: %w", err)
	}

	return nil
}
