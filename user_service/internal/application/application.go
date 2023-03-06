package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/graphql"
	httpserver "github.com/vladimir-kopaliani/tweets/user_service/internal/http_server"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/http_server/handlers"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/http_server/middlewares"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/logger"
	pgrepo "github.com/vladimir-kopaliani/tweets/user_service/internal/repository/users"
	rpcserver "github.com/vladimir-kopaliani/tweets/user_service/internal/rpc_server"
	"github.com/vladimir-kopaliani/tweets/user_service/internal/service"

	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	// logger
	Logger interfaces.Logger
	// servers
	httpServer *httpserver.HTTPServer
	rpcServer  *rpcserver.GRPCServer
	// repositories
	pgTemplateRepository interfaces.UsersRepositorier
	// service
	service interfaces.Servicer
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

	// repositories
	if err = app.setUsersTemplateRepository(ctx, cfg); err != nil {
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

	// creating gRPC server
	if err = app.setGRPCServer(ctx, cfg); err != nil {
		panic(err)
	}

	return app
}

func (app Application) Launch(ctx context.Context) (err error) {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return app.httpServer.Launch(ctx)
	})

	group.Go(func() error {
		return app.rpcServer.Launch(ctx)
	})

	return group.Wait()
}

func (app Application) Shutdown(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if app.pgTemplateRepository == nil {
			return nil
		}

		return app.pgTemplateRepository.Close(ctx)
	})

	group.Go(func() error {
		if app.httpServer == nil {
			return nil
		}

		return app.httpServer.Shutdown(ctx)
	})

	group.Go(func() error {
		if app.rpcServer == nil {
			return nil
		}

		return app.rpcServer.Shutdown(ctx)
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
				Path: "/api/v1/",
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

func (app *Application) setGRPCServer(ctx context.Context, cfg configuration) error {
	server, err := rpcserver.New(ctx, rpcserver.Configuration{
		Address: cfg.GRPCServerPort,
		Logger:  app.Logger,
		Service: app.service,
	})
	if err != nil {
		return fmt.Errorf("creating gRPC server error: %w", err)
	}

	app.rpcServer = server

	return nil
}

func (app *Application) setService(ctx context.Context, cfg configuration) error {
	serv, err := service.New(ctx, service.Configuration{
		Logger:       app.Logger,
		PGRepository: app.pgTemplateRepository,
		JWTSecret:    cfg.JWTSecret,
	})
	if err != nil {
		return fmt.Errorf("creating service error: %w", err)
	}

	app.service = serv

	return nil
}

func (app *Application) setUsersTemplateRepository(ctx context.Context, cfg configuration) (err error) {
	app.pgTemplateRepository, err = pgrepo.NewPGRepo(ctx, pgrepo.Configuration{
		Logger:  app.Logger,
		Address: cfg.UsersPostgresConnectionURL,
	})
	if err != nil {
		return fmt.Errorf("creating new users repository error: %w", err)
	}

	return nil
}
