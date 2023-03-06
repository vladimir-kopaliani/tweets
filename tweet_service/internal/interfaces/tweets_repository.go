package interfaces

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/tweet_service/internal/entities"
)

type TweetsRepositorier interface {
	SaveNewTweet(ctx context.Context, tweet *entities.Tweet) error
	GetTweetByID(ctx context.Context, id string) (*entities.Tweet, error)
	GetTweetsByUserID(ctx context.Context, userID string) ([]*entities.Tweet, error)
	DeleteTweetByID(ctx context.Context, id string) error

	Close(ctx context.Context) error
}
