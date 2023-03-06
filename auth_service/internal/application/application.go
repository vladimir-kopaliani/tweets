package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/graphql"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/http_server/handlers"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/http_server/middlewares"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/logger"
	userservice "github.com/vladimir-kopaliani/tweets/auth_service/internal/repository/user_service"
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/service"

	httpserver "github.com/vladimir-kopaliani/tweets/auth_service/internal/http_server"

	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	// logger
	Logger interfaces.Logger
	// servers
	httpServer *httpserver.HTTPServer
	// service
	service interfaces.Servicer
	// repositories
	userServiceRepository interfaces.UserServicer
}

func New(ctx context.Context) Application {
	var err error
	app := Application{}

	// setting application
	cfg := newConfiguration()

	err = cfg.loadConfigurationFromEnvironment()
	if err != nil {
		panic(err)
	}

	err = cfg.loadConfigurationFromEnvFile()
	if err != nil {
		panic(err)
	}

	// creating logger
	err = app.setLogger(ctx, cfg)
	if err != nil {
		panic(err)
	}

	if err = app.setUserServiceRepository(ctx, cfg); err != nil {
		panic(err)
	}

	// service
	if err = app.setService(ctx, cfg); err != nil {
		panic(err)
	}

	// creating HTTP server
	if err = app.setHTTPServer(ctx, cfg); err != nil {
		panic(err)
	}

	return app
}

func (app Application) Launch(ctx context.Context) (err error) {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return app.httpServer.Launch(ctx)
	})

	return group.Wait()
}

func (app Application) Shutdown(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if app.userServiceRepository == nil {
			return nil
		}

		return app.userServiceRepository.Close(ctx)
	})

	group.Go(func() error {
		if app.httpServer == nil {
			return nil
		}

		return app.httpServer.Shutdown(ctx)
	})

	return group.Wait()
}

func (app *Application) setLogger(ctx context.Context, cfg configuration) error {
	app.Logger = logger.NewLogger(
		logger.Configuration{
			IsDebugMode: !cfg.IsProductionMode,
		},
	)

	return nil
}

func (app *Application) setHTTPServer(ctx context.Context, cfg configuration) error {
	graphqlServer, err := graphql.NewGraphQLServer(
		graphql.Configuration{
			Service:          app.service,
			Logger:           app.Logger,
			JWTSecret:        cfg.JWTSecret,
			IsProductionMode: cfg.isProducationMode(),
		},
	)
	if err != nil {
		return fmt.Errorf("creating GraphQL server error: %w", err)
	}

	restServer, err := handlers.New(handlers.Configuration{
		IsProductionMode: cfg.IsProductionMode,
		Service:          app.service,
		JWTSecret:        cfg.JWTSecret,
	})
	if err != nil {
		return fmt.Errorf("creating REST server error: %w", err)
	}

	// TODO: remove CORS
	corsHandler := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead, http.MethodGet, http.MethodPost,
				http.MethodPut, http.MethodPatch, http.MethodDelete,
			},
			Debug: true,
		},
	)

	l := log.New(os.Stdout, "[cors] ", log.LstdFlags)
	corsHandler.Log = l

	server, err := httpserver.New(ctx, httpserver.Configuration{
		Address: cfg.HTTPServerPort,
		Logger:  app.Logger,
		Service: app.service,
		Handlers: []httpserver.Handler{
			{
				Path: "/api/v1/auth",
				Handler: corsHandler.Handler(
					middlewares.InjectUserContext(
						restServer,
					),
				),
			},
			{
				Path: "/graphql",
				Handler: corsHandler.Handler(
					middlewares.InjectUserContext(
						graphqlServer,
					),
				),
			},
			{
				Path:    "/",
				Handler: http.NotFoundHandler(),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("creating HTTP server error: %w", err)
	}

	app.httpServer = server

	return nil
}

func (app *Application) setService(ctx context.Context, cfg configuration) error {
	serv, err := service.New(ctx, service.Configuration{
		Logger:      app.Logger,
		JWTSecret:   cfg.JWTSecret,
		UserService: app.userServiceRepository,
	})
	if err != nil {
		return fmt.Errorf("creating service error: %w", err)
	}

	app.service = serv

	return nil
}

func (app *Application) setUserServiceRepository(ctx context.Context, cfg configuration) (err error) {
	app.userServiceRepository, err = userservice.New(ctx, userservice.Configuration{
		Logger:  app.Logger,
		Address: cfg.UserServiceAddress,
	})
	if err != nil {
		return fmt.Errorf("creating new postgres_template repository error: %w", err)
	}

	return nil
}
