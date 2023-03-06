package graphql

//go:generate go run -mod=mod github.com/99designs/gqlgen generate

import (
	"context"
	"fmt"

	apperrors "github.com/vladimir-kopaliani/tweets/user_service/internal/app_errors"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewGraphQLServer(conf Configuration) (*handler.Server, error) {
	err := conf.validate()
	if err != nil {
		return nil, fmt.Errorf("validate GraphQL server configuration error: %w", err)
	}

	cfg := Config{
		Resolvers: &Resolver{
			Service: conf.Service,
			Logger:  conf.Logger,
		},
		Directives: DirectiveRoot{
			RequiredAuthentication: requiredAuthentication(conf.JWTSecret),
		},
	}

	srv := handler.NewDefaultServer(NewExecutableSchema(cfg))

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		conf.Logger.Error(ctx, fmt.Sprintf("GraphQL error: %v", err))

		// show only public errors
		if conf.IsProductionMode {
			if e, ok := err.(error); ok && !apperrors.IsPublicError(e) {
				err = apperrors.ErrInternal
			}
		}

		return &gqlerror.Error{
			Message: fmt.Sprint(err),
		}
	})

	return srv, nil
}

func requiredAuthentication(jwtSecret string) func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		userContext, ok := ctx.Value(entities.UserContextKey).(entities.UserContext)
		if !ok {
			return nil, apperrors.ErrExtractUserContext
		}

		if userContext.AcessToken == "" {
			return nil, apperrors.ErrAuthenticationRequired
		}

		token, err := jwt.Parse(userContext.AcessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			return nil, fmt.Errorf("error parsing JWT token: %w", err)
		}

		if !token.Valid {
			return nil, apperrors.ErrJWTInvalid
		}

		return next(ctx)
	}
}
