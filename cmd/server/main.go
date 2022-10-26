package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			log.Println("Config file not found; using default config")
			cfg = config.DefaultConfig
			return
		} else {
			panic(fmt.Errorf("fatal error config file: %v", err))
		}
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}

func main() {
	gamesrv := game.NewServer(&cfg)
	httpsrv := http.NewServer(&cfg)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := gamesrv.Start(); err != nil {
			log.Fatalln("[game.Server] Failed to listen, error:", err)
		}
	}()
	log.Println("[game.Server] Server is running on", cfg.GameServer.Addr)
	go func() {
		if err := httpsrv.Start(); err != nil {
			log.Fatalln("[http.Server] Failed to listen, error:", err)
		}
	}()
	log.Println("[http.Server] Server is running on", cfg.HTTPServer.Addr)

	gamesrv.UpdateSeed()

	<-done

	{
		log.Println("[http.Server] Server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()
		if err := httpsrv.Shutdown(ctx); err != nil {
			log.Fatalln("[http.Server] Failed to gracefully shutdown, error:", err)
		}
		log.Println("[http.Server] Server stopped")
	}
	{
		log.Println("[game.Server] Server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()
		if err := gamesrv.Shutdown(ctx); err != nil {
			log.Fatalln("[game.Server] Failed to gracefully shutdown, error:", err)
		}
		log.Println("[game.Server] Server stopped")
	}
}
