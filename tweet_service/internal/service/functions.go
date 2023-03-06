package service

import (
	"context"
	"fmt"
	"time"

	apperrors "github.com/vladimir-kopaliani/tweets/tweet_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/tweet_service/internal/entities"

	"github.com/google/uuid"
)

func (s Service) CreateNewTweet(ctx context.Context, tweet *entities.Tweet) (*entities.Tweet, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("coudn't generate UUID: %w", err)
	}

	userID := getUserIDFromContext(ctx)

	tweet.ID = uuid.String()
	tweet.AuthorID = userID
	tweet.CreatedAt = time.Now()

	err = s.tweetsRepository.SaveNewTweet(ctx, tweet)
	if err != nil {
		return nil, fmt.Errorf("saving new tweet error: %w", err)
	}

	return tweet, err
}

func (s Service) GetTweetByID(ctx context.Context, id string) (*entities.Tweet, error) {
	tweet, err := s.tweetsRepository.GetTweetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting tweet by id error: %w", err)
	}

	return tweet, nil
}

func (s Service) GetTweetsByUserID(ctx context.Context, userID string) ([]*entities.Tweet, error) {
	tweets, err := s.tweetsRepository.GetTweetsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("getting tweets by user id error: %w", err)
	}

	return tweets, nil
}

func (s Service) DeleteTweetByID(ctx context.Context, id string) error {
	tweet, err := s.tweetsRepository.GetTweetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get tweet by id error: %w", err)
	}

	userID := getUserIDFromContext(ctx)

	if tweet.AuthorID != userID {
		return apperrors.ErrWrongCredentials
	}

	err = s.tweetsRepository.DeleteTweetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("deleting tweet by id error: %w", err)
	}

	return nil
}

func getUserIDFromContext(ctx context.Context) (userID string) {
	userContext, ok := ctx.Value(entities.UserContextKey).(entities.UserContext)
	if !ok {
		return ""
	}

	return userContext.UserID
}
