package app

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type Application struct {
	server *Server
}

func NewApplication(address string) *Application {
	return &Application{
		server: NewServer(address),
	}
}

func (a *Application) Start(ctx context.Context) error {
	log.Info().Msg("Starting the server on " + a.server.Address)

	go func() {
		if err := a.server.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start the server")
		}
	}()

	<-ctx.Done() // listen for the context to be canceled

	if a.server.IsRunning() {
		log.Info().Msg("Shutting down gracefully...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := a.server.Stop(shutdownCtx); err != nil {
			log.Fatal().Err(err).Msg("Failed to shutdown the server")
		}
	}

	return nil
}
