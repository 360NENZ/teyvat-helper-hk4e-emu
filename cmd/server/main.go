package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-emu/pkg/game"
	"github.com/teyvat-helper/hk4e-emu/pkg/http"
)

func main() {
	cfg := config.LoadConfig()

	gamesrv := game.NewServer(&cfg)
	httpsrv := http.NewServer(&cfg)

	if err := httpsrv.LoadSecret(); err != nil {
		panic(err)
	}
	if err := gamesrv.LoadSecret(); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := gamesrv.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start GAME server")
		}
	}()
	go func() {
		if err := httpsrv.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	<-done

	{
		log.Info().Msg("HTTP server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()
		if err := httpsrv.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("Failed to gracefully shutdown HTTP server")
		}
		log.Info().Msg("HTTP server stopped")
	}
	{
		log.Info().Msg("GAME server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()
		if err := gamesrv.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("Failed to gracefully shutdown GAME server")
		}
		log.Info().Msg("GAME server stopped")
	}
}
