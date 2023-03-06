package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"
)

// RegisteredAt is the resolver for the registeredAt field.
func (r *userResolver) RegisteredAt(ctx context.Context, obj *entities.User) (*entities.DateTime, error) {
	return (*entities.DateTime)(&obj.CreatedAt), nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
