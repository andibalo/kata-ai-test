package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	pokemon_be "pokemon-be"
	"pokemon-be/internal/config"
	"pokemon-be/pkg/db"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.InitConfig()

	database := db.InitDB(cfg)

	server := pokemon_be.NewServer(cfg, database)

	cfg.Logger().Info().Msg(fmt.Sprintf("Server starting at port %s", cfg.AppAddress()))

	go func() {
		if err := server.Start(cfg.AppAddress()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			cfg.Logger().Fatal().Err(err).Msg("failed to start server")
		}
	}()

	<-ctx.Done()

	stop()
	cfg.Logger().Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		cfg.Logger().Fatal().Err(err).Msg("Server force to shutdown")
	}

	cfg.Logger().Info().Msg("Server exiting")
}
