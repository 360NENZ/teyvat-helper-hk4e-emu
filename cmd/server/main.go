package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-emu/pkg/game"
	"github.com/teyvat-helper/hk4e-emu/pkg/http"
)

var cfg config.Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/hk4e-emu")
	viper.AddConfigPath("$HOME/.hk4e-emu")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config file not found, using the default config")
			cfg = config.DefaultConfig
			return
		} else {
			log.Panic().Err(err).Msg("Failed to read config file")
		}
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panic().Err(err).Msg("Failed to decode config file")
	}
}

func main() {
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
	log.Info().Str("listen_addr", cfg.GameServer.Addr).Msg("GAME server is running")
	go func() {
		if err := httpsrv.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()
	log.Info().Str("listen_addr", cfg.HTTPServer.Addr).Msg("HTTP server is running")

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
