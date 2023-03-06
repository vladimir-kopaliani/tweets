package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/application"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// handle interrupt signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// launch application
	app := application.New(ctx)

	go func() {
		<-signalCh
		cancel()

		if err := app.Shutdown(ctx); err != nil {
			app.Logger.Error(ctx, fmt.Sprintf("error during shutdown: %s", err.Error()))
		}
	}()

	if err := app.Launch(ctx); err != nil {
		app.Logger.Error(ctx, fmt.Sprintf("error during shutdown: %s", err.Error()))
	}

	<-ctx.Done()
}
