package httpserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/interfaces"
)

const (
	readTimeoutSeconds  = 15
	writeTimeoutSeconds = 60
)

// HTTPServer represent http server
type HTTPServer struct {
	log     interfaces.Logger
	server  *http.Server
	service Servicer
}

// New returns instance of HTTP server
func New(ctx context.Context, conf Configuration) (*HTTPServer, error) {
	srv := &HTTPServer{
		log:     conf.Logger,
		service: conf.Service,
	}

	err := conf.validate()
	if err != nil {
		return nil, fmt.Errorf("validate HTTP server configuration: %w", err)
	}

	baseContext := func(net.Listener) context.Context {
		return ctx
	}

	srv.server = &http.Server{
		Addr:         ":" + conf.Address,
		ReadTimeout:  readTimeoutSeconds * time.Second,
		WriteTimeout: writeTimeoutSeconds * time.Second,
		BaseContext:  baseContext,
	}

	router := http.NewServeMux()

	for i := range conf.Handlers {
		router.Handle(conf.Handlers[i].Path, conf.Handlers[i].Handler)
	}

	srv.server.Handler = router

	return srv, nil
}

// Launch starts HTTP server
func (s HTTPServer) Launch(ctx context.Context) error {
	s.log.Info(ctx, fmt.Sprintf("HTTP is listening on %s", s.server.Addr))

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("launching HTTP server error: %w", err)
	}

	return nil
}

// Shutdown graceful closing HTTP Server
func (s HTTPServer) Shutdown(ctx context.Context) error {
	s.log.Info(ctx, "HTTP server is shutting down...")

	if err := s.server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf(": %w", err)
	}

	s.log.Info(ctx, "HTTP server is off.")

	return nil
}
